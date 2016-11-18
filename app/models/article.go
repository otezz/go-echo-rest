package models

import (
	"github.com/otezz/echo-rest/db"
	"time"
)

type (
	Article struct {
		BaseModel
		User    User   `json:"author"`
		UserID  int    `json:"user_id" gorm:"type:integer(11)"`
		Title   string `json:"title" gorm:"type:varchar(100)"`
		Content string `json:"content" gorm:"type:text"`
	}
)

func CreateArticle(m *Article) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m *Article) UpdateArticle(data *Article) error {
	var err error

	m.UpdatedAt = time.Now()
	m.Title = data.Title
	m.Content = data.Content

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m Article) DeleteArticle() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func GetArticleById(id uint64) (Article, error) {
	var (
		article Article
		err     error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&article, id).Error; err != nil {
		tx.Rollback()
		return article, err
	}
	tx.Commit()

	return article, err
}

func GetArticles() ([]Article, error) {
	var (
		articles []Article
		err      error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&articles).Error; err != nil {
		tx.Rollback()
		return articles, err
	}
	tx.Commit()

	return articles, err
}
