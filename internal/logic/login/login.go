package login

import (
	"GF_Shop/internal/dao"
	"GF_Shop/internal/model"
	"GF_Shop/internal/model/entity"
	"GF_Shop/internal/service"
	"GF_Shop/utility"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) (err error) {
	adminInfo := entity.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或密码不正确")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线 for session
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}
