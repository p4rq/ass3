package group

import (
	"time"

	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/group"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.CreateGroup(groupCreate)
}

func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.UpdateGroup(groupUpdate.ID(), func(oldGroup *group.Group) (*group.Group, error) {
		return group.NewWithID(oldGroup.ID(), oldGroup.CreatedAt(), time.Now().UTC(), groupUpdate.Name(), groupUpdate.Description(), oldGroup.ContactCount()), nil
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteGroup(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	return uc.adapterStorage.ListGroup(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	return uc.adapterStorage.ReadGroupByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountGroup()
}
