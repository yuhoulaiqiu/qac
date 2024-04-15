package models

type CommentModel struct {
	Model
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`         //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"` //父评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                                      //夫评论id
	Content            string          `gorm:"size:256" json:"content"`                                //评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`                     //点赞数
	CommentCount       int             `gorm:"size8;default:0" json:"comment_count"`                   //子评论数
	Answer             AnswerModel     `gorm:"foreignKey:AnswerID" json:"-"`                           //关联的回答
	AnswerID           uint            `json:"answer_id"`                                              //文章id
	User               UserModel       `json:"user"`                                                   //关联的用户
	UserID             uint            `json:"user_id"`                                                //评论的用户
}
