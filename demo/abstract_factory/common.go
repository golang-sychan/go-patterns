package abstract_factory

type AbstractPhone interface {
	Call()
}

type AbstractFactory interface {
	MakePhone() AbstractPhone
	MakeComputer() AbstractComputer
}

//------

type AbstractComputer interface {
	Screen()
}
