package jsonhelper

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{"W":1561651}`

var obj struct {
	W int
}

// 对比下 IJson 和 系统自带的 json 的效率
// go test -bench=_QE_ -benchmem -run=^$
// -benchtime 默认为1秒  -benchmem 获得内存分配的统计数据
func Benchmark_QE_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Unmarshal([]byte(jsonStr), &obj)
	}
}

func Benchmark_QE_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal([]byte(jsonStr), &obj)
	}
}

// Benchmark_QE_1-6         2252055               531.2 ns/op           240 B/op          6 allocs/op
// Benchmark_QE_2-6         7650109               158.7 ns/op            16 B/op          1 allocs/op
// 确实快了很多  不知道复杂的结构会不会不一样
