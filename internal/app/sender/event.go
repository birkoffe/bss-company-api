package sender

import (
	"github.com/ozonmp/bss-company-api/internal/model"
)

type EventSender interface {
	Send(Company *model.CompanyEvent) error
}
