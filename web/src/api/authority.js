import service from '@/utils/request'
// 获取角色列表
export const getAuthorityList = (data) => {
  return service({
    url: '/authorityTree',
    method: 'GET',
    data
  })
}

// 删除角色
export const deleteAuthority = (data) => {
  return service({
    url: '/authority/deleteAuthority',
    method: 'DELETE',
    data
  })
}

// 创建角色
export const createAuthority = (data) => {
  return service({
    url: '/authority/createAuthority',
    method: 'POST',
    data
  })
}

// 拷贝角色
export const copyAuthority = (data) => {
  return service({
    url: '/authority/copyAuthority',
    method: 'POST',
    data
  })
}

// 设置角色资源权限
export const setDataAuthority = (data) => {
  return service({
    url: '/authority/setDataAuthority',
    method: 'POST',
    data
  })
}

// 修改角色
export const updateAuthority = (data) => {
  return service({
    url: '/authority/updateAuthority',
    method: 'PUT',
    data
  })
}
