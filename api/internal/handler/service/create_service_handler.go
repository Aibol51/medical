package service

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/service"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /service/create service CreateService
//
// Create service information | 创建Service信息
//
// Create service information | 创建Service信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ServiceInfo
//
// Responses:
//  200: BaseMsgResp

func CreateServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ServiceInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := service.NewCreateServiceLogic(r.Context(), svcCtx)
		resp, err := l.CreateService(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
