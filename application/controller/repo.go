package controller

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/shfd/repo/model"
	"github.com/shfd/repo/service"
	"github.com/shfd/repo/util"
	"github.com/sirupsen/logrus"
)

type RepoController struct {
	fileService service.FileService
}

func NewRepoController() *RepoController {
	return &RepoController{service.GetFileService()}
}

type AppListForm struct {
	AppList []AppList `json:"app_list"`
	// RootPath string    `json:"root_path"`
}

type AppList struct {
	Id      uint
	Version string
}

func (c *RepoController) PostExport(ctx iris.Context) *HttpResult {
	// appList := &AppListForm{}
	form := &AppListForm{}
	if err := ctx.ReadJSON(form); err != nil {
		logrus.Errorln(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	files := []model.File{}
	appVersions := []model.AppVersion{}
	if len(form.AppList) == 0 {
		model.DB().Model(&model.AppVersion{}).Find(&appVersions)
		// model.DB().versions
		for _, appVersion := range appVersions {
			form.AppList = append(form.AppList, AppList{Id: appVersion.AppId, Version: appVersion.Version})
		}
		model.DB().Where("id in (?)", model.DB().Table("app_versions").Select("file_id")).Find(&files)
	} else {
		model.DB().Where("(app_id, version) in ?", form.AppList).Find(&files)
	}
	// 创建应用信息文件
	file, err := os.Create(c.fileService.GetExportPath() + string(filepath.Separator) + "info.json")
	if err != nil {
		return FailedMessage(err.Error())
	}
	defer file.Close()
	data, err := json.Marshal(form.AppList)
	if err != nil {
		return Failed()
	}
	file.Write(data)
	for _, file := range files {
		util.CopyFile(c.fileService.GetAbsPath(&file), filepath.Join(c.fileService.GetExportPath(), string(filepath.Separator), file.Path))
	}

	// 以下代码创建压缩文件
	absFileName := c.fileService.GetExportPath() + string(filepath.Separator) + "exoprt.zip"
	os.RemoveAll(absFileName)
	exportFile, _ := os.Create(absFileName)
	defer exportFile.Close()
	archive := zip.NewWriter(exportFile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(c.fileService.GetExportPath(), func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == c.fileService.GetExportPath() || path == absFileName {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, c.fileService.GetExportPath()+string(filepath.Separator))

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	ctx.SendFile(absFileName, "export.zip")
	return nil
}
