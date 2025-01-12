package broker

import (
	"fmt"
	"sync"
)

type Message struct {
	ID      int    `json: "id" omitempty`
	Message string `json: "message"`
}

type Queue struct {
	name        string
	mu          sync.RWMutex
	arr         []*Message
	t           int
	h           int
	size        int
	maxSubs     int
	subscribers []chan *Message
}

func NewQueue(name string, size, maxSubs int) *Queue {
	return &Queue{
		name:        name,
		arr:         make([]*Message, size+1),
		t:           0,
		h:           0,
		size:        size,
		maxSubs:     maxSubs,
		subscribers: make([]chan *Message, 0),
	}
}

func (q *Queue) isEmpty() bool {
	if q.h == q.t {
		return true
	}

	return false
}

func (q *Queue) isFull() bool {
	if q.h == 0 && q.t == q.size || q.t == q.h-1 {
		return true
	}

	return false
}

func (q *Queue) Push(val *Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isFull() {
		return fmt.Errorf("Stack overflow")
	}
	q.arr[q.t] = val
	if q.t == q.size {
		q.t = 0
	} else {
		q.t++
	}
	fmt.Println("Added new val:", val, "Now len is", len(q.arr))

	if len(q.subscribers) > 0 {
		fmt.Println("Subscribing:", q.subscribers)

		for _, sub := range q.subscribers {
			fmt.Println("subscribe")
			sub <- val
			fmt.Println("Sended to sub:", sub, "Now len is:", len(q.arr), "Message: ", q.arr[q.t])
			q.Pop()
		}
		fmt.Println("Send to all subs:", len(q.subscribers), "Now len is:", len(q.arr))
	}

	return nil
}

func (q *Queue) Pop() (*Message, error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}

	x := q.arr[q.h]
	if q.h == q.size {
		q.h = 0
	} else {
		q.h++
	}

	return x, nil
}

func (q *Queue) Top() (*Message, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if q.isEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}

	return q.arr[q.h], nil
}

func (q *Queue) AddSubscriber(sub chan *Message) error {
	if len(q.subscribers) >= q.maxSubs {
		return fmt.Errorf("maximum number of subscribers reached")
	}

	fmt.Println("messages in queue:", len(q.arr))
	q.subscribers = append(q.subscribers, sub)
	fmt.Println("Subscriber has been added: sub: ", sub, "Now len is: ", len(q.arr))

	if len(q.subscribers) > 0 {
		fmt.Println("Subscribing:", q.subscribers)

		for _, s := range q.subscribers {
			message, err := q.Pop()
			if err != nil {
				return err
			}
			fmt.Println("subscribe")
			s <- message
			fmt.Println("Sended to sub:", s, "Now len is:", len(q.arr), "Message: ", message)
		}
		fmt.Println("Send to all subs:", len(q.subscribers), "Now len is:", len(q.arr))
	}

	q.h = q.t

	return nil
}
