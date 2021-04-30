package util

import (
	"lingye-gin/src/config"
)

func SelectOne(target interface{}, sql string, values ...interface{}) {
	config.SqlExcutor.Raw(sql, values).Scan(target)
}

func SelectList(target interface{}, sql string, values ...interface{}) {
	config.SqlExcutor.Raw(sql, values).Scan(target)
}

func Create() {

}
