import service from '@/utils/request'

// 获取配置文件内容
export const getSystemConfig = () => {
  return service({
    url: '/system/getSystemConfig',
    method: 'post'
  })
}

// 设置配置文件内容
export const setSystemConfig = (data) => {
  return service({
    url: '/system/setSystemConfig',
    method: 'post',
    data
  })
}

// 获取服务器运行状态
export const getSystemState = () => {
  return service({
    url: '/system/getServerInfo',
    method: 'post',
    donNotShowLoading: true
  })
}
