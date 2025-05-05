<template>
  <div class="login-view">
    <h1>Login</h1>
    <button v-if="!loading" @click="loginWithMicrosoft">Login with Microsoft</button>
    <div v-if="loading">Logging in...</div>
    <div v-if="error" class="error">{{ error }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

const loading = ref(false)
const error = ref('')
const router = useRouter()
const route = useRoute()
const { setAuth } = useAuth()

function loginWithMicrosoft() {
  window.location.href = 'http://localhost:8080/auth/login'
}

onMounted(() => {
  // Handle redirect from Microsoft OAuth with token in query
  const token = route.query.token
  const name = route.query.name
  const id = route.query.id
  if (token && name && id) {
    setAuth(token, { id: Number(id), name })
    router.replace('/')
  }
})
</script>

<style scoped>
.error { color: red; margin-top: 1em; }
</style>
