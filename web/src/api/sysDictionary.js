import service from '@/utils/request'
// @Tags Dictionary
// @Summary 创建Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "创建Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dictionary/createDictionary [post]
export const createDictionary = (data) => {
  return service({
    url: '/Dictionary/createDictionary',
    method: 'post',
    data
  })
}

// @Tags Dictionary
// @Summary 删除Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "删除Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Dictionary/deleteDictionary [delete]
export const deleteDictionary = (data) => {
  return service({
    url: '/Dictionary/deleteDictionary',
    method: 'delete',
    data
  })
}

// @Tags Dictionary
// @Summary 更新Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "更新Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Dictionary/updateDictionary [put]
export const updateDictionary = (data) => {
  return service({
    url: '/Dictionary/updateDictionary',
    method: 'put',
    data
  })
}

// @Tags Dictionary
// @Summary 用id查询Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dictionary true "用id查询Dictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Dictionary/findDictionary [get]
export const findDictionary = (params) => {
  return service({
    url: '/Dictionary/findDictionary',
    method: 'get',
    params
  })
}

// @Tags Dictionary
// @Summary 分页获取Dictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Dictionary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dictionary/getDictionaryList [get]
export const getDictionaryList = (params) => {
  return service({
    url: '/Dictionary/getDictionaryList',
    method: 'get',
    params
  })
}
