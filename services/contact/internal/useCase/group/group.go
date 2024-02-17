package group

import (
	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/group"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	panic("implement me")
}
func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	panic("implement me")
}
