package storage

import (
	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)
	DeleteContact(ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error)
	CountContact( /*Тут можно передавать фильтр*/ ) (uint64, error)
}

type Group interface {
	CreateGroup(group *group.Group) (*group.Group, error)
	UpdateGroup(ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error)
	DeleteGroup(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	ListGroup(parameter queryParameter.QueryParameter) ([]*group.Group, error)
	ReadGroupByID(ID uuid.UUID) (*group.Group, error)
	CountGroup( /*Тут можно передавать фильтр*/ ) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error)
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
	AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error
}
