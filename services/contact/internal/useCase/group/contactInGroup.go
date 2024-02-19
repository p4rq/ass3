package group

import (
	"github.com/google/uuid"

	"architecture_go/services/contact/internal/domain/contact"
)

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContactIntoGroup(groupID, contacts...)
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.AddContactsToGroup(groupID, contactID)
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.DeleteContactFromGroup(groupID, contactID)
}
