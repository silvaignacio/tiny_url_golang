package repository

import (
	"gorm.io/gorm"
	"pet-system/domain/model"
	"pet-system/usecase/repository"
)

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) FindByUrl(u string) (*model.TinyUrl, error) {
	finded := &model.TinyUrl{}
	ur.db.Where(&model.TinyUrl{ShortUrl: u}).First(&finded)
	return finded, nil
}
func (ur *userRepository) Create(u *model.TinyUrl) (*model.TinyUrl, error) {
	//TODO implement me
	err := ur.db.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func NewPetRepository(db *gorm.DB) repository.PetRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.TinyUrl) ([]*model.TinyUrl, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
