package utils

// 自定义类型实现时间格式化 https://blog.csdn.net/LW1314QS/article/details/125605988
// 自定义类型必须实现的接口 https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LocalTime time.Time

// 时间格式化：json格式化时，会自动调用 MarshalJSON()
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	t2 := time.Time(*t)
	s := t2.Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("\"%v\"", s)), nil
}

// sql: converting argument $1 type: unsupported type utils.LocalTime, a struct
// （自定义类型）实现 Scanner 接口：从数据库读取数据，将数据转换为自定义类型
func (n *LocalTime) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		*n = LocalTime(t)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}

// （自定义类型）实现 Valuer 接口：保存数据到数据库，将自定义类型转换为数据库认识的类型
func (n LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	t := time.Time(n)
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t, nil
}

// 自定义格式化 Gorm Model
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt *LocalTime     `json:"created_at"`
	UpdatedAt *LocalTime     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // gorm.DeletedAt 才有软删除功能
}
