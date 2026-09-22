export interface ApiUser {
  ID: number
  Code: string
  Name: string
  Password: string
}

export interface CreateUserInput {
  code: string
  name: string
  password: string
}

export interface UpdateUserInput {
  code: string
  name: string
}

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL?.replace(/\/+$/, '')

if (!apiBaseUrl) {
  throw new Error('VITE_API_BASE_URL is not configured')
}

const USERS_API_URL = `${apiBaseUrl}/api/users`

export async function getUsers(): Promise<ApiUser[]> {
  const response = await fetch(USERS_API_URL)

  if (!response.ok) {
    throw new Error(`ユーザーの取得に失敗しました: ${response.status}`)
  }

  return response.json() as Promise<ApiUser[]>
}

async function request(url: string, options: RequestInit, errorMessage: string) {
  const response = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  })

  if (!response.ok) {
    throw new Error(`${errorMessage}: ${response.status}`)
  }

  return response
}

export async function getUser(id: number): Promise<ApiUser> {
  const response = await request(
    `${USERS_API_URL}/${id}`,
    { method: 'GET' },
    'ユーザーの取得に失敗しました',
  )
  return response.json() as Promise<ApiUser>
}

export async function createUser(input: CreateUserInput): Promise<ApiUser> {
  const response = await request(
    USERS_API_URL,
    { method: 'POST', body: JSON.stringify(input) },
    'ユーザーの作成に失敗しました',
  )
  return response.json() as Promise<ApiUser>
}

export async function updateUser(
  id: number,
  input: UpdateUserInput,
): Promise<ApiUser> {
  const response = await request(
    `${USERS_API_URL}/${id}`,
    { method: 'PUT', body: JSON.stringify(input) },
    'ユーザーの更新に失敗しました',
  )
  return response.json() as Promise<ApiUser>
}

export async function deleteUser(id: number): Promise<void> {
  await request(
    `${USERS_API_URL}/${id}`,
    { method: 'DELETE' },
    'ユーザーの削除に失敗しました',
  )
}
