package main

import (
	"log"

	"go.uber.org/fx"
)

func main() {
	// t := Title("goodbye")
	// p := NewPublisher(&t)
	// m := NewMainService(p)
	// m.Run()
	fx.New(
		fx.Provide(NewMainService),
		fx.Provide(NewPublisher),
		fx.Provide(func() *Title {
			t := Title("goodbye")
			return &t
		}),
		fx.Invoke(func(service *MainService) {
			service.Run()
		}),
	).Run()
}

type MainService struct {
	publisher *Publisher
}

func NewMainService(publisher *Publisher) *MainService {
	return &MainService{
		publisher: publisher,
	}
}
func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("main program")
}

// Dependency
type Publisher struct {
	title *Title
}

func NewPublisher(title *Title) *Publisher {
	return &Publisher{
		title: title,
	}
}
func (publisher *Publisher) Publish() {
	log.Print("publisher:", *publisher.title)
}

// Dependency of publisher
type Title string
