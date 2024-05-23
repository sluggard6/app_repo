package service

import "github.com/shfd/repo/model"

//GetByID 统一的根据id查询对象方法
func GetByID(modelEntity interface{}, id uint) (interface{}, error) {
	return model.GetById(modelEntity, id)
}
