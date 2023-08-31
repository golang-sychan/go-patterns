package main

import (
	"demo-workspace/go-patterns/demo/abstract_factory"
	"demo-workspace/go-patterns/pkg/builder"
	"fmt"
)

func main() {
	xf := abstract_factory.NewXiaomiFactory()
	lf := abstract_factory.NewLenovoFactory()
	xf.MakeComputer().Screen()
	lf.MakeComputer().Screen()
	xf.MakePhone().Call()
	lf.MakePhone().Call()
	//-------builder-----------
	mpvBuilder := builder.GetCarBuilder("mpv")
	suvBuilder := builder.GetCarBuilder("suv")
	director := builder.NewDirector(mpvBuilder)
	mpvcar := director.BuildCar()
	fmt.Printf("MPV 引擎:%s\n", mpvcar.GetEngineType())
	fmt.Printf("MPV 座椅类型:%s\n", mpvcar.GetSeatsType())
	fmt.Printf("MPV 座椅数量:%d\n", mpvcar.GetNum())
	director.SetBuilder(suvBuilder)
	suvCar := director.BuildCar()
	fmt.Printf("SUV 引擎:%s\n", suvCar.GetEngineType())
	fmt.Printf("SUV 座椅类型:%s\n", suvCar.GetSeatsType())
	fmt.Printf("SUV 座椅数量:%d\n", suvCar.GetNum())
}
