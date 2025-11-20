<template>
  <div class="login-page">
    <div class="login-container">
      <!-- 背景装饰 -->
      <div class="login-background">
        <div class="background-shapes">
          <div class="shape shape-1"></div>
          <div class="shape shape-2"></div>
          <div class="shape shape-3"></div>
        </div>
      </div>

      <!-- 登录表单 -->
      <div class="login-form">
        <div class="login-header">
          <div class="logo">
            <div class="logo-icon">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" stroke="currentColor" stroke-width="2"/>
                <path d="M9 12l2 2 4-4" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </div>
            <h1 class="logo-text">IFaceConf</h1>
          </div>
          <p class="login-subtitle">接口配置化框架管理后台</p>
        </div>

        <form @submit.prevent="handleLogin" class="form">
          <div class="form-group">
            <label for="username" class="form-label">用户名</label>
            <div class="input-wrapper">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
                <circle cx="12" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
              </svg>
              <input
                id="username"
                v-model="form.username"
                type="text"
                class="form-input"
                placeholder="请输入用户名"
                required
                :disabled="loading"
              />
            </div>
          </div>

          <div class="form-group">
            <label for="password" class="form-label">密码</label>
            <div class="input-wrapper">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2" stroke="currentColor" stroke-width="2"/>
                <path d="M7 11V7a5 5 0 0110 0v4" stroke="currentColor" stroke-width="2"/>
              </svg>
              <input
                id="password"
                v-model="form.password"
                type="password"
                class="form-input"
                placeholder="请输入密码"
                required
                :disabled="loading"
              />
            </div>
          </div>

          <div class="form-options">
            <label class="checkbox-label">
              <input type="checkbox" v-model="rememberMe" />
              <span class="checkmark"></span>
              记住我
            </label>
            <a href="#" class="forgot-link">忘记密码？</a>
          </div>

          <button type="submit" class="login-btn" :disabled="loading || !formValid">
            <span v-if="loading" class="loading-spinner"></span>
            <span v-else>登录</span>
          </button>

          <div v-if="errorMessage" class="error-message">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
              <path d="M12 8v4" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <path d="M12 16h.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            {{ errorMessage }}
          </div>
        </form>

        <div class="login-footer">
          <p>© 2024 IFaceConf. 保留所有权利</p>
          <p>版本 v1.0.0</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../utils/auth.js'

const router = useRouter()
const { login } = useAuth()

const form = ref({
  username: '',
  password: ''
})

const rememberMe = ref(false)
const loading = ref(false)
const errorMessage = ref('')

const formValid = computed(() => {
  return form.value.username.trim() && form.value.password.trim()
})

const handleLogin = async () => {
  if (!formValid.value) return

  loading.value = true
  errorMessage.value = ''

  try {
    const result = await login({ username: form.value.username, password: form.value.password })
    
    if (result && result.success) {
      // 登录成功，跳转到首页
      router.push('/')
    } else {
      errorMessage.value = result.error || '登录失败'
    }
  } catch (error) {
    errorMessage.value = error.message || '登录失败，请检查网络连接'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(ellipse at top, var(--bg-secondary) 0%, var(--bg-primary) 100%);
  position: relative;
  overflow: hidden;
}

.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 1200px;
  padding: var(--space-xl);
}

.login-background {
  flex: 1;
  position: relative;
  height: 600px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.background-shapes {
  position: relative;
  width: 400px;
  height: 400px;
}

.shape {
  position: absolute;
  border-radius: 50%;
  background: var(--gradient-primary);
  opacity: 0.1;
  animation: float 6s ease-in-out infinite;
}

.shape-1 {
  width: 200px;
  height: 200px;
  top: 0;
  left: 0;
  animation-delay: 0s;
}

.shape-2 {
  width: 150px;
  height: 150px;
  top: 100px;
  right: 50px;
  animation-delay: 2s;
}

.shape-3 {
  width: 100px;
  height: 100px;
  bottom: 50px;
  left: 150px;
  animation-delay: 4s;
}

@keyframes float {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(180deg); }
}

.login-form {
  background: var(--bg-surface);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-xl);
  padding: var(--space-3xl);
  width: 100%;
  max-width: 400px;
  box-shadow: var(--shadow-xl);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 10;
}

.login-header {
  text-align: center;
  margin-bottom: var(--space-2xl);
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-md);
  margin-bottom: var(--space-sm);
}

.logo-icon {
  width: 48px;
  height: 48px;
  background: var(--gradient-primary);
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-icon svg {
  color: white;
}

.logo-text {
  font-size: 2rem;
  font-weight: 700;
  background: var(--gradient-primary);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.login-subtitle {
  color: var(--text-secondary);
  font-size: 0.875rem;
}

.form {
  display: flex;
  flex-direction: column;
  gap: var(--space-lg);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
}

.form-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-primary);
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-wrapper svg {
  position: absolute;
  left: var(--space-md);
  color: var(--text-muted);
  z-index: 1;
}

.form-input {
  width: 100%;
  padding: var(--space-md) var(--space-md) var(--space-md) calc(var(--space-md) * 2 + 20px);
  background: var(--bg-card);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  color: var(--text-primary);
  font-size: 0.875rem;
  transition: all var(--transition-fast);
}

.form-input:focus {
  outline: none;
  border-color: var(--primary-500);
  box-shadow: 0 0 0 3px rgba(76, 187, 76, 0.1);
}

.form-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.form-options {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 0.875rem;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  display: none;
}

.checkmark {
  width: 16px;
  height: 16px;
  border: 1px solid var(--border-medium);
  border-radius: var(--radius-sm);
  position: relative;
  transition: all var(--transition-fast);
}

.checkbox-label input[type="checkbox"]:checked + .checkmark {
  background: var(--primary-500);
  border-color: var(--primary-500);
}

.checkbox-label input[type="checkbox"]:checked + .checkmark::after {
  content: '';
  position: absolute;
  left: 4px;
  top: 1px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.forgot-link {
  color: var(--primary-500);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.forgot-link:hover {
  color: var(--primary-400);
}

.login-btn {
  width: 100%;
  padding: var(--space-md);
  background: var(--gradient-primary);
  border: none;
  border-radius: var(--radius-lg);
  color: white;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all var(--transition-fast);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-sm);
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.error-message {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  padding: var(--space-md);
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: var(--radius-md);
  color: var(--error);
  font-size: 0.875rem;
}

.login-footer {
  margin-top: var(--space-2xl);
  text-align: center;
  font-size: 0.75rem;
  color: var(--text-muted);
}

.login-footer p {
  margin-bottom: var(--space-xs);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-container {
    flex-direction: column;
    padding: var(--space-md);
  }
  
  .login-background {
    display: none;
  }
  
  .login-form {
    padding: var(--space-xl);
    max-width: 100%;
  }
}
</style>