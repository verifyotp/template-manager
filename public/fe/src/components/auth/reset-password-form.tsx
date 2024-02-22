"use client";

import * as React from "react"

import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"

import { useRouter } from 'next/navigation'
import { useToast } from "@/components/ui/use-toast"



export async function initiatePasswordReset(email: string,): Promise<Response> {
  const requestData = {
    email,
  };

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(requestData)
  };

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/users/reset-password`, requestOptions);
    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }
    return data as Response;
  } catch (error : any) {
    throw new Error(error.message);
  }
}

interface Response<T = any> {
  status: boolean;
  message: string;
  data?: T;
}


interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserResetPasswordForm({ className, ...props }: UserAuthFormProps) {
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const { toast } = useToast()
  const router = useRouter(); // Initialize useRouter
  const [email, setEmail] = React.useState<string>("");
  

  async function onSubmit(event: React.SyntheticEvent) {
    event.preventDefault()
    setIsLoading(true)

    initiatePasswordReset(email)
      .then((data) => {
        setIsLoading(false);
        toast({
          title: "Success",
          description: data.message,
        })
        
        setTimeout(() => {
          router.push('/auth/login');
        }, 3000)
      })
      .catch((error) => {
        setIsLoading(false);
        toast({
          variant: "destructive",
          title: "Error",
          description: error.message,
        })
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
              disabled={isLoading}
            />
          </div>
          <Button disabled={isLoading}>
            {isLoading && (
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