package model

// company - company entity.
type company struct {
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}
