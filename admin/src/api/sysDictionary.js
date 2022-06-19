import service from '@/utils/request'

// 创建SysDictionary
export const createSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/createSysDictionary',
    method: 'post',
    data
  })
}

// 删除SysDictionary
export const deleteSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/deleteSysDictionary',
    method: 'delete',
    data
  })
}

// 更新SysDictionary
export const updateSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/updateSysDictionary',
    method: 'put',
    data
  })
}

// 用id查询SysDictionary
export const findSysDictionary = (params) => {
  return service({
    url: '/sysDictionary/findSysDictionary',
    method: 'get',
    params
  })
}

// 分页获取SysDictionary列表
export const getSysDictionaryList = (params) => {
  return service({
    url: '/sysDictionary/getSysDictionaryList',
    method: 'get',
    params
  })
}
