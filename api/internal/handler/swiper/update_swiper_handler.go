package swiper

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/swiper"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /swiper/update swiper UpdateSwiper
//
// Update swiper information | 更新Swiper信息
//
// Update swiper information | 更新Swiper信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SwiperInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateSwiperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SwiperInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := swiper.NewUpdateSwiperLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSwiper(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
