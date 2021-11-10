package model

import "time"

type Company struct {
	ID      uint64
	Name    string
	Address string
	Removed bool
	Created time.Time
	Updated time.Time
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
