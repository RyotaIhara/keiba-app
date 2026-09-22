<script setup lang="ts">
defineEmits<{
  reload: []
}>()
</script>

<template>
  <button
    class="ml-2 underline"
    type="button"
    @click="$emit('reload')"
  >
    再読み込み
  </button>
</template>
