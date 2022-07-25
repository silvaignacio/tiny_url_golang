package repository

import (
	"gorm.io/gorm"
	"pet-system/domain/model"
	"pet-system/usecase/repository"
)

type userRepository struct {
	db *gorm.DB
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
