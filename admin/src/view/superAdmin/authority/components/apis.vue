<template>
  <div>
    <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button class="fl-right" size="small" type="primary" @click="authApiEnter">确 定</el-button>
    </div>
    <div class="tree-content">
      <el-tree ref="apiTree" :data="apiTreeData" :default-checked-keys="apiTreeIds" :props="apiDefaultProps"
        default-expand-all highlight-current node-key="id" show-checkbox :filter-node-method="filterNode"
        @check="nodeChange" />
    </div>
  </div>
</template>
<script>
export default {
  name: 'Apis',
}
</script>

<script setup>
import { addApiAuthority } from '@/api/authority'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
const props = defineProps({
  row: {
    default: function () {
      return {}
    },
    type: Object
  },
  apiTreeData: {
    default: function () {
      return {}
    },
    type: Object
  },
  apiTreeIds: {
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
const apiTreeData = ref([])
const apiTreeIds = ref([])
const init = async () => {
  // 回显API树
  apiTreeData.value = props.apiTreeData
  // 回显API选中
  apiTreeIds.value = props.apiTreeIds
}
init()

const needConfirm = ref(false)
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authApiEnter()
}

// 关联关系确定
const apiTree = ref(null)
const authApiEnter = async () => {
  let checkArr = apiTree.value.getCheckedKeys(true, false)
  let apiIds = checkArr.map(item => String(item))
  const res = await addApiAuthority({
    data: apiIds,
    id: props.row.id
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'api设置成功' })
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
