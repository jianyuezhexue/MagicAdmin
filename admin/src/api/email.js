import service from '@/utils/request'

// 发送测试邮件
export const emailTest = (data) => {
  return service({
    url: '/email/emailTest',
    method: 'post',
    data
  })
}
