package producer

import (
	"sync"
	"testing"

	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-company-api/internal/mocks"
	"github.com/ozonmp/bss-company-api/internal/model"
)

func TestProducer(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)
	workerpool := workerpool.New(1)

	events := make(chan model.CompanyEvent, 10)
	defer close(events)

	event_1 := model.CompanyEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Company{
			ID:      1,
			Name:    "Company_1",
			Address: "Unknown_1",
		},
	}

	event_2 := model.CompanyEvent{
		ID:     2,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Company{
			ID:      2,
			Name:    "Company_2",
			Address: "Unknown_2",
		},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	s1 := sender.EXPECT().Send(&event_1).Return(nil).Times(1)
	s2 := sender.EXPECT().Send(&event_2).Return(nil).Times(1)

	gomock.InOrder(
		s1,
		s2,
	)

	repo.EXPECT().Update([]uint64{1}).Do(
		func(_ []uint64) {
			wg.Done()
		},
	).Times(1)
	repo.EXPECT().Update([]uint64{2}).Do(
		func(_ []uint64) {
			wg.Done()
		},
	).Times(1)

	producer := NewKafkaProducer(
		1,
		sender,
		repo,
		events,
		workerpool,
	)

	producer.Start()

	events <- event_1
	events <- event_2

	wg.Wait()
	producer.Close()
}
