package postgres

import (
	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/group"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) UpdateGroup(ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListGroup(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (r *Repository) ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) CountGroup() (uint64, error) {
	panic("implement me")
}
