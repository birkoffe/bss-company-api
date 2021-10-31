package consumer

import (
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-company-api/internal/mocks"
	"github.com/ozonmp/bss-company-api/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)

	events := make(chan model.CompanyEvent, 10)
	defer close(events)

	event := model.CompanyEvent{
		ID:     0,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Company{
			ID:      0,
			Name:    "Company_1",
			Address: "Unknown_1",
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	repo.EXPECT().Lock(uint64(1)).Return([]model.CompanyEvent{event}, nil).Do(
		func(_ uint64) {
			wg.Done()
		},
	).Times(1)

	consumer := NewDbConsumer(
		1,
		1,
		100*time.Millisecond,
		repo,
		events,
	)

	consumer.Start()
	wg.Wait()
	consumer.Close()

	select {
	case inEvents, ok := <-events:
		if !ok {
			t.Fatal("Error reading from channel")
		}

		assert.Equal(t, event, inEvents, "In and out events must be equal")
	default:
		t.Fatal("No events in channel")
	}
}
