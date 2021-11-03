package sender

import (
	"github.com/ozonmp/bss-cbssany-api/internal/model"
)

type EventSender interface {
	Send(Company *model.CompanyEvent) error
}
