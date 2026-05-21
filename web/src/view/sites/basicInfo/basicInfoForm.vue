
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="域名:" prop="domain">
    <el-input v-model="formData.domain" :clearable="true" placeholder="请输入域名" />
</el-form-item>
        <el-form-item label="链接:" prop="url">
    <el-input v-model="formData.url" :clearable="true" placeholder="请输入链接" />
</el-form-item>
        <el-form-item label="token:" prop="apiToken">
    <el-input v-model="formData.apiToken" :clearable="true" placeholder="请输入token" />
</el-form-item>
        <el-form-item label="cms:" prop="cmsType">
    <el-input v-model="formData.cmsType" :clearable="true" placeholder="请输入cms" />
</el-form-item>
        <el-form-item label="站点类型:" prop="siteType">
    <el-input v-model="formData.siteType" :clearable="true" placeholder="请输入站点类型" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-tree-select v-model="formData.status" placeholder="请选择状态" :data="statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="分类:" prop="categoryIds">
    <el-input v-model="formData.categoryIds" :clearable="true" placeholder="请输入分类" />
</el-form-item>
        <el-form-item label="作者:" prop="authorIds">
    <el-input v-model="formData.authorIds" :clearable="true" placeholder="请输入作者" />
</el-form-item>
        <el-form-item label="ip:" prop="iP">
    <el-input v-model="formData.iP" :clearable="true" placeholder="请输入ip" />
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
  createBasicInfo,
  updateBasicInfo,
  findBasicInfo
} from '@/api/sites/basicInfo'

defineOptions({
    name: 'BasicInfoForm'
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
const statusOptions = ref([])
const formData = ref({
            domain: '',
            url: '',
            apiToken: '',
            cmsType: '',
            siteType: '',
            status: '',
            categoryIds: '',
            authorIds: '',
            iP: '',
        })
// 验证规则
const rule = reactive({
               domain : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               siteType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBasicInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
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
               res = await createBasicInfo(formData.value)
               break
             case 'update':
               res = await updateBasicInfo(formData.value)
               break
             default:
               res = await createBasicInfo(formData.value)
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
