import service from '@/utils/request'

// 初始化用户数据库
export const initDB = (data) => {
  return service({
    url: '/init/initdb',
    method: 'post',
    data
  })
}

// 初始化用户数据库
export const checkDB = () => {
  return service({
    url: '/init/checkdb',
    method: 'post'
  })
}
