package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server
	Database Database
	Stroe    Store
}

type Server struct {
	Host        string
	Port        int
	ContextPath string `yaml:"contextPath"`
}

type Store struct {
	DataRoot string
}

type Database struct {
	Type     dbType
	Url      string
	Username string
	Password string
}

type dbType string
type fileType string

type ConfigException struct {
	Message string
}

func (e *ConfigException) Error() string {
	return e.Message
}

const (
	DefaultConfigPath string   = "conf/application.yml"
	Mysql             dbType   = "mysql"
	Sqlite            dbType   = "sqlite"
	Yml               fileType = "yml"
	Json              fileType = "json"
)

var config Config = Config{
	Server: Server{
		Host:        "0.0.0.0",
		Port:        5678,
		ContextPath: "apprepo",
	},
	Database: Database{
		Type:     Sqlite,
		Url:      "reop.db",
		Username: "",
		Password: "",
	},
	Stroe: Store{
		DataRoot: "file-data",
	},
}

func GetConfig() *Config {
	return &config
}

func New(config string) Config {
	// return Config{"127.0.0.1", 5678}
	return LoadConfig(DefaultConfigPath)
}

func LoadConfig(path string) Config {
	ext := filepath.Ext(path)
	if ext == ".json" {
		return loadJsonConfig(path)
	} else if ext == ".yml" || ext == ".yaml" {
		return loadYmlConfig(path)
	} else {
		log.Error("unknow file " + path)
		return config
	}
}

func LoadConfig2(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	} else {
		ext := filepath.Ext(path)
		switch ext {
		case ".json":
			err = json.Unmarshal(data, &config)
		case ".yml":
			err = yaml.Unmarshal(data, &config)
		case ".yaml":
			err = yaml.Unmarshal(data, &config)
		default:
			panic(&ConfigException{"unknow file exception"})
		}
		if err != nil {
			panic(err)
		}
	}
	return config
}

func loadJsonConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	} else {
		//读取的数据为json格式，需要进行解码
		err = json.Unmarshal(data, &config)
		if err != nil {
			panic(err)
		}
	}
	return config
}

func loadYmlConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	log.Debug(config)
	return config
}
