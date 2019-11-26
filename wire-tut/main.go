package main

import (
	"errors"
	//"github.com/google/wire"
	"log"
	"time"
)

type Message string

type greeter struct {
	Grumpy bool
	Message Message
}

type Event struct {
	Greeter greeter
}

// provider - returns Message
func NewMessage(phrase string) Message {
	return Message(phrase)
}

// provider - depends on Message
func NewGreeter(m Message) greeter {
	var grumpy bool
	if time.Now().Unix() % 5 == 0 {
		grumpy = true
	}
	return greeter{Message: m, Grumpy: grumpy}
}

type Greeter interface {
	Greet() string
}

type MyGreeter string

func (g *MyGreeter) provideMyGreeter() *MyGreeter {
	s := new(MyGreeter)
	*s = "Hello, World"
	return s
}

func (g greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// provider - depends on greeter
func NewEvent(g greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Println(msg)
}

func main() {
	log.Printf("Hello, Wire!")
	event, err := InitializeEvent("Hi")
	if err != nil {
		log.Panic(err)
	}

	//var Set = wire.NewSet(
	//	provideMyGreeter,
	//	wire.Bind(new(Greeter), new(*MyGreeter)),
	//)

	event.Start()
}

func NewEventNumber() int  {
	return 1
}

//func InitializeEvent() Event {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//	return event
//}
