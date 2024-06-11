package model

type App struct {
	Model
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type AppVersion struct {
	Model
	AppId   uint   `json:"app_id"`
	App     App    `gorm:"foreignKey:AppId" json:"app"`
	FileId  uint   `json:"file_id"`
	File    File   `gorm:"foreignKey:FileId" json:"file"`
	Version string `json:"version"`
}
