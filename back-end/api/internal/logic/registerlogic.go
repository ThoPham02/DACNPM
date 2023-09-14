package logic

import (
	"context"
	"time"

	"back-end/api/internal/svc"
	"back-end/api/internal/types"
	"back-end/api/model"
	"back-end/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	l.Logger.Info("Register req: ", req)

	var userModel *model.UserTbl
	var newUserModel *model.UserTbl
	var user types.User

	var hashPassword string
	var token string
	var secret = l.svcCtx.Config.Auth.AccessSecret
	var msTimeNow = time.Now().Unix()
	var expire = l.svcCtx.Config.Auth.AccessExpire

	userModel, err = l.svcCtx.UserModel.FindByName(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel != nil {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.USER_EXITED_ERR_CODE,
				Message: common.USER_EXITED_ERR_MESS,
			},
		}, nil
	}

	hashPassword, err = HashPassword(req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOW_ERR_CODE,
				Message: common.UNKNOW_ERR_MESS,
			},
		}, nil
	}

	newUserModel = &model.UserTbl{
		FullName: req.Fullname,
		UserName: req.Username,
		Password: hashPassword,
		Email:    req.Email,
		Role:     2,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, newUserModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		ID:       newUserModel.Id,
		Name:     newUserModel.UserName,
		FullName: newUserModel.FullName,
		Email:    newUserModel.Email,
		Role:     newUserModel.Role,
	}

	token, err = GenToken(secret, msTimeNow, expire, user.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOW_ERR_CODE,
				Message: common.UNKNOW_ERR_MESS,
			},
		}, nil
	}

	return &types.RegisterRes{
		Token: token,
		User:  user,
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
