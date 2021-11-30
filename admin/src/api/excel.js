import service from '@/utils/request'
import { ElMessage } from 'element-plus'

const handleFileError = (res, fileName) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    var downloadUrl = window.URL.createObjectURL(new Blob([res]))
    var a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = fileName
    var event = new MouseEvent('click')
    a.dispatchEvent(event)
  }
}

// 导出Excel
export const exportExcel = (tableData, fileName) => {
  service({
    url: '/excel/exportExcel',
    method: 'post',
    data: {
      fileName: fileName,
      infoList: tableData
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, fileName)
  })
}

// 导入Excel文件
export const loadExcelData = () => {
  return service({
    url: '/excel/loadExcel',
    method: 'get'
  })
}

// 下载模板
export const downloadTemplate = (fileName) => {
  return service({
    url: '/excel/downloadTemplate',
    method: 'get',
    params: {
      fileName: fileName
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, fileName)
  })
}
