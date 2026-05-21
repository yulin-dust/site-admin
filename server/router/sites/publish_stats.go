package sites

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PublishStatsRouter struct{}

func (s *PublishStatsRouter) InitPublishStatsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	publishStatsRouter := Router.Group("publishStats").Use(middleware.OperationRecord())
	publishStatsRouterWithoutRecord := Router.Group("publishStats")
	publishStatsRouterWithoutAuth := PublicRouter.Group("publishStats")
	{
		publishStatsRouter.POST("createPublishStats", publishStatsApi.CreatePublishStats)
		publishStatsRouter.DELETE("deletePublishStats", publishStatsApi.DeletePublishStats)
		publishStatsRouter.DELETE("deletePublishStatsByIds", publishStatsApi.DeletePublishStatsByIds)
		publishStatsRouter.PUT("updatePublishStats", publishStatsApi.UpdatePublishStats)
	}
	{
		publishStatsRouterWithoutRecord.GET("findPublishStats", publishStatsApi.FindPublishStats)
		publishStatsRouterWithoutRecord.GET("getPublishStatsList", publishStatsApi.GetPublishStatsList)
	}
	{
		publishStatsRouterWithoutAuth.GET("getPublishStatsDataSource", publishStatsApi.GetPublishStatsDataSource)
		publishStatsRouterWithoutAuth.GET("getPublishStatsPublic", publishStatsApi.GetPublishStatsPublic)
		publishStatsRouterWithoutAuth.GET("addPublishCount", middleware.StaticTokenAuth(), publishStatsApi.AddPublishCount)
	}
}
