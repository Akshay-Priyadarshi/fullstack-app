package models

import "github.com/google/uuid"

type IDConstraint interface {
	~string | ~int | uuid.UUID
}

type Identity[ID IDConstraint] struct {
	Id ID `json:"id"`
}
