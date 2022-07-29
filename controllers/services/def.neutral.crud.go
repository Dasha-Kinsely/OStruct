package services

import (
	"errors"
	"strconv"

	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/models/entities"
)

type NeutralCRUDService interface {
	CreateNeutralMonster(req dto.CreateNeutralMonsterRequest) (entities.Monster, error)
	GetAllNeutralMonster() ([]entities.Monster, error)
	GetOneNeutralMonster(name string) (entities.Monster, error)
	UpdateOneNeutralMonster(req dto.UpdateOneNeutralMonster, name string) (entities.Monster, error)
	DeleteOneNeutralMonster(name string) (error)
	// ------------------------------------------------
	CreateSubordinateMonster(req dto.CreateSubordinateMonsterRequest) (entities.Subordinate, error)
	GetAllSubordinateMonster() ([]entities.Subordinate, error)
	GetOneSubordinateMonster(name string) (entities.Subordinate, error)
	UpdateOneSubordinateMonster(req dto.UpdateOneSubordinateMonster, name string) (entities.Subordinate, error)
	DeleteOneSubordinateMonster(name string) (error)
}

type neutralCRUDService struct {
	crudRepo repos.NeutralRepo
}

func NewNeutralCRUDService(repo repos.NeutralRepo) NeutralCRUDService{
	return &neutralCRUDService{
		crudRepo: repo,
	}
}

func (repo *neutralCRUDService) CreateNeutralMonster(req dto.CreateNeutralMonsterRequest) (entities.Monster, error) {
	o := entities.Monster{}
	o.MonsterType = req.MonsterType
	parsed, err := strconv.ParseInt(req.Tankiness, 10, 64)
	if err != nil {
		// should return format error
		return o, err
	}
	o.Tankiness = parsed
	o.AttackPattern = req.AttackPattern
	o.PatrolPattern = req.PatrolPattern
	o.Respawn = req.Respawn
	o.Reward = req.Reward
	inserted, err := repo.crudRepo.InsertNeutralMonster(o)
	return inserted, err
}

func (repo *neutralCRUDService) GetAllNeutralMonster() ([]entities.Monster, error) {
	oList := repo.crudRepo.FindAllNeutralMonster()
	if oList == nil {
		return oList, errors.New("No record found, database table is empty...")
	}
	return oList, nil
}

func (repo *neutralCRUDService) GetOneNeutralMonster(name string) (entities.Monster, error) {
	o := repo.crudRepo.FindOneNeutralMonster(name)
	if o.MonsterType == "" {
		return o, errors.New("No record with the specified monster name was found")
	}
	return o, nil
}

func (repo *neutralCRUDService) UpdateOneNeutralMonster(req dto.UpdateOneNeutralMonster, name string) (entities.Monster, error) {
	o := entities.Monster{}
	o.MonsterType = req.MonsterType
	parsed, err := strconv.ParseInt(req.Tankiness, 10, 64)
	if err != nil {
		// should return format error
		return o, err
	}
	o.Tankiness = parsed
	o.AttackPattern = req.AttackPattern
	o.PatrolPattern = req.PatrolPattern
	o.Respawn = req.Respawn
	o.Reward = req.Reward
	result, err := repo.crudRepo.EditNeutralMonster(o, name)
	return result, err
}

func (repo *neutralCRUDService) DeleteOneNeutralMonster(name string) (error) {
	err := repo.crudRepo.DeleteNeutralMonster(name)
	return err
}

//--------------------------------------------------------------------------------------------------------------------------------

func (repo *neutralCRUDService) CreateSubordinateMonster(req dto.CreateSubordinateMonsterRequest) (entities.Subordinate, error) {
	o := entities.Subordinate{}
	o.MonsterType = req.MonsterType
	parsed, err := strconv.ParseInt(req.Tankiness, 10, 64)
	if err != nil {
		// should return format error
		return o, err
	}
	o.Tankiness = parsed
	o.AttackPattern = req.AttackPattern
	o.Reward = req.Reward
	inserted, err := repo.crudRepo.InsertSubordinateMonster(o)
	return inserted, err
}

func (repo *neutralCRUDService) GetAllSubordinateMonster() ([]entities.Subordinate, error) {
	oList := repo.crudRepo.FindAllSubordinateMonster()
	if oList == nil {
		return oList, errors.New("No record found, database table is empty...")
	}
	return oList, nil
}

func (repo *neutralCRUDService) GetOneSubordinateMonster(name string) (entities.Subordinate, error) {
	o := repo.crudRepo.FindOneSubordinateMonster(name)
	if o.MonsterType == "" {
		return o, errors.New("No record with the specified monster name was found")
	}
	return o, nil
}

func (repo *neutralCRUDService) UpdateOneSubordinateMonster(req dto.UpdateOneSubordinateMonster, name string) (entities.Subordinate, error) {
	o := entities.Subordinate{}
	o.MonsterType = req.MonsterType
	parsed, err := strconv.ParseInt(req.Tankiness, 10, 64)
	if err != nil {
		// should return format error
		return o, err
	}
	o.Tankiness = parsed
	o.AttackPattern = req.AttackPattern
	o.Reward = req.Reward
	result, err := repo.crudRepo.EditSubordinateMonster(o, name)
	return result, err
}

func (repo *neutralCRUDService) DeleteOneSubordinateMonster(name string) (error) {
	err := repo.crudRepo.DeleteSubordinateMonster(name)
	return err
}