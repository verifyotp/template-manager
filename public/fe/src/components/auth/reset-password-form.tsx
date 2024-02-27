"use client";

import * as React from "react"

import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/new-york/input"
import { Label } from "@/components/ui/new-york/label"
import { useTransition } from "react";
import { useRouter } from 'next/navigation'
import { useToast } from "@/components/ui/use-toast"
import {initiatePasswordReset} from "@/actions/reset-password"




interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserResetPasswordForm({ className, ...props }: UserAuthFormProps) {
  const [isPending, startTransition] = useTransition();
  const router = useRouter(); // Initialize useRouter
  const [email, setEmail] = React.useState<string>("");
  const { toast } = useToast()
  

  async function onSubmit(event: React.SyntheticEvent) {
    event.preventDefault()
 
    startTransition(() => {
      initiatePasswordReset(email)
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
    <div className={cn("grid gap-6", className)} {...props}>
      <form onSubmit={onSubmit}>
        <div className="grid gap-5">
          <div className="grid gap-1">
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
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              disabled={isPending}
            />
          </div>
          <Button disabled={isPending}>
            {isPending && (
              <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
            )}
            Reset Password
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