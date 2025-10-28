<template>
  <div class="app-container">
    <Sidebar @open-login="openLoginModal" />
    <div class="main-content">
      <ChatMain />
      <ChatInput />
    </div>

    <!-- 登录弹窗 -->
    <LoginModal v-model="showLoginModal" @login-success="handleLoginSuccess" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from '@/components/layout/Sidebar.vue'
import ChatMain from '@/components/layout/ChatMain.vue'
import ChatInput from '@/components/layout/ChatInput.vue'
import LoginModal from '@/components/common/LoginModal.vue'

// 响应式数据
const showLoginModal = ref(false)
const currentUser = ref(null)

// 方法
const openLoginModal = () => {
  showLoginModal.value = true
}

const handleLoginSuccess = (userInfo) => {
  currentUser.value = userInfo
  console.log('登录成功:', userInfo)
}
</script>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background: #fff;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: #fff;
  position: relative;
}
</style>
