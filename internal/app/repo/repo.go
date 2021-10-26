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
			comp.Status = model.Processed
			ac.company[i].Status = model.Processed

			ret = append(ret, comp)
		}

		if uint64(len(ret)) == n {
			break
		}
	}

	return ret, nil
}

func (ac *companyArchive) Unlock(eventIDs []uint64) error {
	return nil
}

func (ac *companyArchive) Remove(eventIDs []uint64) error {
	return nil
}

func (ac *companyArchive) Update(eventIDs []uint64) error {
	return nil
}
