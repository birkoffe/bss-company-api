package repo

import (
	"github.com/ozonmp/bss-company-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.CompanyEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.CompanyEvent) error
	Remove(eventIDs []uint64) error
}
