package subscriber

import "fmt"

type Sub struct {
	Name string
}

type Subscriber interface {
	Notify(message string)
}

func (s *Sub) Notify(message string) {

	fmt.Printf(" message recieved on %s \n", s.Name)

}
