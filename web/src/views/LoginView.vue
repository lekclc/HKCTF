<template>
    <el-form
      ref="ruleFormRef"
      style="max-width: 600px"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="auto"
      class="demo-ruleForm"
    >
    <el-form-item label="Name" prop="name">
      <el-input v-model="ruleForm.name" />
    </el-form-item>
      <el-form-item label="Password" prop="pass">
        <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleFormRef)">
          Submit
        </el-button>
        <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
      </el-form-item>
    </el-form>
  </template>
  
  <script lang="ts" setup>
  import { reactive, ref } from 'vue'
  import type { FormInstance, FormRules } from 'element-plus'
import axios from 'axios';

  
  const ruleFormRef = ref<FormInstance>()
  
  
  const ruleForm = reactive({
    name: '',
    password: '',
  })
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    name: [{ required: true, message: 'Please input the username', trigger: 'blur' }],
    password: [{ required: true, message: 'Please input the password', trigger: 'blur' }],
  })
  
  const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
      if (valid) {
        //post发送请求
        const datas = new FormData();
        datas.append('name', ruleForm.name);
        datas.append('password', ruleForm.password);

        axios({
          method: 'post',
          url: 'http://192.168.110.132:8000/login',
          data: datas,
        }).then((res) => {
          const res_ = JSON.stringify(res.data.data)
          const token = res_.split('"')[3]
          localStorage.setItem('token', token)
        }).catch((err) => {
          console.log(err)
          resetForm(formEl)
        })
        console.log('submit!')
      } else {
        console.log('error submit!')
      }
    })
  }
  
  const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
  }
  </script>