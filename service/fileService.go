package service

import (
	"io"
	"path/filepath"

	"github.com/shfd/repo/model"
	"github.com/shfd/repo/store"
)

type FileService interface {
	SaveFile(reader io.Reader, name string) (*model.File, error)
	GetAbsPath(file *model.File) string
	GetExportPath() string
}

type myFileSer struct {
	store store.Store
}

var myFileService *myFileSer

// NewFileService 根据存储类型创建文件服务
func NewFileService(store store.Store) FileService {
	if myFileService == nil {
		myFileService = &myFileSer{store}
	}
	return myFileService
	// return (*FileService)(unsafe.Pointer(myFileService))
}

func GetFileService() FileService {
	return myFileService
}

func (s *myFileSer) GetExportPath() string {
	return s.store.GetExportPath()
}

func (s *myFileSer) SaveFile(reader io.Reader, name string) (*model.File, error) {
	// if err := folderService.CheckFileName(folder.ID, name); err != nil {
	// 	return nil, err
	// }
	storeFile, err := s.store.SaveFile(reader, name)
	if err != nil {
		return nil, err
	}

	file := &model.File{Name: name, Ext: filepath.Ext(name), Path: storeFile.Path, Sha: storeFile.Sha, Size: uint64(storeFile.Size)}
	if err := file.CheckOrCreat(); err != nil {
		return nil, err
	}
	// _, err = model.Create(file)
	return file, err
}

func (s *myFileSer) GetAbsPath(file *model.File) string {
	return s.store.GetAbsPath(file.Path)
}
