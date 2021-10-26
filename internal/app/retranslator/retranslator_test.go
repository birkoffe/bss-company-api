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
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	count := 1
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
		repo.EXPECT().Lock(uint64(10)).Return(nil, nil).MinTimes(1),
		// sender.EXPECT().Send(&companies[0]).Return(nil).AnyTimes(),
		// repo.EXPECT().Update([]uint64{companies[0].ID}).Return(nil).AnyTimes(),
	)

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	retranslator.Close()
}
