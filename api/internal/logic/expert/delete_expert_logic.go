package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExpertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteExpertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExpertLogic {
	return &DeleteExpertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteExpertLogic) DeleteExpert(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteExpert(l.ctx, &core.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
