package sites

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicInfoRouter struct{}

func (s *BasicInfoRouter) InitBasicInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	basicInfoRouter := Router.Group("basicInfo").Use(middleware.OperationRecord())
	basicInfoRouterWithoutRecord := Router.Group("basicInfo")
	basicInfoRouterWithoutAuth := PublicRouter.Group("basicInfo")
	{
		basicInfoRouter.POST("createBasicInfo", basicInfoApi.CreateBasicInfo)
		basicInfoRouter.DELETE("deleteBasicInfo", basicInfoApi.DeleteBasicInfo)
		basicInfoRouter.DELETE("deleteBasicInfoByIds", basicInfoApi.DeleteBasicInfoByIds)
		basicInfoRouter.PUT("updateBasicInfo", basicInfoApi.UpdateBasicInfo)
	}
	{
		basicInfoRouterWithoutRecord.GET("findBasicInfo", basicInfoApi.FindBasicInfo)
		basicInfoRouterWithoutRecord.GET("getBasicInfoList", basicInfoApi.GetBasicInfoList)
	}
	{
		basicInfoRouterWithoutAuth.GET("getBasicInfoPublic", basicInfoApi.GetBasicInfoPublic)
		basicInfoRouterWithoutAuth.GET("getAllSites", middleware.StaticTokenAuth(), basicInfoApi.GetAllSites)
	}
}
