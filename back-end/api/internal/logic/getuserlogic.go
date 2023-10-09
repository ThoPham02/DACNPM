package logic

import (
	"context"

	"back-end/api/internal/svc"
	"back-end/api/internal/types"
	"back-end/api/model"
	"back-end/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRes, err error) {
	l.Logger.Info("GetUser req: ", req)

	var userModel *model.UserTbl
	var user types.User

	userModel, err = l.svcCtx.UserModel.FindOne(l.ctx, req.UserID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if userModel == nil {
		return &types.GetUserRes{
			Result: types.Result{
				Code:    common.USER_NOT_EXIT_CODE,
				Message: common.USER_NOT_EXIT_MESS,
			},
		}, nil
	}

	user = types.User{
		ID:       userModel.Id,
		Name:     userModel.Name,
		FullName: userModel.FullName,
		Email:    userModel.Email,
		Role:     userModel.Role.Int64,
	}

	l.Logger.Info("GetUser Success: ", user)
	return &types.GetUserRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		User: user,
	}, nil
}
