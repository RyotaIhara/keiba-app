<script setup lang="ts">
import type { RaceListItem } from '@/mappers/raceMapper'

defineProps<{ races: RaceListItem[] }>()
const emit = defineEmits<{
  detail: [race: RaceListItem]
  entrants: [race: RaceListItem]
  edit: [race: RaceListItem]
  delete: [race: RaceListItem]
}>()
</script>

<template>
  <div class="overflow-x-auto">
    <table class="min-w-full border-collapse text-left">
      <thead class="bg-gray-700 text-white">
        <tr>
          <th class="px-4 py-2">ID</th>
          <th class="px-4 py-2">開催日</th>
          <th class="px-4 py-2">競馬場</th>
          <th class="px-4 py-2">レース順</th>
          <th class="px-4 py-2">レース名</th>
          <th class="px-4 py-2">発走</th>
          <th class="px-4 py-2">操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="race in races" :key="race.id" class="border-b border-gray-200">
          <td class="px-4 py-2">{{ race.id }}</td><td class="px-4 py-2">{{ race.raceDate }}</td>
          <td class="px-4 py-2">{{ race.raceCourseName }}</td>
          <td class="px-4 py-2">{{ race.raceNumber }}R</td>
          <td>{{ race.raceName }}</td>
          <td class="px-4 py-2">{{ race.startTime }}</td>
          <td class="space-x-3 px-4 py-2">
            <button type="button"
              class="rounded bg-gray-600 px-3 py-1.5 text-white"
              @click="emit('detail', race)">
                詳細
            </button>
            <button type="button"
              class="rounded bg-green-700 px-3 py-1.5 text-white"
              @click="emit('entrants', race)">
                出走一覧
            </button>
            <button type="button"
              class="rounded bg-blue-600 px-3 py-1.5 text-white"
              @click="emit('edit', race)">
                編集
            </button>
            <button type="button"
              class="rounded bg-red-600 px-3 py-1.5 text-white"
              @click="emit('delete', race)">
                削除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
