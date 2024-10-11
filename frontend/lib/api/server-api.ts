import ky from "ky";
import {
  GetAllListsBySpaceIDResponseDto,
  GetAllSpacesResponseDto, GetAllTasksBySpaceListIDResponseDto,
  LoginUserResponseDto,
  SuccessResponse,
} from "@/types/server-api-types";
import { LoginUserRequestDto } from "@/schemas/user-schemas";
import { cookies } from "next/headers";

export const serverApiClient = ky.extend({
  prefixUrl: "http://localhost:8081/api",
  hooks: {
    beforeRequest: [
      request => {
        const token = cookies().get("auth_token")

        if (token) {
          request.headers.set("Authorization", `Bearer ${token.value}`);
        }

        return request;
      }
    ]
  }
});

export const loginUserRequest = (dto: LoginUserRequestDto) =>
  serverApiClient.post("user/login", {
    json: dto
  }).json<SuccessResponse<LoginUserResponseDto>>();

export const registerUserRequest = (dto: any) =>
  serverApiClient.post("user/register", {
    json: dto
  }).json();

export const getAllSpacesRequest = () =>
  serverApiClient.get("space").json<SuccessResponse<GetAllSpacesResponseDto>>();

export const getAllListsBySpaceIDRequest = (spaceID: string) =>
  serverApiClient.get(`space/${spaceID}/list`).json<SuccessResponse<GetAllListsBySpaceIDResponseDto>>();

export const getAllTasksBySpaceListIDRequest = (spaceID: string, listID: string) => 
  serverApiClient.get(`space/${spaceID}/list/${listID}/task`).json<SuccessResponse<GetAllTasksBySpaceListIDResponseDto>>()
