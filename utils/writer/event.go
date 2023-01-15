package writer

type Event struct {
	Host       string      `json:"host,omitempty"`
	Index      string      `json:"index,omitempty"`
	Source     string      `json:"source,omitempty"`
	SourceType string      `json:"sourcetype,omitempty"`
	Time       string      `json:"time,omitempty"`
	Event      interface{} `json:"event"`
}

func NewEvent(event interface{}) *Event {
	return &Event{Event: event}
}
