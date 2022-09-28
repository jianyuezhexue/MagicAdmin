import service from '@/utils/request'

// 分页获取字典目录列表
export const getDictionaryList = (params) => {
  return service({
    url: '/dictionary',
    method: 'GET',
    params
  })
}

// 创建字典目录
export const createDictionary = (data) => {
  return service({
    url: '/dictionary',
    method: 'POST',
    data
  })
}

// 删除字典目录
export const deleteDictionary = (data) => {
  return service({
    url: '/dictionary',
    method: 'DELETE',
    data
  })
}

// 更新字典目录
export const updateDictionary = (data) => {
  return service({
    url: '/dictionary',
    method: 'PUT',
    data
  })
}

// 用id查询字典目录
export const findDictionary = (id) => {
  return service({
    url: `/dictionary/${id}`,
    method: 'GET'
  })
}
