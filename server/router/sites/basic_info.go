package sites

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicInfoRouter struct {}

// InitBasicInfoRouter 初始化 基础信息 路由信息
func (s *BasicInfoRouter) InitBasicInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	basicInfoRouter := Router.Group("basicInfo").Use(middleware.OperationRecord())
	basicInfoRouterWithoutRecord := Router.Group("basicInfo")
	basicInfoRouterWithoutAuth := PublicRouter.Group("basicInfo")
	{
		basicInfoRouter.POST("createBasicInfo", basicInfoApi.CreateBasicInfo)   // 新建基础信息
		basicInfoRouter.DELETE("deleteBasicInfo", basicInfoApi.DeleteBasicInfo) // 删除基础信息
		basicInfoRouter.DELETE("deleteBasicInfoByIds", basicInfoApi.DeleteBasicInfoByIds) // 批量删除基础信息
		basicInfoRouter.PUT("updateBasicInfo", basicInfoApi.UpdateBasicInfo)    // 更新基础信息
	}
	{
		basicInfoRouterWithoutRecord.GET("findBasicInfo", basicInfoApi.FindBasicInfo)        // 根据ID获取基础信息
		basicInfoRouterWithoutRecord.GET("getBasicInfoList", basicInfoApi.GetBasicInfoList)  // 获取基础信息列表
	}
	{
	    basicInfoRouterWithoutAuth.GET("getBasicInfoPublic", basicInfoApi.GetBasicInfoPublic)  // 基础信息开放接口
	}
}
