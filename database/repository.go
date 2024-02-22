// database/repository.go

package database

import "gorm.io/gorm"

type CRUDRepository[T any] interface {
	Create(tx *gorm.DB, e *T) error
	Update(tx *gorm.DB, e *T) error
	Delete(tx *gorm.DB, e *T) error
	FindByID(tx *gorm.DB, id uint, e *T) error
}
