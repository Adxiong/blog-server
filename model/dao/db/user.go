package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserDao interface {
	AddUser()
	UpdateUserByUID()
	DeleteUserByUID()
	FindUserByUID()
	FindUserList()
}

func NewUser() *User {
	return &User{}
}

func NewUserList() *UserList {
	return &UserList{}
}

func (user *User) AddUser(ctx context.Context) (*User, error) {
	err := GlobalDb.Table(user.TableName()).Create(user).Error
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return user, nil
}

func (user *User) UpdateUserByUID(ctx context.Context, uid uint64, values map[string]interface{}) (int64, error) {
	var rowsAffected int64
	res := GlobalDb.Table(user.TableName()).Where("uid = ?", uid).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return rowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func (user *User) DeleteUserByUID(ctx context.Context, uid uint64) (int64, error) {
	var rowsAffected int64

	res := GlobalDb.Table(user.TableName()).Where("uid = ?", uid).Update(UserColumns.IsDel, 1)
	if res.Error != nil {
		log.Println("err", res.Error)
		return rowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func (user *User) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	params := map[string]interface{}{
		UserColumns.Email: email,
	}
	res := GlobalDb.Table(user.TableName()).Where(params).Limit(1).Find(user)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func (user *User) FindUserList(ctx context.Context, pn int, num int, params map[string]interface{}) (*UserList, error) {
	userList := NewUserList()
	offset := pn * num

	res := GlobalDb.Table(user.TableName()).Offset(offset).Limit(num).Where(params).Find(userList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return userList, nil
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	user.CreateAt = &t
	user.UpdateAt = &t

	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[UserColumns.UpdateAt]; !ok {
			t := time.Now()
			values[UserColumns.UpdateAt] = &t
		}

		if _, ok := values[UserColumns.Version]; !ok {
			values[UserColumns.UpdateAt] = gorm.Expr(fmt.Sprintf("%s + ?", UserColumns.Version), 1)
		}
	}

	return nil
}
