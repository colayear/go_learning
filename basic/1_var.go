package main

import "fmt"

// ============= 变量声明及赋值 ==================
// 此文件展示变量的声明及赋值和go语言的特性
func main() {

	// 标准声明方式
	var name string = "张三"
	var age int = 25
	var score float64 = 98.5
	var isStudent bool = true
	// 打印变量值和类型（%T可以输出变量类型）
	fmt.Printf("标准声明 - name: %v, 类型: %T\n", name, name)
	fmt.Printf("标准声明 - age: %v, 类型: %T\n", age, age)
	fmt.Printf("标准声明 - score: %v, 类型: %T\n", score, score)
	fmt.Printf("标准声明 - isStudent: %v, 类型: %T\n\n", isStudent, isStudent)

	// 类型推导声明（自动推导）
	var address = "北京市海淀区"
	var height = 180.5
	var hasBook = false
	fmt.Printf("类型推导 - address: %v, 类型: %T\n", address, address)
	fmt.Printf("类型推导 - height: %v, 类型: %T\n", height, height)
	fmt.Printf("类型推导 - hasBook: %v, 类型: %T\n\n", hasBook, hasBook)

	// 短变量声明
	subject := "Go语言"
	grade := 99
	isPass := true
	fmt.Printf("短变量声明 - subject: %v, 类型: %T\n", subject, subject)
	fmt.Printf("短变量声明 - grade: %v, 类型: %T\n", grade, grade)
	fmt.Printf("短变量声明 - isPass: %v, 类型: %T\n\n", isPass, isPass)

	// 批量声明变量
	// 方式1：批量标准声明
	var (
		teacherName string  = "李老师"
		classCount  int     = 5
		classTime   float64 = 45.0
	)
	fmt.Printf("批量声明 - teacherName: %v, classCount: %v, classTime: %v\n\n", teacherName, classCount, classTime)
	// 方式2：批量类型推导声明
	var (
		city    = "上海"
		zipCode = 200000
	)
	fmt.Printf("批量推导 - city: %v, zipCode: %v\n\n", city, zipCode)

	// go语言有很多特点赋值方法
	// 多重赋值
	// 交换两个变量的值
	var a int = 10
	var b int = 20
	fmt.Printf("交换前 - a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后 - a: %d, b: %d\n\n", a, b)

	// 函数返回多个值时的多重赋值 定义一个返回两个值的函数
	getScore := func() (int, int) {
		return 85, 90 // 返回语文和数学成绩
	}
	chinese, math := getScore()
	fmt.Printf("多重赋值接收返回值 - 语文: %d, 数学: %d\n\n", chinese, math)

	// 变量的零值特性
	// 知识点：声明变量但不赋值时，Go会自动赋予对应类型的"零值"，避免空指针等问题（C/C++没有此特性）
	var emptyString string
	var emptyInt int
	var emptyFloat float64
	var emptyBool bool
	var emptyPtr *int
	fmt.Println("=== 变量零值展示 ===")
	fmt.Printf("string零值: '%v' (长度: %d)\n", emptyString, len(emptyString))
	fmt.Printf("int零值: %d\n", emptyInt)
	fmt.Printf("float64零值: %f\n", emptyFloat)
	fmt.Printf("bool零值: %t\n", emptyBool)
	fmt.Printf("指针零值: %v\n\n", emptyPtr)

	//  短变量重声明
	// 知识点：短变量声明":="允许同名变量重声明，但有严格条件：必须在同一作用域，且至少有一个新变量
	var num int = 100
	fmt.Printf("初始num: %d\n", num)
	// 合法重声明：同一作用域，同时声明新变量newNum，允许num重声明
	num, newNum := 200, 300
	fmt.Printf("重声明后 - num: %d, newNum: %d\n", num, newNum)
	// 错误示例（注释掉，运行时取消注释会报错）：
	// num := 400 // 报错：no new variables on left side of :=
	// 原因：没有新变量，单纯重声明同名变量不允许

	// 类型不兼容赋值
	var intVar int = 10
	// 错误示例（注释掉，运行时取消注释会报错）：
	// floatVar := intVar // 报错：cannot use intVar (type int) as type float64 in assignment
	// 正确做法：显式类型转换
	floatVar := float64(intVar)
	fmt.Printf("\n类型转换后 - intVar(%d) -> floatVar(%f)\n", intVar, floatVar)

	// 常量与变量的区别
	// 常量用const声明，值不可修改，编译期确定，常用于固定配置
	const PI = 3.1415926
	fmt.Printf("\n常量PI: %f\n", PI)
	// PI = 3.14 // 报错：cannot assign to PI (untyped float constant)
}
