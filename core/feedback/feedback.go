package feedback

import (
	"find/pkg/glog"
	"go.uber.org/zap"
	"sync"
	"time"
)

// Feedback  自适应协程数量
// 反馈为尚有空闲true 时 增加线程
// 反馈为已满false 时 减少线程
type Feedback struct {
	closeCh chan struct{} // 关闭信号

	pNum      int        // 协程数
	pNumMutex sync.Mutex // 锁

	succeedNumAll uint64    // 总成功的个数
	startTime     time.Time // 总成功的个数

	succeedNum int        // 成功的个数
	failedNum  int        // 失败的个数
	sfMutex    sync.Mutex // 锁

	workFun func() bool // 需要执行的任务函数 要求返回是否执行成功

	initPNum   int           // 初始协程数
	failedRate float32       // 失败率阈值(允许的最大失败率) 超过此值将开始关闭线程
	interval   time.Duration // 调度器间隔时间
	from       string        // 来源 用于日志打印

}

func NewFeedback(workFun func() bool, from string, initPNum int, failedRate float32) *Feedback {
	return &Feedback{
		closeCh:   make(chan struct{}),
		interval:  time.Second * 5,
		startTime: time.Now(),

		from:       from,
		initPNum:   initPNum,
		failedRate: failedRate,

		workFun: workFun,
	}
}

// Run 运行 阻塞式
func (f *Feedback) Run() {
	for i := 0; i < f.initPNum; i++ {
		go f.createGoroutine()
	}
	go f.FeedbackFun()
	go func() {
		t := time.NewTicker(time.Hour)
		for {
			<-t.C
			hours := time.Now().Sub(f.startTime).Hours()
			glog.Info(`调试-最佳初始协程数:`,
				zap.Uint64(`总成功`, f.succeedNumAll),
				zap.Int(`协程`, f.pNum),
				zap.Uint64(`每小时平均`, f.succeedNumAll/uint64(hours)),
				zap.String(`资源`, f.from),
			)
		}

	}()

	select {}
}

// 设置协程数
func (f *Feedback) setPNum(n int) {
	f.pNumMutex.Lock()
	defer f.pNumMutex.Unlock()
	f.pNum += n
}

// 读取协程数
func (f *Feedback) getPNum() int {
	f.pNumMutex.Lock()
	defer f.pNumMutex.Unlock()
	return f.pNum
}

// 成功++
func (f *Feedback) addSucceedNum() {
	f.sfMutex.Lock()
	defer f.sfMutex.Unlock()
	f.succeedNum++
	f.succeedNumAll++
}

// 失败++
func (f *Feedback) addFailedNum() {
	f.sfMutex.Lock()
	defer f.sfMutex.Unlock()
	f.failedNum++
}

// 返回成功和失败的数量并清空
func (f *Feedback) getSucceedFailedNum() (int, int) {
	f.sfMutex.Lock()
	defer f.sfMutex.Unlock()
	defer func() {
		f.succeedNum = 0
		f.failedNum = 0
	}()
	return f.succeedNum, f.failedNum
}

// 创建协程
func (f *Feedback) createGoroutine() {
	f.setPNum(1)
	for {
		select {
		case <-f.closeCh: // 关闭此线程信号
			if f.getPNum() == 1 { // 至少保留一个协程
				hours := time.Now().Sub(f.startTime).Hours()
				var average uint64
				if uint64(hours) == 0 {
					average = f.succeedNumAll
				} else {
					average = f.succeedNumAll / uint64(hours)
				}
				glog.Info(`可能是无效的:`,
					zap.Uint64(`总成功`, f.succeedNumAll),
					zap.String(`资源`, f.from),
					zap.Uint64(`每小时平均`, average),
				)
				time.Sleep(time.Minute)
				continue
			}
			f.setPNum(-1)
			return // 结束该协程
		default:
			if f.workFun() { // 任务执行成功
				f.addSucceedNum()
			} else {
				f.addFailedNum()
			}
		}
	}
}

// FeedbackFun 返回成功和失败的数量并清空
func (f *Feedback) FeedbackFun() (int, int) {
	t := time.NewTicker(f.interval)
	for {
		<-t.C
		succeedNum, failedNum := f.getSucceedFailedNum()
		glog.Debug(`调度:`,
			zap.Int(`协程`, f.pNum),
			zap.Int(`成功`, succeedNum),
			zap.Int(`成功`, succeedNum),
			zap.Int(`失败`, failedNum),
			zap.String(`资源`, f.from),
		)

		// 没有数据 保持不动
		if succeedNum == 0 && failedNum == 0 {
			continue
		}
		// 没有成功的
		if succeedNum == 0 {
			f.closeCh <- struct{}{}
			continue
		}
		// 没有失败的
		if failedNum == 0 {
			go f.createGoroutine()
			continue
		}
		// 失败率过高
		failedRate := float32(failedNum) / float32(succeedNum)
		if failedRate > f.failedRate {
			f.closeCh <- struct{}{}
		}
	}
}
