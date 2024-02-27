"use client";

import * as React from "react"
import { useTransition } from "react";
import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"
import { signUpRequest as signUp } from "@/actions/signup"
import { useRouter } from 'next/navigation'
import { useToast } from "@/components/ui/use-toast"



interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserSignupForm({ className, ...props }: UserAuthFormProps) {
    
  const router = useRouter(); // Initialize useRouter
  const [email, setEmail] = React.useState<string>("");
  const [isPending, startTransition] = useTransition();
  const { toast } = useToast()
  
  async function onSubmit(event: React.SyntheticEvent) {
    event.preventDefault()
    startTransition(() => {
      signUp(email)
        .then((data) => {
          toast({
            title: "Success",
            description: data.message,
          })
          router.push('/auth/login');
        })
        .catch((error) => {
          toast({
            variant: "destructive",
            title: "Error",
            description: error.message,
          })
        });
    });
  }

  return (
    <div className={cn("grid gap-9", className)} {...props}>
      <form onSubmit={onSubmit}>
        <div className="grid gap-5">
          <div className="grid gap-5">
            <Label className="sr-only" htmlFor="email">
              Email
            </Label>
            <Input
              id="email"
              placeholder="name@example.com"
              type="email"
              autoCapitalize="none"
              autoComplete="email"
              autoCorrect="off"
              disabled={isPending}
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <Button disabled={isPending}>
            {isPending && (
              <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
            )}
            Sign Up
          </Button>
        </div>
      </form>
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t" />
        </div>
      </div>
    </div>
  )
}