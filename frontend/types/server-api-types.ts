export interface LoginUserResponseDto {
  token: string;
}

export interface SuccessResponse<T> {
  code: number;
  status: string;
  data: T;
}

export interface SpaceDto {
  id: string;
  name: string;
  description: string;
}

export interface SpaceWithRoleDto {
  space: SpaceDto;
  role: string;
}

export interface GetAllSpacesResponseDto {
  spaces: SpaceWithRoleDto[]
}

export interface ListDto {
  id: string;
  name: string;
  description: string;
  emoji: string;
}

export interface GetAllListsBySpaceIDResponseDto {
  lists: ListDto[];
}

export interface TaskDto {
  id: string;
  content: string;
  full_content: string;
  completed: boolean;
  order: number;
}

export interface GetAllTasksBySpaceListIDResponseDto {
  tasks: TaskDto[];
}