package builder

type Car struct {
	seatsType  string
	engineType string
	number     int
}

func (c Car) GetSeatsType() string {
	return c.seatsType
}

func (c Car) GetEngineType() string {
	return c.engineType
}

func (c Car) GetNum() int {
	return c.number
}

type CarBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

func GetCarBuilder(builderType string) CarBuilder {
	switch builderType {
	case "mpv":
		return NewMPVBuilder()
	case "suv":
		return NewSUVBuilder()
	}
	return nil
}

type MPVBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func NewMPVBuilder() *MPVBuilder {
	return &MPVBuilder{}
}

func (mpv *MPVBuilder) SetSeatsType() {
	mpv.SeatsType = "aircraft"
}

func (mpv *MPVBuilder) SetEngineType() {
	mpv.EngineType = "oil"
}

func (mpv *MPVBuilder) SetNumber() {
	mpv.Number = 7
}

func (mpv *MPVBuilder) GetCar() Car {
	return Car{
		seatsType:  mpv.SeatsType,
		engineType: mpv.EngineType,
		number:     mpv.Number,
	}
}

type SUVBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func NewSUVBuilder() *SUVBuilder {
	return &SUVBuilder{}
}

func (suv *SUVBuilder) SetSeatsType() {
	suv.SeatsType = "common"
}

func (suv *SUVBuilder) SetEngineType() {
	suv.EngineType = "electric"
}

func (suv *SUVBuilder) SetNumber() {
	suv.Number = 5
}

func (suv *SUVBuilder) GetCar() Car {
	return Car{
		seatsType:  suv.SeatsType,
		engineType: suv.EngineType,
		number:     suv.Number,
	}
}
