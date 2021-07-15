package insert

import (
	"belajar-go-xorm/model"

	"xorm.io/xorm"
)

func Insert(engine xorm.EngineInterface) (int64, error) {
	user := new(model.User)
	user.Username = "test1@gmail.com"
	user.Password = "123123"
	return engine.Insert(user)
}

func InsertMany(engine xorm.EngineInterface) (int64, error) {
	users := make([]*model.User, 0)
	users = append(users, &model.User{
		Username: "test2@gmail.com",
		Password: "123123",
	})
	users = append(users, &model.User{
		Username: "test3@gmail.com",
		Password: "123123",
	})
	users = append(users, &model.User{
		Username: "test4@gmail.com",
		Password: "123123",
	})
	users = append(users, &model.User{
		Username: "test5@gmail.com",
		Password: "123123",
	})
	return engine.Insert(users)
}
