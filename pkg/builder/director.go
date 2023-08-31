package builder

type Director struct {
	builder CarBuilder
}

func NewDirector(b CarBuilder) *Director {
	return &Director{
		b,
	}
}

func (d *Director) SetBuilder(b CarBuilder) {
	d.builder = b
}

func (d *Director) BuildCar() Car {
	d.builder.SetNumber()
	d.builder.SetEngineType()
	d.builder.SetSeatsType()
	return d.builder.GetCar()
}
