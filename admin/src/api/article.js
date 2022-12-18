import service from '@/utils/request'

// 创建基础api
export const createArticle = (data) => {
    return service({
        url: '/article',
        method: 'POST',
        data
    })
}