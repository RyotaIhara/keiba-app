import type { ApiUser } from '@/api/users'

export interface UserListItem {
  id: number
  code: string
  name: string
}

export function toUserListItem(user: ApiUser): UserListItem {
  return {
    id: user.ID,
    code: user.Code,
    name: user.Name,
  }
}
