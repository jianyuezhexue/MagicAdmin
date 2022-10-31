import service from '@/utils/request'
// 删除操作记录
export const deleteSysOperationRecord = (data) => {
  return service({
    url: '/recode',
    method: 'delete',
    data
  })
}

// 批量删除操作记录
export const deleteSysOperationRecordByIds = (data) => {
  return service({
    url: '/recode',
    method: 'delete',
    data
  })
}

// 分页查询操作记录
export const getSysOperationRecordList = (params) => {
  return service({
    url: '/recode',
    method: 'GET',
    params
  })
}
