package abstractions

type Validatable interface {
	Validate() error
}
