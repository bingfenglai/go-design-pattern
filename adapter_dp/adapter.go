package adapter_dp

type UserInfo struct {
	Uid    int64
	Name   string
	Age    int
	Gender string
}

type UserSimpleInfo struct {
	Uid  int64
	Name string
}

type IUserService interface {
	GetUserInfo(uid int64) *UserInfo
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUserInfo(uid int64) *UserInfo {
	return &UserInfo{Uid: uid, Name: "韩立", Age: 9999, Gender: "男"}
}

type UserSimpleInfoService interface {
	GetUserSimpleInfo(uid int64) *UserSimpleInfo
}

type UserServiceAdapter struct {
	UserService IUserService
}

func (receiver *UserServiceAdapter) GetUserSimpleInfo(uid int64) *UserSimpleInfo {

	userInfo := receiver.UserService.GetUserInfo(uid)

	return &UserSimpleInfo{
		Uid:  userInfo.Uid,
		Name: userInfo.Name,
	}
}
