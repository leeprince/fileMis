package service

import "fileMis/src/model"

type UserService struct {

}

var (
	user *model.UserModel
)

func init() {
	user = model.NewUserModel()
}

func (this *UserService) GetList() []*model.UserModel  {
	return user.All()
}
