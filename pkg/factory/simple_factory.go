package factory

// Factory 抽象
type Factory interface {
	MakeProduct() Product
}

// Product 抽象
type Product interface {
	GetSize() int
}

//具体实现

//产品

// Anta
type Anta struct {
}

func (a *Anta) GetSize() int {
	return 4
}

// ClothFactory 服装厂
type ClothFactory struct {
	name string
}

func (cf *ClothFactory) MakeProduct(name string) Product {
	switch name {
	case "anta":
		return &Anta{}
	}
	return nil
}
