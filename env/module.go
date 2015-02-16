package env

// Module interface represents a functional package
type Module interface {
	// Register takes a Registrar
	Register(r Registrar)
	// Spec() *Spec
}
