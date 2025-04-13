package service

import (
	"ginDemo/model"
	"ginDemo/model/vo"
)

type SheepService interface {

	//接口方法 便于实现
	GetSheepByCode(value string) ([]model.SheepInfos, error)

	InsertSheepService(sheep model.SheepInfos) error

	GetSheepByList(req vo.SheepQueryRequest) ([]model.SheepInfos, error)
}
