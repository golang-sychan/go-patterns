package main

import (
	"demo-workspace/go-patterns/demo/abstract_factory"
)

func main() {
	xf := abstract_factory.NewXiaomiFactory()
	lf := abstract_factory.NewLenovoFactory()
	xf.MakeComputer().Screen()
	lf.MakeComputer().Screen()
	xf.MakePhone().Call()
	lf.MakePhone().Call()
}
