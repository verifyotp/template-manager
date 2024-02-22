"use client";

import * as React from "react"
import { FormEvent } from 'react'
import { useRouter } from 'next/navigation'

import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"

import { useToast } from "@/components/ui/use-toast"


interface UserLoginFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserLoginForm({ className, ...props }: UserLoginFormProps) {
  const router = useRouter(); // Initialize useRouter
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const [email, setEmail] = React.useState<string>("");
  const [password, setPassword] = React.useState<string>("");
  const { toast } = useToast()
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);

    loginUser(email, password)
      .then((data) => {
        setIsLoading(false);
        toast({
          title: "Success",
          description: data.message,
        })

        // save token to local storage
        localStorage.setItem('session', JSON.stringify(data.data?.session));
        localStorage.setItem('account', JSON.stringify(data.data?.account));
        localStorage.setItem('authToken', data.data?.session.token as string) ;
        router.push('/dashboard');
      })
      .catch((error) => {
        setIsLoading(false);
        toast({
          variant: "destructive",
          title: "Error",
          description: error.message,
        })
      });
  };

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <form onSubmit={handleSubmit}>
        <div className="grid gap-3">
          <div className="grid gap-3">
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
              required={true}
            />
          </div>
          <div className="grid gap-1">
            <Label className="sr-only" htmlFor="password">
              Password
            </Label>
            <Input
              id="password"
              placeholder="44asfeYY"
              type="password"
              autoCapitalize="none"
              autoComplete="password"
              autoCorrect="off"
              value={password}
              disabled={isLoading}
              onChange={(e) => setPassword(e.target.value)}
              required={true}
            />
          </div>

          <Button type="submit" disabled={isLoading}>
            {isLoading && (
              <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
            )}
            Login
          </Button>

          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <input
                id="remember-me"
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary"
              />
              <Label htmlFor="remember-me" className="ml-2 text-sm">
                Remember me
              </Label>
            </div>

            <div className="text-sm">
              <a href="/auth/reset-password" className="font-medium text-primary">
                Forgot your password ?
              </a>
            </div>
          </div>
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

export async function loginUser(email: string, password: string): Promise<Response<LoginResponse>> {
  const loginData = {
    email,
    password,
  };

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(loginData)
  };

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/users/login`, requestOptions);
    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }
    return data as Response<LoginResponse>;
  } catch (error: any) {
    throw new Error(error.message);
  }
}


interface Session {
  id: string;
  account_id: string;
  device: any; // Define the Device type if needed
  token: string;
  expires_at: string;
  last_active: string;
  created_at: string;
}

interface Account {
  id: string;
  email: string;
  verified_at: string | null; // This can be a string or null
  created_at: string;
  updated_at: string | null; // This can be a string or null
}

interface LoginResponse {
  account: Account;
  session: Session;
}

interface Response<T = any> {
  status: boolean;
  message: string;
  data?: T;
}

