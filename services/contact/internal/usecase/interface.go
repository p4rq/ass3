package useCase

import (
	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type Contact interface {
	Create(contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(contact contact.Contact) (*contact.Contact, error)
	Delete(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	ContactReader
}

type ContactReader interface {
	List(parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadByID(ID uuid.UUID) (response *contact.Contact, err error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}

type Group interface {
	Create(domainGroup *group.Group) (*group.Group, error)
	Update(group *group.Group) (*group.Group, error)
	Delete(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	List(parameter queryParameter.QueryParameter) ([]*group.Group, error)
	ReadByID(ID uuid.UUID) (*group.Group, error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
