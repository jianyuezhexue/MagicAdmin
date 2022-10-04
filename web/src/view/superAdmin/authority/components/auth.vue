<template>
    <div class="continer">
        <div>
            <el-input v-model="filterText" class="fitler" placeholder="关键词搜索" />
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-row>
            <el-col :span="5">
                <el-divider content-position="left">菜单权限</el-divider>
                <el-tree ref="treeRef" class="filter-tree" show-checkbox :data="data" :props="menuDefaultProps"
                    :filter-node-method="filterNode" node-key="id" :default-checked-keys="selected"
                    :highlight-current="true" :accordion="true" @check-change="handleCheckChange" />
            </el-col>
            <el-col :span="1">
            </el-col>
            <el-col :span="18">
                <el-divider content-position="left">操作/数据/按钮权限</el-divider>
                <!-- 回显API -->
            </el-col>
        </el-row>

    </div>

</template>
  
<script setup>
import { ref, watch } from 'vue'
import { ElTree } from 'element-plus'
import { getBaseMenuTree } from '@/api/menu'

// 接收角色信息
const props = defineProps({
    row: {
        default: function () {
            return {}
        },
        type: Object
    }
})

const filterText = ref('')
const treeRef = ref()    // 关联树
const data = ref([])     // 菜单数据
const selected = ref([]) // 已有菜单权限

// 初始化查询所有菜单和权限
const init = async () => {
    // 回显菜单数据
    const res = await getBaseMenuTree()
    let menus = res.data
    data.value = menus

    // 回显选中
    let selectedArr = props.row.menuIds.split(",")
    menus.forEach(item => {
        if (item.children.length > 0) { // 防止父级选中子集全选
            selectedArr = selectedArr.filter(val => val != item.id)
        }
    })
    selected.value = selectedArr

    // 回显API，数据，按钮
}
init()

// 取值说明
const menuDefaultProps = ref({
    children: 'children',
    label: function (data) {
        return data.meta.title
    }
})
// 监听&关键字过滤
watch(filterText, (val) => {
    treeRef.value.filter(val)
})
const filterNode = (value, data) => {
    if (!value) return true
    return data.meta.title.indexOf(value) !== -1
}
// 监听选择
const handleCheckChange = (data, checked, indeterminate) => {
}

// 请求设置权限
const relation = async () => {
    alert("发请求接口")
    // const checkArr = menuTree.value.getCheckedNodes(false, true)
    // const res = await addMenuAuthority({
    //     menus: checkArr,
    //     authorityId: props.row.authorityId
    // })
    // if (res.code === 0) {
    //     ElMessage({
    //         type: 'success',
    //         message: '菜单设置成功!'
    //     })
    // }
}

</script>
  
<script>
export default { name: 'Auth' }
</script>

<style lang="scss" scope>
@import "@/style/button.scss";

.continer {
    padding: 0 1.5rem;
}

.el-drawer__title {
    font-size: 1.5rem;
}
</style>
    