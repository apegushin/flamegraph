package events

type Event struct {
	name     string
	start    float64
	duration float64
	children []Event
}

type EventTree struct {
	Root Event
}
