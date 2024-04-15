package page_list

import (
	"QAComm/global"
	"QAComm/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// PageList 是一个通用的分页查询函数，用于从数据库中查询指定模型的数据并进行分页处理。
// 参数说明：
// - model 是要查询的模型对象，可以是任意类型。
// - option 是用于配置分页查询选项的结构体。
// 返回值说明：
// - list 是查询结果的列表，类型为传入的模型类型。
// - count 是符合查询条件的记录总数。
// - err 是查询过程中可能发生的错误。
func PageList[T any](model T, option Option) (list []T, count int64, err error) {
	// 获取全局数据库实例
	DB := global.DB
	// 如果设置了 Debug 标志，开启调试模式，将日志输出到全局 MysqlLog 中
	if option.Debug {
		DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 如果未设置排序条件，则默认按照创建时间倒序排序
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	// 查询符合条件的记录总数，并将结果存储到 count 中
	count = DB.Select("id").Find(&list, model).RowsAffected
	// 计算偏移量，用于分页查询
	offset := option.Limit * (option.Page - 1)
	if offset <= 0 {
		offset = 0
	}
	// 如果未设置限制数量，则执行无限制查询；否则执行带限制数量的查询
	if option.Limit == 0 {
		// 执行排序后的查询，并将结果存储到 list 中
		err = DB.Order(option.Sort).Find(&list, model).Error
	} else {
		// 执行带有限制数量和偏移量的查询，并将结果存储到 list 中
		err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list, model).Error
	}
	// 返回查询结果列表、记录总数和可能发生的错误
	return list, count, err
}

func PageListByQuestId[T any](model T, option Option, id int) (list []T, count int64, err error) {
	// 获取全局数据库实例
	DB := global.DB
	// 如果设置了 Debug 标志，开启调试模式，将日志输出到全局 MysqlLog 中
	if option.Debug {
		DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 如果未设置排序条件，则默认按照创建时间倒序排序
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	// 查询符合条件的记录总数，并将结果存储到 count 中
	count = DB.Select("quest_id").Where("quest_id = ?", id).Find(&list, model).RowsAffected
	// 计算偏移量，用于分页查询
	offset := option.Limit * (option.Page - 1)
	if offset <= 0 {
		offset = 0
	}
	// 如果未设置限制数量，则执行无限制查询；否则执行带限制数量的查询
	if option.Limit == 0 {
		// 执行排序后的查询，并将结果存储到 list 中
		err = DB.Order(option.Sort).Where("quest_id = ?", id).Find(&list, model).Error
	} else {
		// 执行带有限制数量和偏移量的查询，并将结果存储到 list 中
		err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Where("quest_id = ?", id).Find(&list, model).Error
	}
	// 返回查询结果列表、记录总数和可能发生的错误
	return list, count, err
}

func PageListByUserId[T any](model T, option Option, id int) (list []T, count int64, err error) {
	// 获取全局数据库实例
	DB := global.DB
	// 如果设置了 Debug 标志，开启调试模式，将日志输出到全局 MysqlLog 中
	if option.Debug {
		DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 如果未设置排序条件，则默认按照创建时间倒序排序
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	// 查询符合条件的记录总数，并将结果存储到 count 中
	count = DB.Select("user_id").Where("user_id = ?", id).Find(&list, model).RowsAffected
	// 计算偏移量，用于分页查询
	offset := option.Limit * (option.Page - 1)
	if offset <= 0 {
		offset = 0
	}
	// 如果未设置限制数量，则执行无限制查询；否则执行带限制数量的查询
	if option.Limit == 0 {
		// 执行排序后的查询，并将结果存储到 list 中
		err = DB.Order(option.Sort).Where("user_id = ?", id).Find(&list, model).Error
	} else {
		// 执行带有限制数量和偏移量的查询，并将结果存储到 list 中
		err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Where("user_id = ?", id).Find(&list, model).Error
	}
	// 返回查询结果列表、记录总数和可能发生的错误
	return list, count, err
}
