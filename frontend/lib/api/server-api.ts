import ky from "ky";
import { LoginUserResponseDto, SuccessResponse } from "@/types/server-api-types";
import { LoginUserRequestDto } from "@/schemas/user-schemas";

export const serverApiClient = ky.extend({
  prefixUrl: "http://localhost:8081/api"

});

export const loginUserRequest = (dto: LoginUserRequestDto) =>
  serverApiClient.post("user/login", {
    json: dto
  }).json<SuccessResponse<LoginUserResponseDto>>();

export const registerUserRequest = (dto: any) =>
  serverApiClient.post("user/register", {
    json: dto
  }).json();