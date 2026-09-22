<script setup lang="ts">
import { onMounted, ref } from 'vue'

import {
  createUser,
  deleteUser,
  getUser,
  getUsers,
  updateUser,
} from '@/api/users'
import LoadingMessage from '@/components/common/LoadingMessage.vue'
import ReloadButton from '@/components/common/ReloadButton.vue'
import UserFormModal from '@/components/user/UserFormModal.vue'
import UserList from '@/components/user/UserList.vue'
import {
  toUserListItem,
  type UserListItem,
} from '@/mappers/userMapper'

const users = ref<UserListItem[]>([])
const isLoading = ref(true)
const errorMessage = ref('')
const modalOpen = ref(false)
const modalMode = ref<'create' | 'update'>('create')
const modalUserId = ref<number | null>(null)
const modalCode = ref('')
const modalName = ref('')
const modalError = ref('')
const isLoadingInitial = ref(false)
const isSubmitting = ref(false)

async function loadUsers() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const apiUsers = await getUsers()
    users.value = apiUsers.map(toUserListItem)
  } catch (error) {
    users.value = []
    errorMessage.value =
      error instanceof Error
        ? error.message
        : 'ユーザーの取得に失敗しました。'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadUsers)

function openCreate() {
  modalMode.value = 'create'
  modalUserId.value = null
  modalCode.value = ''
  modalName.value = ''
  modalError.value = ''
  isLoadingInitial.value = false
  modalOpen.value = true
}

async function openUpdate(user: UserListItem) {
  modalMode.value = 'update'
  modalUserId.value = user.id
  modalCode.value = ''
  modalName.value = ''
  modalError.value = ''
  isLoadingInitial.value = true
  modalOpen.value = true
  try {
    const detail = await getUser(user.id)
    modalCode.value = detail.Code
    modalName.value = detail.Name
  } catch (error) {
    modalError.value = error instanceof Error ? error.message : 'ユーザーの取得に失敗しました。'
  } finally {
    isLoadingInitial.value = false
  }
}

function closeModal() {
  if (!isSubmitting.value) modalOpen.value = false
}

async function saveUser(input: { code: string; name: string; password?: string }) {
  modalError.value = ''
  isSubmitting.value = true
  try {
    if (modalMode.value === 'create') {
      await createUser({ code: input.code, name: input.name, password: input.password ?? '' })
    } else if (modalUserId.value !== null) {
      await updateUser(modalUserId.value, { code: input.code, name: input.name })
    }
    modalOpen.value = false
    await loadUsers()
  } catch (error) {
    modalError.value = error instanceof Error ? error.message : 'ユーザーの保存に失敗しました。'
  } finally {
    isSubmitting.value = false
  }
}

async function removeUser(user: UserListItem) {
  if (!window.confirm(`「${user.name}」を削除しますか？`)) return
  errorMessage.value = ''
  try {
    await deleteUser(user.id)
    await loadUsers()
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : 'ユーザーの削除に失敗しました。'
  }
}
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold">
      ユーザーリスト
    </h1>
    <button class="mt-4 rounded bg-green-600 px-4 py-2 text-white"
      @click="openCreate"
    >
      ユーザーを作成
    </button>

    <LoadingMessage v-if="isLoading" />
    <p v-else-if="errorMessage" class="mt-4 text-red-600">
      {{ errorMessage }}
      <ReloadButton @reload="loadUsers" />
    </p>
    <p v-else-if="users.length === 0" class="mt-4">
      ユーザーが見つかりませんでした。
    </p>
    <UserList
      v-if="!isLoading && !errorMessage && users.length > 0"
      class="mt-4"
      :users="users"
      @edit="openUpdate"
      @delete="removeUser"
    />
    <UserFormModal
      :open="modalOpen"
      :mode="modalMode"
      :initial-code="modalCode"
      :initial-name="modalName"
      :is-loading-initial="isLoadingInitial"
      :is-submitting="isSubmitting"
      :error-message="modalError"
      @cancel="closeModal"
      @submit="saveUser"
    />
  </div>
</template>
