import service from '@/utils/request'
// @Tags DictionaryDetail
// @Summary 创建DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "创建DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/createDictionaryDetail [post]
export const createDictionaryDetail = (data) => {
  return service({
    url: '/DictionaryDetail/createDictionaryDetail',
    method: 'post',
    data
  })
}

// @Tags DictionaryDetail
// @Summary 删除DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "删除DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DictionaryDetail/deleteDictionaryDetail [delete]
export const deleteDictionaryDetail = (data) => {
  return service({
    url: '/DictionaryDetail/deleteDictionaryDetail',
    method: 'delete',
    data
  })
}

// @Tags DictionaryDetail
// @Summary 更新DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "更新DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DictionaryDetail/updateDictionaryDetail [put]
export const updateDictionaryDetail = (data) => {
  return service({
    url: '/DictionaryDetail/updateDictionaryDetail',
    method: 'put',
    data
  })
}

// @Tags DictionaryDetail
// @Summary 用id查询DictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DictionaryDetail true "用id查询DictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DictionaryDetail/findDictionaryDetail [get]
export const findDictionaryDetail = (params) => {
  return service({
    url: '/DictionaryDetail/findDictionaryDetail',
    method: 'get',
    params
  })
}

// @Tags DictionaryDetail
// @Summary 分页获取DictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DictionaryDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DictionaryDetail/getDictionaryDetailList [get]
export const getDictionaryDetailList = (params) => {
  return service({
    url: '/DictionaryDetail/getDictionaryDetailList',
    method: 'get',
    params
  })
}
