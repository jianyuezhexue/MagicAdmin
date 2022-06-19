import service from '@/utils/request'

// 删除SysOperationRecord
export const deleteSysOperationRecord = (data) => {
    return service({
        url: '/sysOperationRecord/deleteSysOperationRecord',
        method: 'delete',
        data
    })
}

// 删除SysOperationRecord
export const deleteSysOperationRecordByIds = (data) => {
    return service({
        url: '/sysOperationRecord/deleteSysOperationRecordByIds',
        method: 'delete',
        data
    })
}

// 分页获取SysOperationRecord列表
export const getSysOperationRecordList = (params) => {
    return service({
        url: '/sysOperationRecord/getSysOperationRecordList',
        method: 'get',
        params
    })
}