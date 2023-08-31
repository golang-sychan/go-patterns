package factory

// AbstractProduct 定义产品类型使用
type AbstractProduct interface {
	GetName() string
}

type AbstractFactory interface {
	CreateProduct() AbstractProduct
}

type ConcreteProduct struct {
}

func (cp *ConcreteProduct) GetName() string {
	return "anta"
}

type ConcreteFactory struct {
}

func NewConcreteFactory() ConcreteFactory {
	return ConcreteFactory{}
}

func (cf *ConcreteFactory) CreateProduct() AbstractProduct {
	return &ConcreteProduct{}
}
