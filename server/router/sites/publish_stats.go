package sites

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PublishStatsRouter struct {}

// InitPublishStatsRouter 初始化 网站更新数据 路由信息
func (s *PublishStatsRouter) InitPublishStatsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	publishStatsRouter := Router.Group("publishStats").Use(middleware.OperationRecord())
	publishStatsRouterWithoutRecord := Router.Group("publishStats")
	publishStatsRouterWithoutAuth := PublicRouter.Group("publishStats")
	{
		publishStatsRouter.POST("createPublishStats", publishStatsApi.CreatePublishStats)   // 新建网站更新数据
		publishStatsRouter.DELETE("deletePublishStats", publishStatsApi.DeletePublishStats) // 删除网站更新数据
		publishStatsRouter.DELETE("deletePublishStatsByIds", publishStatsApi.DeletePublishStatsByIds) // 批量删除网站更新数据
		publishStatsRouter.PUT("updatePublishStats", publishStatsApi.UpdatePublishStats)    // 更新网站更新数据
	}
	{
		publishStatsRouterWithoutRecord.GET("findPublishStats", publishStatsApi.FindPublishStats)        // 根据ID获取网站更新数据
		publishStatsRouterWithoutRecord.GET("getPublishStatsList", publishStatsApi.GetPublishStatsList)  // 获取网站更新数据列表
	}
	{
	    publishStatsRouterWithoutAuth.GET("getPublishStatsDataSource", publishStatsApi.GetPublishStatsDataSource)  // 获取网站更新数据数据源
	    publishStatsRouterWithoutAuth.GET("getPublishStatsPublic", publishStatsApi.GetPublishStatsPublic)  // 网站更新数据开放接口
	}
}
