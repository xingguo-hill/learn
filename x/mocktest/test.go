package mocktest

type A struct {
	y string
}

type B struct {
	a A
}

func NewBObj(abj A) MyInterfaceB {
	return &B{a: abj}
}

type MyInterfaceB interface {
	SomeMethodB(append string) string
}

func (b *B) SomeMethodB(append string) string {
	return b.a.y + append
}
