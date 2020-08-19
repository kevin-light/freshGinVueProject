package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"` // char 字段定长
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"` // varchar 字段变长
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}

// 用 hook 在文章创建之前给 uuid赋值
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
