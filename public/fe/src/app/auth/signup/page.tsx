import { Metadata } from "next"
import { SignupForm } from "@/components/auth/signup-form"

export const metadata: Metadata = {
  title: "Sign Up",
  description: "Sign up for an account",
}

export default function SignUpPage() {
  return (
    <>
     
            <SignupForm />
            
    </>
  )
}