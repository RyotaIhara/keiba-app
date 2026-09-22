<script setup lang="ts">
const navigationItems = [
  { label: 'ホーム', to: '/' },
  { label: 'ユーザー', to: '/user' },
  { label: '競馬場', to: '/race-course' },
  { label: 'レース', to: '/race' },
]
</script>

<template>
  <header class="border-b bg-white shadow-sm">
    <div class="mx-auto flex max-w-7xl flex-wrap items-center gap-4 px-4 py-4 sm:px-6 lg:px-8">
      <RouterLink to="/" class="text-xl font-bold text-green-700">
        競馬管理
      </RouterLink>

      <nav aria-label="メインナビゲーション" class="flex flex-wrap gap-2">
        <RouterLink
          v-for="item in navigationItems"
          :key="item.to"
          :to="item.to"
          class="rounded px-3 py-2 text-sm font-medium text-gray-600 transition hover:bg-green-50 hover:text-green-700"
          active-class="bg-green-100 text-green-800"
          exact-active-class="bg-green-100 text-green-800"
        >
          {{ item.label }}
        </RouterLink>
      </nav>
    </div>
  </header>
</template>
