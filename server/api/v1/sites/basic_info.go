package sites

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
	sitesReq "github.com/flipped-aurora/gin-vue-admin/server/model/sites/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BasicInfoApi struct{}

// CreateBasicInfo 创建基础信息
// @Tags BasicInfo
// @Summary 创建基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.BasicInfo true "创建基础信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /basicInfo/createBasicInfo [post]
func (basicInfoApi *BasicInfoApi) CreateBasicInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var basicInfo sites.BasicInfo
	err := c.ShouldBindJSON(&basicInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = basicInfoService.CreateBasicInfo(ctx, &basicInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBasicInfo 删除基础信息
// @Tags BasicInfo
// @Summary 删除基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.BasicInfo true "删除基础信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /basicInfo/deleteBasicInfo [delete]
func (basicInfoApi *BasicInfoApi) DeleteBasicInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := basicInfoService.DeleteBasicInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBasicInfoByIds 批量删除基础信息
// @Tags BasicInfo
// @Summary 批量删除基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /basicInfo/deleteBasicInfoByIds [delete]
func (basicInfoApi *BasicInfoApi) DeleteBasicInfoByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := basicInfoService.DeleteBasicInfoByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBasicInfo 更新基础信息
// @Tags BasicInfo
// @Summary 更新基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body sites.BasicInfo true "更新基础信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /basicInfo/updateBasicInfo [put]
func (basicInfoApi *BasicInfoApi) UpdateBasicInfo(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var basicInfo sites.BasicInfo
	err := c.ShouldBindJSON(&basicInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = basicInfoService.UpdateBasicInfo(ctx, basicInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBasicInfo 用id查询基础信息
// @Tags BasicInfo
// @Summary 用id查询基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询基础信息"
// @Success 200 {object} response.Response{data=sites.BasicInfo,msg=string} "查询成功"
// @Router /basicInfo/findBasicInfo [get]
func (basicInfoApi *BasicInfoApi) FindBasicInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rebasicInfo, err := basicInfoService.GetBasicInfo(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebasicInfo, c)
}

// GetBasicInfoList 分页获取基础信息列表
// @Tags BasicInfo
// @Summary 分页获取基础信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query sitesReq.BasicInfoSearch true "分页获取基础信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /basicInfo/getBasicInfoList [get]
func (basicInfoApi *BasicInfoApi) GetBasicInfoList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo sitesReq.BasicInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := basicInfoService.GetBasicInfoInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBasicInfoPublic 不需要鉴权的基础信息接口
// @Tags BasicInfo
// @Summary 不需要鉴权的基础信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /basicInfo/getBasicInfoPublic [get]
func (basicInfoApi *BasicInfoApi) GetBasicInfoPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	basicInfoService.GetBasicInfoPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的基础信息接口信息",
	}, "获取成功", c)
}

// GetAllSites 获取所有正常状态的网站信息
// @Tags BasicInfo
// @Summary 获取正常状态的网站信息；pending 为真时仅返回今日发布数未达到计划数的站点
// @Accept application/json
// @Produce application/json
// @Param pending query bool false "为 true/1 时仅返回 today_count < plan_count 的待发布站点"
// @Param token header string true "静态校验 token"
// @Success 200 {object} response.Response{data=[]response.BasicInfoBrief,msg=string} "成功"
// @Router /basicInfo/getAllSites [GET]
func (basicInfoApi *BasicInfoApi) GetAllSites(c *gin.Context) {
	ctx := c.Request.Context()

	// 不传 / 空值 / 解析失败均视为 false，即返回所有正常站点
	onlyPending, _ := strconv.ParseBool(c.Query("pending"))

	list, err := basicInfoService.GetAllSites(ctx, onlyPending)
	if err != nil {
		global.GVA_LOG.Error("获取所有站点失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
