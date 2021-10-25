package sender

import (
	"github.com/ozonmp/bss-company-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.SubdomainEvent) error
}
