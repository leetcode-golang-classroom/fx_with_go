# fx_with_go

FX is a dependency injection system for Go

https://github.com/uber-go/fx

透過 fx framework 可以幫注 我們已一種特定的方式實現相依性注入

來做到服務隔離

## 舉例如下

```go
package main

import (
	"log"

	// "go.uber.org/fx"
)

func main() {
	t := Title("goodbye")
	p := NewPublisher(&t)
	m := NewMainService(p)
	m.Run()
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

```

假設有兩個 Services 彼此的相依關係為

MainService 相依於 PublishService

PublishService 又相依於 Title 屬性

所以啟用 MainService 時，就需要逐步從 Title 屬性建立起每個 Service 並且注入


## 使用 fx

```go
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

```

## 透過 fx.Anotation 可以解隅 Service 到 Interface

透過 fx.Anotation 可以指定 Provide 的 Instance 為哪個行別

```go
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
		fx.Provide(
			fx.Annotate(
				NewPublisher,
				fx.As(new(IPublisher)),
			),
		),
		fx.Provide(func() *Title {
			t := Title("hello")
			return &t
		}),
		fx.Invoke(func(service *MainService) {
			service.Run()
		}),
	).Run()
}

type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{
		publisher: publisher,
	}
}
func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("main program")
}

// Dependency
type IPublisher interface {
	Publish()
}
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

```