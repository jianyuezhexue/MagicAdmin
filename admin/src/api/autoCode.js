import service from '@/utils/request'

export const preview = (data) => {
  return service({
    url: '/autoCode/preview',
    method: 'post',
    data
  })
}

export const createTemp = (data) => {
  return service({
    url: '/autoCode/createTemp',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

// 获取当前所有数据库
export const getDB = () => {
  return service({
    url: '/autoCode/getDB',
    method: 'get'
  })
}

// 获取当前数据库所有表
export const getTable = (params) => {
  return service({
    url: '/autoCode/getTables',
    method: 'get',
    params
  })
}

// 获取当前数据库所有表
export const getColumn = (params) => {
  return service({
    url: '/autoCode/getColumn',
    method: 'get',
    params
  })
}

export const getSysHistory = (data) => {
  return service({
    url: '/autoCode/getSysHistory',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return service({
    url: '/autoCode/rollback',
    method: 'post',
    data
  })
}

export const getMeta = (data) => {
  return service({
    url: '/autoCode/getMeta',
    method: 'post',
    data
  })
}

export const delSysHistory = (data) => {
  return service({
    url: '/autoCode/delSysHistory',
    method: 'post',
    data
  })
}
