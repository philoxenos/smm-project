<template>
  <button @click="onLike" :disabled="loading">
    Like ({{ likes }})
  </button>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  likes: { type: Number, required: true },
  onLike: { type: Function, required: true },
  loading: { type: Boolean, default: false }
})
</script>
