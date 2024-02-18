package group

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/group"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	createdGroup, err := uc.adapterStorage.CreateGroup(groupCreate)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		return nil, err
	}
	return createdGroup, nil
}
func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	updateFunc := func(g *group.Group) (*group.Group, error) {
		*g = *groupUpdate
		return g, nil
	}
	updatedGroup, err := uc.adapterStorage.UpdateGroup(groupUpdate.ID(), updateFunc)
	if err != nil {
		log.Printf("Failed to update group with ID %s: %v", groupUpdate.ID(), err)
		return nil, err
	}
	return updatedGroup, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return errors.New("adapterStorage is not initialized")
	}
	err := uc.adapterStorage.DeleteGroup(ID)
	if err != nil {
		log.Printf("Failed to delete group with ID %s: %v", ID, err)
		return err
	}
	return nil
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	groups, err := uc.adapterStorage.ListGroup(parameter)
	if err != nil {
		log.Printf("Failed to list groups: %v", err)
		return nil, err
	}
	return groups, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return nil, errors.New("adapterStorage is not initialized")
	}
	group, err := uc.adapterStorage.ReadGroupByID(ID)
	if err != nil {
		log.Printf("Failed to read group with ID %s: %v", ID, err)
		return nil, err
	}
	return group, nil
}

func (uc *UseCase) Count() (uint64, error) {
	if uc.adapterStorage == nil {
		log.Println("adapterStorage is not initialized")
		return 0, errors.New("adapterStorage is not initialized")
	}
	count, err := uc.adapterStorage.CountGroup()
	if err != nil {
		log.Printf("Failed to count groups: %v", err)
		return 0, err
	}
	return count, nil
}
