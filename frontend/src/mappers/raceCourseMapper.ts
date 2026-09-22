import type { ApiRaceCourse } from '@/api/raceCourses'

export interface RaceCourseListItem {
  id: number
  code: string
  name: string
}

export function toRaceCourseListItem(
  raceCourse: ApiRaceCourse,
): RaceCourseListItem {
  return {
    id: raceCourse.ID,
    code: raceCourse.Code,
    name: raceCourse.Name,
  }
}
