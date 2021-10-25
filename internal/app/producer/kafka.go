package producer

import (
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-company-api/internal/app/repo"
	"github.com/ozonmp/bss-company-api/internal/app/sender"
	"github.com/ozonmp/bss-company-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender  sender.EventSender
	updater repo.EventRepo
	cleaner repo.EventRepo
	events  <-chan model.CompanyEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

// todo for students: add repo
func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	updater repo.EventRepo,
	cleaner repo.EventRepo,
	events <-chan model.CompanyEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		updater:    updater,
		cleaner:    cleaner,
		events:     events,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							// error in sending to Kafka, unlock event in DB
							log.Printf("Unable send to Kafka: %v", err)
							batch := []uint64{event.ID}
							if err := p.cleaner.Unlock(batch); err != nil {
								log.Printf("Unable to unlock event in db: %v", err)
							}
						})
					} else {
						p.workerPool.Submit(func() {
							// event was sended to Kafka, update DB
							batch := []uint64{event.ID}
							if err := p.updater.Update(batch); err != nil {
								log.Printf("Unable to update db: %v", err)
							}
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
