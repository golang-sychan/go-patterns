package abstract_factory

import "fmt"

type XiaomiPhone struct {
}

func (x *XiaomiPhone) Call() {
	fmt.Printf("小米拯救世界!\n")
}

type XiaomiFactory struct {
}

func (xc *XiaomiFactory) MakePhone() AbstractPhone {
	return &XiaomiPhone{}
}

type XiaomiComputer struct {
}

func (xc *XiaomiComputer) Screen() {
	fmt.Printf("欢迎使用小米笔记本电脑\n")
}

func NewXiaomiFactory() AbstractFactory {
	return &XiaomiFactory{}
}

func (xc *XiaomiFactory) MakeComputer() AbstractComputer {
	return &XiaomiComputer{}
}
