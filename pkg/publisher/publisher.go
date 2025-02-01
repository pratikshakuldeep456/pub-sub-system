package publisher

import (
	"errors"
	"fmt"
	"pratikshakuldeep456/pub-sub-system/pkg/topic"
)

type Publisher interface {
	RegisterTopic(topic topic.Topic) error
	DeRegisterTopic(topic topic.Topic) error
	Publish(message string) error
}

type Pub struct {
	Topics []*topic.Topic
}

func (p *Pub) RegisterTopic(t *topic.Topic) error {
	if t == nil {
		return errors.New(" topic cant be empty")
	}
	p.Topics = append(p.Topics, t)
	return nil
}

func (p *Pub) DeRegisterTopic(t *topic.Topic) error {
	if t == nil {
		return errors.New(" topic cant be empty")
	}
	len := len(p.Topics)
	for key := range len {
		if p.Topics[key] == t {
			p.Topics[len-1], p.Topics[key] = p.Topics[key], p.Topics[len-1]
		}

	}
	p.Topics = p.Topics[:len-1]

	fmt.Println(" available topics after removing from list is", p.Topics)
	return nil
}

func (p *Pub) Publish(message string) error {

	for _, topic := range p.Topics {
		for _, sub := range topic.Subscribers {
			sub.Notify(message)
		}

	}

	return nil
}
