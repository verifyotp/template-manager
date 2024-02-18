"use client";

import * as React from "react"
import { FormEvent } from 'react'
import { useRouter } from 'next/navigation'

import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/registry/new-york/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"



interface UserLoginFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserLoginForm({ className, ...props }: UserLoginFormProps) {
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const [apiError, setApiError] = React.useState<string | null>(null);
  const router = useRouter(); // Initialize useRouter

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget as HTMLFormElement);

    const email = formData.get("email") as string;
    const password = formData.get("password") as string;

    console.log('Form Data:', formData);
    console.log(email, password);

    setIsLoading(true);
    setApiError(null);

    try {
      await loginUser(email, password);
      router.push('/dashboard');
    } catch (error) {
      setApiError('An error occurred');
    } finally {
      setIsLoading(false);
    }
  }

  // alert error message
  React.useEffect(() => {
    if (apiError) {
      alert(apiError);
    }
  setApiError(null);
  }, [apiError]);

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
              disabled={isLoading}
            />
          </div>
          <div className="grid gap-1">
            <Label className="sr-only" htmlFor="email">
              Password
            </Label>
            <Input
              id="password"
              placeholder="44asfeYY"
              type="password"
              autoCapitalize="none"
              autoComplete="password"
              autoCorrect="off"
              disabled={isLoading}
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


export async function loginUser(email: string, password: string) : Promise<Response<LoginResponse>> {
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
    const response = await fetch('http://localhost:9000/api/user/login', requestOptions);
    if (!response.ok) {
      const errorMessage = await response.text();
      throw new Error(errorMessage || 'Failed to login');
    }

    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }

    // save token to local storage
    localStorage.setItem('session', data.session);
    localStorage.setItem('account', data.account);
    return data;
  } catch (error ) {
    throw new Error(error as string || 'An error occurred');
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

