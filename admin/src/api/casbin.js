import service from '@/utils/request'

// 更改角色api权限
export const UpdateCasbin = (data) => {
  return service({
    url: '/casbin/updateCasbin',
    method: 'post',
    data
  })
}

// 获取权限列表
export const getPolicyPathByAuthorityId = (data) => {
  return service({
    url: '/casbin/getPolicyPathByAuthorityId',
    method: 'post',
    data
  })
}
