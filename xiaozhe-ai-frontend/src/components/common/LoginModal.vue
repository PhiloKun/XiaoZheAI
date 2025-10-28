<template>
    <a-modal v-model:visible="visible" title="" :footer="false" :width="400" :mask-closable="true" @cancel="handleClose"
        class="login-modal">
        <div class="login-container">
            <!-- 标题 -->
            <div class="login-header">
                <h2 class="login-title">{{ isLogin ? '登录' : '注册' }}</h2>
                <p class="login-subtitle">{{ isLogin ? '欢迎回来' : '创建新账户' }}</p>
            </div>

            <!-- 表单 -->
            <div class="login-form">
                <a-form :model="formData" layout="vertical">
                    <!-- 手机号 -->
                    <a-form-item label="手机号">
                        <a-input v-model="formData.phone" placeholder="请输入手机号" size="large">
                            <template #prefix>
                                <IconPhone />
                            </template>
                        </a-input>
                    </a-form-item>

                    <!-- 验证码 -->
                    <a-form-item label="验证码">
                        <a-input v-model="formData.verifyCode" placeholder="请输入验证码" size="large">
                            <template #prefix>
                                <IconSafe />
                            </template>
                            <template #suffix>
                                <a-button type="text" size="small" :disabled="!canSendCode || codeCountdown > 0"
                                    @click="sendVerifyCode" class="code-btn">
                                    {{ codeCountdown > 0 ? `${codeCountdown}s后重发` : '发送验证码' }}
                                </a-button>
                            </template>
                        </a-input>
                    </a-form-item>



                    <!-- 登录按钮 -->
                    <a-form-item style="margin-bottom: 12px;">
                        <a-button type="primary" size="large" long :loading="loading" @click="handleSubmit"
                            class="login-btn">
                            {{ isLogin ? '登录' : '注册' }}
                        </a-button>
                    </a-form-item>

                    <!-- 切换登录/注册 -->
                    <div class="login-switch">
                        <span class="switch-text">
                            {{ isLogin ? '还没有账户？' : '已有账户？' }}
                        </span>
                        <a-button type="text" @click="toggleMode" class="switch-btn">
                            {{ isLogin ? '立即注册' : '立即登录' }}
                        </a-button>
                    </div>
                </a-form>
            </div>

            <!-- 第三方登录 -->
            <!-- <div class="social-login">
                <a-divider>
                    <span class="divider-text">其他方式登录</span>
                </a-divider>
                <div class="social-buttons">
                    <a-button type="outline" size="large" class="social-btn">
                        <template #icon>
                            <IconGithub />
                        </template>
                        GitHub
                    </a-button>
                    <a-button type="outline" size="large" class="social-btn">
                        <template #icon>
                            <IconGoogle />
                        </template>
                        Google
                    </a-button>
                </div>
            </div> -->
        </div>
    </a-modal>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
    IconPhone,
    IconSafe
} from '@arco-design/web-vue/es/icon'

// Props
const props = defineProps({
    modelValue: {
        type: Boolean,
        default: false
    }
})

// Emits
const emit = defineEmits(['update:modelValue', 'login-success'])

// 响应式数据
const isLogin = ref(true)
const loading = ref(false)
const codeCountdown = ref(0) // 验证码倒计时
const formData = ref({
    phone: '',
    verifyCode: ''
})

// 计算属性
const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
})

// 是否可以发送验证码
const canSendCode = computed(() => {
    return formData.value.phone && /^1[3-9]\d{9}$/.test(formData.value.phone)
})

// 方法
const handleClose = () => {
    visible.value = false
    resetForm()
}

const resetForm = () => {
    formData.value = {
        phone: '',
        verifyCode: ''
    }
    isLogin.value = true
    codeCountdown.value = 0
}

const toggleMode = () => {
    isLogin.value = !isLogin.value
    formData.value.verifyCode = ''
}

// 发送验证码
const sendVerifyCode = async () => {
    if (!canSendCode.value) return

    try {
        // 模拟发送验证码 API
        await new Promise(resolve => setTimeout(resolve, 500))

        // 开始倒计时
        codeCountdown.value = 60
        const timer = setInterval(() => {
            codeCountdown.value--
            if (codeCountdown.value <= 0) {
                clearInterval(timer)
            }
        }, 1000)

        console.log('验证码已发送到:', formData.value.phone)
    } catch (error) {
        console.error('发送验证码失败:', error)
    }
}

const handleSubmit = async () => {
    // 验证表单
    if (!formData.value.phone || !formData.value.verifyCode) {
        return
    }

    loading.value = true

    try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 1000))

        // 登录成功
        emit('login-success', {
            phone: formData.value.phone,
            loginType: 'phone',
            isLogin: isLogin.value
        })

        handleClose()
    } catch (error) {
        console.error('登录失败:', error)
    } finally {
        loading.value = false
    }
}
</script>

<style scoped>
:deep(.arco-modal) {
    border-radius: 12px;
}

:deep(.arco-modal-header) {
    display: none;
}

:deep(.arco-modal-body) {
    padding: 32px;
}

.login-container {
    width: 100%;
}

.login-header {
    text-align: center;
    margin-bottom: 32px;
}

.login-title {
    font-size: 24px;
    font-weight: 600;
    color: #1d2129;
    margin: 0 0 8px 0;
}

.login-subtitle {
    font-size: 14px;
    color: #86909c;
    margin: 0;
}

.login-form {
    margin-bottom: 24px;
}

.code-btn {
    font-size: 12px;
    color: #165dff;
    white-space: nowrap;
}

.code-btn:disabled {
    color: #c9cdd4;
}

:deep(.arco-form-item-label) {
    font-weight: 500;
    color: #1d2129;
}

:deep(.arco-input-wrapper) {
    border-radius: 8px;
}

:deep(.arco-input-prefix) {
    color: #86909c;
}

.login-btn {
    height: 44px;
    border-radius: 8px;
    font-weight: 500;
    font-size: 16px;
}

.login-switch {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 4px;
}

.switch-text {
    font-size: 14px;
    color: #86909c;
}

.switch-btn {
    padding: 0;
    height: auto;
    font-size: 14px;
    color: #165dff;
}

.switch-btn:hover {
    color: #0e42d2;
}

.social-login {
    margin-top: 24px;
}

.divider-text {
    font-size: 12px;
    color: #c9cdd4;
}

.social-buttons {
    display: flex;
    gap: 12px;
    margin-top: 16px;
}

.social-btn {
    flex: 1;
    height: 44px;
    border-radius: 8px;
    font-weight: 500;
}

:deep(.arco-divider-text) {
    background: #fff;
    padding: 0 16px;
}
</style>