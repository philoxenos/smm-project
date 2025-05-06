<template>
  <form @submit.prevent="handleSubmit" enctype="multipart/form-data">
    <div>
      <label>Content</label><br />
      <textarea v-model="content" required rows="4" cols="40"></textarea>
    </div>
    <div>
      <label>Image (optional)</label><br />
      <ImageUploader v-model="file" :initial-url="initialImageUrl" @update:modelValue="handleImageChange" />
    </div>
    <div style="margin-top:1em;">
      <button type="submit">{{ submitLabel }}</button>
      <button type="button" @click="$emit('cancel')">Cancel</button>
    </div>
    <div v-if="error" class="error">{{ error }}</div>
  </form>
</template>

<script setup>
import { ref, watch } from 'vue'
import ImageUploader from './ImageUploader.vue'

const props = defineProps({
  initialContent: { type: String, default: '' },
  initialImageUrl: { type: String, default: '' },
  submitLabel: { type: String, default: 'Submit' },
  error: { type: String, default: '' }
})
const emit = defineEmits(['submit', 'cancel'])

const content = ref(props.initialContent)
const file = ref(null)
const imageRemoved = ref(false)

watch(() => props.initialContent, val => { content.value = val })
watch(() => props.initialImageUrl, val => {
  if (val) imageRemoved.value = false
})

function handleImageChange(newFile) {
  file.value = newFile
  if (!newFile && props.initialImageUrl) {
    imageRemoved.value = true
  } else if (newFile) {
    imageRemoved.value = false
  }
}

function handleSubmit() {
  emit('submit', { content: content.value, file: file.value, imageRemoved: imageRemoved.value })
}
</script>

<style scoped>
.error { color: red; margin-top: 1em; }
</style>
