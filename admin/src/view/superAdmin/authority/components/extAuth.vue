<template>
  <div>
    <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button class="fl-right" size="small" type="primary" @click="authExtEnter">确 定</el-button>
    </div>
    <div class="tree-content">
      <el-tree ref="apiTree" :data="extAuthTreeData" :default-checked-keys="extAuthTreeIds" :props="apiDefaultProps"
        default-expand-all highlight-current node-key="id" show-checkbox :filter-node-method="filterNode"
        @check="nodeChange" />
    </div>
  </div>
</template>
<script>
export default {
  name: 'ExtAuth',
}
</script>
<script setup>
import { addExtAuthAuthority } from '@/api/authority'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
const props = defineProps({
  row: {
    default: function () {
      return {}
    },
    type: Object
  },
  extAuthTreeData: {
    default: function () {
      return {}
    },
    type: Object
  },
  extAuthTreeIds: {
    default: function () {
      return {}
    },
    type: Array
  }
})

const apiDefaultProps = ref({
  children: 'children',
  label: 'name'
})
const filterText = ref('')
const extAuthTreeData = ref([])
const extAuthTreeIds = ref([])
const init = async () => {
  // 区分按钮和数据权限
  let newTreeData = []
  props.extAuthTreeData.forEach(item => {
    let children = []
    item.children.forEach(child => {
      if (child.type == "1") { // 按钮权限
        child.name = "按钮-" + child.name
      } else {                // 数据权限
        child.name = "数据-" + child.name
      }
      children.push(child)
    });
    item.children = children
    newTreeData.push(item)
  });

  // 回显拓展权限树
  extAuthTreeData.value = newTreeData
  // 回显拓展权限选中
  extAuthTreeIds.value = props.extAuthTreeIds
}
init()

const needConfirm = ref(false)
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authExtEnter()
}

// 关联关系确定
const apiTree = ref(null)
const authExtEnter = async () => {
  let checkArr = apiTree.value.getCheckedKeys(true, false)
  let extIds = checkArr.map(item => String(item))
  const res = await addExtAuthAuthority({
    data: extIds,
    id: props.row.id
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '拓展权限设置成功' })
  }
}

defineExpose({
  needConfirm,
  enterAndNext
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.name.indexOf(value) !== -1
}
watch(filterText, (val) => {
  apiTree.value.filter(val)
})

</script>

<style lang="scss" scoped>
@import "@/style/button.scss";
</style>
