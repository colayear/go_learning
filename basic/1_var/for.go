package main

import "fmt"

func main() {
	forSection1()
	forSection2()
}
func forSection1() {
	fmt.Println("===循环的写法===")
	// 普通写法
	fmt.Println("普通写法")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// while
	fmt.Println("while 写法")
	j := 1
	for j <= 10 {
		j++
		fmt.Println(j)
	}

	// 无限循环
	fmt.Println("无限循环写法")
	count := 1
	for {
		count++
		fmt.Println(count)
		if count >= 5 {
			break
		}
	}

	//range 写法遍历数组、map等
	fmt.Println("range 写法遍历数组、map等")
	fruits := []string{"apple", "banana", "orange"}
	// i：索引，v：对应的值；不需要的变量可以用_忽略
	for i, v := range fruits {
		fmt.Printf("for-range：索引%d，值%s\n", i, v)
	}

	fmt.Println("")
}

func forSection2() {
	fmt.Println("===循环的终止、退出===")
	// ========== 1. break：退出当前循环 ==========
	fmt.Println("===== break示例 =====")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Printf("i=%d，触发break，退出循环\n", i)
			break
		}
		fmt.Printf("i=%d\n", i)
	}

	// ========== 2. continue：跳过本次循环 ==========
	fmt.Println("\n===== continue示例 =====")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Printf("i=%d，触发continue，跳过本次循环\n", i)
			continue // 跳过本次循环，直接执行i++，不打印下面的内容
		}
		fmt.Printf("i=%d\n", i)
	}

	// ========== 3. 标签+break：退出多层循环（实战重点） ==========
	fmt.Println("===== 无标签break（仅退出内层循环） =====")
	// 去掉标签，仅用普通break
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 6 {
				fmt.Printf("i=%d, j=%d，触发普通break，仅退出内层循环\n", i, j)
				break
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
		// 内层循环退出后，外层循环会继续执行下一个i
		fmt.Printf("内层循环退出，外层循环i=%d继续\n", i)
	}

	fmt.Println("\n===== 标签+break（退出多层循环） =====")
	// 定义标签：标签名后加冒号，必须在循环上方
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 6 {
				fmt.Printf("i=%d, j=%d，触发outerLoop break，退出所有循环\n", i, j)
				break outerLoop // 退出标签对应的外层循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// ========== 4. goto：跳转到标签（慎用，仅特殊场景） ==========
	fmt.Println("\n===== goto示例 =====")
	num := 1
	for {
		if num >= 4 {
			goto loopEnd // 跳转到loopEnd标签位置
		}
		fmt.Printf("num=%d\n", num)
		num++
	}
	// 定义goto的目标标签
loopEnd:
	fmt.Println("goto跳转到这里，循环终止")

	// ========== 5. return：退出函数（循环也终止） ==========
	fmt.Println("\n===== return示例（调用函数） =====")
	testReturnInLoop()
}

// 测试循环内return
func testReturnInLoop() {
	for i := 1; i <= 5; i++ {
		if i == 2 {
			fmt.Printf("i=%d，触发return，退出函数\n", i)
			return // 直接退出函数，后续代码都不执行
		}
		fmt.Printf("i=%d\n", i)
	}
	fmt.Println("这段代码不会执行")
}
