package url_frontier

import (
	"errors"
)

// Interface to define different types of queues out UrlFrontier can implement
type Queue interface {
	AddUrl(string) error
	PopUrl() (string, error)
	Size() (int, error)
}

// Main struct to represent our UrlFrontier (queue) to be used in implementation
type UrlFrontier struct {
	queue Queue
}

// InMemoryQueue struct to implemenet the Queue interface. We choose a doubly
// linked list for queue implemtnation.
type InMemoryQueue struct {
	MaxSize int
	size    int
	first   *inMemoryQueueNode
	last    *inMemoryQueueNode
}

type inMemoryQueueNode struct {
	Url 	string
	Next    *inMemoryQueueNode
}

// InMemoryQueue METHODS

func (q *InMemoryQueue) AddUrl(url string) error {
	if q.MaxSize > 0 && q.size >= q.MaxSize  {
		return errors.New("queue is full")
	}
	// add to end of doubly linked list
	i := &inMemoryQueueNode{Url: url}
	if q.size == 0 {
		q.first = i
	} else {
		q.last.Next = i
	}
	q.last = i
	q.size++
	return nil
}

func (q *InMemoryQueue) PopUrl() (string, error) {
	if q.size == 0 {
		return "", errors.New("no item to pop, UrlFrontier is empty")
	}
	// pop item from the start
	i := q.first.Url
	q.first = q.first.Next
	q.size--
	return i, nil
}

func (q InMemoryQueue) Size() (int, error) {
	return q.size, nil
}

// UrlFrontier METHODS

func New() (*UrlFrontier, error) {
	q := &InMemoryQueue{MaxSize: 100000}
	return &UrlFrontier{
		queue: q,
	}, nil
}

func (url_frontier *UrlFrontier) AddUrl(url string) error {
	url_frontier.queue.AddUrl(url)
	return nil
}

func (url_frontier *UrlFrontier) PopUrl() error {
	url_frontier.queue.PopUrl()
	return nil
}

func (url_frontier *UrlFrontier) Size() (int, error) {
	return url_frontier.queue.Size()
}
