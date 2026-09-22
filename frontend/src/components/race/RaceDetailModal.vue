<script setup lang="ts">
import type { RaceListItem } from '@/mappers/raceMapper'

defineProps<{
  open: boolean
  race: RaceListItem | null
  isLoading?: boolean
  errorMessage?: string
}>()

const emit = defineEmits<{
  close: []
}>()
</script>

<template>
  <div
    v-if="open"
    class="fixed inset-0 z-10 flex items-center justify-center bg-black/50 p-4"
    role="dialog"
    aria-modal="true"
    aria-label="レース詳細"
  >
    <div class="w-full max-w-2xl rounded bg-white p-6 shadow-lg">
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-bold">レース詳細</h2>
        <button
          type="button"
          class="rounded border px-3 py-1"
          @click="emit('close')"
        >
          閉じる
        </button>
      </div>

      <p v-if="isLoading" class="mt-4">詳細を読み込んでいます...</p>
      <p v-else-if="errorMessage" class="mt-4 text-red-600">
        {{ errorMessage }}
      </p>
      <dl v-else-if="race" class="mt-4 grid grid-cols-2">
        <div class="space-y-3 border-r border-gray-300 pr-4">
          <div>
            <dt class="font-semibold text-gray-600">ID</dt>
            <dd>{{ race.id }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">開催日</dt>
            <dd>{{ race.raceDate }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">競馬場</dt>
            <dd>{{ race.raceCourseName }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">レース順</dt>
            <dd>{{ race.raceNumber }}R</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">レース名</dt>
            <dd>{{ race.raceName }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">発走時刻</dt>
            <dd>{{ race.startTime }}</dd>
          </div>
        </div>
        <div class="space-y-3 pl-4">
          <div>
            <dt class="font-semibold text-gray-600">馬場</dt>
            <dd>{{ race.surfaceLabel }}</dd>
          </div>
            <div>
            <dt class="font-semibold text-gray-600">距離</dt>
            <dd>{{ race.distance }}m</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">方向</dt>
            <dd>{{ race.directionLabel }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">天候</dt>
            <dd>{{ race.weatherLabel }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">馬場状態</dt>
            <dd>{{ race.trackConditionLabel }}</dd>
          </div>
          <div>
            <dt class="font-semibold text-gray-600">出走条件</dt>
            <dd>{{ race.raceConditions }}</dd>
          </div>
        </div>
      </dl>
    </div>
  </div>
</template>
