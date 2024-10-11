import { PropsWithChildren } from "react";
import { BaseHeader } from "@/components/base-header";

export default function DefaultLayout({ children }: PropsWithChildren) {
  return (
    <div className="flex flex-col min-h-screen bg-gradient-to-br from-teal-50 to-cyan-100">
      <BaseHeader />
      <main className="flex-1">
        {children}
      </main>
    </div>
  )
}