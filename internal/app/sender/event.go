package sender

import (
	"github.com/ozonmp/bss-company-api/internal/model"
)

type EventSender interface {
	Send(company *model.CompanyEvent) error
}
