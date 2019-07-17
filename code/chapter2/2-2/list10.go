package main

import "fmt"

/*
	将字典（map）类型内置，可直接从运行时层面获得性能优化
*/
func main() {
	m := make(map[string]string) // 创建字典类型对象

	m["key1"] = "value1" // 添加或者更新值

	v, ok := m["key"] // 使用ok-idiom获取值，可以知道key/value是否存在

	fmt.Println(v, ok)

	fmt.Println(m)
	delete(m, "key1")
	fmt.Println(m)
}

/*
	所谓okidiom模式，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功。因为很多操作默认返回零值，所以须额外说明
 */