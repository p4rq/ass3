package group

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"

	"architecture_go/services/contact/internal/domain/contact"
)

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	createdContacts, err := uc.adapterStorage.CreateContactIntoGroup(groupID, contacts...)
	if err != nil {
		log.Printf("Failed to create contacts in group: %v", err)
		return nil, err
	}
	return createdContacts, nil
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return errors.New("adapterStorage is not initialized")
	}
	err := uc.adapterStorage.AddContactsToGroup(groupID, contactID)
	if err != nil {
		log.Printf("Failed to add contact to group: %v", err)
		return err
	}
	return nil
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return errors.New("adapterStorage is not initialized")
	}
	err := uc.adapterStorage.DeleteContactFromGroup(groupID, contactID)
	if err != nil {
		log.Printf("Failed to delete contact from group: %v", err)
		return err
	}
	return nil
}
