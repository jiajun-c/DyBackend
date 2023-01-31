package db

import "context"

type User struct {
	uid      string `gorm:"index:user_index;primaryKey"`
	Name     string
	Password string
}

func (u User) TableName() string {
	return "User"
}

// CreateUser
//
//	@Description: 创建用户
//	@param ctx	传入的上下文
//	@param users 用户列表
//	@return error	错误信息
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// CheckUser
//
//	@Description: 检查用户是否存在
//	@param ctx	上下文
//	@param username	用户名
//	@param password	密码
//	@return bool true存在，false不存在
func CheckUser(ctx context.Context, username string, password string) bool {
	return DB.WithContext(ctx).Where("name = ? and password = ?", username, password).Error == nil
}
