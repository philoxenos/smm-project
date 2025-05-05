<template>
  <div class="post-card">
    <div class="post-header">
      <strong>{{ post.user?.name || 'Unknown' }}</strong>
      <span class="date">{{ formatDate(post.created_at) }}</span>
    </div>
    <div class="post-content">{{ post.content }}</div>
    <img v-if="post.image_url" :src="backendUrl + post.image_url" class="post-image" />
    <div class="post-actions">
      <button @click="$emit('view', post.id)">View</button>
      <LikeButton :likes="post.likes" :onLike="() => $emit('like', post.id)" />
      <button v-if="isOwner" @click="$emit('edit', post.id)">Edit</button>
      <button v-if="isOwner" @click="$emit('delete', post.id)">Delete</button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import LikeButton from './LikeButton.vue'

const props = defineProps({
  post: { type: Object, required: true },
  userId: { type: Number, default: null },
  backendUrl: { type: String, required: true }
})

const isOwner = computed(() => props.userId && props.post.user_id === props.userId)

function formatDate(ts) {
  if (!ts) return ''
  const d = new Date(ts * 1000)
  return d.toLocaleString()
}
</script>

<style scoped>
.post-card { border: 1px solid #ccc; border-radius: 8px; margin: 1em 0; padding: 1em; }
.post-header { display: flex; justify-content: space-between; margin-bottom: 0.5em; }
.post-content { margin-bottom: 0.5em; }
.post-image { max-width: 100%; max-height: 300px; margin-bottom: 0.5em; }
.post-actions button { margin-right: 0.5em; }
.date { color: #888; font-size: 0.9em; }
</style>
