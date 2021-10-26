package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/bss-company-api/internal/app/repo"
	"github.com/ozonmp/bss-company-api/internal/app/retranslator"
	"github.com/ozonmp/bss-company-api/internal/app/sender"
	"github.com/ozonmp/bss-company-api/internal/model"
)

func main() {

	sigs := make(chan os.Signal, 1)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
		ConsumeTimeout: 100 * time.Millisecond,
		Repo:           &repo.AC,
		Sender:         &sender.AC,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	count := 2
	companies := make([]model.CompanyEvent, count)

	for i := 0; i < count; i++ {
		companies[i] = model.CompanyEvent{
			ID:     uint64(i),
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.Company{
				ID:      uint64(i),
				Name:    fmt.Sprintf("Company_%d", i),
				Address: fmt.Sprintf("Unknown_%d", i),
			},
		}
	}

	repo.AC.Add(companies)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
