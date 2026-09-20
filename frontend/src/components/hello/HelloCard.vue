<script setup lang="ts">
defineProps<{
  title: string
}>()
</script>

<template>
  <div class="rounded-lg border p-4 shadow-sm">
    <h2 class="text-xl font-bold">
      {{ title }}
    </h2>

    <p class="mt-2 text-gray-600">
      これはcomponentから表示しています。
    </p>
  </div>
</template>
