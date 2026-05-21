import service from '@/utils/request'
// @Tags BasicInfo
// @Summary 创建基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BasicInfo true "创建基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /basicInfo/createBasicInfo [post]
export const createBasicInfo = (data) => {
  return service({
    url: '/basicInfo/createBasicInfo',
    method: 'post',
    data
  })
}

// @Tags BasicInfo
// @Summary 删除基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BasicInfo true "删除基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicInfo/deleteBasicInfo [delete]
export const deleteBasicInfo = (params) => {
  return service({
    url: '/basicInfo/deleteBasicInfo',
    method: 'delete',
    params
  })
}

// @Tags BasicInfo
// @Summary 批量删除基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicInfo/deleteBasicInfo [delete]
export const deleteBasicInfoByIds = (params) => {
  return service({
    url: '/basicInfo/deleteBasicInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags BasicInfo
// @Summary 更新基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BasicInfo true "更新基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicInfo/updateBasicInfo [put]
export const updateBasicInfo = (data) => {
  return service({
    url: '/basicInfo/updateBasicInfo',
    method: 'put',
    data
  })
}

// @Tags BasicInfo
// @Summary 用id查询基础信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.BasicInfo true "用id查询基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicInfo/findBasicInfo [get]
export const findBasicInfo = (params) => {
  return service({
    url: '/basicInfo/findBasicInfo',
    method: 'get',
    params
  })
}

// @Tags BasicInfo
// @Summary 分页获取基础信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取基础信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicInfo/getBasicInfoList [get]
export const getBasicInfoList = (params) => {
  return service({
    url: '/basicInfo/getBasicInfoList',
    method: 'get',
    params
  })
}

// @Tags BasicInfo
// @Summary 不需要鉴权的基础信息接口
// @Accept application/json
// @Produce application/json
// @Param data query sitesReq.BasicInfoSearch true "分页获取基础信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /basicInfo/getBasicInfoPublic [get]
export const getBasicInfoPublic = () => {
  return service({
    url: '/basicInfo/getBasicInfoPublic',
    method: 'get',
  })
}
