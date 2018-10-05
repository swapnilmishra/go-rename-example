package greeting

type Hello interface {
	SayHello() *FullName
}

type FullName struct {
	Name     string
	LastName string
}
