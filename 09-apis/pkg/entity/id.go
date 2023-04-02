package entity

import "github.com/google/uuid"

type ID = uuid.UUID

var NilID = uuid.Nil

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return ID(id), err
}
