// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"errors"
	"fmt"
	"time"
)

// Injectors from wire.go:

// 声明injector注入器
func InitializeEvent() (Event, error) {
	message := NewMessage()
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// wire.go:

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

// 宴会
type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

//创建消息
func NewMessage() Message {
	return Message("Hi there!")
}

// 创建招待人
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// 绑定欢迎方法
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// 创建一场宴会
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy暴躁的")
	}
	return Event{Greeter: g}, nil
}

// 宴会开始
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
