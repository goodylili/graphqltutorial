package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type BioRepository interface {
	CreatePerson(bioData BioData) (BioData, error)
	UpdatePerson(bioData BioData, id int) error
	DeletePerson(id int) error
	GetAllPersons() ([]BioData, error)
}

func Database() *gorm.DB {
	// function body goes here
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&BioData{})
	if err != nil {
		log.Fatal("failed to migrate")
	}

	return db
}

func NewBookService(db *gorm.DB) *BioService {
	return &BioService{
		db: db,
	}
}

type BioService struct {
	db *gorm.DB
}

func (bio *BioService) CreatePerson(bioData *BioData) (*BioData, error) {
	bioInstance := &BioData{
		ID:   bioData.ID,
		Name: bioData.Name,
		Age:  bioData.Age,
	}

	err := bio.db.Create(&bioInstance).Error
	return bioInstance, err
}

func (bio *BioService) UpdatePerson(bioData *BioData, id int) error {
	bioInstance := BioData{
		ID:   bioData.ID,
		Name: bioData.Name,
		Age:  bioData.Age,
	}
	err := bio.db.Model(&bioData).Where("id = ?", id).Updates(bioInstance).Error
	return err
}

func (bio *BioService) DeletePerson(id int) error {
	book := BioData{}
	err := bio.db.Delete(book, id).Error
	return err
}

func (bio *BioService) GetAllPersons() ([]*BioData, error) {
	var bioInstance []*BioData
	err := bio.db.Find(&bioInstance).Error
	return bioInstance, err
}
