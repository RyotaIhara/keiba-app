<script setup lang="ts">
import type { RaceCourseListItem } from '@/mappers/raceCourseMapper'

defineProps<{
  raceCourses: RaceCourseListItem[]
}>()

const emit = defineEmits<{
  edit: [raceCourse: RaceCourseListItem]
  delete: [raceCourse: RaceCourseListItem]
}>()
</script>

<template>
  <div class="overflow-x-auto">
    <table class="min-w-full border-collapse text-left">
      <thead class="bg-gray-700 text-white">
        <tr>
          <th class="px-4 py-2" scope="col">ID</th>
          <th class="px-4 py-2" scope="col">Code</th>
          <th class="px-4 py-2" scope="col">競馬場名</th>
          <th class="px-4 py-2" scope="col">操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="raceCourse in raceCourses" :key="raceCourse.id" class="border-b border-gray-200">
          <td class="px-4 py-2">{{ raceCourse.id }}</td>
          <td class="px-4 py-2">{{ raceCourse.code }}</td>
          <td class="px-4 py-2">{{ raceCourse.name }}</td>
          <td class="space-x-3 px-4 py-2">
            <button type="button" class="rounded bg-blue-600 px-3 py-1.5 text-white" @click="emit('edit', raceCourse)">
              編集
            </button>
            <button type="button" class="rounded bg-red-600 px-3 py-1.5 text-white" @click="emit('delete', raceCourse)">
              削除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
