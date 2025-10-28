<template>
  <div class="sidebar">
    <!-- 顶部Logo区域 -->
    <div class="sidebar-header">
      <div class="logo-container">
        <div class="logo-image">
          <router-link to="/" class="logo-link">
            <a-image :preview="false" src="/xiaozhe-ai-logo.svg" alt="小哲AI Logo" />
          </router-link>
        </div>
      </div>
      <div class="header-actions">
        <a-button type="text" class="action-btn" @click="toggleSidebar">
          <template #icon>
            <IconMenuFold v-if="isCollapsed" />
            <IconMenuUnfold v-else />
          </template>
        </a-button>
      </div>
    </div>

    <!-- 导航菜单 -->
    <div class="nav-menu">
      <div class="nav-item active">
        <div class="nav-icon">
          <IconPlus />
        </div>
        <span class="nav-text">新建会话</span>
        <div class="nav-shortcut">Ctrl K</div>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconRefresh />
        </div>
        <span class="nav-text">小哲+</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon ppt-icon">
          <IconFile />
        </div>
        <span class="nav-text">PPT 助手</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconClockCircle />
        </div>
        <span class="nav-text">历史会话</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconApps />
        </div>
        <span class="nav-text">查看全部</span>
      </div>
    </div>

    <!-- 底部用户信息 -->
    <div class="user-section">
      <div class="user-info" @click="handleLoginClick">
        <div class="user-profile">
          <a-avatar class="user-avatar" :style="{ backgroundColor: '#6b7280' }">
            <IconUser />
          </a-avatar>
          <span class="username">登录</span>
        </div>
        <div class="dropdown-action">
          <IconDown class="dropdown-icon" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  IconApps,
  IconClockCircle,
  IconDown,
  IconFile,
  IconMenuFold,
  IconMenuUnfold,
  IconPlus,
  IconRefresh,
  IconUser
} from '@arco-design/web-vue/es/icon'

// 定义事件
const emit = defineEmits(['open-login'])

// 定义响应式变量
const isCollapsed = ref(false)

// 切换侧边栏折叠状态
const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
}

// 处理登录点击
const handleLoginClick = () => {
  emit('open-login')
}
</script>

<style scoped>
.sidebar {
  width: 280px;
  background: #2c2c2c;
  display: flex;
  flex-direction: column;
  height: 100vh;
  color: #fff;
}

.sidebar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #404040;
}

.logo-container {
  display: flex;
  align-items: center;
}


.header-actions {
  display: flex;
  align-items: center;
}

.action-btn {
  color: #fff;
  border: none;
  background: transparent;
}

.action-btn:hover {
  background: #404040;
}

.nav-menu {
  flex: 1;
  padding: 16px 0;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  margin: 4px 0;
  cursor: pointer;
  border-radius: 8px;
  margin: 4px 16px;
  position: relative;
}

.nav-item:hover {
  background: #404040;
}

.nav-item.active {
  background: #404040;
}

.nav-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.ppt-icon {
  color: #ff4d4f;
}

.nav-text {
  flex: 1;
  font-size: 14px;
  color: #fff;
}

.nav-shortcut {
  background: #404040;
  color: #ccc;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.user-section {
  padding: 16px 20px;
  border-top: 1px solid #404040;
}

.user-info {
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  padding: 10px 12px;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.user-info:hover {
  background: #404040;
}

.user-profile {
  display: flex;
  align-items: center;
  flex: 1;
}

.user-avatar {
  margin-right: 12px;
  width: 32px;
  height: 32px;
}

.username {
  font-size: 14px;
  color: #e5e7eb;
  font-weight: 500;
}

.dropdown-action {
  display: flex;
  align-items: center;
  padding: 4px;
  border-radius: 4px;
  transition: background 0.2s ease;
}

.dropdown-action:hover {
  background: #4b5563;
}

.dropdown-icon {
  color: #9ca3af;
  font-size: 14px;
  transition: transform 0.2s ease;
}

.user-info:hover .dropdown-icon {
  transform: rotate(180deg);
}
</style>