<template>
  <div id="app">
    <nav class="navbar">
      <router-link to="/">Home</router-link>
      <router-link to="/create" v-if="isLoggedIn">Create Post</router-link>
      <span class="spacer"></span>
      <span v-if="isLoggedIn">
        <span class="user">{{ user?.name }}</span>
        <button @click="logout">Logout</button>
      </span>
      <router-link to="/login" v-else>Login</router-link>
    </nav>
    <router-view />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useAuth } from '@/composables/useAuth'

const { user, logout } = useAuth()
const isLoggedIn = computed(() => !!user.value)
</script>

<style scoped>
#app { font-family: Arial, sans-serif; }
.navbar {
  display: flex;
  align-items: center;
  background: #2c3e50;
  color: #fff;
  padding: 1em;
  margin-bottom: 2em;
}
.navbar a {
  color: #fff;
  margin-right: 1em;
  text-decoration: none;
}
.navbar .spacer {
  flex: 1;
}
.user {
  margin-right: 1em;
  font-weight: bold;
}
button {
  background: #42b983;
  color: #fff;
  border: none;
  padding: 0.5em 1em;
  border-radius: 4px;
  cursor: pointer;
}
button:hover {
  background: #369870;
}
</style>
