import service from '@/utils/request'

// jwt加入黑名单
export const jsonInBlacklist = () => {
  return service({
    url: '/jwt/jsonInBlacklist',
    method: 'post'
  })
}
