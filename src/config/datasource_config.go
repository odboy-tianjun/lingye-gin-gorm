package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"lingye-gin/src/modules/system/entity"
	"strings"
)

var SqlExcutor *gorm.DB

type DataSourcePool struct{}

func (v DataSourcePool) Connect() DataSourcePool {
	if strings.Compare(AppProps.DataBase.Dialect, "") == 0 {
		panic("yaml file dialect property is nil")
	}

	var err error
	if strings.Compare(AppProps.DataBase.Dialect, "mysql") == 0 {
		connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			AppProps.DataBase.DataSources.MySQL.Username,
			AppProps.DataBase.DataSources.MySQL.Password,
			AppProps.DataBase.DataSources.MySQL.Addr,
			AppProps.DataBase.DataSources.MySQL.Database)
		SqlExcutor, err = gorm.Open(AppProps.DataBase.Dialect, connectStr)
		// 创建表时添加表后缀
		if SqlExcutor != nil {
			SqlExcutor.Set("gorm:table_options", "ENGINE=InnoDB")
		}
	} else if strings.Compare(AppProps.DataBase.Dialect, "sqlite3") == 0 {
		// windows上会报'cgo: exec gcc: exec'异常, 注掉吧
		// SqlExcutor, err = gorm.Open(AppProps.DataBase.Dialect, AppProps.DataBase.DataSources.SQLite3.DBFile)
	}
	if err != nil {
		Logger.Error("db connect error: " + err.Error())
		panic(err)
	}
	if SqlExcutor != nil {
		// 全局禁用表名复数
		// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
		SqlExcutor.SingularTable(true)
		// 设置为true之后控制台会输出对应的SQL语句
		SqlExcutor.LogMode(true)
	}
	return v
}

// 统一在这里注册数据表
func (DataSourcePool) LoadEntity() {
	// 自动迁移模式
	if SqlExcutor != nil {
		SqlExcutor.AutoMigrate(&entity.User{})
	}
}
