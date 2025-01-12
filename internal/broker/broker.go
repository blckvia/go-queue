package broker

import (
	"errors"
	"fmt"
	"sync"
)

type Broker struct {
	queues map[string]*Queue
	mu     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		queues: make(map[string]*Queue),
	}
}

func (b *Broker) CreateQueue(name string, size, maxSubs int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.queues[name]; exists {
		return errors.New("queue already exists")
	}

	b.queues[name] = NewQueue(name, size, maxSubs)
	return nil
}

func (b *Broker) SendMessage(queueName string, message *Message) error {
	b.mu.RLock()
	queue, exists := b.queues[queueName]
	b.mu.RUnlock()

	if !exists {
		return errors.New("queue not found")
	}
	fmt.Println("Sending message to queue:", queueName, "Message:", message)
	fmt.Println("Subscribers count:", len(queue.subscribers))

	return queue.Push(message)
}

func (b *Broker) AddSubscriber(queueName string, subscriber chan *Message) error {
	b.mu.RLock()
	queue, exists := b.queues[queueName]
	b.mu.RUnlock()

	if !exists {
		return errors.New("queue not found")
	}

	return queue.AddSubscriber(subscriber)
}
