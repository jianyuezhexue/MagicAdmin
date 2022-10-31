import service from '@/utils/request'

// 分页获取字典详情列表
export const getDictionaryDetailList = (params) => {
  return service({
    url: '/dictionaryDetail',
    method: 'GET',
    params
  })
}

// 创建字典详情
export const createDictionaryDetail = (data) => {
  return service({
    url: '/dictionaryDetail',
    method: 'POST',
    data
  })
}

// 删除字典详情
export const deleteDictionaryDetail = (id) => {
  return service({
    url: `/dictionaryDetail/${id}`,
    method: 'DELETE'
  })
}

// 更新字典详情
export const updateDictionaryDetail = (data) => {
  return service({
    url: '/dictionaryDetail',
    method: 'PUT',
    data
  })
}

// 用id查询字典详情
export const findDictionaryDetail = (id) => {
  return service({
    url: `/dictionaryDetail/${id}`,
    method: 'GET'
  })
}
