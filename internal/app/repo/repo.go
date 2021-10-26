package repo

import (
	"sync"

	"github.com/ozonmp/bss-company-api/internal/model"
)

type companyArchive struct {
	company []model.CompanyEvent
	mu      sync.Mutex
}

var AC companyArchive

func (ac *companyArchive) Add(company []model.CompanyEvent) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	ac.company = append(ac.company, company...)

	return nil
}

func (ac *companyArchive) Lock(n uint64) ([]model.CompanyEvent, error) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	var ret []model.CompanyEvent

	for i, comp := range ac.company {
		if comp.Type == model.Created && comp.Status == model.Deferred {
			comp.Status = model.Captured
			ac.company[i].Status = model.Captured

			ret = append(ret, comp)
		}

		if uint64(len(ret)) == n {
			break
		}
	}

	return ret, nil
}

func (ac *companyArchive) Unlock(eventIDs []uint64) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	for _, event_id := range eventIDs {
		ac.company[event_id].Status = model.Deferred
	}

	return nil
}

func (ac *companyArchive) Remove(eventIDs []uint64) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	for _, event_id := range eventIDs {
		ac.company = append(ac.company[0:event_id], ac.company[event_id+1:]...)
	}
}

func (ac *companyArchive) Update(eventIDs []uint64) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	for _, event_id := range eventIDs {
		ac.company[event_id].Status = model.Processed
	}

	return nil
}
