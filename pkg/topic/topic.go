package topic

import "pratikshakuldeep456/pub-sub-system/pkg/subscriber"

type Topic struct {
	Subscribers []subscriber.Subscriber
}

func (t *Topic) RegisterSubscriber(s subscriber.Subscriber) {

	t.Subscribers = append(t.Subscribers, s)

}
