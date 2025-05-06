<template>
  <div class="home-view">
    <div class="home-header">
      <h1>My Feed</h1>
    </div>
    <div v-if="loading" class="center-text">Loading posts...</div>
    <div v-if="error" class="error center-text">{{ error }}</div>
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
    <div v-if="!loading && posts.length === 0" class="center-text">No posts yet.</div>
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
.home-view {
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
.home-header {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 2em;
}
.center-text {
  text-align: center;
  margin: 2em 0;
  width: 100%;
}
.error {
  color: #e74c3c;
  margin-top: 1em;
  text-align: center;
}
@media (max-width: 800px) {
  .home-view {
    padding: 1.2em 0.5em;
    max-width: 98vw;
  }
}
</style>
