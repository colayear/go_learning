package main

import (
	"fmt"
	"time"
)

func main() {
	ptrSection1()
	ptrSection2()
	ptrSection3()
}

func ptrSection1() {
	fmt.Println("===指针基本知识===")
	// ========== 1. 基础：变量与指针的关系 ==========
	// 定义普通整型变量a，赋值为10
	// 变量a会占用一块内存，比如地址是 0xc00001a088，值是10
	a := 10
	fmt.Println("变量a的值：", a)
	fmt.Println("变量a的内存地址：", &a)

	// ========== 2. 声明并初始化指针变量 ==========
	// 声明指针变量p，类型为*int（表示指向int类型的指针）
	// 把变量a的地址赋值给指针p，此时p存储的是a的内存地址
	var p *int = &a
	fmt.Println("指针p的值（即a的地址）：", p)
	fmt.Println("指针p指向的值（即a的值）：", *p)

	// ========== 3. 通过指针修改原变量的值 ==========
	// 解引用后赋值：修改指针指向地址的内容，原变量a的值会同步变化
	*p = 20
	fmt.Println("修改后a的值：", a)
	fmt.Println("修改后*p的值：", *p)

	// ========== 4. 空指针：未指向任何有效地址的指针 ==========
	// 声明指针但未初始化，默认值为nil（空指针）
	var emptyP *int
	fmt.Println("空指针的值：", emptyP)

	// 注意：空指针不能直接解引用
	// 空指针安全处理：使用前先判断
	if emptyP != nil {
		fmt.Println(*emptyP)
	} else {
		fmt.Println("emptyP是空指针，无法解引用")
	}
	fmt.Println()
}

func ptrSection2() {
	fmt.Println("===值传递与指针传递===")
	// 测试基础类型传参
	num := 10
	fmt.Println("调用前num的值：", num) // 输出：10
	modifyByValue(num)
	fmt.Println("调用modifyByValue后num的值：", num) // 输出：10（未被修改）

	modifyByPointer(&num)
	fmt.Println("调用modifyByPointer后num的值：", num) // 输出：100（被修改）
	fmt.Println()
}

// 接收int类型参数（值传递）：函数内修改的是拷贝的临时变量，原变量不受影响
func modifyByValue(x int) {
	x = 100
	fmt.Println("函数内modifyByValue的x值：", x)
}

// 接收*int类型参数（指针传递）：函数内修改的是原变量的地址指向的值
func modifyByPointer(x *int) {
	// 解引用后修改原变量的值
	*x = 100
	fmt.Println("函数内modifyByPointer的*x值：", *x)
}

func ptrSection3() {
	fmt.Println("===指针传递与值传递的性能测试===")
	// 执行性能测试
	testPerformance()

	fmt.Println("\n小类型的传参优化")
	num := 10
	// 对于int、bool、float等小类型（≤8字节），值传递可能比指针更快！
	// 原因：指针需要解引用（额外CPU操作），而值传递直接拷贝8字节，开销更低
	testSmallType(num, &num)
	fmt.Println()
}

type BigStruct struct {
	Name string
	Age  int
	// 放大数据体积，让测试结果更明显
	Data [1024 * 1024]int
}

// 性能：每次调用需拷贝4MB+数据，耗时且占内存
func processByValue(bs BigStruct) {
	bs.Age = 30
	// 模拟业务处理（放大耗时差异）
	for i := 0; i < len(bs.Data); i++ {
		bs.Data[i] = i
	}
}

// 性能：仅拷贝8字节地址，所有操作直接作用于原内存，无额外拷贝
func processByPointer(bs *BigStruct) {
	bs.Age = 30
	for i := 0; i < len(bs.Data); i++ {
		bs.Data[i] = i
	}
}

func testPerformance() {

	var bigData0 BigStruct
	bigData0.Name = "test_data"
	bigData0.Age = 20
	fmt.Println("\n【数据修改验证】")
	fmt.Println("原结构体初始Age：20")
	processByValue(bigData0)
	fmt.Println("值传递后Age：", bigData0.Age) // 输出20（未修改）
	processByPointer(&bigData0)
	fmt.Println("指针传递后Age：", bigData0.Age) // 输出30（已修改）
	fmt.Println()

	var bigData BigStruct
	bigData.Name = "test_data"
	bigData.Age = 20

	testCount := 100
	startValue := time.Now()
	for i := 0; i < testCount; i++ {
		processByValue(bigData)
	}
	durationValue := time.Since(startValue)

	startPointer := time.Now()
	for i := 0; i < testCount; i++ {
		processByPointer(&bigData)
	}
	durationPointer := time.Since(startPointer)

	fmt.Printf("【性能对比】测试次数：%d次\n", testCount)
	fmt.Printf("值传递总耗时：%v\n", durationValue)
	fmt.Printf("指针传递总耗时：%v\n", durationPointer)
	fmt.Printf("指针传递比値传递快：%.2f倍\n", float64(durationValue.Nanoseconds())/float64(durationPointer.Nanoseconds()))
}

func testSmallType(val int, ptr *int) {
	count := 100000000

	startVal := time.Now()
	for i := 0; i < count; i++ {
		_ = val + 1
	}
	durVal := time.Since(startVal)

	startPtr := time.Now()
	for i := 0; i < count; i++ {
		_ = *ptr + 1
	}
	durPtr := time.Since(startPtr)
	fmt.Printf("测试次数：%v\n", count)
	fmt.Printf("小类型（int）值传递耗时：%v\n", durVal)
	fmt.Printf("小类型（int）指针传递耗时：%v\n", durPtr)
	fmt.Println("小类型优先值传递，大类型优先指针传递")
}
