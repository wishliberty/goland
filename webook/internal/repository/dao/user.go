package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突") //预定义错误
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// 操作数据库
type UserDao struct {
	db *gorm.DB
}

// 具体的表字段
type User struct {
	Id       int64  `gorm:"primaryKey autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64 //创建时间 不用time.time的原因是因为不同服务器时区不同，统一用UTC 0的毫秒数
	Utime    int64 //更新时间
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}
func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli() //毫秒数
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062 //（邮箱）重复冲突状态码1062
		if me.Number == duplicateErr {
			//用户冲突,邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {

	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}
