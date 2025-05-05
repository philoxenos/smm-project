<template>
  <div>
    <input type="file" @change="onFileChange" accept="image/*" />
    <div v-if="previewUrl">
      <img :src="previewUrl" alt="Preview" style="max-width:200px;max-height:200px;margin-top:8px;" />
      <button type="button" @click="clearImage">Remove</button>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: File,
  initialUrl: { type: String, default: '' }
})
const emit = defineEmits(['update:modelValue'])

const previewUrl = ref(props.initialUrl)

function onFileChange(e) {
  const f = e.target.files[0]
  emit('update:modelValue', f)
  if (f) {
    previewUrl.value = URL.createObjectURL(f)
  } else {
    previewUrl.value = ''
  }
}

function clearImage() {
  emit('update:modelValue', null)
  previewUrl.value = ''
}

watch(() => props.initialUrl, (val) => {
  if (val) previewUrl.value = val
})
</script>
