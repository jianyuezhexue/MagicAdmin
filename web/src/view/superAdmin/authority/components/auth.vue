<template>
    <div>
        <div class="clearfix sticky-button">
            <el-input v-model="filterText" class="fitler" placeholder="关键词搜索" />
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-tree ref="treeRef" class="filter-tree" show-checkbox :data="data" :props="menuDefaultProps"
            :filter-node-method="filterNode" @check-change="handleCheckChange" />
    </div>
</template>
  
<script setup>
import { ref, watch } from 'vue'
import { ElTree } from 'element-plus'

// API接口
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'

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
const treeRef = ref()
const data = ref([]) // 数据

// 初始化查询所有菜单和权限
const init = async () => {
    const res = await getBaseMenuTree()
    // TODO:重新处理数据:API,其他权限拿处理

    data.value = res.data
    console.log(res)
    //   menuTreeData.value = res.data.menus
    //   const res1 = await getMenuAuthority({ authorityId: props.row.authorityId })
    //   const menus = res1.data.menus
    //   const arr = []
    //   menus.forEach(item => {
    //     // 防止直接选中父级造成全选
    //     if (!menus.some(same => same.parentId === item.menuId)) {
    //       arr.push(Number(item.menuId))
    //     }
    //   })
    //   menuTreeIds.value = arr
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
    const checkArr = menuTree.value.getCheckedNodes(false, true)
    const res = await addMenuAuthority({
        menus: checkArr,
        authorityId: props.row.authorityId
    })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '菜单设置成功!'
        })
    }
}

</script>
  
<script>
export default { name: 'Auth' }
</script>

<style lang="scss" scope>
@import "@/style/button.scss";

.continer {
    padding: 1.5rem 0;
}
</style>
    