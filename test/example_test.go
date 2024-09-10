package test

import (
	"demo_test/say"
)

// 测试函数就是一个普通的函数，
// 不过示例测试主要是由Output注释来体现的，待测试函数只有一行输出时，使用Output注释来检测输出
func ExampleHello() {
	say.Hello()
	// Output:
	// hello
}

func ExampleGoodBye() {
	say.GoodBye()
	// Output:
	// bye
}

func ExampleSay() {
	say.Hello()
	say.GoodBye()
	// Output:
	// hello
	// bye
}
