package deepcopy

import "testing"

// 此处注意，声明的结构体一定要属性为public，否则无法序列化反序列化，json的编码解码是相同道理
type Person struct {
	Name string
	Age  int
}

// TestStructCopy 测试结构体深度复制
func TestStructCopy(t *testing.T) {
	src := Person{
		Name: "Jedeft",
		Age:  23,
	}
	var dst Person
	if err := DeepCopy(&dst, src); err != nil {
		t.Fatalf("err :%v ", err)
		return
	}
	if &dst == &src {
		t.Fatal("dst and src are the same ptr")
		return
	}
}

// TestMapCopy 测试map深度复制
func TestMapCopy(t *testing.T) {
	src := make(map[string]int)
	src["Jedeft"] = 23
	var dst map[string]int
	if err := DeepCopy(&dst, src); err != nil {
		t.Fatalf("err :%v ", err)
		return
	}
	if &dst == &src {
		t.Fatal("dst and src are the same ptr")
		return
	}
}

// TestIntCopy 测试基本类型深度复制
func TestIntCopy(t *testing.T) {
	src := 3
	var dst int
	if err := DeepCopy(&dst, src); err != nil {
		t.Fatalf("err :%v ", err)
		return
	}
	if &dst == &src {
		t.Fatal("dst and src are the same ptr")
		return
	}
}
