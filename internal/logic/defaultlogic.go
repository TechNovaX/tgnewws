package logic

import (
	"context"

	"tgnewws/internal/svc"
	"tgnewws/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DefaultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DefaultLogic {
	return &DefaultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DefaultLogic) Default(req *types.DefaultReq) (resp *types.DefaultResp, err error) {
	// todo: add your logic here and delete this line

	return
}
