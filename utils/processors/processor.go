package processors

type Processor interface {
	Greet() string
	// UnmarshalSettings([]byte) Processor
}
