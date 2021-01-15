package split

import (
	"reflect"
	"testing"
)

// 测试

//// 简单版，单个测试
//func TestSplit(t *testing.T) {
//	got := Split("我爱你你爱她她爱他他爱它", "爱")
//	want := []string{"我", "你", "你", "她", "她", "他", "他", "它"}
//	//got := Split("abc", "b")
//	//want := []string{"a", "c"}
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("want:%v got:%v", want, got)
//	}
//}


// 进阶版
func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep string
		want []string
	}
	tests := map[string]test {
		"chinese1":test{input: "包青天", sep: "青", want: []string{"包", "天"}},
		"chinese2":test{input: "牛肉面", sep: "面", want: []string{"牛肉", ""}},
		"english1":test{input: "shapeofyou", sep: "of", want: []string{"shape", "you"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s failed, want: %v got: %v", name, tc.want, got)
			}
		})
	}
}


// 性能基准测试
/*
结果：
goos: darwin
goarch: amd64
pkg: demo/34_test/split
BenchmarkSplit-12        6299355               177 ns/op             112 B/op          3 allocs/op
具体执行的线程              测试量               每次操作执行时间        每次操作消耗的内存    每次操作会做 3 次内存申请
PASS
ok      demo/34_test/split      3.546s
*/
func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i := 0; i < b.N; i++ {
		Split("吃葡萄不吐葡萄皮", "葡萄")
	}
}