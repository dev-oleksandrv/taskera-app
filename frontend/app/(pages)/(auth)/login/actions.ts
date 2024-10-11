"use server"

import { loginUserRequest } from "@/lib/api/server-api";
import { cookies } from "next/headers";
import { ServerActionResult } from "@/types/server-actions-types";
import { loginUserSchema } from "@/schemas/user-schemas";

export const loginUser = async (_: ServerActionResult | null, formData: FormData): Promise<ServerActionResult> => {
  try {
    const input = {
      email: formData.get("email"),
      password: formData.get("password")
    }

    const result = loginUserSchema.safeParse(input);

    if (!result.success) {
      return { success: false, fieldErrors: result.error.formErrors.fieldErrors }
    }

    const response = await loginUserRequest({
      email: "test@test.com",
      password: "test"
    });

    if (response.code === 200) {
      cookies().set("auth_token", response.data.token)
      return { success: true }
    }

    return { success: false, errors: [response.status] }
  } catch (error: any) {
    return { success: false, errors: [error?.message] }
  }
}