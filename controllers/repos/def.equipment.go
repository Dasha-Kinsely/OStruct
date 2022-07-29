package repos

import (
	"errors"

	"github.com/dasha-kinsely/ostruct/models/entities"
	"gorm.io/gorm"
)

type EquipmentRepo interface {
	InsertEquipment(equipment entities.Equipment) (entities.Equipment, error)
	FindAllEquipment() ([]entities.Equipment)
	FindAllEquipmentByCategory(category string) ([]entities.Equipment)
	FindOneEquipmentByName(name string) (entities.Equipment)
	EditOneEquipment(equipment entities.Equipment, name string) (entities.Equipment, error)
	SmartUpdateEquipLimit(category string, limit int64) (error)
	DeleteEquipment(name string) (error)
}

func NewEquipmentRepo(db *gorm.DB) EquipmentRepo {
	return &MySQLClient{
		database: db,
	}
}

func (db *MySQLClient) InsertEquipment(equipment entities.Equipment) (entities.Equipment, error) {
	if err := db.database.Save(&equipment).Error; err != nil {
		return equipment, err
	} 
	return equipment, nil
}

func (db *MySQLClient) FindAllEquipment() []entities.Equipment {
	var oList []entities.Equipment
	db.database.Find(&oList)
	return oList
}

func (db *MySQLClient) FindAllEquipmentByCategory(category string) []entities.Equipment {
	var oList []entities.Equipment
	db.database.Where("category=?", category).Find(&oList)
	return oList
}

func (db *MySQLClient) FindOneEquipmentByName(name string) entities.Equipment {
	o := entities.Equipment{}
	db.database.Where("equipment_type=?", name).Find(&o)
	return o
}

func (db *MySQLClient) EditOneEquipment(equipment entities.Equipment, name string) (entities.Equipment, error) {
	if err := db.database.Model(&entities.Equipment{}).Where("equipment_type=?", name).Updates(&equipment).Error; err != nil {
		return equipment, err
	}
	return equipment, nil
}

func (db *MySQLClient) DeleteEquipment(name string) (error) {
	o := entities.Equipment{}
	db.database.Where("equipment_type=?", name).Find(&o)
	if (o.ID == 0 && o.EquipmentType == "") {
		return errors.New("the record with the specified \"EquipmentType\" does not exist")
	}
	if err := db.database.Model(entities.Monster{}).Where("equipment_type=?", name).Update("equipment_type", nil).Error; err != nil {
		return err
	}
	// NOTE: this does not delete the record from SQL entirely but rather set "deleted_at" column to current time and \n
	// forbids anyone other than db admins from accessing. It will be removed during the routine back_up phases using the caching and logging db.
	result := db.database.Where("ID = ?", o.ID).Delete(&o)
	if result.Error == nil {
		return nil
	} else {
		return result.Error
	}
}

func (db *MySQLClient) SmartUpdateEquipLimit(category string, limit int64) error {
	if err := db.database.Model(&entities.Equipment{}).Where("category=?", category).Update("equip_limit", limit).Error; err != nil {
		return err
	}
	return nil
}