package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// ===================== 1. 有符号整数（可表示正数、负数、0） =====================
	// int8: 占1字节，范围 -128 ~ 127
	// var num8Err int8 = 128 // 报错：constant 128 overflows int8
	// int16: 占2字节，范围 -32768 ~ 32767
	// int32: 占4字节，范围 -2147483648 ~ 2147483647（常用作整型ID）
	// int64: 占8字节，范围 -9223372036854775808 ~ 9223372036854775807
	// int: 占用字节随系统变化（32位系统4字节，64位系统8字节），日常开发默认使用
	var num8 int8 = 127
	fmt.Printf("int8 类型 - 值：%d，占用字节：%d，取值范围：-128 ~ 127\n", num8, unsafe.Sizeof(num8))
	var num16 int16 = -32768
	fmt.Printf("int16 类型 - 值：%d，占用字节：%d，取值范围：-32768 ~ 32767\n", num16, unsafe.Sizeof(num16))
	var num32 int32 = 2147483647
	fmt.Printf("int32 类型 - 值：%d，占用字节：%d，取值范围：-2^31 ~ 2^31-1\n", num32, unsafe.Sizeof(num32))
	var num64 int64 = -9223372036854775808
	fmt.Printf("int64 类型 - 值：%d，占用字节：%d，取值范围：-2^63 ~ 2^63-1\n", num64, unsafe.Sizeof(num64))
	var num int = 100
	fmt.Printf("int 类型 - 值：%d，占用字节：%d（64位系统）\n\n", num, unsafe.Sizeof(num))

	// 2. 无符号整数（仅表示0和正数）
	// uint8 (byte): 占1字节，范围 0 ~ 255（常用作字节表示）
	// uint16: 占2字节，范围 0 ~ 65535
	// uint32: 占4字节，范围 0 ~ 4294967295
	// uint64: 占8字节，范围 0 ~ 18446744073709551615
	// uint: 占用字节随系统变化，无符号版int
	var u8 uint8 = 255
	fmt.Printf("uint8(byte) 类型 - 值：%d，占用字节：%d，取值范围：0 ~ 255\n", u8, unsafe.Sizeof(u8))
	var u16 uint16 = 65535
	fmt.Printf("uint16 类型 - 值：%d，占用字节：%d，取值范围：0 ~ 65535\n", u16, unsafe.Sizeof(u16))
	var u32 uint32 = 4294967295
	fmt.Printf("uint32 类型 - 值：%d，占用字节：%d，取值范围：0 ~ 2^32-1\n", u32, unsafe.Sizeof(u32))
	var u64 uint64 = 18446744073709551615
	fmt.Printf("uint64 类型 - 值：%d，占用字节：%d，取值范围：0 ~ 2^64-1\n", u64, unsafe.Sizeof(u64))
	var u uint = 100
	fmt.Printf("uint 类型 - 值：%d，占用字节：%d\n\n", u, unsafe.Sizeof(u))

	// ===================== 1. 浮点数基础声明 =====================
	// float32：占4字节，精度约6-7位有效数字
	var f32 float32 = 3.141592653589793
	// float64：占8字节，精度约15-16位有效数字
	var f64 float64 = 3.141592653589793

	fmt.Println("=== 浮点数精度对比 ===")
	fmt.Printf("float32 - 值：%.15f，占用字节：%d（精度丢失）\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("float64 - 值：%.15f，占用字节：%d（精度更高）\n\n", f64, unsafe.Sizeof(f64))

	// 2. 科学计数法声明
	// 场景：表示极大/极小的数（如天文数字、微观数值）
	var bigNum float64 = 1.23e9    // 1.23 × 10^9 = 1230000000
	var smallNum float64 = 4.56e-6 // 4.56 × 10^-6 = 0.00000456
	fmt.Println("=== 科学计数法 ===")
	fmt.Printf("1.23e9 = %f\n", bigNum)
	fmt.Printf("4.56e-6 = %f\n\n", smallNum)

	// 3. 浮点数核心：精度误差
	// 知识点1：0.1 + 0.2 ≠ 0.3（二进制存储导致的精度丢失）
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	sum := a + b
	fmt.Println("=== 浮点数精度坑点 ===")
	fmt.Printf("0.1 + 0.2 = %f\n", sum)
	fmt.Printf("0.3 = %f\n", c)
	// 错误写法：直接用 == 比较浮点数
	if sum == c {
		fmt.Println("sum == c（错误）")
	} else {
		fmt.Println("sum != c（精度误差导致）")
	}
	// 正确写法：判断差值是否小于极小值（如1e-9）
	// 1e-9 是工程中常用的“精度阈值”，可根据场景调整（如1e-6、1e-12）
	if math.Abs(sum-c) < 1e-9 {
		fmt.Println("sum 和 c 实际相等（差值 < 1e-9）\n")
	}
	// 金融场景示例：用整型存储金额（分）
	var amountCent int64 = 1001 // 10.01元
	amountYuan := float64(amountCent) / 100
	fmt.Printf("金融场景 - 分转元：%d 分 = %.2f 元\n\n", amountCent, amountYuan)

	// ===================== 不同进制的声明方式 =====================
	// 十进制：默认写法，无前缀
	var dec int = 100
	// 二进制：前缀 0b 或 0B
	var bin int = 0b1100100 // 二进制1100100 = 十进制100
	// 八进制：前缀 0o 或 0O（Go1.13+）
	var oct int = 0o144 // 八进制144 = 十进制100
	// 十六进制：前缀 0x 或 0X，支持0-9、a-f、A-F
	var hex int = 0x64 // 十六进制64 = 十进制100

	fmt.Println("=== 不同进制声明同一数值 ===")
	fmt.Printf("十进制 100 = %d\n", dec)
	fmt.Printf("二进制 0b1100100 = %d\n", bin)
	fmt.Printf("八进制 0o144 = %d\n", oct)
	fmt.Printf("十六进制 0x64 = %d\n\n", hex)

	// 进制格式化输出
	var number int = 255
	fmt.Println("=== 同一数值的不同进制输出 ===")
	fmt.Printf("十进制：%d\n", number)      // %d 十进制输出
	fmt.Printf("二进制：%b\n", number)      // %b 二进制输出
	fmt.Printf("八进制：%o\n", number)      // %o 八进制输出
	fmt.Printf("十六进制（小写）：%x\n", number) // %x 十六进制小写
	fmt.Printf("十六进制（大写）：%X\n", number) // %X 十六进制大写

	// 进制转换
	// 场景1：文件权限表示（八进制）
	// Linux文件权限如 0755：所有者读/写/执行，组和其他读/执行
	var filePerm uint16 = 0o755
	fmt.Printf("\n文件权限 0o755 十进制：%d，八进制：%o\n", filePerm, filePerm)

	// 场景2：颜色值表示（十六进制）
	// 十六进制0xFF0000 表示红色（RGB）
	var redColor int = 0xFF0000
	fmt.Printf("红色值 0xFF0000 十进制：%d，十六进制：%X\n", redColor, redColor)
}
