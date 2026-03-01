package main

import "fmt"

func main() {
	functionSection1()
	functionSection2()
	functionSection3()
}

// 匿名函数：没有函数名的函数，是 “一次性 / 临时性” 的函数，无法单独定义，只能 “定义时调用” 或 “赋值给变量后调用
func functionSection1() {
	// ========== 方式1：定义后立即调用（一次性使用） ==========
	// 场景：临时执行一段简单逻辑，无需复用
	func(a, b int) {
		fmt.Println("【匿名函数-立即调用】计算a+b：", a+b)
	}(5, 3) // 括号直接传参，执行后销毁，输出：8

	// ========== 方式2：赋值给变量，多次调用（常用） ==========
	// 场景：逻辑需要复用，但不想定义全局普通函数（减少命名污染）
	addFunc := func(a, b int) int {
		return a + b
	}
	// 调用变量指向的匿名函数
	fmt.Println("【匿名函数-变量调用】10+20 =", addFunc(10, 20)) // 输出：30
	fmt.Println("【匿名函数-变量调用】20+30 =", addFunc(20, 30)) // 输出：50

	// 匿名函数可以访问外层变量（如main中的num）
	num := 100
	func() {
		num += 50
		fmt.Println("【匿名函数-访问外层变量】num =", num) // 输出：150
	}()
}

// 高阶函数：接收函数作为参数
func functionSection2() {
	// 传入普通函数作为处理逻辑
	res1 := processNum(5, double)
	fmt.Println("【高阶函数】翻倍结果：", res1) // 输出：10

	res2 := processNum(5, square)
	fmt.Println("【高阶函数】平方结果：", res2) // 输出：25

	// 直接传入匿名函数（更灵活，无需定义普通函数）
	res3 := processNum(5, func(num int) int {
		return num + 10 // 自定义逻辑：加10
	})
	fmt.Println("【高阶函数】加10结果：", res3)

	// 场景：创建“加5”的函数
	add5 := createAdder(5)
	fmt.Println("【高阶函数-返回函数】10+5 =", add5(10))
	fmt.Println("【高阶函数-返回函数】20+5 =", add5(20))

	// 场景：创建“加10”的函数（复用createAdder逻辑）
	add10 := createAdder(10)
	fmt.Println("【高阶函数-返回函数】10+10 =", add10(10))
}

// 功能：对一个数字执行自定义处理逻辑
// 参数1：要处理的数字；参数2：处理逻辑（函数类型：func(int) int）
func processNum(num int, handleFunc func(int) int) int {
	fmt.Println("【高阶函数】开始处理数字：", num)
	// 调用传入的函数，执行自定义逻辑
	return handleFunc(num)
}

// 定义两个普通函数，作为处理逻辑
// 逻辑1：数字翻倍
func double(num int) int {
	return num * 2
}

// 逻辑2：数字平方
func square(num int) int {
	return num * num
}

// ========== 高阶函数：返回一个函数 ==========
// 功能：创建一个“指定加数”的加法函数（函数工厂）
// 参数：要加的固定数值
// 返回值：一个加法函数（接收int，返回int）
func createAdder(fixedNum int) func(int) int {
	// 返回一个匿名函数，作为结果
	return func(num int) int {
		return num + fixedNum
	}
}

// 闭包函数： 匿名函数 + 该函数捕获的外部变量（即使外部函数执行完毕，变量也不会销毁）
func functionSection3() {
	// 场景1：创建第一个计数器（闭包1）
	c1 := counter()
	// 调用闭包，count持续累加（背包里的count独立）
	fmt.Println("【闭包-计数器1】第1次调用：", c1())
	fmt.Println("【闭包-计数器1】第2次调用：", c1())
	fmt.Println("【闭包-计数器1】第3次调用：", c1())

	// 场景2：创建第二个计数器（闭包2）
	// 闭包2有自己的count变量（独立背包），和c1互不干扰
	c2 := counter()
	fmt.Println("【闭包-计数器2】第1次调用：", c2())
	fmt.Println("【闭包-计数器1】第4次调用：", c1())
}

// ========== 外部函数：生成闭包 ==========
func counter() func() int {
	// 外部变量：被闭包捕获（背包里的东西）
	count := 0 // 这个变量不会随counter函数执行完毕而销毁

	// 返回匿名函数（闭包）：携带count变量
	return func() int {
		count++ // 闭包可以修改捕获的变量
		return count
	}
}

// 为什么要有闭包函数：
// 特性			通俗解释											实战价值
// 变量私有化		捕获的变量（如 count）只能通过闭包修改，外部无法直接访问	模拟 “私有变量”，避免全局变量污染
// 环境独立性		每个闭包实例的变量独立（c1 和 c2 的 count 互不干扰）	批量创建独立状态的函数（如多个计数器）
// 生命周期延长	外部函数执行完毕后，捕获的变量不会被 GC 销毁			保持状态（如计数器、累加器）
