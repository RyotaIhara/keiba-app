<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import { getRaceCourses, type ApiRaceCourse } from '@/api/raceCourses'
import type { RaceInput } from '@/api/races'

const props = withDefaults(defineProps<{
  open: boolean
  mode: 'create' | 'update'
  initialInput?: Partial<RaceInput>
  isLoadingInitial?: boolean
  isSubmitting?: boolean
  errorMessage?: string
}>(), { initialInput: () => ({}), isLoadingInitial: false, isSubmitting: false, errorMessage: '' })

const emit = defineEmits<{ cancel: []; submit: [input: RaceInput] }>()
const input = ref<Partial<RaceInput>>({})
const raceCourses = ref<ApiRaceCourse[]>([])
const isLoadingCourses = ref(false)
const courseError = ref('')
const title = computed(() => props.mode === 'create' ? 'レースを作成' : 'レースを編集')
const options = {
  surface: [{ value: 'turf', label: '芝' }, { value: 'dirt', label: 'ダート' }],
  direction: [{ value: 'right', label: '右回り' }, { value: 'left', label: '左回り' }],
  weather: [{ value: 'sunny', label: '晴れ' }, { value: 'cloudy', label: '曇り' }, { value: 'rainy', label: '雨' }, { value: 'snowy', label: '雪' }],
  trackCondition: [{ value: 'firm', label: '良' }, { value: 'good', label: '稍良' }, { value: 'yield', label: '重' }, { value: 'soft', label: '不良' }],
} as const

async function loadRaceCourses() {
  isLoadingCourses.value = true
  courseError.value = ''
  try { raceCourses.value = await getRaceCourses() }
  catch (error) { raceCourses.value = []; courseError.value = error instanceof Error ? error.message : '競馬場の取得に失敗しました。' }
  finally { isLoadingCourses.value = false }
}

watch(() => [props.open, props.mode, props.initialInput], () => {
  if (props.open) { input.value = { ...props.initialInput }; void loadRaceCourses() }
}, { immediate: true })

function submit() {
  emit('submit', input.value as RaceInput)
}
</script>

<template>
  <div v-if="open" class="fixed inset-0 z-10 flex items-center justify-center overflow-y-auto bg-black/50 p-4" role="dialog" aria-modal="true" :aria-label="title">
    <form class="w-full max-w-2xl rounded bg-white p-6 shadow-lg" @submit.prevent="submit">
      <h2 class="text-xl font-bold">{{ title }}</h2>
      <p v-if="errorMessage" class="mt-3 text-red-600">{{ errorMessage }}</p>
      <p v-if="isLoadingInitial" class="mt-3">詳細を読み込んでいます...</p>
      <div class="grid gap-3 sm:grid-cols-2">
        <label class="mt-4 block"><span class="mb-1 block">開催日</span><input v-model="input.race_date" required type="date" :disabled="isLoadingInitial" class="w-full rounded border px-3 py-2" /></label>
        <label class="mt-4 block"><span class="mb-1 block">競馬場</span>
          <select v-model.number="input.race_course_id" required :disabled="isLoadingInitial || isLoadingCourses || !!courseError" class="w-full rounded border px-3 py-2">
            <option :value="undefined" disabled>競馬場を選択</option><option v-for="course in raceCourses" :key="course.ID" :value="course.ID">{{ course.Name }}</option>
          </select>
          <span v-if="isLoadingCourses" class="text-sm text-gray-600">競馬場を読み込んでいます...</span>
          <span v-if="courseError" class="text-sm text-red-600">{{ courseError }}</span>
        </label>
        <label class="block"><span class="mb-1 block">レース番号</span><input v-model.number="input.race_number" required min="1" type="number" class="w-full rounded border px-3 py-2" /></label>
        <label class="block"><span class="mb-1 block">レース名</span><input v-model="input.race_name" required class="w-full rounded border px-3 py-2" /></label>
        <label class="block"><span class="mb-1 block">発走時刻</span><input v-model="input.start_time" required type="time" step="1" class="w-full rounded border px-3 py-2" /></label>
        <label class="block"><span class="mb-1 block">距離（m）</span><input v-model.number="input.distance" required min="1" type="number" class="w-full rounded border px-3 py-2" /></label>
        <label class="block"><span class="mb-1 block">馬場</span><select v-model="input.surface" required class="w-full rounded border px-3 py-2"><option v-for="item in options.surface" :key="item.value" :value="item.value">{{ item.label }}</option></select></label>
        <label class="block"><span class="mb-1 block">方向</span><select v-model="input.direction" required class="w-full rounded border px-3 py-2"><option v-for="item in options.direction" :key="item.value" :value="item.value">{{ item.label }}</option></select></label>
        <label class="block"><span class="mb-1 block">天候</span><select v-model="input.weather" required class="w-full rounded border px-3 py-2"><option v-for="item in options.weather" :key="item.value" :value="item.value">{{ item.label }}</option></select></label>
        <label class="block"><span class="mb-1 block">馬場状態</span><select v-model="input.track_condition" required class="w-full rounded border px-3 py-2"><option v-for="item in options.trackCondition" :key="item.value" :value="item.value">{{ item.label }}</option></select></label>
        <label class="block sm:col-span-2"><span class="mb-1 block">出走条件</span><input v-model="input.race_conditions" required class="w-full rounded border px-3 py-2" /></label>
      </div>
      <div class="mt-6 flex justify-end gap-3"><button type="button" class="rounded border px-4 py-2" :disabled="isSubmitting || isLoadingInitial" @click="emit('cancel')">キャンセル</button><button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white disabled:opacity-50" :disabled="isSubmitting || isLoadingInitial || isLoadingCourses || !!courseError">{{ isSubmitting ? '保存中...' : '保存' }}</button></div>
    </form>
  </div>
</template>
