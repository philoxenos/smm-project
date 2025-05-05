<template>
  <div>
    <div v-if="posts.length === 0">No posts yet.</div>
    <PostCard
      v-for="post in posts"
      :key="post.id"
      :post="post"
      :user-id="userId"
      :backend-url="backendUrl"
      @view="$emit('view', $event)"
      @like="$emit('like', $event)"
      @edit="$emit('edit', $event)"
      @delete="$emit('delete', $event)"
    />
  </div>
</template>

<script setup>
import PostCard from './PostCard.vue'

const props = defineProps({
  posts: { type: Array, required: true },
  userId: { type: Number, default: null },
  backendUrl: { type: String, required: true }
})
</script>
