// database/gorm_repository.go

package database

import "gorm.io/gorm"

type GormRepository[T any] struct{}

func (repo *GormRepository[T]) Create(tx *gorm.DB, e *T) error {
	return tx.Create(e).Error
}

func (repo *GormRepository[T]) Update(tx *gorm.DB, e *T) error {
	return tx.Save(e).Error
}

func (repo *GormRepository[T]) Delete(tx *gorm.DB, e *T) error {
	return tx.Delete(e).Error
}

func (repo *GormRepository[T]) FindByID(tx *gorm.DB, id uint, e *T) error {
	return tx.First(e, id).Error
}
