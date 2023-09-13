package logic

import (
	"context"

	"back-end/api/internal/svc"
	"back-end/api/internal/types"
	"back-end/api/model"
	"back-end/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	l.Logger.Info("Loign req: ", req)

	var userModel *model.UserTbl
	var user types.User
	var token string

	userModel, err = l.svcCtx.UserModel.FindByName(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel == nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.USER_NOT_EXIT_CODE,
				Message: common.USER_NOT_EXIT_MESS,
			},
		}, nil
	}

	err = ComparePassword(userModel.Password, req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.WRONG_PASSWORD_ERROR_CODE,
				Message: common.WRONG_PASSWORD_ERROR_MESS,
			},
		}, nil
	}

	token, err = GenToken()
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.UNKNOW_ERR_CODE,
				Message: common.UNKNOW_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		ID:       userModel.Id,
		Name:     userModel.UserName,
		Fullname: userModel.FullName,
		Email:    userModel.Email,
		Role:     userModel.Role,
	}

	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		User:  user,
		Token: token,
	}, nil
}
