<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addMenu(0)">新增根菜单</el-button>
      </div>
      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="id">
        <!-- <el-table-column align="left" label="id" min-width="100" prop="id" /> -->
        <el-table-column align="left" label="展示名称" min-width="150" prop="authorityName">
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="图标" min-width="180" prop="authorityName">
          <template #default="scope">
            <div v-if="scope.row.meta.icon" class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="路由Name" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" label="路由Path" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" label="是否隐藏" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden?"隐藏":"显示" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="父节点" min-width="90" prop="parentId" />
        <el-table-column align="left" label="排序" min-width="70" prop="sort" />
        <el-table-column align="left" label="文件路径" min-width="360" prop="component" />
        <el-table-column align="left" fixed="right" label="操作" width="300">
          <template #default="scope">
            <el-button size="small" type="primary" link icon="plus" @click="addMenu(scope.row.id)">添加子菜单</el-button>
            <el-button size="small" type="primary" link icon="edit" @click="editMenu(scope.row.id)">编辑</el-button>
            <el-button size="small" type="primary" link icon="delete" @click="deleteMenu(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle">
      <warning-bar title="新增菜单，需要在角色管理内配置权限才可使用" />
      <el-form v-if="dialogFormVisible" ref="menuForm" :inline="true" :model="form" :rules="rules" label-position="top"
        label-width="85px">
        <el-form-item label="图标" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta" style="width:100%" />
        </el-form-item>
        <el-form-item label="展示名称" prop="meta.title" style="width:30%">
          <el-input v-model="form.meta.title" autocomplete="off" placeholder="中文名" />
        </el-form-item>
        <el-form-item label="是否普通页面" style="width:30%">
          <el-select v-model="form.hidden" placeholder="是否在列表隐藏">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="父节点ID" style="width:30%">
          <el-cascader v-model="form.parentId" style="width:100%" :disabled="!isEdit" :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'id',disabled:'disabled',emitPath:false}"
            :show-all-levels="false" filterable />
        </el-form-item>
        <el-form-item label="路由Name" prop="path" style="width:30%">
          <el-input v-model="form.name" autocomplete="off" placeholder="唯一英文字符串" @change="changeName" />
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <template #label>
            <div style="display:inline-flex">
              路由Path
              <el-checkbox v-model="checkFlag" style="float:right;margin-left:20px;">自定义Path</el-checkbox>
            </div>
          </template>
          <el-input v-model="form.path" :disabled="!checkFlag" autocomplete="off" placeholder="建议只在后方拼接参数" />
        </el-form-item>
        <el-form-item prop="component" style="width:63%">
          <template #label>
            <div style="display:inline-flex">
              文件路径
              <span @click="form.component = 'view/routerHolder.vue'" style="margin-left:30px;">点我设置默认路径</span>
            </div>
          </template>
          <el-input v-model="form.component" autocomplete="off" placeholder="页面:view/xxx/xx.vue 插件:plugin/xx/xx.vue"
            @blur="fmtComponent" />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort" style="width:30%">
          <el-input v-model.number="form.sort" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div>
        <el-button size="small" type="primary" icon="edit" @click="addApi(form)">新增API权限</el-button>
        <el-table :data="form.api" style="width: 100%">
          <el-table-column align="left" prop="type" label="请求类型" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.method" placeholder="请选择">
                <el-option key="GET" value="GET" label="GET" />
                <el-option key="POST" value="POST" label="POST" />
                <el-option key="PUT" value="PUT" label="PUT" />
                <el-option key="PATCH" value="PATCH" label="PATCH" />
                <el-option key="DELETE" value="DELETE" label="DELETE" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="route" label="路由地址" width="340">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.route" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="路由名称">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button type="danger" size="small" icon="delete" @click="deApi(form.api,scope.$index)">删除
                </el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <el-button style="margin-top:12px" size="small" type="primary" icon="edit" @click="addExtAuth(form)">
          新增拓展权限
        </el-button>
        <el-table :data="form.extAuth" style="width: 100%">
          <el-table-column align="left" prop="type" label="请求类型" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="请选择">
                <el-option key="按钮权限" value="1" label="按钮权限" />
                <el-option key="数据权限" value="2" label="数据权限" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="权限中文名" width="250">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="权限英文名" width="240">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.val" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button type="danger" size="small" icon="delete" @click="delExtAuth(form.extAuth,scope.$index)">删除
                </el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
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
import {
  updateMenu,
  getMenuList,
  createMenu,
  delMenu,
  getBaseMenuById
} from '@/api/menu'
import { deleteApi, deleteExtAuth } from '@/api/api'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

// 验证规则
const rules = reactive({
  path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
  component: [
    { required: true, message: '请输入文件路径', trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// 查询列表
const getTableData = async () => {
  const table = await getMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 初始化表格数据
getTableData()

// 新增API配置
const addApi = (form) => {
  if (!form.api) {
    form.api = []
  }
  form.api.push({
    id: 0,
    method: 'GET',
    route: '',
    name: '',
  })
}
// 删除API配置
const deApi = async (api, index) => {
  let apiId = api[index].id
  if (apiId === 0) {                 // 刚加的直接删
    api.splice(index, 1)
    return
  }
  const res = await deleteApi(apiId) // 库里的先删库里的
  if (res.code === 0) {
    api.splice(index, 1)
    return
  }
}

// 新增菜单拓展[按钮+数据]权限
const addExtAuth = (form) => {
  if (!form.extAuth) {
    form.extAuth = []
  }
  form.extAuth.push({
    id: 0,
    type: '1',
    name: '',
    val: '',
  })
}
// 删除菜单数据权限
const delExtAuth = async (expands, index) => {
  let extAuthId = expands[index].id
  if (extAuthId === 0) {                  // 刚加的直接删
    expands.splice(index, 1)
    return
  }
  const res = await deleteExtAuth(extAuthId) // 库里的先删库里的
  if (res.code === 0) {
    expands.splice(index, 1)
    return
  }
}

// 表单数据
const form = ref({
  id: 0,
  path: '',
  name: '',
  hidden: false,
  parentId: 0,
  component: 'view/routerHolder.vue',
  meta: {
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  sort: 50,
  api: [],
  extAuth: []
})

// 修改path名字
const changeName = () => {
  form.value.path = form.value.name
}

// 格式化文件路径
const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// 关闭操作
const handleClose = (done) => {
  initForm()
  done()
}

// 删除菜单
const deleteMenu = (id) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await delMenu(id)
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
  }).catch((err) => {
    ElMessage({
      type: 'info',
      message: '已取消删除'
    })
  })
}

// 初始化弹窗内表格方法
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    id: 0,
    path: "",
    name: "",
    hidden: false,
    parentId: 0,
    component: "view/routerHolder.vue",
    sort: 50,
    meta: {
      title: "",
      icon: "",
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    },
    api: [],
    extAuth: []
  }
}

// 关闭弹窗
const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

// 添加menu
const enterDialog = async () => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (isEdit.value) {
        form.value.parentId = Number(form.value.parentId)
        res = await updateMenu(form.value)
      } else {
        res = await createMenu(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '编辑成功' : '添加成功!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    id: 0,
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      id: "0",
      title: '根目录'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData && menuData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        title: item.meta.title,
        id: String(item.id),
        disabled: disabled || item.id === form.value.id,
        children: []
      }
      setMenuOptions(
        item.children,
        option.children,
        disabled || item.id === form.value.id
      )
      optionsData.push(option)
    } else {
      const option = {
        title: item.meta.title,
        id: String(item.id),
        disabled: disabled || item.id === form.value.id
      }
      optionsData.push(option)
    }
  })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('新增菜单')
const addMenu = (id) => {
  dialogTitle.value = '新增菜单'
  form.value.parentId = id
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async (id) => {
  dialogTitle.value = '编辑菜单'
  const res = await getBaseMenuById(id)
  form.value = res.data
  form.value.parentId = String(form.value.parentId)

  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}

.icon-column {
  display: flex;
  align-items: center;

  .el-icon {
    margin-right: 8px;
  }
}
</style>
