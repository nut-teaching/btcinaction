package main

import "fmt"

// 可以为当前包内的任意类型定义方法
// 名称前的参数称作receiver
func (u *user) TheBirthDay() {
	u.age++
}

// 还可直接调用匿名字段的方法，这种方式可实现与继承类似的功能
func (u user) ToString() string {
	return fmt.Sprintf(
		"my name is %v, i am %v year old", u.name, u.age)
}

func main() {
	var s student

	s.name = "zhangsan"
	s.age = 10

	fmt.Println(s.ToString())

	var sp Speaker

	s.Speak()

	sp = s //只要包含接口所需的全部方法，即表示实现了该接口
	sp.Speak()
}

/*
	接口
	接口采用了duck type方式，也就是说无须在实现类型上添加显式声明
	当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。”
	我们并不关心对象是什么类型，到底是不是鸭子，只关心行为

	只要包含接口所需的全部方法，即表示实现了该接口
*/
// 定义一个接口类型
type Speaker interface {
	Speak()
}

func (u student) Speak() {
	fmt.Println("i am the king of the world")
}
