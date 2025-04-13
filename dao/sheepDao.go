package dao

import (
	"ginDemo/config"
	"ginDemo/model"
	"ginDemo/model/vo"
	"ginDemo/utils"
	"github.com/google/uuid"
	"time"
)

type SheepDao struct{}

/*
(s *SheepDao)    表示这个函数是 SheepDao 类型的**“方法”**，而不是普通函数。
(value string)   函数的参数
([]model.SheepModel, error) 函数的返回值
根据条件值查询多条数据
*/

// 构造函数
func NewSheepDao(cxf *config.ApplicationContext) interface{} {
	return &SheepDao{}
}

func (s *SheepDao) GetSheepByValue(value string) ([]model.SheepInfos, error) {
	db := config.GetMysqlDb()
	var infos []model.SheepInfos
	result := db.Debug().Where("code = ?", value).Find(&infos)
	if result.Error != nil {
		utils.GetLogger().Error("查询失败: %v", result.Error)
		return []model.SheepInfos{}, result.Error
	}
	return infos, nil
}

func (s *SheepDao) InsertSheep(sheep *model.SheepInfos) error {
	db := config.GetMysqlDb()
	info := model.SheepInfos{
		UUID:       uuid.New().String(),
		Code:       sheep.Code,
		Category:   sheep.Category,
		Age:        sheep.Age,
		Birth:      time.Now(),
		Avatar:     []byte(uuid.New().String()),
		CreateTime: time.Now(),
		CreateUser: "admin",
		UpdateTime: time.Now(),
		UpdateUser: "",
	}
	return db.Create(info).Error
}

func (dao *SheepDao) GetSheepList(offset int, req vo.SheepQueryRequest) ([]model.SheepInfos, int64, error) {
	db := config.GetMysqlDb()

	tx := db.Model(&model.SheepInfos{})

	// 多条件过滤
	if req.Sheep.Code != "" {
		tx = tx.Where("code = ?", req.Sheep.Code)
	}
	if req.Sheep.Category != "" {
		tx = tx.Where("category = ?", req.Sheep.Category)
	}
	if req.Sheep.Age != 0 {
		tx = tx.Where("age = ?", req.Sheep.Age)
	}

	// 查询总数（一定要在分页之前 count）
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		utils.GetLogger().Error("统计总数失败: %v", err)
		return nil, 0, err
	}

	// 分页查询
	var list []model.SheepInfos
	result := tx.Offset(offset).Limit(req.PageParam.PageSize).Find(&list)
	if result.Error != nil {
		utils.GetLogger().Error("分页查询失败: %v", result.Error)
		return nil, 0, result.Error
	}

	return list, total, nil
}
