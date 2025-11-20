<template>
  <div class="forbidden-container">
    <div class="forbidden-card">
      <div class="forbidden-icon">
        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="currentColor" stroke-width="2"/>
          <path d="M15 9L9 15" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          <path d="M9 9L15 15" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
        </svg>
      </div>
      
      <div class="forbidden-content">
        <h2>权限不足</h2>
        <p>抱歉，您没有权限访问此页面</p>
        <p class="forbidden-detail">需要管理员权限才能访问此功能</p>
        
        <div class="forbidden-actions">
          <button @click="goBack" class="btn btn-secondary">返回上一页</button>
          <button @click="goHome" class="btn btn-primary">返回首页</button>
          <button @click="logout" class="btn btn-danger">退出登录</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuth } from '../utils/auth.js'

const router = useRouter()
const { logout: authLogout } = useAuth()

const goBack = () => {
  router.go(-1)
}

const goHome = () => {
  router.push('/')
}

const logout = () => {
  authLogout()
  router.push('/login')
}
</script>

<style scoped>
.forbidden-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 20px;
}

.forbidden-card {
  background: white;
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 500px;
  width: 100%;
}

.forbidden-icon {
  color: #ff4757;
  margin-bottom: 1.5rem;
}

.forbidden-content h2 {
  margin: 0 0 1rem 0;
  color: #2d3436;
  font-size: 2rem;
  font-weight: 600;
}

.forbidden-content p {
  margin: 0 0 0.5rem 0;
  color: #636e72;
  font-size: 1.1rem;
}

.forbidden-detail {
  font-size: 0.9rem !important;
  color: #b2bec3 !important;
  margin-bottom: 2rem !important;
}

.forbidden-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  min-width: 120px;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover {
  background-color: #2980b9;
}

.btn-secondary {
  background-color: #95a5a6;
  color: white;
}

.btn-secondary:hover {
  background-color: #7f8c8d;
}

.btn-danger {
  background-color: #e74c3c;
  color: white;
}

.btn-danger:hover {
  background-color: #c0392b;
}

@media (max-width: 600px) {
  .forbidden-card {
    padding: 2rem;
    margin: 1rem;
  }
  
  .forbidden-actions {
    flex-direction: column;
    align-items: center;
  }
  
  .btn {
    width: 100%;
    max-width: 200px;
  }
}
</style>