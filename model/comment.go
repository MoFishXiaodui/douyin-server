package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID           uint      `gorm:"primary_key"`
	Content      string    `gorm:"type:text"`
	Time         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Likes        int       `gorm:"default:0"`
	Reply        string    `gorm:"type:text"`
	IP           string    `gorm:"type:varchar(15);unique_index"`
	CommentCount int       `gorm:"default:0"`
}

type CommentDao interface {
	InsertComment(comment Comment) error
	GetComments() ([]Comment, error)
	DeleteComment(id uint) error
}

type CommentDaoImpl struct {
	db *gorm.DB
}

func NewCommentDaoImpl(db *gorm.DB) *CommentDaoImpl {
	return &CommentDaoImpl{db: db}
}

func (dao *CommentDaoImpl) InsertComment(comment Comment) error {
	return dao.db.Create(&comment).Error
}

func (dao *CommentDaoImpl) GetComments() ([]Comment, error) {
	var comments []Comment
	if err := dao.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (dao *CommentDaoImpl) DeleteComment(id uint) error {
	return dao.db.Delete(&Comment{}, id).Error
}
