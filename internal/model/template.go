package model

// company - company entity.
type company struct {
	ID      uint64 `db:"id"`
	Name    string `db:"Name"`
	Address string `db:"Address"`
}
