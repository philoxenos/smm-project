<template>
  <div class="post-detail-view">
    <div v-if="loading">Loading post...</div>
    <div v-if="error" class="error">{{ error }}</div>
    <PostCard
      v-if="post && !loading"
      :post="post"
      :user-id="user?.id"
      :backend-url="backendUrl"
      @like="likePost"
      @edit="goToEdit"
      @delete="deletePost"
      @view="goBack"
    />
    <button v-if="!loading && post" @click="goBack">Back</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import PostCard from '@/components/PostCard.vue'

const route = useRoute()
const router = useRouter()
const { token, user } = useAuth()
const backendUrl = 'http://localhost:8080'

const post = ref(null)
const loading = ref(false)
const error = ref('')

function goToEdit(id) {
  router.push(`/edit/${id}`)
}
function goBack() {
  router.back()
}

async function fetchPost() {
  loading.value = true
  error.value = ''
  try {
    const res = await fetch(`${backendUrl}/posts/`)
    const posts = await res.json()
    const found = posts.find(p => p.id == route.params.id)
    if (found) {
      post.value = found
    } else {
      error.value = 'Post not found.'
    }
  } catch {
    error.value = 'Failed to load post.'
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
    fetchPost()
  } catch {
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
    router.push('/')
  } catch {
    error.value = 'Failed to delete post.'
  }
}

onMounted(fetchPost)
</script>

<style scoped>
.post-detail-view { max-width: 600px; margin: 0 auto; }
.error { color: red; margin-top: 1em; }
</style>
