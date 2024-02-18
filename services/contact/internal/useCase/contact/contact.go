package contact

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	createdContacts, err := uc.adapterStorage.CreateContact(contacts...)
	if err != nil {
		log.Printf("Failed to create contacts: %v", err)
		return nil, err
	}
	return createdContacts, nil
}

func (uc *UseCase) Update(contactUpdate contact.Contact) (*contact.Contact, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	updateFunc := func(c *contact.Contact) (*contact.Contact, error) {
		*c = contactUpdate
		return c, nil
	}
	updatedContact, err := uc.adapterStorage.UpdateContact(contactUpdate.ID(), updateFunc)
	if err != nil {
		log.Printf("Failed to update contact with ID %s: %v", contactUpdate.ID(), err)
		return nil, err
	}
	return updatedContact, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return errors.New("adapterStorage is not initialized")
	}
	err := uc.adapterStorage.DeleteContact(ID)
	if err != nil {
		log.Printf("Failed to delete contact with ID %s: %v", ID, err)
		return err
	}
	return nil
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	contacts, err := uc.adapterStorage.ListContact(parameter)
	if err != nil {
		log.Printf("Failed to list contacts: %v", err)
		return nil, err
	}
	return contacts, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *contact.Contact, err error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	contact, err := uc.adapterStorage.ReadContactByID(ID)
	if err != nil {
		log.Printf("Failed to read contact with ID %s: %v", ID, err)
		return nil, err
	}
	return contact, nil
}

func (uc *UseCase) Count() (uint64, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return 0, errors.New("adapterStorage is not initialized")
	}
	count, err := uc.adapterStorage.CountContact()
	if err != nil {
		log.Printf("Failed to count contacts: %v", err)
		return 0, err
	}
	return count, nil
}
