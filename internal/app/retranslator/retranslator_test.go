package retranslator

import (
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

	companies := []model.CompanyEvent{
		{
			ID:     1,
			Type:   model.Created,
			Status: model.Processed,
			Entity: &model.Company{
				ID:      1,
				Name:    "Company 1",
				Address: "unknown",
			},
		},
	}

	gomock.InOrder(
		repo.EXPECT().Lock(uint64(5)).Return(companies, nil).AnyTimes(),
		sender.EXPECT().Send(&companies[0]).Return(nil).AnyTimes(),
		repo.EXPECT().Remove([]uint64{companies[0].ID}).Return(nil).AnyTimes(),
	)

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	retranslator.Close()
}
