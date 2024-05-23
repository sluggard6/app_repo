package model

//Policy 存储策略
type Policy struct {
	Model
	Type PolicyType
	Path string
	Sha  string
}

//PolicyType 策略类型
type PolicyType string

const (
	//MyFile 自带的文件存储方案
	MyFile PolicyType = "myfile"
)

//File 文件
type File struct {
	Model
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	FolderID uint   `json:"-"`
	// Folder   *Folder `gorm:"foreignKey:FolderID" json:"-"`
	Size     uint64  `json:"size"`
	PolicyID uint    `json:"-"`
	Policy   *Policy `gorm:"foreignKey:PolicyID" json:"-"`
}

//GetFilesByFolderID 查询目录下的所有文件
func (file *File) GetFilesByFolderID() (files *[]File, err error) {
	files = &[]File{}
	err = db.Where("folder_id=?", file.FolderID).Find(files).Error
	return
}

// CheckOrCreat 检查文件是否存在，如果不存在就创建文件
func (policy *Policy) CheckOrCreat() (err error) {
	if err = db.Where("sha=?", policy.Sha).Find(policy).Error; err != nil {
		return
	}
	if policy.ID == 0 {
		err = db.Create(policy).Error
	}
	return
}
