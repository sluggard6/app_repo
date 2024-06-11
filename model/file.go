package model

//File 文件
type File struct {
	Model
	Name string `json:"name"`
	Ext  string `json:"ext"`
	// AppId uint   `json:"-"`
	// Folder   *Folder `gorm:"foreignKey:FolderID" json:"-"`
	Size uint64 `json:"size"`
	Path string `json:"path"`
	Sha  string `json:"sha"`
	// PolicyID uint    `json:"-"`
	// Policy   *Policy `gorm:"foreignKey:PolicyID" json:"-"`
}

//GetFilesByFolderID 查询目录下的所有文件
// func (file *File) GetFilesByFolderID() (files *[]File, err error) {
// 	files = &[]File{}
// 	err = db.Where("app_id=?", file.AppId).Find(files).Error
// 	return
// }

// CheckOrCreat 检查文件是否存在，如果不存在就创建文件
func (file *File) CheckOrCreat() (err error) {
	if err = db.Where("sha=?", file.Sha).Find(file).Error; err != nil {
		return
	}
	if file.ID == 0 {
		err = db.Create(file).Error
	}
	return
}
