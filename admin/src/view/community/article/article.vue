<template>
  <div>
    <!-- <warning-bar title="获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释" /> -->
    <!-- 搜索条件 -->
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="字典名（中）">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="字典名（英）">
          <el-input v-model="searchInfo.value" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="超管权限" prop="status">
          <el-select v-model="searchInfo.super" clear placeholder="请选择">
            <el-option key="true" label="是" value="1" />
            <el-option key="false" label="否" value="2" />
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="描述">
            <el-input v-model="searchInfo.desc" placeholder="搜索条件" />
          </el-form-item> -->
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- 数据列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="字典名（中）" prop="name" width="160" />
        <el-table-column align="left" label="字典名（英）" prop="value" width="120" />
        <el-table-column align="left" label="超管权限" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.super) }}</template>
        </el-table-column>
        <el-table-column align="left" label="描述" prop="desc" width="280" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button size="small" icon="document" type="primary" link @click="toDetail(scope.row)">详情</el-button>
            <el-button size="small" icon="edit" type="primary" link @click="updateDictionaryFunc(scope.row)">变更
            </el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="deleteDictionaryFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" size="small" style="margin-left: 10px"
                  @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <!-- 新增编辑 -->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="dialogForm" :model="formData" :rules="rules" :inline="true" size="default" label-width="110px">
        <el-form-item label="所属栏目" prop="categoryList" style="width:25%">
          <el-select v-model="formData.categoryList" placeholder="请选择栏目">
            <el-option v-for="item in categoryList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="文章类型" style="width:25%">
          <el-select v-model="formData.type" placeholder="是否在列表隐藏">
            <el-option :value="1" label="图文" />
            <el-option :value="2" label="视频" />
          </el-select>
        </el-form-item>
        <el-form-item label="文章标记" prop="markList" style="width:40%">
          <el-select v-model="formData.markList" multiple collapse-tags placeholder="请选择栏目">
            <el-option v-for="item in markList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="文章标题" prop="title" style="width:100%">
          <el-input v-model="formData.title" placeholder="请输入标题" clearable />
        </el-form-item>
        <el-form-item label="文章摘要" prop="summary" style="width:100%">
          <el-input v-model="formData.summary" type="textarea" placeholder="请输入文章摘要" />
        </el-form-item>
        <el-form-item label="文章内容" prop="content" style="width:100%">
          <el-input v-model="formData.content" type="textarea" :rows="3" placeholder="请输入文章内容" />
        </el-form-item>
        <el-form-item label="图片地址" prop="picList" style="width:100%">
          <el-input v-model="formData.picList" placeholder="图片地址" />
        </el-form-item>
        <el-form-item label="视频地址" prop="videoList" style="width:100%">
          <el-input v-model="formData.videoList" placeholder="视频地址" />
        </el-form-item>
        <el-form-item label="文章话题" prop="topicList" style="width:100%">
          <el-select v-model="formData.topicList" multiple placeholder="请选择话题" style="width: 750px">
            <el-option v-for="item in topicList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="文章标签" prop="tagList" style="width:100%">
          <el-select v-model="formData.tagList" multiple placeholder="请选择标签" style="width: 750px">
            <el-option v-for="item in tagList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
  
<script>
export default {
  name: 'Article',
}
</script>
  
<script setup>
import {
  createArticle,
  // deleteDictionary,
  // updateDictionary,
  // findDictionary,
  // getDictionaryList,
} from '@/api/article' // 此处请自行替换地址
import { getDict } from '@/utils/dictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'

// 路由
const router = useRouter()

// 初始化选项
var categoryList = ref([])
var markList = ref([])
var topicList = ref([])
var tagList = ref([])
const initPage = async () => {
  getDict('category').then((res) => { categoryList = res });
  getDict('mark').then((res) => { markList = res });
  getDict('topic').then((res) => { topicList = res });
  getDict('tag').then((res) => { tagList = res });
}
initPage()

// 表单数据
const formData = ref({
  title: "",
  summary: "",
  content: "",
  type: 1,
  picList: "",
  videoList: "",
  categoryList: "",
  markList: "",
  topicList: [],
  tagList: []
})
const rules = ref({
  // name: [
  //   {
  //     required: true,
  //     message: '请输入字典名（中）',
  //     trigger: 'blur',
  //   },
  // ],
  // value: [
  //   {
  //     required: true,
  //     message: '请输入字典名（英）',
  //     trigger: 'blur',
  //   },
  // ],
  // desc: [
  //   {
  //     required: true,
  //     message: '请输入描述',
  //     trigger: 'blur',
  //   },
  // ],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}

// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getDictionaryList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// getTableData()

// 打开详情页
const toDetail = (row) => {
  router.push({
    name: 'dictionaryDetail',
    params: {
      id: row.id,
    },
  })
}

const dialogFormVisible = ref(false)
const type = ref('')
const updateDictionaryFunc = async (row) => {
  const res = await findDictionary(row.id)
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    pid: 0,
    super: 0,
    name: "",
    value: "",
    sort: 50,
    desc: "",
  }
}
const deleteDictionaryFunc = async (row) => {
  row.visible = false
  const res = await deleteDictionary(row.id)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async () => {
  dialogForm.value.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createArticle(formData.value)
        break
      case 'update':
        res = await updateDictionary(formData.value)
        break
      default:
        res = await createArticle(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}
</script>
  
<style>

</style>
  