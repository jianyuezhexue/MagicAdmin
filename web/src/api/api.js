import service from '@/utils/request'

// 分页获取角色列表
export const getApiList = (params) => {
  return service({
    url: '/api',
    method: 'GET',
    params
  })
}

// 创建基础api
export const createApi = (data) => {
  return service({
    url: '/api/createApi',
    method: 'post',
    data
  })
}

// 根据id获取菜单
export const getApiById = (data) => {
  return service({
    url: '/api/getApiById',
    method: 'post',
    data
  })
}

// 更新api
export const updateApi = (data) => {
  return service({
    url: '/api/updateApi',
    method: 'post',
    data
  })
}

// 更新api
export const setAuthApi = (data) => {
  return service({
    url: '/api/setAuthApi',
    method: 'post',
    data
  })
}

// 获取所有的Api 不分页
export const getAllApis = (data) => {
  return service({
    url: '/api/getAllApis',
    method: 'post',
    data
  })
}

// 删除指定api
export const deleteApi = (id) => {
  return service({
    url: `/api/${id}`,
    method: 'DELETE'
  })
}

// 删除选中Api
export const deleteApisByIds = (data) => {
  return service({
    url: '/api/deleteApisByIds',
    method: 'delete',
    data
  })
}
