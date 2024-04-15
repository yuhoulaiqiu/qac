package models

type AnswerModel struct {
	Model
	QuestID uint   `json:"quest_id"`                      // 问题id
	Content string `gorm:"size:128" json:"content"`       // 回答内容
	UserID  uint   `json:"user_id"`                       // 回答用户id
	IsAdopt bool   `gorm:"default:false" json:"is_adopt"` // 是否采纳
}
