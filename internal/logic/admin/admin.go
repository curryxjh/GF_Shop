package admin

import (
	"GF_Shop/internal/dao"
	"GF_Shop/internal/model"
	"GF_Shop/internal/service"
	"GF_Shop/utility"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/util/grand"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(&in); err != nil {
		return out, err
	}
	// 处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt

	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertID)}, err
}

func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err := ghtml.SpecialCharsMapOrStruct(&in); err != nil {
			return err
		}
		if in.Password != "" {
			UserSalt := grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, UserSalt)
			in.UserSalt = UserSalt
		}
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			OmitEmpty().
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	m := dao.AdminInfo.Ctx(ctx)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	out.Total, err = m.Count()

	if err != nil || out.Total == 0 {
		return out, err
	}

	out.List = make([]model.AdminGetListOutputItem, 0, in.Size)

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
