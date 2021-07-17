package delete

import (
	"belajar-go-xorm/model"

	"xorm.io/xorm"
)

func Delete(engine xorm.EngineInterface) (int64, error) {
	return engine.Where("id > 0").Unscoped().Delete(&model.User{})
}
