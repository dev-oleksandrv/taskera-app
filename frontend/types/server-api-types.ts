export interface LoginUserResponseDto {
  token: string;
}

export interface SuccessResponse<T> {
  code: number;
  status: string;
  data: T;
}