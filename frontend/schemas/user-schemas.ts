import { z } from "zod";

export const loginUserSchema = z.object({
  email: z.string().email(),
  password: z.string(),
})

export type LoginUserRequestDto = z.infer<typeof loginUserSchema>