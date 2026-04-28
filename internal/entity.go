package entity

import "time"

type Entity struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Entity) NewEntity(id int) *Entity {
	return &Entity{
		ID: id,
	}
}
