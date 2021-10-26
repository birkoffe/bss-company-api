package sender

import (
	"sync"

	"github.com/ozonmp/bss-company-api/internal/model"
)

type companyArchive struct {
	company []model.CompanyEvent
	mu      sync.Mutex
}

var AC companyArchive

func (ac *companyArchive) Send(company *model.CompanyEvent) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	if ac.company == nil {
		ac.company = make([]model.CompanyEvent, 10)
	}

	ac.company = append(ac.company, *company)

	return nil
}
