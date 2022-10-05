import service from '@/utils/request'

// 用户登录 获取动态路由
export const asyncMenu = () => {
  return service({
    method: 'GET',
    url: '/myMenu'
  })
}

// 获取menu列表
export const getMenuList = (data) => {
  return service({
    method: 'GET',
    url: '/menus',
    data
  })
}

// 根据id查询菜单
export const getBaseMenuById = (id) => {
  return service({
    url: `/menu/${id}`,
    method: 'GET'
  })
}

// 新增menu
export const createMenu = (data) => {
  return service({
    url: '/menu',
    method: 'POST',
    data
  })
}

// 修改menu列表
export const updateMenu = (data) => {
  return service({
    url: '/menu',
    method: 'PUT',
    data
  })
}

// 根据id删除菜单
export const delMenu = (id) => {
  return service({
    url: `/menu/${id}`,
    method: 'DELETE'
  })
}

// 获取某角色的路由和已有权限
export const getBaseMenuTree = (params) => {
  return service({
    url: '/menuTree',
    method: 'GET',
    params
  })
}


// --------------------------------------


// 获取用户menu关联关系
export const getMenuAuthority = (data) => {
  return service({
    url: '/menu/getMenuAuthority',
    method: 'POST',
    data
  })
}