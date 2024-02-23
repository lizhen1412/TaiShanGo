// database/gorm_repository.go

package database

import (
	"gorm.io/gorm"
)

// GormRepository 是一个实现了 CRUDRepository 接口的结构体，
// 使用了 GORM 库来操作数据库。
type GormRepository[T any] struct{}

// Create 方法实现了 CRUDRepository 接口中的 Create 方法，
// 用于向数据库中添加新的记录。
// 参数 tx 是数据库事务对象，e 是要添加的记录的指针。
func (repo *GormRepository[T]) Create(tx *gorm.DB, e *T) error {
	return tx.Create(e).Error
}

// Update 方法实现了 CRUDRepository 接口中的 Update 方法，
// 用于更新数据库中的记录。
// 参数 tx 是数据库事务对象，e 是要更新的记录的指针。
func (repo *GormRepository[T]) Update(tx *gorm.DB, e *T) error {
	return tx.Save(e).Error
}

// Delete 方法实现了 CRUDRepository 接口中的 Delete 方法，
// 用于从数据库中删除记录。
// 参数 tx 是数据库事务对象，e 是要删除的记录的指针。
func (repo *GormRepository[T]) Delete(tx *gorm.DB, e *T) error {
	return tx.Delete(e).Error
}

// FindByID 方法实现了 CRUDRepository 接口中的 FindByID 方法，
// 用于通过记录的唯一标识符从数据库中检索记录。
// 参数 tx 是数据库事务对象，id 是要检索的记录的唯一标识符，e 是要填充检索到的记录的指针。
func (repo *GormRepository[T]) FindByID(tx *gorm.DB, id uint, e *T) error {
	return tx.First(e, id).Error
}
