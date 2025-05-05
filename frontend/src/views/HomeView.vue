<template>
  <div class="home-view">
    <h1>Social Feed</h1>
    <button @click="goToCreate">Create Post</button>
    <div v-if="loading">Loading posts...</div>
    <div v-if="error" class="error">{{ error }}</div>
    <PostList
      v-if="!loading && posts.length > 0"
      :posts="posts"
      :user-id="user?.id"
      :backend-url="backendUrl"
      @view="goToDetail"
      @like="likePost"
      @edit="goToEdit"
      @delete="deletePost"
    />
    <div v-if="!loading && posts.length === 0">No posts yet.</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import PostList from '@/components/PostList.vue'

const posts = ref([])
const loading = ref(false)
const error = ref('')
const router = useRouter()
const { token, user } = useAuth()
const backendUrl = 'http://localhost:8080'

function formatDate(ts) {
  if (!ts) return ''
  const d = new Date(ts * 1000)
  return d.toLocaleString()
}

function isOwner(post) {
  return user.value && post.user_id === user.value.id
}

function goToDetail(id) {
  router.push(`/post/${id}`)
}
function goToEdit(id) {
  router.push(`/edit/${id}`)
}
function goToCreate() {
  router.push('/create')
}

async function fetchPosts() {
  loading.value = true
  error.value = ''
  try {
    const res = await fetch(`${backendUrl}/posts/`)
    posts.value = await res.json()
  } catch (e) {
    error.value = 'Failed to load posts.'
  } finally {
    loading.value = false
  }
}

async function likePost(id) {
  if (!token.value) {
    router.push('/login')
    return
  }
  try {
    await fetch(`${backendUrl}/posts/${id}/like`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${token.value}` },
    })
    fetchPosts()
  } catch (e) {
    error.value = 'Failed to like post.'
  }
}

async function deletePost(id) {
  if (!token.value) {
    router.push('/login')
    return
  }
  if (!confirm('Delete this post?')) return
  try {
    await fetch(`${backendUrl}/posts/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${token.value}` },
    })
    fetchPosts()
  } catch (e) {
    error.value = 'Failed to delete post.'
  }
}

onMounted(fetchPosts)
</script>

<style scoped>
.home-view { max-width: 600px; margin: 0 auto; }
.post-card { border: 1px solid #ccc; border-radius: 8px; margin: 1em 0; padding: 1em; }
.post-header { display: flex; justify-content: space-between; margin-bottom: 0.5em; }
.post-content { margin-bottom: 0.5em; }
.post-image { max-width: 100%; max-height: 300px; margin-bottom: 0.5em; }
.post-actions button { margin-right: 0.5em; }
.error { color: red; margin-top: 1em; }
.date { color: #888; font-size: 0.9em; }
</style>
