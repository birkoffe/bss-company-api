package retranslator

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-company-api/internal/app/repo"
	"github.com/ozonmp/bss-company-api/internal/mocks"
	"github.com/ozonmp/bss-company-api/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.AssignableToTypeOf(uint64(0))).AnyTimes()
	sender.EXPECT().Send(gomock.AssignableToTypeOf(&model.CompanyEvent{})).AnyTimes()

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	retranslator.Close()
}

func TestPipeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  1,
		ConsumeSize:    1,
		ConsumeTimeout: 100 * time.Millisecond,
		ProducerCount:  3,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	count := 10
	companies := make([]model.CompanyEvent, 0, count)

	var wg = sync.WaitGroup{} // !!!

	for i := 0; i < count; i++ {
		companies = append(companies, model.CompanyEvent{
			ID:     uint64(i),
			Type:   model.Created,
			Status: model.Processed,
			Entity: &model.Company{
				ID:      uint64(i),
				Name:    fmt.Sprintf("Company_%d", i),
				Address: fmt.Sprintf("Unknown_%d", i),
			},
		})
	}

	repo.EXPECT().Lock(uint64(1)).Return(companies, nil).Times(1)
	for _, event := range companies {
		_ = event
		wg.Add(1)
		repo.EXPECT().Update([]uint64{event.ID}).Return(nil).Times(1)
		sender.EXPECT().Send(gomock.AssignableToTypeOf(&model.CompanyEvent{})).Do(func(ce *model.CompanyEvent) {
			wg.Done()
		}).Times(1)
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	defer retranslator.Close()

	wg.Wait() // here is waiting that all wg done before closing retranslator

	// time.Sleep(time.Second)
}

func TestSendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  1,
		ConsumeSize:    1,
		ConsumeTimeout: 100 * time.Millisecond,
		ProducerCount:  3,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	count := 10
	companies := make([]model.CompanyEvent, 0, count)

	var wg = sync.WaitGroup{} // !!!

	for i := 0; i < count; i++ {
		companies = append(companies, model.CompanyEvent{
			ID:     uint64(i),
			Type:   model.Created,
			Status: model.Processed,
			Entity: &model.Company{
				ID:      uint64(i),
				Name:    fmt.Sprintf("Company_%d", i),
				Address: fmt.Sprintf("Unknown_%d", i),
			},
		})
	}

	repo.EXPECT().Lock(uint64(1)).Return(companies, nil).Times(1)
	for idx, event := range companies {
		_ = event
		if idx == 0 {
			wg.Add(1)
			repo.EXPECT().Update(gomock.AssignableToTypeOf([]uint64{})).Return(nil).Times(0)
			repo.EXPECT().Unlock(gomock.AssignableToTypeOf([]uint64{})).Return(nil).Times(1)
			sender.EXPECT().Send(gomock.AssignableToTypeOf(&model.CompanyEvent{})).DoAndReturn(
				func(_ *model.CompanyEvent) error {
					defer wg.Done()
					return errors.New("Ooops")
				},
			)
		} else {
			wg.Add(1)
			repo.EXPECT().Update(gomock.AssignableToTypeOf([]uint64{})).Return(nil).Times(1)
			sender.EXPECT().Send(gomock.AssignableToTypeOf(&model.CompanyEvent{})).Do(func(_ *model.CompanyEvent) {
				wg.Done()
			}).Times(1)
		}
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	defer retranslator.Close()

	wg.Wait() // here is waiting that all wg done before closing retranslator
}

func TestOnlyCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	// repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    1,
		ConsumeTimeout: 100 * time.Millisecond,
		ProducerCount:  3,
		WorkerCount:    2,
		Repo:           &repo.AC,
		Sender:         sender,
	}

	count := 3
	companies := make([]model.CompanyEvent, 0, count)

	var wg = sync.WaitGroup{} // !!!

	companies = append(companies,
		model.CompanyEvent{
			ID:     0,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.Company{},
		},
		model.CompanyEvent{
			ID:     1,
			Type:   model.Removed,
			Status: model.Deferred,
			Entity: &model.Company{},
		},
		model.CompanyEvent{
			ID:     2,
			Type:   model.Removed,
			Status: model.Deferred,
			Entity: &model.Company{},
		},
	)

	repo.AC.Add(companies)

	wg.Add(1)
	c := companies[0]
	c.Status = model.Captured
	sender.EXPECT().Send(&c).Do(
		func(_ *model.CompanyEvent) {
			wg.Done()
		},
	).Times(1)

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	defer retranslator.Close()

	wg.Wait() // here is waiting that all wg done before closing retranslator

	time.Sleep(time.Second)

	assert.Equal(t, model.Processed, repo.AC.Get(0).Status, "Should be preccesed")
	assert.Equal(t, model.Deferred, repo.AC.Get(1).Status, "Should be not preccesed")
	assert.Equal(t, model.Deferred, repo.AC.Get(2).Status, "Should be not preccesed")
}
