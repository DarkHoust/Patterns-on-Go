package main

import "fmt"

type Observer interface {
	Notify()
}

type Observable interface {
	AddSubscriber(o Observer)
	NotifySubscribers()
}

type Channel struct {
	Observers []Observer
}

func (ch *Channel) AddSubscriber(o Observer) {
	ch.Observers = append(ch.Observers, o)
}

func (ch *Channel) NotifySubscribers() {
	for _, i := range ch.Observers {
		i.Notify()
	}
}

type Subscriber struct {
	Name string
}

func (s *Subscriber) Notify() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Hello", s.Name, ", you have an offer:")
	fmt.Println("-----------------------------------------")
}

type Moderator struct {
	Username string
}

func (m *Moderator) Notify() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Hello", m.Username, ", you have to check this offer:")
	fmt.Println("-----------------------------------------")
}

func main() {
	ch1 := Channel{}
	Alex := Subscriber{Name: "Alex"}
	Adam := Subscriber{Name: "Adam"}
	Aron := Moderator{Username: "Anonymus"}
	ch1.AddSubscriber(&Alex)
	ch1.AddSubscriber(&Adam)
	ch1.AddSubscriber(&Aron)
	ch1.NotifySubscribers()
}
