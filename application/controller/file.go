package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/shfd/repo/model"
	"github.com/shfd/repo/service"
	"github.com/shfd/repo/store"
	"github.com/sirupsen/logrus"
)

// FileController 文件控制器
type FileController struct {
	fileService service.FileService
}

// NewFileController 创建
func NewFileController(store store.Store) *FileController {
	return &FileController{service.NewFileService(store)}
}

// BeforeActivation 控制器配置
func (c *FileController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:uint}/{name:string}", "LoadFile")
}

// LoadFile 加载文件
func (c *FileController) LoadFile(ctx iris.Context, id uint, name string) *HttpResult {
	file := &model.File{}
	_, err := service.GetByID(file, id)
	if err != nil {
		return nil
	}
	policy := &model.Policy{}
	_, err = service.GetByID(policy, file.PolicyID)
	if err != nil {
		return nil
	}
	file.Policy = policy
	// folder := &model.Folder{}
	// _, err = service.GetByID(folder, file.FolderID)
	// if err != nil {
	// 	return nil
	// }
	// sess := sessions.Get(ctx)
	// user := sess.Get("user")
	// if hasRole, _, _ := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole {
	// 	return FailedForbidden(ctx)
	// }
	_, err = service.GetByID(file.Policy, file.PolicyID)
	// src, err := os.Open(file.Policy.Path)
	if err != nil {
		return nil
	}
	// defer src.Close()
	// io.Copy(ctx.ResponseWriter(), src)
	abspath := c.fileService.GetAbsPath(file)
	logrus.Debug(abspath, err)
	ctx.SendFile(abspath, file.Name)
	return nil
}

// PostUpload 上传文件
func (c *FileController) PostUpload(ctx iris.Context) *HttpResult {
	file, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return FailedMessage(err.Error())
	}
	// folder := &model.Folder{}
	// folderID, err := strconv.Atoi(ctx.FormValue("folderId"))
	if err != nil {
		return FailedCode(PARAM_ERROR)
	}
	logrus.Debugln(file)
	logrus.Debugln(fileHeader)
	// logrus.Debugln(folderID)
	// model.GetById(folder, uint(folderID))
	// sess := sessions.Get(ctx)
	// user := sess.Get("user")
	// if hasRole, role, _ := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole || role == model.Read {
	// 	return FailedForbidden(ctx)
	// }
	dbFile, err := c.fileService.SaveFile(file, fileHeader.Filename)
	if err != nil {
		return FailedMessage(err.Error())
	}
	return Success(dbFile)
}

// PostUploadFolder 上传文件夹
func (c *FileController) PostUploadFolder(ctx iris.Context) *HttpResult {
	if err := ctx.Request().ParseMultipartForm(ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()); err != nil {
		return FailedMessage(err.Error())
	}
	return nil
}
