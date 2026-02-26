<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <span>用户登录</span>
        </div>
      </template>
      <el-form :model="loginForm" :rules="rules" ref="loginFormRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="loginForm.password" placeholder="请输入密码" show-password></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin(loginFormRef)">登录</el-button>
          <el-button @click="goToRegister">注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import userApi from '@/api/user.js';

const router = useRouter();
const loginFormRef = ref(null);

const loginForm = reactive({
  username: '',
  password: '',
});

const rules = reactive({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
});

const handleLogin = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      try {
        const res = await userApi.login(loginForm);
        if (res.code === 0) {
          ElMessage.success('登录成功');
          // 保存token等信息
          localStorage.setItem('token', res.data.token || '');
          window.location.href = '/products';
        } else {
          ElMessage.error(res.msg || '登录失败');
        }
      } catch (e) {
        ElMessage.error('登录请求失败');
      }
    } else {
      ElMessage.error('请输入用户名和密码');
      return false;
    }
  });
};

const goToRegister = () => {
  router.push('/register');
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f7fa;
}

.login-card {
  width: 450px;
}

.card-header {
  text-align: center;
  font-size: 20px;
}
</style> 