package retranslator

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-company-api/internal/mocks"
	"github.com/ozonmp/bss-company-api/internal/model"
)

func TestStart(t *testing.T) {

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()
	sender.EXPECT().Send(gomock.Any()).AnyTimes()

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
		ConsumerCount:  2,
		ConsumeSize:    1,
		ConsumeTimeout: 100 * time.Microsecond,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	count := 10
	companies := make([]model.CompanyEvent, count)

	for i := 0; i < count; i++ {
		companies[i] = model.CompanyEvent{
			ID:     uint64(i),
			Type:   model.Created,
			Status: model.Processed,
			Entity: &model.Company{
				ID:      uint64(1),
				Name:    fmt.Sprintf("Company_%d", i),
				Address: fmt.Sprintf("Unknown_%d", i),
			},
		}
	}

	gomock.InOrder(
		repo.EXPECT().Lock(gomock.Any()).Return(companies, nil).Times(1),
		// sender.EXPECT().Send(gomock.Any()).Return(nil).Times(2),
		// repo.EXPECT().Update(gomock.Any()).Return(nil).Times(0),
	)

	retranslator := NewRetranslator(cfg)
	retranslator.Start()

	// !!!
	time.Sleep(5 * time.Second)

	retranslator.Close()
}
