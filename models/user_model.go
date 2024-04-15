package models

type UserModel struct {
	Model
	UserName      string        `gorm:"size:36" json:"user_name"`      //用户名
	Password      string        `gorm:"size:258" json:"password"`      //密码
	Token         string        `gorm:"size:512" json:"token"`         //token
	RefreshToken  string        `gorm:"size:512" json:"refresh_token"` //refresh token
	QuestionModel QuestionModel `gorm:"foreignKey:UserID" json:"-"`
}
