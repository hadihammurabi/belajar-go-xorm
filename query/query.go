package query

// DOCS:
// https://gobook.io/read/gitea.com/xorm/manual-en-US/chapter-05/4.find.html

import (
	"belajar-go-xorm/model"
	"errors"

	"xorm.io/xorm"
)

func Get(engine xorm.EngineInterface) (*model.User, error) {
	user := &model.User{Username: "test2@gmail.com"}
	has, err := engine.Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New("record not found")
	}

	return user, nil
}

func Find(engine xorm.EngineInterface) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := engine.Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindWhere(engine xorm.EngineInterface) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := engine.Where("id > ?", "3").Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
