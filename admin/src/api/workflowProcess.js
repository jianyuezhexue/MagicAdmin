import service from '@/utils/request'

// 创建WorkflowProcess
export const createWorkflowProcess = (data) => {
    return service({
        url: '/workflowProcess/createWorkflowProcess',
        method: 'post',
        data
    })
}

// 删除WorkflowProcess
export const deleteWorkflowProcess = (data) => {
    return service({
        url: '/workflowProcess/deleteWorkflowProcess',
        method: 'delete',
        data
    })
}

// 删除WorkflowProcess
export const deleteWorkflowProcessByIds = (data) => {
    return service({
        url: '/workflowProcess/deleteWorkflowProcessByIds',
        method: 'delete',
        data
    })
}

// 更新WorkflowProcess
export const updateWorkflowProcess = (data) => {
    return service({
        url: '/workflowProcess/updateWorkflowProcess',
        method: 'put',
        data
    })
}

// 用id查询WorkflowProcess
export const findWorkflowProcess = (params) => {
    return service({
        url: '/workflowProcess/findWorkflowProcess',
        method: 'get',
        params
    })
}

// 分页获取WorkflowProcess列表
export const getWorkflowProcessList = (params) => {
    return service({
        url: '/workflowProcess/getWorkflowProcessList',
        method: 'get',
        params
    })
}

// 用id查询工作流步骤
export const findWorkflowStep = (params) => {
    return service({
        url: '/workflowProcess/findWorkflowStep',
        method: 'get',
        params
    })
}

// 创建ExaWfLeave
export const startWorkflow = (data, params = { businessType: data.wf.businessType }) => {
    return service({
        url: '/workflowProcess/startWorkflow',
        method: 'post',
        data,
        params
    })
}

// 创建ExaWfLeave
export const completeWorkflowMove = (data, params = { businessType: data.wf.businessType }) => {
    return service({
        url: '/workflowProcess/completeWorkflowMove',
        method: 'post',
        data,
        params
    })
}

// 我发起的工作流
export const getMyStated = () => {
    return service({
        url: '/workflowProcess/getMyStated',
        method: 'get'
    })
}

// 我发起的工作流
export const getMyNeed = () => {
    return service({
        url: '/workflowProcess/getMyNeed',
        method: 'get'
    })
}

// 根据id获取当前节点详情和历史
export const getWorkflowMoveByID = (params) => {
    return service({
        url: '/workflowProcess/getWorkflowMoveByID',
        method: 'get',
        params
    })
}