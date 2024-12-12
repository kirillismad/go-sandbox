package gof

type CookingProgramm interface {
	Cook()
}

type MeatCookingProgramm struct{}

func (m *MeatCookingProgramm) Cook() {
	println("Meat is cooked")
}

type FishCookingProgramm struct{}

func (f *FishCookingProgramm) Cook() {
	println("Fish is cooked")
}

type Program int

const (
	Meat Program = iota
	Fish
)

func NewCookingProgramm(kind Program) CookingProgramm {
	switch kind {
	case Meat:
		return &MeatCookingProgramm{}
	case Fish:
		return &FishCookingProgramm{}
	default:
		panic("invalid kind")
	}
}
