import service from '@/utils/request'
// @Tags PublishStats
// @Summary 创建网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PublishStats true "创建网站更新数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /publishStats/createPublishStats [post]
export const createPublishStats = (data) => {
  return service({
    url: '/publishStats/createPublishStats',
    method: 'post',
    data
  })
}

// @Tags PublishStats
// @Summary 删除网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PublishStats true "删除网站更新数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /publishStats/deletePublishStats [delete]
export const deletePublishStats = (params) => {
  return service({
    url: '/publishStats/deletePublishStats',
    method: 'delete',
    params
  })
}

// @Tags PublishStats
// @Summary 批量删除网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网站更新数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /publishStats/deletePublishStats [delete]
export const deletePublishStatsByIds = (params) => {
  return service({
    url: '/publishStats/deletePublishStatsByIds',
    method: 'delete',
    params
  })
}

// @Tags PublishStats
// @Summary 更新网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PublishStats true "更新网站更新数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /publishStats/updatePublishStats [put]
export const updatePublishStats = (data) => {
  return service({
    url: '/publishStats/updatePublishStats',
    method: 'put',
    data
  })
}

// @Tags PublishStats
// @Summary 用id查询网站更新数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PublishStats true "用id查询网站更新数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /publishStats/findPublishStats [get]
export const findPublishStats = (params) => {
  return service({
    url: '/publishStats/findPublishStats',
    method: 'get',
    params
  })
}

// @Tags PublishStats
// @Summary 分页获取网站更新数据列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取网站更新数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /publishStats/getPublishStatsList [get]
export const getPublishStatsList = (params) => {
  return service({
    url: '/publishStats/getPublishStatsList',
    method: 'get',
    params
  })
}
// @Tags PublishStats
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /publishStats/findPublishStatsDataSource [get]
export const getPublishStatsDataSource = () => {
  return service({
    url: '/publishStats/getPublishStatsDataSource',
    method: 'get',
  })
}

// @Tags PublishStats
// @Summary 不需要鉴权的网站更新数据接口
// @Accept application/json
// @Produce application/json
// @Param data query sitesReq.PublishStatsSearch true "分页获取网站更新数据列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /publishStats/getPublishStatsPublic [get]
export const getPublishStatsPublic = () => {
  return service({
    url: '/publishStats/getPublishStatsPublic',
    method: 'get',
  })
}
// AddPublishCount 调用就自动添加当天的更新即总更新数量
// @Tags PublishStats
// @Summary 调用就自动添加当天的更新即总更新数量
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /publishStats/addPublishCount [GET]
export const addPublishCount = () => {
  return service({
    url: '/publishStats/addPublishCount',
    method: 'GET'
  })
}
