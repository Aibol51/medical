package service

import (
	"context"

    "github.com/suyuan32/simple-admin-core/rpc/ent/service"
    "github.com/suyuan32/simple-admin-core/rpc/internal/svc"
    "github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteServiceLogic) DeleteService(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Service.Delete().Where(service.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
