import { Metadata } from "next"
import Link from "next/link"
import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"
import { LoginForm } from "@/components/auth/login-form"

export const metadata: Metadata = {
  title: "Login",
  description: "Login to your account",
}


export default function LoginPage() {
  return (
    <>
      <LoginForm />
    </>
  )
}