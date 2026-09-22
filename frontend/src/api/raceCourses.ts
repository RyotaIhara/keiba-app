export interface ApiRaceCourse {
  ID: number
  Code: string
  Name: string
}

export interface CreateRaceCourseInput {
  code: string
  name: string
}

export type UpdateRaceCourseInput = CreateRaceCourseInput

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL?.replace(/\/+$/, '')

if (!apiBaseUrl) {
  throw new Error('VITE_API_BASE_URL is not configured')
}

const RACE_COURSES_API_URL = `${apiBaseUrl}/api/race_courses`

async function request(
  url: string,
  options: RequestInit,
  errorMessage: string,
): Promise<Response> {
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

export async function getRaceCourses(): Promise<ApiRaceCourse[]> {
  const response = await request(
    RACE_COURSES_API_URL,
    { method: 'GET' },
    '競馬場の取得に失敗しました',
  )
  const raceCourses = (await response.json()) as ApiRaceCourse[] | null
  return raceCourses ?? []
}

export async function getRaceCourse(id: number): Promise<ApiRaceCourse> {
  const response = await request(
    `${RACE_COURSES_API_URL}/${id}`,
    { method: 'GET' },
    '競馬場の取得に失敗しました',
  )
  return response.json() as Promise<ApiRaceCourse>
}

export async function createRaceCourse(
  input: CreateRaceCourseInput,
): Promise<ApiRaceCourse> {
  const response = await request(
    RACE_COURSES_API_URL,
    { method: 'POST', body: JSON.stringify(input) },
    '競馬場の作成に失敗しました',
  )
  return response.json() as Promise<ApiRaceCourse>
}

export async function updateRaceCourse(
  id: number,
  input: UpdateRaceCourseInput,
): Promise<ApiRaceCourse> {
  const response = await request(
    `${RACE_COURSES_API_URL}/${id}`,
    { method: 'PUT', body: JSON.stringify(input) },
    '競馬場の更新に失敗しました',
  )
  return response.json() as Promise<ApiRaceCourse>
}

export async function deleteRaceCourse(id: number): Promise<void> {
  await request(
    `${RACE_COURSES_API_URL}/${id}`,
    { method: 'DELETE' },
    '競馬場の削除に失敗しました',
  )
}
