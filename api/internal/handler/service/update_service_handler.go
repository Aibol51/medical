package service

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/service"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /service/update service UpdateService
//
// Update service information | 更新Service信息
//
// Update service information | 更新Service信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ServiceInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ServiceInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := service.NewUpdateServiceLogic(r.Context(), svcCtx)
		resp, err := l.UpdateService(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
