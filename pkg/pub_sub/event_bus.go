package pub_sub

type Event struct {
	EventType string
}

type PubSub struct {
	subscribers map[string][]string
}

func (ps *PubSub) Pub(event Event) {

}

func (ps *PubSub) Sub() {

}

func NewPubSub() *PubSub {
	return &PubSub{}
}
