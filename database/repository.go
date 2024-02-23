// database/repository.go

package database

import "gorm.io/gorm"

// CRUDRepository 是一个通用的接口，定义了对数据库进行增删改查操作的方法。
// 该接口使用了泛型，可以适用于任何类型的数据结构。
type CRUDRepository[T any] interface {

	// Create 方法用于向数据库中添加新的记录。
	// 参数 tx 是数据库事务对象，e 是要添加的记录的指针。
	Create(tx *gorm.DB, e *T) error

	// Update 方法用于更新数据库中的记录。
	// 参数 tx 是数据库事务对象，e 是要更新的记录的指针。
	Update(tx *gorm.DB, e *T) error

	// Delete 方法用于从数据库中删除记录。
	// 参数 tx 是数据库事务对象，e 是要删除的记录的指针。
	Delete(tx *gorm.DB, e *T) error

	// FindByID 方法用于通过记录的唯一标识符从数据库中检索记录。
	// 参数 tx 是数据库事务对象，id 是要检索的记录的唯一标识符，e 是要填充检索到的记录的指针。
	FindByID(tx *gorm.DB, id uint, e *T) error
}
