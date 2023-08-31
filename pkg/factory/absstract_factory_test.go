package factory

import "testing"

func TestNewConcreteFactory(t *testing.T) {
	factory := NewConcreteFactory()
	product := factory.CreateProduct()
	anta := product.GetName()
	if anta != "anta" {
		t.Fatal("not spec product")
	}
	t.Log("ok")
}
