<script setup lang="ts">
import type { UserListItem } from '@/mappers/userMapper'

defineProps<{
  users: UserListItem[]
}>()

const emit = defineEmits<{
  edit: [user: UserListItem]
  delete: [user: UserListItem]
}>()
</script>

<template>
  <div class="overflow-x-auto">
    <table class="min-w-full border-collapse text-left">
      <thead class="bg-gray-700 text-white">
        <tr>
          <th class="px-4 py-2" scope="col">ID</th>
          <th class="px-4 py-2" scope="col">Code</th>
          <th class="px-4 py-2" scope="col">ユーザー名</th>
          <th class="px-4 py-2" scope="col">操作</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="user in users"
          :key="user.id"
          class="border-b border-gray-200"
        >
          <td class="px-4 py-2">{{ user.id }}</td>
          <td class="px-4 py-2">{{ user.code }}</td>
          <td class="px-4 py-2">{{ user.name }}</td>
          <td class="space-x-3 px-4 py-2">
            <button
              type="button"
              class="rounded bg-blue-600 px-3 py-1.5 text-white"
              @click="emit('edit', user)"
            >
              編集
            </button>
            <button
              type="button"
              class="rounded bg-red-600 px-3 py-1.5 text-white"
              @click="emit('delete', user)"
            >
              削除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
