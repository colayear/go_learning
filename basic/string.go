package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	// ===================== 字符串的声明方式 =====================
	// 方式1：标准声明（显式类型）
	var str1 string = "Hello Go语言"
	// 方式2：类型推导
	var str2 = "Hello World！"
	// 方式3：短变量声明（函数内常用）
	str3 := `多行字符串
	支持换行、制表符\t
	无需转义反斜杠\` // 反引号：原生字符串，不解析转义符，适合SQL/配置

	fmt.Println("=== 字符串声明 ===")
	fmt.Printf("str1: %s\n", str1)
	fmt.Printf("str2: %s\n", str2)
	fmt.Printf("str3:\n%s\n\n", str3)

	// ===================== 字符串的底层本质 =====================
	// 知识点1：字符串是字节序列，len()返回字节数（不是字符数）
	// 中文UTF-8占3字节，英文占1字节
	str4 := "Go语言"
	fmt.Println("=== 字符串底层特性 ===")
	fmt.Printf("字符串：%s\n", str4)
	fmt.Printf("字节数(len)：%d\n", len(str4))            // 2(Go) + 3*2(语言) = 8
	fmt.Printf("占用内存大小：%d 字节\n", unsafe.Sizeof(str4)) // 字符串变量本身占16字节（指针+长度）

	// 知识点2：字符串不可变（直接修改字符会报错）
	// 错误示例（注释掉，运行会报错）：
	// str4[0] = 'g' // 报错：cannot assign to str4[0]

	// 知识点3：通过索引访问的是字节，不是字符（中文会乱码）
	fmt.Printf("索引0的字节：%c（正确，G的ASCII码）\n", str4[0])
	fmt.Printf("索引2的字节：%c（乱码，'语'的第一个字节）\n\n", str4[2])

	// ===================== 字符类型（rune）：处理中文等多字节字符 =====================
	// rune是int32的别名，代表一个UTF-8字符（无论占几个字节）
	// 遍历字符串的正确方式：for range（按rune遍历）
	fmt.Println("=== 按rune遍历字符串（正确处理中文） ===")
	for index, char := range str4 {
		fmt.Printf("索引：%d，字符：%c（rune值：%d）\n", index, char, char)
	}

	// 统计字符数（不是字节数）
	charCount := 0
	for range str4 {
		charCount++
	}
	fmt.Printf("字符数：%d\n\n", charCount) // 输出4（G、o、语、言）

	// ===================== 字符串拼接 =====================
	fmt.Println("=== 字符串拼接 ===")
	// 方式1：+ 拼接（简单场景，少量拼接）
	str5 := "Hello" + " " + "Go"
	fmt.Printf("+拼接：%s\n", str5)

	// 方式2：fmt.Sprintf（灵活，支持格式化，中等量拼接）
	name := "张三"
	age := 25
	str6 := fmt.Sprintf("姓名：%s，年龄：%d", name, age)
	fmt.Printf("Sprintf拼接：%s\n", str6)

	// 方式3：strings.Builder（高性能，大量拼接/循环拼接）
	// 工程化推荐：循环拼接时用Builder，避免+拼接的内存浪费
	var builder strings.Builder
	for i := 0; i < 3; i++ {
		builder.WriteString("Go")
		builder.WriteByte(' ') // 写入单个字节
	}
	str7 := builder.String()
	fmt.Printf("Builder拼接：%s\n\n", str7)

	// ===================== 字符串截取（切片） =====================
	// 知识点：截取基于字节索引，需注意中文边界（避免截断UTF-8字符）
	fmt.Println("=== 字符串截取 ===")
	str8 := "Go语言编程"
	// 截取前2个字节（Go）
	sub1 := str8[0:2] // 等价于 str8[:2]
	fmt.Printf("截取前2字节：%s\n", sub1)

	// 截取"语言"："语"从第2字节开始，占3字节；"言"占3字节 → 2~8字节
	sub2 := str8[2:8]
	fmt.Printf("截取语言：%s\n", sub2)

	// 错误示例：截断中文（第3字节开始，只取2字节 → 乱码）
	sub3 := str8[2:4]
	fmt.Printf("错误截取（截断中文）：%s\n\n", sub3)

	// ===================== 字符串修改（间接修改） =====================
	// 知识点：字符串不可变，需转为[]byte/[]rune修改后转回
	fmt.Println("=== 字符串修改 ===")
	str9 := "Hello Go"
	// 方式1：转为[]byte（适合纯ASCII字符）
	byteSlice := []byte(str9)
	byteSlice[0] = 'h' // 把H改为h
	newStr1 := string(byteSlice)
	fmt.Printf("修改ASCII字符：%s\n", newStr1)

	// 方式2：转为[]rune（适合含中文的字符串）
	str10 := "Go语言"
	runeSlice := []rune(str10)
	runeSlice[2] = '文' // 把"语"改为"文"
	newStr2 := string(runeSlice)
	fmt.Printf("修改中文字符：%s\n\n", newStr2)

	// ===================== 字符串分割与合并 =====================
	fmt.Println("=== 字符串分割与合并 ===")
	// 分割：strings.Split（按分隔符拆分）
	str11 := "苹果,香蕉,橙子"
	fruits := strings.Split(str11, ",")
	fmt.Printf("分割结果：%v\n", fruits)

	// 合并：strings.Join（高性能拼接切片）
	joinStr := strings.Join(fruits, "|")
	fmt.Printf("合并结果：%s\n\n", joinStr)

	str := "Hello Go语言, Hello World"
	fmt.Printf("原始字符串：%s\n\n", str)

	// ===================== 查找与包含 =====================
	fmt.Println("=== 查找与包含 ===")
	// 判断是否包含子串
	hasGo := strings.Contains(str, "Go")
	fmt.Printf("是否包含Go：%t\n", hasGo)

	// 判断是否以指定前缀/后缀开头/结尾
	isStartWithHello := strings.HasPrefix(str, "Hello")
	isEndWithWorld := strings.HasSuffix(str, "World")
	fmt.Printf("是否以Hello开头：%t\n", isStartWithHello)
	fmt.Printf("是否以World结尾：%t\n", isEndWithWorld)

	// 查找子串首次/末次出现的索引（-1表示未找到）
	indexGo := strings.Index(str, "Go")
	lastIndexHello := strings.LastIndex(str, "Hello")
	fmt.Printf("Go首次出现索引：%d\n", indexGo)
	fmt.Printf("Hello末次出现索引：%d\n\n", lastIndexHello)

	// ===================== 替换 =====================
	fmt.Println("=== 替换 ===")
	// 替换所有匹配的子串
	replaceAll := strings.ReplaceAll(str, "Hello", "Hi")
	fmt.Printf("替换所有Hello为Hi：%s\n", replaceAll)

	// 替换指定次数的子串（n=-1表示全部）
	replaceN := strings.Replace(str, "Hello", "Hi", 1)
	fmt.Printf("替换1次Hello为Hi：%s\n\n", replaceN)

	// ===================== 大小写转换 =====================
	fmt.Println("=== 大小写转换 ===")
	upperStr := strings.ToUpper("hello go")
	lowerStr := strings.ToLower("HELLO GO")
	fmt.Printf("转大写：%s\n", upperStr)
	fmt.Printf("转小写：%s\n\n", lowerStr)

	// ===================== 去除空白字符 =====================
	fmt.Println("=== 去除空白字符 ===")
	spaceStr := "  \tGo语言 \n  "
	// TrimSpace：去除首尾的空格、制表符、换行符
	trimSpace := strings.TrimSpace(spaceStr)
	fmt.Printf("去除首尾空白：'%s'\n", trimSpace)

	// Trim：去除首尾指定字符
	trimChar := strings.Trim("###Go###", "#")
	fmt.Printf("去除首尾#：%s\n\n", trimChar)

	// ===================== 其他高频操作 =====================
	fmt.Println("=== 其他高频操作 ===")
	// 重复字符串
	repeatStr := strings.Repeat("Go", 3)
	fmt.Printf("重复3次Go：%s\n", repeatStr)

	// 统计子串出现次数
	countHello := strings.Count(str, "Hello")
	fmt.Printf("Hello出现次数：%d\n", countHello)

	// ===================== 字符串 ↔ 整数 =====================
	fmt.Println("=== 字符串 ↔ 整数 ===")
	// 字符串转整数：Atoi（ParseInt的简化版，默认十进制）
	numStr := "123"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("字符串转整数失败：", err)
	} else {
		fmt.Printf("字符串'%s'转整数：%d（类型：%T）\n", numStr, num, num)
	}

	// 整数转字符串：Itoa（FormatInt的简化版）
	num2 := 456
	str01 := strconv.Itoa(num2)
	fmt.Printf("整数%d转字符串：'%s'（类型：%T）\n\n", num2, str01, str01)

	// 进阶：指定进制转换（如二进制/十六进制）
	hexStr := "64"
	// ParseInt(字符串, 进制, 位数)：位数0表示自动适配，64表示int64
	hexNum, _ := strconv.ParseInt(hexStr, 16, 0)
	fmt.Printf("十六进制'%s'转整数：%d\n", hexStr, hexNum)

	// ===================== 字符串 ↔ 浮点数 =====================
	fmt.Println("\n=== 字符串 ↔ 浮点数 ===")
	floatStr := "3.1415926"
	// ParseFloat(字符串, 位数)：64表示float64，32表示float32
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("字符串转浮点数失败：", err)
	} else {
		fmt.Printf("字符串'%s'转浮点数：%f（类型：%T）\n", floatStr, floatNum, floatNum)
	}

	// 浮点数转字符串：FormatFloat（指定格式和精度）
	float2 := 2.71828
	// FormatFloat(浮点数, 格式, 精度, 位数)
	// 格式：'f'=普通浮点，'e'=科学计数法
	str02 := strconv.FormatFloat(float2, 'f', 4, 64)
	fmt.Printf("浮点数%f转字符串（保留4位）：'%s'\n\n", float2, str02)

	// ===================== 字符串 ↔ 布尔值 =====================
	fmt.Println("=== 字符串 ↔ 布尔值 ===")
	// 字符串转布尔值：仅"true"/"false"（大小写敏感）
	boolStr1 := "true"
	boolVal1, _ := strconv.ParseBool(boolStr1)
	fmt.Printf("字符串'%s'转布尔值：%t\n", boolStr1, boolVal1)

	boolStr2 := "False" // 小写false才有效
	_, err = strconv.ParseBool(boolStr2)
	if err != nil {
		fmt.Printf("字符串'%s'转布尔值失败：%s\n", boolStr2, err)
	}

	// 布尔值转字符串
	boolVal3 := false
	str03 := strconv.FormatBool(boolVal3)
	fmt.Printf("布尔值%t转字符串：'%s'\n\n", boolVal3, str03)

	// ===================== 转换失败处理 =====================
	fmt.Println("=== 转换失败处理 ===")
	// 错误场景：非数字字符串转整数
	badStr1 := "abc123"
	badNum, err := strconv.Atoi(badStr1)
	if err != nil {
		// 业务中必须处理错误，避免程序崩溃
		fmt.Printf("转换失败：%s → 错误信息：%v\n", badStr1, err)
	} else {
		fmt.Println("转换结果：", badNum)
	}

	// ===================== 高性能拼接 =====================
	fmt.Println("=== 高性能拼接 ===")
	// 反例：循环中用+拼接（每次创建新字符串，内存浪费）
	var badStr2 string
	for i := 0; i < 1000; i++ {
		badStr2 += "a" // 性能差！
	}

	// 正例：用strings.Builder（内存复用，性能提升10倍+）
	var builder1 strings.Builder
	for i := 0; i < 1000; i++ {
		builder1.WriteByte('a')
	}
	goodStr := builder1.String()
	fmt.Printf("Builder拼接结果长度：%d\n\n", len(goodStr))

	// ===================== 中文编码兼容 =====================
	fmt.Println("=== 中文编码兼容 ===")
	// 正确遍历中文：for range（按rune）
	chineseStr := "Go语言编程"
	for _, c := range chineseStr {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

	// 正确截取中文：先转rune切片，再截取，最后转回字符串
	runeSlice1 := []rune(chineseStr)
	// 截取前3个字符（Go语）
	chineseSub := string(runeSlice1[:3])
	fmt.Printf("截取前3个中文字符：%s\n\n", chineseSub)

	// ===================== 空字符串判断 =====================
	fmt.Println("=== 空字符串判断 ===")
	// 正确写法：直接用 == ""（最简单、性能最高）
	emptyStr := ""
	if emptyStr == "" {
		fmt.Println("字符串为空")
	}

	// 错误写法：len(str) == 0（虽然结果对，但可读性差）
	if len(emptyStr) == 0 {
		fmt.Println("len判断字符串为空")
	}

	// 注意：空格字符串不是空字符串
	spaceStr1 := "   "
	if spaceStr1 == "" {
		fmt.Println("空格字符串为空")
	} else {
		fmt.Println("空格字符串不为空（需用TrimSpace处理）")
	}
}
