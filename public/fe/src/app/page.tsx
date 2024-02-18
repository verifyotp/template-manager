import { Metadata } from "next"
import Link from "next/link"

import { cn } from "@/lib/utils"
import { buttonVariants } from "@/registry/new-york/ui/button"
export const metadata: Metadata = {
  title: "Template Manager",
  description: "Manage your templates with ease",
}

export default function LandingPage() {
  return (
    <div className="container h-screen flex justify-center items-center relative">

      {/* Second component */}
      <div className="max-w-lg p-8 bg-white rounded-lg shadow-lg absolute top-1/2 transform -translate-y-1/2">
        <h1 className="text-3xl font-semibold mb-4">Template Manager</h1>
        <p className="text-lg text-gray-600 mb-6">Please login to access your account.</p>
        <Link
          href="/auth/signup"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            "block w-full text-center py-3 px-4 rounded-md shadow-md"
          )}
        >
          Sign Up
        </Link>
      </div>
    </div>
  )
}