export interface ApiRaceDetail {
  ID: number
  Race: {
    ID: number
    RaceDate: string
    Racecourse: {
      ID: number
      Code: string
      Name: string
    }
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
  HorseNumber: number
  FrameNumber: number
  HorseName: string
  Sex: 'colt' | 'filly'
  Age: number
  Weight: number
  Jockey: string
  Stable: string
  BodyWeight: number
  BodyWeightChange: number
  Odds: number
  Popularity: number
}

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL?.replace(/\/+$/, '')

if (!apiBaseUrl) {
  throw new Error('VITE_API_BASE_URL is not configured')
}

async function request(url: string): Promise<Response> {
  const response = await fetch(url, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  })

  if (!response.ok) throw new Error(`出走馬の取得に失敗しました: ${response.status}`)
  return response
}

export async function getRaceDetails(raceId: number): Promise<ApiRaceDetail[]> {
  const response = await request(`${apiBaseUrl}/api/races/${raceId}/details`)
  return response.json() as Promise<ApiRaceDetail[]>
}
