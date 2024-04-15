package models

type QuestionModel struct {
	Model
	Title         string `gorm:"size:128" json:"title"`           // 问题标题
	Content       string `json:"content"`                         // 问题内容
	AnswerCount   int    `gorm:"default:0" json:"answer_count"`   // 回答量
	LookCount     int    `gorm:"default:0" json:"look_count"`     // 浏览量
	DiggCount     int    `gorm:"default:0" json:"digg_count"`     // 点赞量
	CollectsCount int    `gorm:"default:0" json:"collects_count"` // 收藏量
	UserID        uint   `json:"user_id"`                         // 提问用户id
	Category      string `gorm:"size:20" json:"category"`         // 问题分类
}
