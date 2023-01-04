package processors

type Processor interface {
	Greet(string) string
	// UnmarshalSettings([]byte) Processor
}
