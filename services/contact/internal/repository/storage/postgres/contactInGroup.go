package postgres

import (
	"github.com/google/uuid"

	"architecture_go/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContactIntoGroup(groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	panic("implement me")
}
