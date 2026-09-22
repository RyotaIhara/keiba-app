export interface ApiUser {
  ID: number
  Code: string
  Name: string
  Password: string
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
