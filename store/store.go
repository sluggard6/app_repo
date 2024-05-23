package store

import "io"

//Store 存储接口
type Store interface {
	SaveFile(reader io.Reader, name string) (*File, error)
	GetAbsPath(path string) string
}
