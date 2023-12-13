package events

type Event struct {
	Name     string
	Start    float64
	Duration float64
	Parent   *Event
	Children []*Event
}

type EventTree struct {
	Root Event
}

func NewEventTree() *EventTree {
	return &EventTree{Root: Event{Name: "e2e", Parent: nil}}
}

func (e *Event) AddChildNode(child *Event) {
	e.Children = append(e.Children, child)
	child.Parent = e
}
