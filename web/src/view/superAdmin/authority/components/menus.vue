<template>
  <div>
    <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
    </div>
    <div class="tree-content">
      <el-tree ref="menuTree" :data="menuTreeData" :default-checked-keys="menuTreeIds" :props="menuDefaultProps"
        highlight-current node-key="id" show-checkbox :filter-node-method="filterNode" @check="nodeChange">
        <template #default="{ node , data }">
          <span class="custom-tree-node">
            <span>{{ node.label }}</span>
          </span>
        </template>
      </el-tree>
    </div>
    <el-dialog v-model="btnVisible" title="分配按钮" destroy-on-close>
      <el-table ref="btnTableRef" :data="btnData" row-key="ID" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column label="按钮名称" prop="name" />
        <el-table-column label="按钮备注" prop="desc" />
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { getBaseMenuTree } from '@/api/menu'
import {addMenuAuthority } from '@/api/authority'
import {
  updateAuthority
} from '@/api/authority'
import { getAuthorityBtnApi, setAuthorityBtnApi } from '@/api/authorityBtn'
import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  row: {
    default: function () {
      return {}
    },
    type: Object
  }
})

const emit = defineEmits(['changeRow'])
const filterText = ref('')
const menuTreeData = ref([])    // 数据
const menuTreeIds = ref([])     // 选中
const needConfirm = ref(false)
const menuDefaultProps = ref({ // 取值规则
  children: 'children',
  label: function (data) {
    return data.meta.title
  }
})

// 初始化接口
const init = async () => {
  // 回显菜单树
  const res = await getBaseMenuTree()
  menuTreeData.value = res.data

  // 回显菜单选中
  let selectedArr = props.row.menuIds.split(",")
  res.data.forEach(item => {
    if (item.children.length > 0) { // 防止父级选中子集全选
      selectedArr = selectedArr.filter(val => val != item.id)
    }
  })
  menuTreeIds.value = selectedArr
}
init()

const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  relation()
}
// 关联树 确认方法
const menuTree = ref(null)
const relation = async () => {
  const checkArr = menuTree.value.getCheckedNodes(false, true)
  let menuIds = checkArr.map(item => String(item.id))
  const res = await addMenuAuthority({
    // menus: checkArr,
    data:menuIds,
    id: props.row.id
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '菜单设置成功!'
    })
  }
}

defineExpose({ enterAndNext, needConfirm })

const btnVisible = ref(false)

const btnData = ref([])
const multipleSelection = ref([])
const btnTableRef = ref()
let menuID = ''
const OpenBtn = async (data) => {
  menuID = data.ID
  const res = await getAuthorityBtnApi({ menuID: menuID, authorityId: props.row.authorityId })
  if (res.code === 0) {
    openDialog(data)
    await nextTick()
    if (res.data.selected) {
      res.data.selected.forEach(id => {
        btnData.value.some(item => {
          if (item.ID === id) {
            btnTableRef.value.toggleRowSelection(item, true)
          }
        })
      })
    }
  }
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const openDialog = (data) => {
  btnVisible.value = true
  btnData.value = data.menuBtn
}

const closeDialog = () => {
  btnVisible.value = false
}
const enterDialog = async () => {
  const selected = multipleSelection.value.map(item => item.ID)
  const res = await setAuthorityBtnApi({
    menuID,
    selected,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置成功' })
    btnVisible.value = false
  }
}

const filterNode = (value, data) => {
  if (!value) return true
  return data.meta.title.indexOf(value) !== -1
}

watch(filterText, (val) => {
  menuTree.value.filter(val)
})

</script>

<script>

export default {
  name: 'Menus'
}
</script>

<style lang="scss" scope>
@import "@/style/button.scss";

.custom-tree-node {
  span+span {
    margin-left: 12px;
  }
}
</style>
