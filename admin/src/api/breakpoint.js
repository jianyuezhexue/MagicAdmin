import service from '@/utils/request'

// 设置角色资源权限
export const findFile = (params) => {
  return service({
    url: '/fileUploadAndDownload/findFile',
    method: 'get',
    params
  })
}

export const breakpointContinueFinish = (params) => {
  return service({
    url: '/fileUploadAndDownload/breakpointContinueFinish',
    method: 'post',
    params
  })
}

export const removeChunk = (data, params) => {
  return service({
    url: '/fileUploadAndDownload/removeChunk',
    method: 'post',
    data,
    params
  })
}
