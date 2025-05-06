<template>
  <div class="my-posts-view">
    <div class="my-posts-header">
      <h1>My Posts</h1>
      <input v-model="search" placeholder="Search posts..." class="search-input" />
    </div>
    <div v-if="loading" class="center-text">Loading posts...</div>
    <div v-if="error" class="error center-text">{{ error }}</div>
    <PostList
      v-if="!loading && filteredPosts.length > 0"
      :posts="filteredPosts"
      :user-id="user?.id"
      :backend-url="backendUrl"
      @view="goToDetail"
      @like="likePost"
      @edit="goToEdit"
      @delete="deletePost"
    />
    <div v-if="!loading && filteredPosts.length === 0" class="center-text">No posts found.</div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import PostList from '@/components/PostList.vue'

const posts = ref([])
const loading = ref(false)
const error = ref('')
const search = ref('')
const router = useRouter()
const { token, user } = useAuth()
const backendUrl = 'http://localhost:8080'

function goToDetail(id) {
  router.push(`/post/${id}`)
}
function goToEdit(id) {
  router.push(`/edit/${id}`)
}

async function fetchPosts() {
  if (!user.value) return
  loading.value = true
  error.value = ''
  try {
    const res = await fetch(`${backendUrl}/posts/?user_id=${user.value.id}`)
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

const filteredPosts = computed(() => {
  if (!search.value) return posts.value
  return posts.value.filter(post =>
    post.content.toLowerCase().includes(search.value.toLowerCase())
  )
})

onMounted(fetchPosts)
</script>

<style scoped>
.my-posts-view {
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
.my-posts-header {
  width: 100%;
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 2em;
  flex-wrap: wrap;
}
.my-posts-header h1 {
  font-size: 2em;
  margin: 0 1em 0 0;
}
.search-input {
  padding: 0.5em 1em;
  border-radius: 6px;
  border: 1px solid #ccc;
  font-size: 1em;
  min-width: 200px;
  margin-top: 0.5em;
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
  .my-posts-view {
    padding: 1.2em 0.5em;
    max-width: 98vw;
  }
}
@media (max-width: 600px) {
  .my-posts-header {
    flex-direction: column;
    align-items: stretch;
  }
  .search-input {
    margin-top: 1em;
    width: 100%;
    min-width: 0;
  }
}
</style>
