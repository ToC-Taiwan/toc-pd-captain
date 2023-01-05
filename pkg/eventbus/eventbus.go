// Package eventbus package eventbus
package eventbus

import (
	"github.com/asaskevich/EventBus"
)

var singleton *Bus

// Bus Bus
type Bus struct {
	bus EventBus.Bus
}

// New New
func New() *Bus {
	if singleton != nil {
		return singleton
	}

	new := &Bus{
		bus: EventBus.New(),
	}

	singleton = new
	return new
}

// PublishTopicEvent PublishTopicEvent
func (c *Bus) PublishTopicEvent(topic string, arg ...interface{}) {
	go c.bus.Publish(topic, arg...)
}

// SubscribeTopic SubscribeTopic
func (c *Bus) SubscribeTopic(topic string, fn ...interface{}) {
	for _, f := range fn {
		err := c.bus.SubscribeAsync(topic, f, false)
		if err != nil {
			panic(err)
		}
	}
}
