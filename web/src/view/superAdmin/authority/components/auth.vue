<template>
    <div class="continer">
        <div>
            <el-input v-model="filterText" class="fitler" placeholder="关键词搜索" />
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-row>
            <!-- 菜单树形权限 -->
            <el-col :span="5">
                <el-divider content-position="left">菜单权限</el-divider>
                <el-tree ref="treeRef" class="filter-tree" show-checkbox :data="data" :props="menuDefaultProps"
                    :filter-node-method="filterNode" node-key="id" :default-checked-keys="selected"
                    :highlight-current="true" :accordion="true" @check-change="handleCheckChange" />
            </el-col>
            <el-col :span="1">
            </el-col>
            <!-- API操作/数据/按钮权限 -->
            <el-col :span="18">
                <el-divider content-position="left">操作/数据/按钮权限</el-divider>
                <div class="tree-content">
                    <div class="apiAuht" v-for="item,index in apiChecks" :key="index">
                        <label>菜单管理-操作权限</label>
                        <el-checkbox-group v-model="item.checkedApis" @change="handleCheckedCitiesChange">
                            <el-checkbox v-for="Api in item.Apis" :key="Api" :label="Api">{{Api}}</el-checkbox>
                        </el-checkbox-group>
                    </div>
                </div>
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

// 菜单树
const filterText = ref('')
const treeRef = ref()    // 关联树
const data = ref([])     // 菜单数据
const selected = ref([]) // 已有菜单权限
// API/数据/按钮
const checkAll = ref(false)
const isIndeterminate = ref(true)
const checkedApis = ref(['Shanghai', 'Beijing'])
const Apis = ['Shanghai', 'Beijing', 'Guangzhou', 'Shenzhen']
const apiChecks = ref([
    {
        title: "菜单管理",
        meunId: 9,
        type: "API",
        checkedApis: ['菜单列表', 'Beijing'],
        Apis: ['菜单列表', 'Beijing', 'Guangzhou1', 'Shenzhen1']
    },
    {
        title: "字典目录",
        meunId: 9,
        type: "filed",
        checkedApis: ['字典列表', 'Beijing'],
        Apis: ['字典列表', 'Beijing', 'Guangzhou2', 'Shenzhen2']
    }
])

// 初始化查询所有菜单和权限
const init = async () => {
    // 回显菜单数据
    // const res = await getBaseMenuTree({ id: props.row.id })
    const res = await getBaseMenuTree({ id: 1 })
    let menus = res.data.treeMenus
    data.value = menus

    // 回显菜单选中
    let selectedArr = res.data.menuIds.split(",")
    menus.forEach(item => {
        if (item.children.length > 0) { // 防止父级选中子集全选
            selectedArr = selectedArr.filter(val => val != item.id)
        }
    })
    selected.value = selectedArr

    // 回显API，数据，按钮
}
init()


// 菜单树function
const menuDefaultProps = ref({ // 取值说明
    children: 'children',
    label: function (data) {
        return data.meta.title
    }
})
watch(filterText, (val) => { // 监听&关键字过滤
    treeRef.value.filter(val)
})
const filterNode = (value, data) => {
    if (!value) return true
    return data.meta.title.indexOf(value) !== -1
}
// 监听菜单选择
const handleCheckChange = (data, checked, indeterminate) => {
}

// API/数据/按钮function
const handleCheckAllChange = (val) => {  // 全选
    checkedApis.value = val ? Apis : []
    isIndeterminate.value = false
}
const handleCheckedCitiesChange = (value) => { // 单选
    const checkedCount = value.length
    checkAll.value = checkedCount === Apis.length
    isIndeterminate.value = checkedCount > 0 && checkedCount < Apis.length
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

    .apiAuht {
        margin-bottom: 1rem;
    }

    .el-checkbox__label {
        line-height: 1.5;
    }
}
</style>
    