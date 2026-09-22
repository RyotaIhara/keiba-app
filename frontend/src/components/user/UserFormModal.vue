<script setup lang="ts">
import { computed, ref, watch } from 'vue'

const props = withDefaults(
  defineProps<{
    open: boolean
    mode: 'create' | 'update'
    initialCode?: string
    initialName?: string
    isLoadingInitial?: boolean
    isSubmitting?: boolean
    errorMessage?: string
  }>(),
  {
    initialCode: '',
    initialName: '',
    isLoadingInitial: false,
    isSubmitting: false,
    errorMessage: '',
  },
)

const emit = defineEmits<{
  cancel: []
  submit: [
    {
      code: string
      name: string
      password?: string
    },
  ]
}>()

const code = ref('')
const name = ref('')
const password = ref('')
const passwordConfirmation = ref('')
const validationMessage = ref('')

const title = computed(() =>
  props.mode === 'create' ? 'ユーザーを作成' : 'ユーザーを編集',
)

function resetForm() {
  code.value = props.initialCode
  name.value = props.initialName
  password.value = ''
  passwordConfirmation.value = ''
  validationMessage.value = ''
}

watch(
  () => [props.open, props.mode, props.initialCode, props.initialName],
  () => {
    if (props.open) resetForm()
  },
  { immediate: true },
)

function submit() {
  validationMessage.value = ''
  if (props.mode === 'create' && password.value !== passwordConfirmation.value) {
    validationMessage.value = 'パスワードが一致しません。'
    return
  }

  emit('submit', {
    code: code.value,
    name: name.value,
    ...(props.mode === 'create' ? { password: password.value } : {}),
  })
}
</script>

<template>
  <div
    v-if="open"
    class="fixed inset-0 z-10 flex items-center justify-center bg-black/50 p-4"
    role="dialog"
    aria-modal="true"
    :aria-label="title"
  >
    <form class="w-full max-w-md rounded bg-white p-6 shadow-lg" @submit.prevent="submit">
      <h2 class="text-xl font-bold">{{ title }}</h2>
      <p v-if="errorMessage" class="mt-3 text-red-600">{{ errorMessage }}</p>
      <p v-if="validationMessage" class="mt-3 text-red-600">{{ validationMessage }}</p>
      <p v-if="isLoadingInitial" class="mt-3">詳細を読み込んでいます...</p>

      <label class="mt-4 block">
        <span class="mb-1 block">Code</span>
        <input v-model="code" required :disabled="isLoadingInitial" class="w-full rounded border px-3 py-2" />
      </label>
      <label class="mt-3 block">
        <span class="mb-1 block">ユーザー名</span>
        <input v-model="name" required :disabled="isLoadingInitial" class="w-full rounded border px-3 py-2" />
      </label>
      <template v-if="mode === 'create'">
        <label class="mt-3 block">
          <span class="mb-1 block">パスワード</span>
          <input v-model="password" required type="password" :disabled="isLoadingInitial" class="w-full rounded border px-3 py-2" />
        </label>
        <label class="mt-3 block">
          <span class="mb-1 block">パスワード（確認）</span>
          <input
            v-model="passwordConfirmation"
            required
            type="password"
            :disabled="isLoadingInitial"
            class="w-full rounded border px-3 py-2"
          />
        </label>
      </template>

      <div class="mt-6 flex justify-end gap-3">
        <button type="button" class="rounded border px-4 py-2" :disabled="isSubmitting || isLoadingInitial" @click="emit('cancel')">
          キャンセル
        </button>
        <button
          type="submit"
          class="rounded bg-blue-600 px-4 py-2 text-white disabled:opacity-50"
          :disabled="isSubmitting || isLoadingInitial"
        >
          {{ isSubmitting ? '保存中...' : '保存' }}
        </button>
      </div>
    </form>
  </div>
</template>
