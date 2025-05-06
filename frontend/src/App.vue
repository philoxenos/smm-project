<template>
  <div id="app">
    <nav v-if="isLoggedIn" class="navbar fixed-navbar">
      <div class="navbar-left">
        <img src="/psa-logo.png" alt="PSA Logo" class="psa-logo-navbar" />
        <span class="logo">Social Media Management</span>
        <div class="nav-buttons">
          <router-link to="/">My Feed</router-link>
          <router-link to="/my-posts" v-if="isLoggedIn">My Posts</router-link>
          <router-link to="/create" v-if="isLoggedIn">Create Post</router-link>
        </div>
      </div>
      <div class="navbar-right">
        <div class="avatar-dropdown" @click="toggleDropdown">
          <div class="avatar-circle">{{ userInitials }}</div>
        </div>
        <div v-if="dropdownOpen" class="dropdown-menu">
          <div class="dropdown-item" @click.stop="openProfileModal">Profile Section</div>
          <div class="dropdown-item" @click.stop="openLogoutModal">Logout</div>
        </div>
      </div>
    </nav>
    <div v-if="modalOpen" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <template v-if="modalType === 'profile'">
          <h2>Profile Section</h2>
          <p><b>Name:</b> {{ user?.name }}</p>
          <button class="btn btn-secondary" @click="closeModal">Close</button>
        </template>
        <template v-else-if="modalType === 'logout'">
          <h2>Confirm Logout</h2>
          <p>Are you sure you want to logout?</p>
          <button class="btn btn-danger" @click="confirmLogout">Logout</button>
          <button class="btn btn-secondary" @click="closeModal">Cancel</button>
        </template>
      </div>
    </div>
    <main :class="['main-content', { 'login-bg': !isLoggedIn }]">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useAuth } from '@/composables/useAuth'
import { useRouter } from 'vue-router'

const { user, logout } = useAuth()
const isLoggedIn = computed(() => !!user.value)
const router = useRouter()

const dropdownOpen = ref(false)
const modalOpen = ref(false)
const modalType = ref('')

const userInitials = computed(() => {
  if (!user.value?.name) return ''
  return user.value.name.split(' ').map(n => n[0]).join('').toUpperCase()
})

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}
function openProfileModal() {
  modalType.value = 'profile'
  modalOpen.value = true
  dropdownOpen.value = false
}
function openLogoutModal() {
  modalType.value = 'logout'
  modalOpen.value = true
  dropdownOpen.value = false
}
function closeModal() {
  modalOpen.value = false
}
function confirmLogout() {
  closeModal()
  handleLogout()
}

function handleLogout() {
  logout()
  router.push('/login')
}
</script>

<style scoped>
#app { font-family: Arial, sans-serif; }
.fixed-navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  z-index: 100;
  box-shadow: 0 2px 8px rgba(0,0,0,0.07);
  background: #2c3e50;
  color: #fff;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
}
.navbar-left {
  margin: 2em;
  display: flex;
  align-items: center;
}
.nav-buttons {
  margin: 10em;
  display: flex;
  align-items: center;
}

.psa-logo-navbar {
  width: 38px;
  height: 38px;
  margin-right: 1em;
  vertical-align: middle;
}
.logo {
  font-weight: bold;
  font-size: 1.15em;
  margin-right: 2em;
  letter-spacing: 1px;
  display: inline-block;
  vertical-align: middle;
}
.navbar a {
  color: #fff;
  margin-right: 1.2em;
  text-decoration: none;
  font-weight: 500;
}
.navbar a.router-link-exact-active {
  text-decoration: underline;
}
.navbar-right {
  margin-right: 2em;
  display: flex;
  align-items: center;
  position: relative;
}
.avatar-dropdown {
  cursor: pointer;
  margin-left: 1em;
  position: relative;
}
.avatar-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #e76f6f;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 1.1em;
  box-shadow: 0 2px 8px rgba(0,0,0,0.10);
  border: 2px solid #fff;
}
.dropdown-menu {
  position: absolute;
  right: 0;
  top: 52px;
  background: #fff;
  color: #222;
  border-radius: 10px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.13);
  min-width: 160px;
  z-index: 200;
  padding: 0.5em 0;
}
.dropdown-item {
  padding: 0.7em 1.2em;
  cursor: pointer;
  font-weight: 500;
  transition: background 0.15s;
}
.dropdown-item:hover {
  background: #f2f2f2;
}
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(30, 40, 50, 0.65);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
}
.modal-content {
  background: #fff;
  border-radius: 18px;
  padding: 2em 2.5em;
  min-width: 320px;
  box-shadow: 0 8px 32px rgba(31,38,135,0.18);
  text-align: center;
}
.btn.btn-danger {
  background: #e76f6f;
  color: #fff;
  border: none;
  padding: 0.6em 1.5em;
  border-radius: 8px;
  margin-right: 1em;
  font-weight: bold;
  cursor: pointer;
}
.btn.btn-secondary {
  background: #eee;
  color: #222;
  border: none;
  padding: 0.6em 1.5em;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
}
.main-content {
  flex: 1 1 auto;
  min-height: calc(100vh - 64px);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 80px;
  width: 100vw;
}
.login-bg {
  margin-top: 0;
  min-height: 100vh;
  min-width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
}
</style>
