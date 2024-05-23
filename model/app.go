package model

type App struct {
	Model
	Name string `json:name`
	Desc string `json:desc`
}

type VersionFile struct {
	Model
	App_ID  uint   `json:app_id`
	App     App    `gorm:"foreignKey:App_ID" json:"app"`
	File_ID uint   `json:file_id`
	File    File   `gorm:"foreignKey:File_ID" json:"file"`
	Version string `json:version`
}
