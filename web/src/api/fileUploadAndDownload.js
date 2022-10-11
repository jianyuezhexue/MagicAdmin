import service from '@/utils/request'
// @Summary 分页文件列表
export const getFileList = (params) => {
  return service({
    url: '/common/fileList',
    method: 'GET',
    params
  })
}

// @Summary 删除文件
export const deleteFile = (data) => {
  return service({
    url: '/fileUploadAndDownload/deleteFile',
    method: 'post',
    data
  })
}

// 编辑文件名或者备注
export const editFileName = (data) => {
  return service({
    url: '/fileUploadAndDownload/editFileName',
    method: 'post',
    data
  })
}
