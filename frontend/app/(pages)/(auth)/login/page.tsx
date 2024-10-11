"use client";

import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { useFormState } from "react-dom";
import { loginUser } from "@/app/(pages)/(auth)/login/actions";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const router = useRouter();

  const [state, formAction] = useFormState(loginUser, null)

  if (state?.success) {
    return router.push("/dashboard")
  }

  return (
    <>
      <form action={formAction} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            placeholder="Enter your email"
            required
            type="email"
            name="email"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            required
            type="password"
            name="password"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <Button
          className="w-full bg-teal-600 text-white hover:bg-teal-700"
          type="submit"
        >
          Login
        </Button>
      </form>

      <div className="mt-4 text-center">
        <Link
          className="text-sm text-teal-600 hover:underline"
          href="/forgot-password"
        >
          Forgot Password?
        </Link>
      </div>
    </>
  );
}
