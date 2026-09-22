<script setup lang="ts">
import AppHeader from '@/layout/common/AppHeader.vue'
</script>

<template>
  <div class="min-h-screen bg-gray-50 text-gray-900">
    <AppHeader />
    <main>
      <slot />
    </main>
  </div>
</template>
