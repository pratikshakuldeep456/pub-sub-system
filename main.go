package main

import (
	"pratikshakuldeep456/pub-sub-system/pkg/publisher"
	"pratikshakuldeep456/pub-sub-system/pkg/subscriber"
	"pratikshakuldeep456/pub-sub-system/pkg/topic"
)

func main() {

	s1 := &subscriber.Sub{Name: "s11"}

	t1 := &topic.Topic{}

	t1.RegisterSubscriber(s1)

	p1 := &publisher.Pub{}

	p1.RegisterTopic(t1)

	p1.Publish("publish message")
}
