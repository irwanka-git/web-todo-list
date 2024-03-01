package entity

import "time"

type Task struct {
	IDTask      int32     `gorm:"column:id_task;primaryKey;autoIncrement:true" json:"id_task"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreateBy    int32     `gorm:"column:create_by" json:"create_by"`
}
