"use client"

import { PropsWithChildren } from "react";
import Link from "next/link";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { usePathname } from "next/navigation";

export default function AuthLayout({children}: PropsWithChildren) {
  const pathname = usePathname()
  const active = pathname.includes("login") ? "login" : "register";

  return (
    <div className="flex-1 flex items-center justify-center">
      <div className="w-full max-w-md">
        <Tabs value={active} className="w-full">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="login" asChild className="text-teal-700">
              <Link href="/login">
                Login
              </Link>
            </TabsTrigger>
            <TabsTrigger value="register" asChild className="text-teal-700">
              <Link href="/register">
                Register
              </Link>
            </TabsTrigger>
          </TabsList>
          <div className="mt-4 rounded-lg border bg-white shadow-sm p-6">
            {children}
          </div>
        </Tabs>
      </div>
    </div>
  )
}