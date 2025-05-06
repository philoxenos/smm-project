<template>
  <div v-if="!pageLoading" class="login-bg">
    <div class="login-card">
      <div class="login-card-left">
        <div class="psa-header">
          <img src="/psa-logo.png" alt="PSA Logo" class="psa-logo" />
          <span class="psa-text">Social Media<br />Management App</span>
        </div>
        <img src="/img2.svg" alt="Decorative" class="decorative-img" style="display:block;margin:2em auto;max-width:120px;" />
        <h1 class="login-title">Join Today</h1>
        <button class="ms-login-btn" :disabled="loading" @click="loginWithMicrosoft">
          <img src="https://upload.wikimedia.org/wikipedia/commons/4/44/Microsoft_logo.svg" alt="Microsoft Logo" class="ms-logo" />
          <span>Sign in with Microsoft</span>
        </button>
        <div v-if="loading" class="loading-text">Logging in...</div>
        <div v-if="error" class="error">{{ error }}</div>
      </div>
      <div class="login-card-right">
        <img src="/img1.jpg" alt="Login Visual" class="login-img" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

const loading = ref(false)
const error = ref('')
const pageLoading = ref(true)
const router = useRouter()
const route = useRoute()
const { setAuth, user } = useAuth()

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
    return
  }
  // If already logged in, redirect to home
  if (user.value) {
    router.replace('/')
    return
  }
  pageLoading.value = false
})
</script>

<style scoped>
html, body {
  font-family: inherit;
  /* Ensure inheritance from global CSS */
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app, .login-bg, .login-card, .login-card-left, .login-title, .ms-login-btn, .psa-text, .psa-header, .psa-logo, .loading-text, .error, .login-card-right, .login-img {
  font-family: 'Franklin Gothic', Inter, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif !important;
}

.login-bg {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url('/img1.jpg') center center / cover no-repeat fixed;
  position: relative;
  overflow: hidden;
}
.login-bg::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(20, 30, 40, 0.55);
  z-index: 1;
}
.login-card {
  position: relative;
  z-index: 2;
  display: flex;
  background: #fff;
  border-radius: 48px;
  box-shadow: 0 12px 48px 0 rgba(31, 38, 135, 0.25), 0 1.5px 8px 0 rgba(0,0,0,0.10);
  border: 2.5px solid rgba(255,255,255,0.85);
  overflow: hidden;
  max-width: 900px;
  width: 95vw;
  min-height: 480px;
  margin: 2em auto;
  backdrop-filter: blur(8px) saturate(120%);
  -webkit-backdrop-filter: blur(8px) saturate(120%);
}
.login-card-left {
  flex: 0 0 40%;
  max-width: 40%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  padding: 2em 3em 2em 3em;
  background: #fff;
  min-width: 240px;
}
.psa-header {
  display: flex;
  align-items: center;
  margin-bottom: 0.5em;
}
.psa-logo {
  width: 50px;
  height: 50px;
  margin-right: 0.7em;
  margin-bottom: 3em;
}
.psa-text {
  font-size: 1.5em;
  font-weight: bold;
  color: #222;
  line-height: 1.1;
  letter-spacing: 0.5px;
  margin-bottom: 2em;
}
.login-title {
  font-size: 1.2em;
  font-weight: bold;
  margin-top: 0;
  margin-bottom: 1em;
  color: #222;
  letter-spacing: 1px;
  line-height: 1.1;
}
.ms-login-btn {
  display: flex;
  align-items: center;
  background: #fff;
  color: #222;
  font-size: 1.1em;
  font-weight: 600;
  border: 2px solid #0a1368;
  border-radius: 30px;
  padding: 0.5em 1.5em;
  cursor: pointer;
  margin-bottom: 1.5em;
  transition: background 0.2s, box-shadow 0.2s, border 0.2s, color 0.2s;
  box-shadow: 0 2px 8px 0 rgba(0,0,0,0.04);
}
.ms-login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
.ms-login-btn:hover:not(:disabled), .ms-login-btn:focus:not(:disabled) {
  background: #0a1368;
  color: #fff;
  border-color: #0096c7;
  box-shadow: 0 4px 16px 0 rgba(0,180,216,0.15);
}
.ms-login-btn:hover:not(:disabled) span,
.ms-login-btn:focus:not(:disabled) span {
  color: #fff !important;
}
.ms-logo {
  width: 28px;
  height: 28px;
  margin-right: 1em;
  background: #fff;
  border-radius: 4px;
  padding: 2px;
}
.loading-text {
  color: #888;
  margin-bottom: 1em;
}
.error {
  color: #e74c3c;
  margin-top: 1em;
  text-align: left;
}
.login-card-right {
  flex: 0 0 60%;
  max-width: 60%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  min-width: 320px;
  padding: 0;
}
.login-img {
  width: 90%;
  height: 90%;
  display: block;
  object-fit: cover;
  border-radius: 32px;
  box-shadow: 0 2px 8px 0 rgba(0,0,0,0.04);
  background: #fff;
}
@media (max-width: 900px) {
  .login-card {
    flex-direction: column;
    max-width: 98vw;
    min-height: unset;
  }
  .login-card-left, .login-card-right {
    min-width: 0;
    width: 100%;
    border-radius: 0;
    padding: 2em 1em;
    max-width: 100%;
    flex: 1 1 0;
  }
  .login-img {
    border-radius: 0 0 40px 40px;
    max-height: 220px;
    width: 95%;
    height: 95%;
  }
}

.decorative-img {
  display: block;
  width: 100px;
  height: 100px;
  margin: 0 auto;
}
</style>
