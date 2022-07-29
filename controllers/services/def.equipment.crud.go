package services

import (
	"errors"
	"strconv"

	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/models/entities"
)

type EquipmentCRUDService interface {
	CreateEquipment(req dto.CreateEquipmentRequest) (entities.Equipment, error)
	GetAllEquipment() ([]entities.Equipment, error)
	GetAllEquipmentByCategory(category string) ([]entities.Equipment, error)
	GetOneEquipment(name string) (entities.Equipment, error)
	UpdateOneEquipment(req dto.UpdateOneEquipmentRequest, name string) (entities.Equipment, error)
	UpdateEquipLimit(req dto.BulkUpdateEquipLimitRequest) (error)
	DeleteOneEquipment(name string) (error)
}

type equipmentCRUDService struct {
	crudRepo repos.EquipmentRepo
}

func NewEquipmentService(repo repos.EquipmentRepo) EquipmentCRUDService {
	return &equipmentCRUDService{
		crudRepo: repo,
	}
}

func (repo *equipmentCRUDService) CreateEquipment(req dto.CreateEquipmentRequest) (entities.Equipment, error) {
	o := entities.Equipment{}
	o.EquipmentType = req.EquipmentType
	o.Category = req.Category
	inserted, err := repo.crudRepo.InsertEquipment(o)
	return inserted, err
}

func (repo *equipmentCRUDService) GetAllEquipment() ([]entities.Equipment, error) {
	oList := repo.crudRepo.FindAllEquipment()
	if oList == nil {
		return oList, errors.New("No record found, database table is empty...")
	}
	return oList, nil
}

func (repo *equipmentCRUDService) GetAllEquipmentByCategory(category string) ([]entities.Equipment, error) {
	oList := repo.crudRepo.FindAllEquipmentByCategory(category)
	if oList == nil {
		return oList, errors.New("No record found, database table is empty...")
	}
	return oList, nil
}

func (repo *equipmentCRUDService) GetOneEquipment(name string) (entities.Equipment, error) {
	o := repo.crudRepo.FindOneEquipmentByName(name)
	if o.EquipmentType == "" {
		return o, errors.New("No record with the specified equipment name was found")
	}
	return o, nil
}

func (repo *equipmentCRUDService) UpdateOneEquipment(req dto.UpdateOneEquipmentRequest, name string) (entities.Equipment, error) {
	o := entities.Equipment{}
	o.Category = req.Category
	o.EquipmentType = req.EquipmentType
	result, err := repo.crudRepo.EditOneEquipment(o, name)
	return result, err
}

func (repo *equipmentCRUDService) UpdateEquipLimit(req dto.BulkUpdateEquipLimitRequest) (error) {
	parsed, err := strconv.ParseInt(req.EquipLimit, 10, 64)
	if err != nil {
		// should return format error
		return err
	}
	err = repo.crudRepo.SmartUpdateEquipLimit(req.Category, parsed)
	return err
}

func (repo *equipmentCRUDService) DeleteOneEquipment(name string) (error) {
	err := repo.crudRepo.DeleteEquipment(name)
	return err
}
