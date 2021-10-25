package model

type Company struct {
	ID      uint64
	Name    string
	Address string
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type CompanyEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Company
}
