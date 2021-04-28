package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

// 配置文件
type AppConfig struct {
	// 环境
	Env string `yaml:"env"`
	// 运行模式
	Mode string `yaml:"mode"`
	// 应用名称
	Name string `yaml:"name"`
}
type ServerConfig struct {
	// 服务端口
	Port int `yaml:"port"`
}
type LogConfig struct {
	// 日志文件配置
	File LogFileConfig `yaml:"file"`
	// 日志级别(debug, info, warn, error)
	Level string `yaml:"level"`
}
type LogFileConfig struct {
	// 日志路径
	Path string `yaml:"path"`
	// 日志名称
	Name string `yaml:"name"`
}

type RedisConfig struct {
	// 连接地址
	Addr string `yaml:"addr"`
	// 密码
	Passwd string `yaml:"passwd"`
	// 库索引
	Database int `yaml:"database"`
	// 最大等待时间
	MaxIdle int `yaml:"max-idle"`
	// 最大连接数
	MaxActive int `yaml:"max-active"`
}

type DataBaseConfig struct {
	// 数据库方言
	Dialect string `yaml:"dialect"`
	// 数据源
	DataSources DataSourceConfig `yaml:"datasources"`
}

type DataSourceConfig struct {
	MySQL   DataSourceMySQLConfig   `yaml:"mysql"`
	SQLite3 DataSourceSQLite3Config `yaml:"sqlite3"`
}

type DataSourceMySQLConfig struct {
	// 连接地址
	Addr string `yaml:"addr"`
	// 用户名
	Username string `yaml:"username"`
	// 密码
	Password string `yaml:"password"`
	// 数据库名称
	Database string `yaml:"database"`
}

type DataSourceSQLite3Config struct {
	// 数据库文件
	DBFile string `yaml:"db-file"`
}

type JWTConfig struct {
	// 加密盐
	Secret string `yaml:"secret"`
	// 过期时间
	Expiry int `yaml:"expiry"`
}

type ApplicationProperties struct {
	// 应用配置
	App AppConfig `yaml:"app"`
	// Gin服务配置
	Server ServerConfig `yaml:"server"`
	// 日志配置
	Log LogConfig `yaml:"log"`
	// redis配置
	Redis RedisConfig `yaml:"redis"`
	// 数据库配置
	DataBase DataBaseConfig `yaml:"database"`
	// jwt配置
	Jwt JWTConfig `yaml:"jwt"`
}

// 根据路径读取yaml文件
func readYaml(path string) ApplicationProperties {
	var result ApplicationProperties
	data, err := ioutil.ReadFile(path)
	if err != nil {
		Logger.Error("File reading error, application.yml not exist!" + err.Error())
		panic(err)
	}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		Logger.Errorf("cannot unmarshal data: %v", err)
		panic(err)
	}
	return result
}

// 根据环境配置取配置明细
func (v *ApplicationProperties) Init() {
	result := readYaml(fmt.Sprintf("%s/application.yml", GetCurrentPath()))
	// 判断环境
	if strings.Compare(result.App.Env, "dev") == 0 {
		AppProps = readYaml(fmt.Sprintf("%s/application-%s.yml", GetCurrentPath(), "dev"))
	}
}

// 获取当前路径，比如：d:/abc
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		Logger.Error(err)
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
