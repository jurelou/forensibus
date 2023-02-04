package writer

type Event struct {
	Host       string `json:"host,omitempty"`
	Index      string `json:"index,omitempty"`
	Source     string `json:"source,omitempty"`
	SourceType string `json:"sourcetype,omitempty"`
	Time       string `json:"time,omitempty"`
	Event      string `json:"event"`
}

func NewEvent(event string) *Event {
	return &Event{Event: event}
}
