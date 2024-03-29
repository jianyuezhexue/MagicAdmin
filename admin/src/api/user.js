import service from '@/utils/request'

// 获取验证码
export const captcha = (data) => {
  return service({
    url: '/user/captcha',
    method: 'post',
    data: data
  })
}

// 用户登录
export const login = (data) => {
  return service({
    url: '/user/login',
    method: 'post',
    data: data
  })
}

// 用户注册
export const register = (data) => {
  return service({
    url: '/user/register',
    method: 'POST',
    data: data
  })
}

// 修改密码
export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// 分页获取用户列表
export const getUserList = (params) => {
  return service({
    url: '/users',
    method: 'GET',
    params
  })
}

// 切换用户权限
export const switchAuth = (data) => {
  return service({
    url: '/user/switchAuth',
    method: 'PATCH',
    data: data
  })
}

// 删除用户
export const deleteUser = (id) => {
  return service({
    url: `/user/${id}`,
    method: 'DELETE',
  })
}

// 更新用户信息
export const setUserInfo = (data) => {
  return service({
    url: '/user',
    method: 'PUT',
    data: data
  })
}

// 设置用户信息
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// 设置用户权限
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/auth',
    method: 'PATCH',
    data: data
  })
}

// 设置用户状态-启用和禁用
export const setUserStatus = (data) => {
  return service({
    url: '/user/status',
    method: 'PATCH',
    data: data
  })
}

// 获取用户信息
export const getUserInfo = () => {
  return service({
    url: '/user/info',
    method: 'get'
  })
}

// 重置密码
export const resetPassword = (id) => {
  return service({
    url: `/user/reSetPwd/${id}`,
    method: 'PATCH',
  })
}
