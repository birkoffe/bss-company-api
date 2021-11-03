package sender

import (
	"github.com/ozonmp/bss-cbssany-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.SubdomainEvent) error
}
