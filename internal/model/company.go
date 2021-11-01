package model

import "fmt"

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
	Captured
	Processed
)

type CompanyEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Company
}

func (c *Company) String() string {
	return fmt.Sprintf("{ ID: %v, Name: %v, Address: %v }", c.ID, c.Name, c.Address)
}
