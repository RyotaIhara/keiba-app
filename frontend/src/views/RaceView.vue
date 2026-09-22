<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import { getRaceCourses } from '@/api/raceCourses'
import {
  createRace,
  deleteRace,
  getRace,
  getRaces,
  updateRace,
  type ApiRace,
  type RaceInput,
  type RaceSearchParams,
} from '@/api/races'
import LoadingMessage from '@/components/common/LoadingMessage.vue'
import ReloadButton from '@/components/common/ReloadButton.vue'
import RaceFormModal from '@/components/race/RaceFormModal.vue'
import RaceDetailModal from '@/components/race/RaceDetailModal.vue'
import RaceList from '@/components/race/RaceList.vue'
import { toRaceCourseListItem, type RaceCourseListItem } from '@/mappers/raceCourseMapper'
import { toRaceInput, toRaceListItem, type RaceListItem } from '@/mappers/raceMapper'

const races = ref<RaceListItem[]>([])
const raceCourses = ref<RaceCourseListItem[]>([])
const searchForm = ref({
  race_date: '',
  race_course_id: '',
})
const appliedSearch = ref<RaceSearchParams>({})
const isLoading = ref(true)
const errorMessage = ref('')
const modalOpen = ref(false)
const modalMode = ref<'create' | 'update'>('create')
const modalRaceId = ref<number | null>(null)
const modalInput = ref<Partial<RaceInput>>({})
const modalError = ref('')
const isLoadingInitial = ref(false)
const isSubmitting = ref(false)
const detailOpen = ref(false)
const detailRace = ref<RaceListItem | null>(null)
const detailError = ref('')
const isLoadingDetail = ref(false)
const router = useRouter()

async function loadRaces() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    races.value = (await getRaces(appliedSearch.value)).map(toRaceListItem)
  } catch (error) {
    races.value = []
    errorMessage.value = error instanceof Error ? error.message : 'レースの取得に失敗しました。'
  } finally {
    isLoading.value = false
  }
}

async function loadRaceCourses() {
  try {
    raceCourses.value = (await getRaceCourses()).map(toRaceCourseListItem)
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : '競馬場の取得に失敗しました。'
  }
}

function buildSearchParams(): RaceSearchParams {
  const params: RaceSearchParams = {}
  if (searchForm.value.race_date) params.race_date = searchForm.value.race_date
  if (searchForm.value.race_course_id) {
    params.race_course_id = Number(searchForm.value.race_course_id)
  }
  return params
}

async function searchRaces() {
  appliedSearch.value = buildSearchParams()
  await loadRaces()
}

async function clearSearch() {
  searchForm.value = { race_date: '', race_course_id: '' }
  appliedSearch.value = {}
  await loadRaces()
}

onMounted(() => {
  void Promise.all([loadRaces(), loadRaceCourses()])
})

function openCreate() {
  modalMode.value = 'create'
  modalRaceId.value = null
  modalInput.value = {
    race_date: new Date().toISOString().slice(0, 10),
    start_time: '12:00:00',
    race_number: 1,
    distance: 1000,
    surface: 'turf',
    direction: 'right',
    weather: 'sunny',
    track_condition: 'firm',
  }
  modalError.value = ''
  isLoadingInitial.value = false
  modalOpen.value = true
}

async function openUpdate(race: RaceListItem) {
  modalMode.value = 'update'
  modalRaceId.value = race.id
  modalInput.value = {}
  modalError.value = ''
  isLoadingInitial.value = true
  modalOpen.value = true
  try {
    modalInput.value = toRaceInput(await getRace(race.id))
  } catch (error) {
    modalError.value = error instanceof Error ? error.message : 'レースの取得に失敗しました。'
  } finally {
    isLoadingInitial.value = false
  }
}

async function openDetail(race: RaceListItem) {
  detailOpen.value = true
  detailRace.value = null
  detailError.value = ''
  isLoadingDetail.value = true
  try {
    const detail: ApiRace = await getRace(race.id)
    detailRace.value = toRaceListItem(detail)
  } catch (error) {
    detailError.value = error instanceof Error ? error.message : 'レースの取得に失敗しました。'
  } finally {
    isLoadingDetail.value = false
  }
}

function closeDetail() {
  detailOpen.value = false
}

function openEntrants(race: RaceListItem) {
  router.push(`/race/${race.id}/details`)
}

function closeModal() {
  if (!isSubmitting.value) modalOpen.value = false
}

async function saveRace(input: RaceInput) {
  modalError.value = ''
  isSubmitting.value = true
  try {
    if (modalMode.value === 'create') await createRace(input)
    else if (modalRaceId.value !== null) await updateRace(modalRaceId.value, input)
    modalOpen.value = false
    await loadRaces()
  } catch (error) {
    modalError.value = error instanceof Error ? error.message : 'レースの保存に失敗しました。'
  } finally {
    isSubmitting.value = false
  }
}

async function removeRace(race: RaceListItem) {
  if (!window.confirm(`「${race.raceName}」を削除しますか？`)) return
  errorMessage.value = ''
  try {
    await deleteRace(race.id)
    await loadRaces()
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : 'レースの削除に失敗しました。'
  }
}
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold">レースリスト</h1>
    <div class="mt-4 rounded-lg border border-gray-200 bg-gray-50 p-4">
      <form class="grid max-w-xl gap-3" @submit.prevent="searchRaces">
        <label class="flex flex-col gap-1 text-sm text-gray-700">
          開催日
          <input
            v-model="searchForm.race_date"
            type="date"
            class="rounded border border-gray-300 bg-white px-3 py-2"
          />
        </label>
        <label class="flex flex-col gap-1 text-sm text-gray-700">
          競馬場
          <select
            v-model="searchForm.race_course_id"
            class="rounded border border-gray-300 bg-white px-3 py-2"
          >
            <option value="">すべて</option>
            <option v-for="raceCourse in raceCourses" :key="raceCourse.id" :value="raceCourse.id">
              {{ raceCourse.name }}
            </option>
          </select>
        </label>
        <div class="flex items-end gap-2">
          <button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white">検索</button>
          <button
            type="button"
            class="rounded border border-gray-300 bg-white px-4 py-2 text-gray-700"
            @click="clearSearch"
          >
            クリア
          </button>
        </div>
      </form>
    </div>
    <button class="mt-4 rounded bg-green-600 px-4 py-2 text-white" @click="openCreate">
      レースを作成
    </button>
    <LoadingMessage v-if="isLoading" />
    <p v-else-if="errorMessage" class="mt-4 text-red-600">
      {{ errorMessage }} <ReloadButton @reload="loadRaces" />
    </p>
    <p v-else-if="races.length === 0" class="mt-4">レースが見つかりませんでした。</p>
    <RaceList
      v-if="!isLoading && !errorMessage && races.length > 0"
      class="mt-4"
      :races="races"
      @detail="openDetail"
      @entrants="openEntrants"
      @edit="openUpdate"
      @delete="removeRace"
    />
    <RaceFormModal
      :open="modalOpen"
      :mode="modalMode"
      :initial-input="modalInput"
      :is-loading-initial="isLoadingInitial"
      :is-submitting="isSubmitting"
      :error-message="modalError"
      @cancel="closeModal"
      @submit="saveRace"
    />
    <RaceDetailModal
      :open="detailOpen"
      :race="detailRace"
      :is-loading="isLoadingDetail"
      :error-message="detailError"
      @close="closeDetail"
      @entrants="detailRace && openEntrants(detailRace)"
    />
  </div>
</template>
