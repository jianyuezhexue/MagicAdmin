import service from '@/utils/request'

// 获取角色列表
export const getAuthorityList = (data) => {
  return service({
    url: '/authority/getAuthorityList',
    method: 'post',
    data
  })
}

// 删除角色
export const deleteAuthority = (data) => {
  return service({
    url: '/authority/deleteAuthority',
    method: 'post',
    data
  })
}

// 创建角色
export const createAuthority = (data) => {
  return service({
    url: '/authority/createAuthority',
    method: 'post',
    data
  })
}

// 拷贝角色
export const copyAuthority = (data) => {
  return service({
    url: '/authority/copyAuthority',
    method: 'post',
    data
  })
}

// 设置角色资源权限
export const setDataAuthority = (data) => {
  return service({
    url: '/authority/setDataAuthority',
    method: 'post',
    data
  })
}

// 修改角色
export const updateAuthority = (data) => {
  return service({
    url: '/authority/updateAuthority',
    method: 'put',
    data
  })
}
