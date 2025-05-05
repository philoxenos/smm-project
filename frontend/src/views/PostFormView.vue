<template>
  <div class="post-form-view">
    <h1>{{ isEdit ? 'Edit Post' : 'Create Post' }}</h1>
    <PostForm
      :initial-content="content"
      :initial-image-url="imageUrl"
      :submit-label="isEdit ? 'Update' : 'Create'"
      :error="error"
      @submit="handleSubmit"
      @cancel="goBack"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import PostForm from '@/components/PostForm.vue'

const router = useRouter()
const route = useRoute()
const { token } = useAuth()
const backendUrl = 'http://localhost:8080'

const isEdit = !!route.params.id
const content = ref('')
const imageUrl = ref('')
const error = ref('')

function goBack() {
  router.back()
}

async function fetchPost() {
  if (!isEdit) return
  try {
    const res = await fetch(`${backendUrl}/posts/`)
    const posts = await res.json()
    const post = posts.find(p => p.id == route.params.id)
    if (post) {
      content.value = post.content
      imageUrl.value = post.image_url ? backendUrl + post.image_url : ''
    }
  } catch {
    error.value = 'Failed to load post.'
  }
}

async function handleSubmit({ content: newContent, file }) {
  error.value = ''
  if (!token.value) {
    router.push('/login')
    return
  }
  const formData = new FormData()
  formData.append('content', newContent)
  if (file) formData.append('image', file)
  try {
    let url = `${backendUrl}/posts/`
    let method = 'POST'
    if (isEdit) {
      url = `${backendUrl}/posts/${route.params.id}`
      method = 'PUT'
    }
    const res = await fetch(url, {
      method,
      headers: { Authorization: `Bearer ${token.value}` },
      body: formData,
    })
    if (!res.ok) throw new Error('Failed to save post')
    router.push('/')
  } catch {
    error.value = 'Failed to save post.'
  }
}

onMounted(fetchPost)
</script>

<style scoped>
.error { color: red; margin-top: 1em; }
</style>
