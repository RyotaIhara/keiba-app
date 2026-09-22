import type { ApiRace, RaceInput } from '@/api/races'

export interface RaceListItem {
  id: number
  raceDate: string
  raceCourseId: number
  raceCourseName: string
  raceNumber: number
  raceName: string
  startTime: string
  surface: ApiRace['Surface']
  surfaceLabel: string
  distance: number
  direction: ApiRace['Direction']
  directionLabel: string
  weather: ApiRace['Weather']
  weatherLabel: string
  trackCondition: ApiRace['TrackCondition']
  trackConditionLabel: string
  raceConditions: string
}

const surfaceLabels: Record<ApiRace['Surface'], string> = {
  turf: '芝',
  dirt: 'ダート',
}

const directionLabels: Record<ApiRace['Direction'], string> = {
  right: '右回り',
  left: '左回り',
}

const weatherLabels: Record<ApiRace['Weather'], string> = {
  sunny: '晴れ',
  cloudy: '曇り',
  rainy: '雨',
  snowy: '雪',
}

const trackConditionLabels: Record<ApiRace['TrackCondition'], string> = {
  firm: '良',
  good: '稍良',
  yield: '重',
  soft: '不良',
}

function dateValue(value: string): string {
  return value.includes('T') ? value.slice(0, 10) : value.slice(0, 10)
}

function timeValue(value: string): string {
  const match = value.match(/(?:T|\s)(\d{2}:\d{2}:\d{2})/)
  return match?.[1] ?? value.slice(0, 8)
}

export function toRaceListItem(race: ApiRace): RaceListItem {
  return {
    id: race.ID,
    raceDate: dateValue(race.RaceDate),
    raceCourseId: race.Racecourse.ID,
    raceCourseName: race.Racecourse.Name,
    raceNumber: race.RaceNumber,
    raceName: race.RaceName,
    startTime: timeValue(race.StartTime),
    surface: race.Surface,
    surfaceLabel: surfaceLabels[race.Surface],
    distance: race.Distance,
    direction: race.Direction,
    directionLabel: directionLabels[race.Direction],
    weather: race.Weather,
    weatherLabel: weatherLabels[race.Weather],
    trackCondition: race.TrackCondition,
    trackConditionLabel: trackConditionLabels[race.TrackCondition],
    raceConditions: race.RaceConditions,
  }
}

export function toRaceInput(race: ApiRace): RaceInput {
  const item = toRaceListItem(race)
  return {
    race_date: item.raceDate,
    race_course_id: item.raceCourseId,
    race_number: item.raceNumber,
    race_name: item.raceName,
    start_time: item.startTime,
    surface: item.surface,
    distance: item.distance,
    direction: item.direction,
    weather: item.weather,
    track_condition: item.trackCondition,
    race_conditions: item.raceConditions,
  }
}
