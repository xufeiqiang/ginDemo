package impl

import (
	"ginDemo/config"
	"ginDemo/dao"
	"ginDemo/model"
	"ginDemo/model/vo"
)

type SheepServiceImpl struct {
	// 定义成员变量
	sheepDao *dao.SheepDao
}

/*
定义构造函数 便于实例化,sheepDao实例注入
*/
func NewSheepServiceImpl(cxf *config.ApplicationContext) interface{} {
	sheepDao := cxf.GetBean("sheepDao").(*dao.SheepDao)
	return &SheepServiceImpl{
		sheepDao: sheepDao,
	}
}

/*
实现 SheepService 接口的方法  GetSheepByValue方法名一样即可实现
*/
func (sheepService *SheepServiceImpl) GetSheepByCode(value string) ([]model.SheepInfos, error) {
	// 调用 dao 层方法，获取数据
	sheepList, err := sheepService.sheepDao.GetSheepByValue(value)
	if err != nil {
		return nil, err
	}
	return sheepList, nil
}

func (sheepService *SheepServiceImpl) InsertSheepService(sheep *model.SheepInfos) error {

	return sheepService.sheepDao.InsertSheep(sheep)
}

func (sheepService *SheepServiceImpl) GetSheepByList(req vo.SheepQueryRequest) ([]model.SheepInfos, int64, error) {
	// 默认分页处理
	if req.PageParam.Page < 1 {
		req.PageParam.Page = 1
	}
	if req.PageParam.PageSize < 1 {
		req.PageParam.PageSize = 25
	}

	offset := (req.PageParam.Page - 1) * req.PageParam.PageSize

	return sheepService.sheepDao.GetSheepList(offset, req)

}
