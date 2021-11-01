package consumer

import (
	"sync"
	"time"

	"github.com/ozonmp/bss-company-api/internal/app/repo"
	"github.com/ozonmp/bss-company-api/internal/model"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.CompanyEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan bool
	wg   *sync.WaitGroup
}

type Config struct {
	n         uint64
	events    chan<- model.CompanyEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
}

func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.CompanyEvent) Consumer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
		done:      done,
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					// log.Printf("Try batch")
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						continue
					}
					// log.Printf("Events (%d): %v", len(events), events)
					for _, event := range events {
						c.events <- event
					}
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
}
