<script setup lang="ts">
import { onMounted, ref } from 'vue'

import { getUsers } from '@/api/users'
import LoadingMessage from '@/components/common/LoadingMessage.vue'
import ReloadButton from '@/components/common/ReloadButton.vue'
import UserList from '@/components/user/UserList.vue'
import {
  toUserListItem,
  type UserListItem,
} from '@/mappers/userMapper'

const users = ref<UserListItem[]>([])
const isLoading = ref(true)
const errorMessage = ref('')

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
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold">
      User一覧のページ
    </h1>

    <LoadingMessage v-if="isLoading" />
    <p v-else-if="errorMessage" class="mt-4 text-red-600">
      {{ errorMessage }}
      <ReloadButton @reload="loadUsers" />
    </p>
    <p v-else-if="users.length === 0" class="mt-4">
      ユーザーが見つかりませんでした。
    </p>
    <UserList v-else class="mt-4" :users="users" />
  </div>
</template>
