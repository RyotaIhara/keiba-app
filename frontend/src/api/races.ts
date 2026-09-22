export interface ApiRaceCourseReference {
  ID: number
  Code: string
  Name: string
}

export interface ApiRace {
  ID: number
  RaceDate: string
  Racecourse: ApiRaceCourseReference
  RaceNumber: number
  RaceName: string
  StartTime: string
  Surface: 'turf' | 'dirt'
  Distance: number
  Direction: 'right' | 'left'
  Weather: 'sunny' | 'cloudy' | 'rainy' | 'snowy'
  TrackCondition: 'firm' | 'good' | 'yield' | 'soft'
  RaceConditions: string
}

export interface RaceInput {
  race_date: string
  race_course_id: number
  race_number: number
  race_name: string
  start_time: string
  surface: ApiRace['Surface']
  distance: number
  direction: ApiRace['Direction']
  weather: ApiRace['Weather']
  track_condition: ApiRace['TrackCondition']
  race_conditions: string
}

export interface RaceSearchParams {
  race_date?: string
  race_course_id?: number
}

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL?.replace(/\/+$/, '')

if (!apiBaseUrl) {
  throw new Error('VITE_API_BASE_URL is not configured')
}

const RACES_API_URL = `${apiBaseUrl}/api/races`

async function request(url: string, options: RequestInit, errorMessage: string): Promise<Response> {
  const response = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  })

  if (!response.ok) throw new Error(`${errorMessage}: ${response.status}`)
  return response
}

export async function getRaces(params: RaceSearchParams = {}): Promise<ApiRace[]> {
  const searchParams = new URLSearchParams()
  if (params.race_date) searchParams.set('race_date', params.race_date)
  if (params.race_course_id !== undefined) {
    searchParams.set('race_course_id', String(params.race_course_id))
  }

  const query = searchParams.toString()
  const response = await request(
    query ? `${RACES_API_URL}?${query}` : RACES_API_URL,
    { method: 'GET' },
    'レースの取得に失敗しました',
  )
  return response.json() as Promise<ApiRace[]>
}

export async function getRace(id: number): Promise<ApiRace> {
  const response = await request(
    `${RACES_API_URL}/${id}`,
    { method: 'GET' },
    'レースの取得に失敗しました',
  )
  return response.json() as Promise<ApiRace>
}

export async function createRace(input: RaceInput): Promise<ApiRace> {
  const response = await request(
    RACES_API_URL,
    { method: 'POST', body: JSON.stringify(input) },
    'レースの作成に失敗しました',
  )
  return response.json() as Promise<ApiRace>
}

export async function updateRace(id: number, input: RaceInput): Promise<ApiRace> {
  const response = await request(
    `${RACES_API_URL}/${id}`,
    { method: 'PUT', body: JSON.stringify(input) },
    'レースの更新に失敗しました',
  )
  return response.json() as Promise<ApiRace>
}

export async function deleteRace(id: number): Promise<void> {
  await request(`${RACES_API_URL}/${id}`, { method: 'DELETE' }, 'レースの削除に失敗しました')
}
