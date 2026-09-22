<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { getRace, type ApiRace } from '@/api/races'
import { getRaceDetails, type ApiRaceDetail } from '@/api/raceDetails'
import LoadingMessage from '@/components/common/LoadingMessage.vue'
import ReloadButton from '@/components/common/ReloadButton.vue'
import { toRaceListItem, type RaceListItem } from '@/mappers/raceMapper'

const route = useRoute()
const router = useRouter()
const race = ref<RaceListItem | null>(null)
const details = ref<ApiRaceDetail[]>([])
const isLoading = ref(true)
const errorMessage = ref('')

function raceId(): number | null {
  const value = Number(route.params.raceId)
  return Number.isInteger(value) && value > 0 ? value : null
}

async function loadRaceDetails() {
  const id = raceId()
  if (id === null) {
    errorMessage.value = 'レースIDが不正です。'
    isLoading.value = false
    return
  }

  isLoading.value = true
  errorMessage.value = ''
  try {
    const [raceResponse, detailResponse] = await Promise.all([
      getRace(id),
      getRaceDetails(id),
    ])
    race.value = toRaceListItem(raceResponse)
    details.value = detailResponse
  } catch (error) {
    race.value = null
    details.value = []
    errorMessage.value =
      error instanceof Error ? error.message : '出走馬の取得に失敗しました。'
  } finally {
    isLoading.value = false
  }
}

function goBack() {
  router.push('/race')
}

onMounted(loadRaceDetails)
</script>

<template>
  <div class="p-8">
    <button type="button" class="rounded border px-4 py-2" @click="goBack">
      レース一覧に戻る
    </button>

    <LoadingMessage v-if="isLoading" />
    <p v-else-if="errorMessage" class="mt-4 text-red-600">
      {{ errorMessage }}
      <ReloadButton @reload="loadRaceDetails" />
    </p>
    <template v-else-if="race">
      <div class="mt-6">
        <h1 class="text-2xl font-bold">
          {{ race.raceName }}（{{ race.raceNumber }}R）出走一覧
        </h1>
        <p class="mt-2 text-gray-600">
          {{ race.raceDate }} / {{ race.raceCourseName }} / {{ race.startTime }}
        </p>
      </div>

      <p v-if="details.length === 0" class="mt-6">
        出走馬が見つかりませんでした。
      </p>
      <div v-else class="mt-6 overflow-x-auto">
        <table class="min-w-full border-collapse text-left">
          <thead class="bg-gray-700 text-white">
            <tr>
              <th class="px-4 py-2">枠番</th>
              <th class="px-4 py-2">馬番</th>
              <th class="px-4 py-2">馬名</th>
              <th class="px-4 py-2">性齢</th>
              <th class="px-4 py-2">斤量</th>
              <th class="px-4 py-2">騎手</th>
              <th class="px-4 py-2">厩舎</th>
              <th class="px-4 py-2">馬体重</th>
              <th class="px-4 py-2">オッズ</th>
              <th class="px-4 py-2">人気</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="detail in details"
              :key="detail.ID"
              class="border-b border-gray-200"
            >
              <td class="px-4 py-2">{{ detail.FrameNumber }}</td>
              <td class="px-4 py-2">{{ detail.HorseNumber }}</td>
              <td class="px-4 py-2 font-semibold">{{ detail.HorseName }}</td>
              <td class="px-4 py-2">
                {{ detail.Sex === 'colt' ? '牡' : '牝' }}{{ detail.Age }}
              </td>
              <td class="px-4 py-2">{{ detail.Weight }}kg</td>
              <td class="px-4 py-2">{{ detail.Jockey }}</td>
              <td class="px-4 py-2">{{ detail.Stable }}</td>
              <td class="px-4 py-2">
                {{ detail.BodyWeight }}kg
                <span
                  :class="detail.BodyWeightChange > 0 ? 'text-red-600' : detail.BodyWeightChange < 0 ? 'text-blue-600' : 'text-gray-600'"
                >
                  ({{ detail.BodyWeightChange > 0 ? '+' : '' }}{{ detail.BodyWeightChange }})
                </span>
              </td>
              <td class="px-4 py-2">{{ detail.Odds.toFixed(1) }}</td>
              <td class="px-4 py-2">{{ detail.Popularity }}番人気</td>
            </tr>
          </tbody>
        </table>
      </div>
    </template>
  </div>
</template>
