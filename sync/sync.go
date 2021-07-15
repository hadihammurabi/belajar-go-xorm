package sync

import (
	"belajar-go-xorm/model"

	"xorm.io/xorm"
)

func Sync(engine xorm.EngineInterface) error {
	err := engine.Sync2(
		new(model.User),
	)
	return err
}
