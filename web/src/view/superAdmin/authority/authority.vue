<template>
  <div class="authority">
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addAuthority(0)">新增角色</el-button>
        <el-button size="small" type="primary" icon="plus" @click="setPower()">新权限配置页面</el-button>
      </div>
      <el-table :data="tableData" :tree-props="{children: 'children', hasChildren: 'hasChildren'}" row-key="id"
        style="width: 100%">
        <el-table-column align="left" label="角色名称" min-width="180" prop="name" />
        <el-table-column label="角色简介" min-width="240" prop="desc" />
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button icon="setting" size="small" type="primary" link @click="opdendrawer(scope.row)">设置权限</el-button>
            <el-button icon="plus" size="small" type="primary" link @click="addAuthority(scope.row.id)">新增子角色
            </el-button>
            <el-button icon="copy-document" size="small" type="primary" link @click="copyAuthorityFunc(scope.row)">拷贝
            </el-button>
            <el-button icon="edit" size="small" type="primary" link @click="editAuthority(scope.row)">编辑</el-button>
            <el-button icon="delete" size="small" type="primary" link @click="deleteAuth(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authorityForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="父级角色" prop="pid">
          <el-cascader v-model="form.pid" style="width:100%" :disabled="dialogType=='add'" :options="AuthorityOption"
            :props="{ checkStrictly: true,label:'name',value:'id',disabled:'disabled',emitPath:false}"
            :show-all-levels="false" filterable />
        </el-form-item>
        <el-form-item label="角色姓名" prop="name">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="角色介绍" prop="desc">
          <el-input v-model="form.desc" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 权限配置弹窗 -->
    <el-drawer v-if="drawer" v-model="drawer" title="权限配置" custom-class="auth-drawer" :with-header="true" size="84.5%">
      <!-- <el-tabs :before-leave="autoEnter" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
      </el-tabs> -->
      <Auth  :row="activeRow" />
    </el-drawer>

    <!-- 新角色配置弹窗 -->
    <el-dialog v-model="dialogFormVisible2" width="100%" title="权限配置">
      <Auth  :row="activeRow" />
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getAuthorityList,
  deleteAuthority,
  createAuthority,
  updateAuthority,
  copyAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus.vue'
import Apis from '@/view/superAdmin/authority/components/apis.vue'
import Auth from '@/view/superAdmin/authority/components/auth.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}

const AuthorityOption = ref([
  {
    id: 0,
    name: '根角色'
  }
])

// 抽屉
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

// 表单
const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

// 权限配置
const setPower = function () {
  dialogFormVisible2.value = true
}
const dialogFormVisible2 = ref(false)
const closeDialog2 = () => {
  dialogFormVisible2.value = false
}

const checkAll = ref(false)
const isIndeterminate = ref(true)
const checkedCities = ref(['字典列表', '新增目录'])
const cities = ['字典列表', '新增目录', '编辑目录', '删除目录']

const handleCheckAllChange = (val) => {
  checkedCities.value = val ? cities : []
  isIndeterminate.value = false
}
const handleCheckedCitiesChange = (value) => {
  const checkedCount = value.length
  checkAll.value = checkedCount === cities.length
  isIndeterminate.value = checkedCount > 0 && checkedCount < cities.length
}
// 权限配置

const form = ref({
  id: 0,
  name: '',
  pid: 0,
  desc: ''
})
const rules = ref({
  id: [
    { required: true, message: '请输入角色ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ],
  name: [
    { required: true, message: '请输入角色名', trigger: 'blur' }
  ],
  pid: [
    { required: true, message: '请选择父角色', trigger: 'blur' },
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// 查询
const getTableData = async () => {
  const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
// 拷贝角色
const copyAuthorityFunc = (row) => {
  setOptions()
  dialogTitle.value = '拷贝角色'
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  dialogFormVisible.value = true
}
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// 删除角色
const deleteAuth = (row) => {
  ElMessageBox.confirm('此操作将永久删除该角色, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteAuthority(row.id)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功!'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '已取消删除'
    })
  })
}
// 初始化表单
const authorityForm = ref(null)
const initForm = () => {
  if (authorityForm.value) {
    authorityForm.value.resetFields()
  }
  form.value = {
    id: 0,
    name: '',
    pid: 0,
    desc: ''
  }
}
// 关闭窗口
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}

// 确定弹窗
const enterDialog = () => {
  form.value.id = Number(form.value.id)
  authorityForm.value.validate(async valid => {
    if (valid) {
      switch (dialogType.value) {
        case 'add':
          {
            const res = await createAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'edit':
          {
            const res = await updateAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '编辑成功!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'copy': {
          const data = {
            authority: {
              id: 0,
              name: '',
              datid: [],
              pid: 0
            },
            oldid: 0
          }
          data.authority.id = form.value.id
          data.authority.name = form.value.name
          data.authority.pid = form.value.pid
          data.authority.dataid = copyForm.value.dataid
          data.oldid = copyForm.value.id
          const res = await copyAuthority(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '复制成功！'
            })
            getTableData()
          }
        }
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      id: 0,
      name: '根角色'
    }
  ]
  setAuthorityOptions(tableData.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  form.value.id = String(form.value.id)
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          id: item.id,
          name: item.name,
          disabled: disabled || item.id === form.value.id,
          children: []
        }
        setAuthorityOptions(
          item.children,
          option.children,
          disabled || item.id === form.value.id
        )
        optionsData.push(option)
      } else {
        const option = {
          id: item.id,
          name: item.name,
          disabled: disabled || item.id === form.value.id
        }
        optionsData.push(option)
      }
    })
}
// 增加角色
const addAuthority = (pid) => {
  initForm()
  dialogTitle.value = '新增角色'
  dialogType.value = 'add'
  form.value.pid = pid
  setOptions()
  dialogFormVisible.value = true
}
// 编辑角色
const editAuthority = (row) => {
  setOptions()
  dialogTitle.value = '编辑角色'
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>

export default {
  name: 'Authority'
}
</script>

<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;

    span {
      display: none;
    }
  }
}

.tree-content {
  overflow: auto;
  height: calc(100vh - 100px);
  margin-top: 10px;
}

.auth-drawer {
  .el-drawer__body {
    overflow: hidden;
  }
}
</style>
