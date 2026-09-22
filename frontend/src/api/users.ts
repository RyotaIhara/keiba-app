export interface ApiUser {
  ID: number
  Code: string
  Name: string
  Password: string
}

const USERS_API_URL = 'http://localhost:3000/api/users'

export async function getUsers(): Promise<ApiUser[]> {
  const response = await fetch(USERS_API_URL)

  if (!response.ok) {
    throw new Error(`ユーザーの取得に失敗しました: ${response.status}`)
  }

  return response.json() as Promise<ApiUser[]>
}
