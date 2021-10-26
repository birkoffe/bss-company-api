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
	// log.Printf("Recieved %v", company)
	ac.mu.Lock()
	defer ac.mu.Unlock()

	if ac.company == nil {
		ac.company = make([]model.CompanyEvent, 10)
	}

	// r := rand.Intn(10)
	// log.Printf("r: %+v", r)

	// if r != 1 {
	// 	ac.company = append(ac.company, *company)
	// } else {
	// 	return errors.New("Oppps! Tray again!")
	// }

	ac.company = append(ac.company, *company)

	return nil
}
