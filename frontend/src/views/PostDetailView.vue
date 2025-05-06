<template>
  <div class="post-detail-view">
    <div v-if="loading" class="center-text">Loading post...</div>
    <div v-if="error" class="error center-text">{{ error }}</div>
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
    <div class="actions-row" v-if="!loading && post">
      <button class="btn btn-secondary" @click="goBack">Back</button>
    </div>
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
.post-detail-view {
  min-height: 70vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 700px;
  margin: 3em auto 0 auto;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 4px 24px rgba(0,0,0,0.10);
  padding: 2.5em 2em 2em 2em;
}
.center-text {
  text-align: center;
  margin: 2em 0;
  width: 100%;
}
.actions-row {
  display: flex;
  justify-content: flex-end;
  margin-top: 1.5em;
  width: 100%;
}
.error {
  color: #e74c3c;
  margin-top: 1em;
  text-align: center;
}
@media (max-width: 800px) {
  .post-detail-view {
    padding: 1.2em 0.5em;
    max-width: 98vw;
  }
}
</style>
