package processors

type Processor interface {
	Configure() error
	Run(string) error
	// UnmarshalSettings([]byte) Processor
}
