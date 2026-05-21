package sites

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/sites"
    sitesReq "github.com/flipped-aurora/gin-vue-admin/server/model/sites/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PublishStatsApi struct {}



// CreatePublishStats 创建网站更新数据
// @Tags PublishStats
// @Summary 创建网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.PublishStats true "创建网站更新数据"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /publishStats/createPublishStats [post]
func (publishStatsApi *PublishStatsApi) CreatePublishStats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var publishStats sites.PublishStats
	err := c.ShouldBindJSON(&publishStats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = publishStatsService.CreatePublishStats(ctx,&publishStats)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePublishStats 删除网站更新数据
// @Tags PublishStats
// @Summary 删除网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.PublishStats true "删除网站更新数据"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /publishStats/deletePublishStats [delete]
func (publishStatsApi *PublishStatsApi) DeletePublishStats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := publishStatsService.DeletePublishStats(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePublishStatsByIds 批量删除网站更新数据
// @Tags PublishStats
// @Summary 批量删除网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /publishStats/deletePublishStatsByIds [delete]
func (publishStatsApi *PublishStatsApi) DeletePublishStatsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := publishStatsService.DeletePublishStatsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePublishStats 更新网站更新数据
// @Tags PublishStats
// @Summary 更新网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.PublishStats true "更新网站更新数据"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /publishStats/updatePublishStats [put]
func (publishStatsApi *PublishStatsApi) UpdatePublishStats(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var publishStats sites.PublishStats
	err := c.ShouldBindJSON(&publishStats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = publishStatsService.UpdatePublishStats(ctx,publishStats)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPublishStats 用id查询网站更新数据
// @Tags PublishStats
// @Summary 用id查询网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询网站更新数据"
// @Success 200 {object} response.Response{data=sites.PublishStats,msg=string} "查询成功"
// @Router /publishStats/findPublishStats [get]
func (publishStatsApi *PublishStatsApi) FindPublishStats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	republishStats, err := publishStatsService.GetPublishStats(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(republishStats, c)
}
// GetPublishStatsList 分页获取网站更新数据列表
// @Tags PublishStats
// @Summary 分页获取网站更新数据列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query sitesReq.PublishStatsSearch true "分页获取网站更新数据列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /publishStats/getPublishStatsList [get]
func (publishStatsApi *PublishStatsApi) GetPublishStatsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo sitesReq.PublishStatsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := publishStatsService.GetPublishStatsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetPublishStatsDataSource 获取PublishStats的数据源
// @Tags PublishStats
// @Summary 获取PublishStats的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /publishStats/getPublishStatsDataSource [get]
func (publishStatsApi *PublishStatsApi) GetPublishStatsDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := publishStatsService.GetPublishStatsDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetPublishStatsPublic 不需要鉴权的网站更新数据接口
// @Tags PublishStats
// @Summary 不需要鉴权的网站更新数据接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /publishStats/getPublishStatsPublic [get]
func (publishStatsApi *PublishStatsApi) GetPublishStatsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    publishStatsService.GetPublishStatsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的网站更新数据接口信息",
    }, "获取成功", c)
}
