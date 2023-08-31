package abstract_factory

import "fmt"

type LenovoPhone struct {
}

func (x *LenovoPhone) Call() {
	fmt.Printf("联想也有手机!\n")
}

type LenovoFactory struct {
}

func (xc *LenovoFactory) MakePhone() AbstractPhone {
	return &LenovoPhone{}
}

type LenovoComputer struct {
}

func (xc *LenovoComputer) Screen() {
	fmt.Printf("欢迎使用lenovo笔记本电脑\n")
}

func NewLenovoFactory() AbstractFactory {
	return &LenovoFactory{}
}

func (xc *LenovoFactory) MakeComputer() AbstractComputer {
	return &LenovoComputer{}
}
