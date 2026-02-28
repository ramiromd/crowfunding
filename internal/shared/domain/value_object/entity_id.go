package value_object

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidEntityIdFormat = errors.New("invalid entity id format")

type EntityId struct {
	value uuid.UUID
}

func NewEntityId(entityId string) (EntityId, error) {
	entityIdAsUuid, err := uuid.Parse(entityId)
	if err != nil {
		return EntityId{}, ErrInvalidEntityIdFormat
	}
	return EntityId{value: entityIdAsUuid}, nil	
}

func (e EntityId) Value() string {
	return e.value.String()
}