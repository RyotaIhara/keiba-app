<script setup lang="ts">
import { onMounted, ref } from 'vue'

import {
  createRaceCourse,
  deleteRaceCourse,
  getRaceCourse,
  getRaceCourses,
  updateRaceCourse,
} from '@/api/raceCourses'
import LoadingMessage from '@/components/common/LoadingMessage.vue'
import ReloadButton from '@/components/common/ReloadButton.vue'
import RaceCourseFormModal from '@/components/race-course/RaceCourseFormModal.vue'
import RaceCourseList from '@/components/race-course/RaceCourseList.vue'
import {
  toRaceCourseListItem,
  type RaceCourseListItem,
} from '@/mappers/raceCourseMapper'

const raceCourses = ref<RaceCourseListItem[]>([])
const isLoading = ref(true)
const errorMessage = ref('')
const modalOpen = ref(false)
const modalMode = ref<'create' | 'update'>('create')
const modalRaceCourseId = ref<number | null>(null)
const modalCode = ref('')
const modalName = ref('')
const modalError = ref('')
const isLoadingInitial = ref(false)
const isSubmitting = ref(false)

async function loadRaceCourses() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    raceCourses.value = (await getRaceCourses()).map(toRaceCourseListItem)
  } catch (error) {
    raceCourses.value = []
    errorMessage.value =
      error instanceof Error ? error.message : '競馬場の取得に失敗しました。'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadRaceCourses)

function openCreate() {
  modalMode.value = 'create'
  modalRaceCourseId.value = null
  modalCode.value = ''
  modalName.value = ''
  modalError.value = ''
  isLoadingInitial.value = false
  modalOpen.value = true
}

async function openUpdate(raceCourse: RaceCourseListItem) {
  modalMode.value = 'update'
  modalRaceCourseId.value = raceCourse.id
  modalCode.value = ''
  modalName.value = ''
  modalError.value = ''
  isLoadingInitial.value = true
  modalOpen.value = true
  try {
    const detail = await getRaceCourse(raceCourse.id)
    modalCode.value = detail.Code
    modalName.value = detail.Name
  } catch (error) {
    modalError.value =
      error instanceof Error ? error.message : '競馬場の取得に失敗しました。'
  } finally {
    isLoadingInitial.value = false
  }
}

function closeModal() {
  if (!isSubmitting.value) modalOpen.value = false
}

async function saveRaceCourse(input: { code: string; name: string }) {
  modalError.value = ''
  isSubmitting.value = true
  try {
    if (modalMode.value === 'create') {
      await createRaceCourse(input)
    } else if (modalRaceCourseId.value !== null) {
      await updateRaceCourse(modalRaceCourseId.value, input)
    }
    modalOpen.value = false
    await loadRaceCourses()
  } catch (error) {
    modalError.value =
      error instanceof Error ? error.message : '競馬場の保存に失敗しました。'
  } finally {
    isSubmitting.value = false
  }
}

async function removeRaceCourse(raceCourse: RaceCourseListItem) {
  if (!window.confirm(`「${raceCourse.name}」を削除しますか？`)) return
  errorMessage.value = ''
  try {
    await deleteRaceCourse(raceCourse.id)
    await loadRaceCourses()
  } catch (error) {
    errorMessage.value =
      error instanceof Error ? error.message : '競馬場の削除に失敗しました。'
  }
}
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold">競馬場リスト</h1>
    <button class="mt-4 rounded bg-green-600 px-4 py-2 text-white" @click="openCreate">
      競馬場を作成
    </button>

    <LoadingMessage v-if="isLoading" />
    <p v-else-if="errorMessage" class="mt-4 text-red-600">
      {{ errorMessage }}
      <ReloadButton @reload="loadRaceCourses" />
    </p>
    <p v-else-if="raceCourses.length === 0" class="mt-4">
      競馬場が見つかりませんでした。
    </p>
    <RaceCourseList
      v-if="!isLoading && !errorMessage && raceCourses.length > 0"
      class="mt-4"
      :race-courses="raceCourses"
      @edit="openUpdate"
      @delete="removeRaceCourse"
    />
    <RaceCourseFormModal
      :open="modalOpen"
      :mode="modalMode"
      :initial-code="modalCode"
      :initial-name="modalName"
      :is-loading-initial="isLoadingInitial"
      :is-submitting="isSubmitting"
      :error-message="modalError"
      @cancel="closeModal"
      @submit="saveRaceCourse"
    />
  </div>
</template>
