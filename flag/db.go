package flag

import (
	"QAComm/global"
	"QAComm/models"
)

func Makemigrations() {
	var err error
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.CommentModel{},
			&models.QuestionModel{},
			&models.AnswerModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功")
}
