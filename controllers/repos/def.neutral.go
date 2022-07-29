package repos

import (
	"errors"

	"github.com/dasha-kinsely/ostruct/models/entities"
	"gorm.io/gorm"
)

type NeutralRepo interface {
	InsertNeutralMonster(neutral entities.Monster) (entities.Monster, error)
	FindAllNeutralMonster() []entities.Monster
	FindOneNeutralMonster(name string) (entities.Monster)
	EditNeutralMonster(neutral entities.Monster, name string) (entities.Monster, error)
	DeleteNeutralMonster(name string) (error)
	// ------------------------------------
	InsertSubordinateMonster(subordinate entities.Subordinate) (entities.Subordinate, error)
	FindAllSubordinateMonster() []entities.Subordinate
	FindOneSubordinateMonster(name string) entities.Subordinate
	EditSubordinateMonster(subordinate entities.Subordinate, name string) (entities.Subordinate, error)
	DeleteSubordinateMonster(name string) error
}

func NewNeutralRepo(db *gorm.DB) NeutralRepo {
	return &MySQLClient{
		database: db,
	}
}

func (db *MySQLClient) InsertNeutralMonster(monster entities.Monster) (entities.Monster, error) {
	if err := db.database.Save(&monster).Error; err != nil {
		return monster, err
	} 
	db.database.Preload("Subordinate").Find(&monster)
	return monster, nil
}

func (db *MySQLClient) FindAllNeutralMonster() []entities.Monster {
	var oList []entities.Monster
	db.database.Find(&oList)
	return oList
}

func (db *MySQLClient) FindOneNeutralMonster(name string) (entities.Monster) {
	o := entities.Monster{}
	db.database.Where("monster_type=?", name).Find(&o)
	return o
}

func (db *MySQLClient) EditNeutralMonster(monster entities.Monster, name string) (entities.Monster, error) {
	if err := db.database.Model(&entities.Monster{}).Where("monster_type=?", name).Updates(&monster).Error; err != nil {
		return monster, err
	}
	return monster, nil
}

func (db *MySQLClient) DeleteNeutralMonster(name string) (error) {
	o := entities.Monster{}
	db.database.Where("monster_type=?", name).Find(&o)
	if (o.ID == 0 && o.MonsterType == "") {
		return errors.New("the record with the specified \"MonsterType\" does not exist")
	}
	if err := db.database.Model(entities.Monster{}).Where("monster_type=?", name).Update("monster_type", nil).Error; err != nil {
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

// ---------------------------------------------------------------------------------
func (db *MySQLClient) InsertSubordinateMonster(subordinate entities.Subordinate) (entities.Subordinate, error) {
	if err := db.database.Save(&subordinate).Error; err != nil {
		return subordinate, err
	}
	return subordinate, nil
}

func (db *MySQLClient) FindAllSubordinateMonster() []entities.Subordinate {
	var oList []entities.Subordinate
	db.database.Find(&oList)
	return oList
}

func (db *MySQLClient) FindOneSubordinateMonster(name string) (entities.Subordinate) {
	o := entities.Subordinate{}
	db.database.Where("monster_type=?", name).Find(&o)
	return o
}

func (db *MySQLClient) EditSubordinateMonster(subordinate entities.Subordinate, name string) (entities.Subordinate, error) {
	if err := db.database.Model(&entities.Subordinate{}).Where("monster_type=?", name).Updates(&subordinate).Error; err != nil {
		return subordinate, err
	}
	return subordinate, nil
}

func (db *MySQLClient) DeleteSubordinateMonster(name string) error {
	o := entities.Subordinate{}
	db.database.Where("monster_type=?", name).Find(&o)
	if (o.ID == 0 || o.MonsterType == "") {
		return errors.New("the record with the specified \"MonsterType\" does not exist")
	} 
	if err := db.database.Model(entities.Subordinate{}).Where("monster_type=?", name).Update("monster_type", nil).Error; err != nil {
		return err
	}
	// NOTE: this does not delete the record from SQL entirely but rather do a soft-delete by setting "deleted_at" column to current time and \n
	// forbids anyone other than db admins from accessing. It will be removed during the routine back_up phases using the caching and logging db.
	result := db.database.Where("ID=?", o.ID).Delete(&o)
	if result.Error == nil {
		return nil
	} else {
		return result.Error
	}
}