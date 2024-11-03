package main

type Binary struct {
	uint64
}

type Stringer interface {
	String() string
}

func (b *Binary) String() string {
	return "Binary"
}

func main() {
	// a := Binary{64}
	a := &Binary{64}
	b := Stringer(a)

	str := b.String()
	println(str)
}
