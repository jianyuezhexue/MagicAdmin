import service from '@/utils/request'

// 用户登录 获取动态路由
export const asyncMenu = () => {
  return service({
    url: '/menu',
    method: 'GET'
  })
}

// 获取menu列表
export const getMenuList = (data) => {
  return service({
    url: '/menu/getMenuList',
    method: 'POST',
    data
  })
}

// 新增基础menu
export const addBaseMenu = (data) => {
  return service({
    url: '/menu/addBaseMenu',
    method: 'POST',
    data
  })
}

// 获取基础路由列表
export const getBaseMenuTree = () => {
  return service({
    url: '/menu/getBaseMenuTree',
    method: 'POST'
  })
}

// 添加用户menu关联关系
export const addMenuAuthority = (data) => {
  return service({
    url: '/menu/addMenuAuthority',
    method: 'POST',
    data
  })
}

// 获取用户menu关联关系
export const getMenuAuthority = (data) => {
  return service({
    url: '/menu/getMenuAuthority',
    method: 'POST',
    data
  })
}

// 获取用户menu关联关系
export const deleteBaseMenu = (data) => {
  return service({
    url: '/menu/deleteBaseMenu',
    method: 'POST',
    data
  })
}

// 修改menu列表
export const updateBaseMenu = (data) => {
  return service({
    url: '/menu/updateBaseMenu',
    method: 'POST',
    data
  })
}

// 根据id获取菜单
export const getBaseMenuById = (data) => {
  return service({
    url: '/menu/getBaseMenuById',
    method: 'POST',
    data
  })
}
