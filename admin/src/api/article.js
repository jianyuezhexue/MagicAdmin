import service from '@/utils/request'

// 创建基础api
export const createArticle = (data) => {
    return service({
        url: '/article',
        method: 'POST',
        data
    })
}

// 分页查询文章列表
export const getArticleList = (params) => {
    return service({
        url: '/article',
        method: 'GET',
        params
    })
}

// 根据id查找文章
export const findArticle = (id) => {
    return service({
        url: `/article/${id}`,
        method: 'GET',
    })
}