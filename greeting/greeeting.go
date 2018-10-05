package greeting

type Greeting struct{}

func NewGreeting() *Greeting {
	return &Greeting{}
}

func (greet *Greeting) SayHello() *FullName {
	return &FullName{
		Name:     "Rob",
		LastName: "Thomson",
	}
}
