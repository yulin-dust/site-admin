
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="站点ID:" prop="siteId">
    <el-select v-model="formData.siteId" placeholder="请选择站点ID" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.siteId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="今日发布数:" prop="todayCount">
    <el-input v-model.number="formData.todayCount" :clearable="true" placeholder="请输入今日发布数" />
</el-form-item>
        <el-form-item label="计划发布数:" prop="planCount">
    <el-input v-model.number="formData.planCount" :clearable="true" placeholder="请输入计划发布数" />
</el-form-item>
        <el-form-item label="总发布数:" prop="totalCount">
    <el-input v-model.number="formData.totalCount" :clearable="true" placeholder="请输入总发布数" />
</el-form-item>
        <el-form-item label="最后发布时间:" prop="lastPublishTime">
    <el-date-picker v-model="formData.lastPublishTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="昨日发布数:" prop="yesterdayCount">
    <el-input v-model.number="formData.yesterdayCount" :clearable="true" placeholder="请输入昨日发布数" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getPublishStatsDataSource,
  createPublishStats,
  updatePublishStats,
  findPublishStats
} from '@/api/sites/publishStats'

defineOptions({
    name: 'PublishStatsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            siteId: undefined,
            todayCount: undefined,
            planCount: undefined,
            totalCount: undefined,
            lastPublishTime: new Date(),
            yesterdayCount: undefined,
        })
// 验证规则
const rule = reactive({
               siteId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getPublishStatsDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPublishStats({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createPublishStats(formData.value)
               break
             case 'update':
               res = await updatePublishStats(formData.value)
               break
             default:
               res = await createPublishStats(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
