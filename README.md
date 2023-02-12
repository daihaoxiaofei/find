# Find
    在以太坊,BSC等网络上寻找有余额的地址和对应的私钥

    原文参考: https://blog.csdn.net/weixin_42608885/article/details/106261263

    大致思路是这样:
        1.先随机生成私钥,然后通过私钥生成公钥,再生成地址.
        2.把生成的地址通过节点获取对应余额,如果有价值则在本地保存当前地址及对应的私钥及余额,写入本地文件.

### 使用方法
    1 编译启动
        go build -a -installsuffix cgo -o find main.go
        ./find

    2 使用docker
        制作新镜像
          docker build -t find .
        启动
          docker-compose up -d

    复制 pkg/config/config.example.yaml 到根目录改名为 config.yaml
    编辑config.yaml  主要配置 Source-Urls下的链接  需要自己去各资源网站申请账号
    其他三个Bsc HECO OKEX默认就好了 若是自己搭建了节点也可以加入进去 
    若是不想扫描某个链直接空着就可以了
    Notify 主要用于配置扫描到结果时接收通知的 邮件 钉钉告警 飞书通知 任选

### 本次更新

    原有问题:
        程序协程和实际需要不匹配的问题:
            有时开的协程数很多,但第三方网站或本地节点限制大部分请求被拒绝
            有时开的协程数很小,没有达到最大使用量,浪费时间
        
    改进:
        增加了负反馈feedback模块 根据任务成功情况自动调整协程数量 以匹配实际用量
