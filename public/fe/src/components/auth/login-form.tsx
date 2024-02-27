"use client";

import * as React from "react"
import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/new-york/input"
import { Label } from "@/components/ui/new-york/label"
import { login } from "@/actions/login"
import { useTransition } from "react";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { FormError } from "@/components/form-error";
import { FormSuccess } from "@/components/form-success";
import { LoginSchema } from "@/schemas";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";


interface LoginFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function LoginForm({ className, ...props }: LoginFormProps) {
  const router = useRouter(); // Initialize useRouter
  const [isPending, startTransition] = useTransition();
  const [error, setError] = useState<string | undefined>("");
  const [success, setSuccess] = useState<string | undefined>("");


  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });



  const onSubmit = (values: z.infer<typeof LoginSchema>) => {
    setError("");
    setSuccess("");

    startTransition(() => {
      login(values)
        .then((data) => {
          if (data?.status === false) {
            setError(data.message);
          }

          if (data?.status) {
            setSuccess(data.message);
            localStorage.setItem('authToken', data.data?.session.token as string);
            // wait for the token to be set
            setTimeout(() => {
              router.push('/dashboard');
            }, 750);
          }
        })
        .catch(() => setError("Something went wrong"));
    });
  };

  return (
    <div className="container relative hidden h-[800px] flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-1 lg:px-0">
      <div className="lg:p-auto">
        <div className="mx-auto flex w-full flex-col justify-center space-y-8 sm:w-[350px]">
          <div className="flex flex-col space-y-2 text-center">
            <h1 className="text-2xl font-semibold tracking-tight">
              Login
            </h1>
            <p className="text-sm text-muted-foreground">
              Enter your email and password to login
            </p>
          </div>
          <div className={cn("grid gap-5", className)} {...props}>
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
              >
                <div className="grid gap-3">
                  <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                      <FormItem>
                        <FormControl>
                          <Input
                            {...field}
                            disabled={isPending}
                            placeholder="name@example.com"
                            type="email"
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <div className="grid gap-1">
                    <FormField
                      control={form.control}
                      name="password"
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              {...field}
                              disabled={isPending}
                              placeholder="******"
                              type="password"
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                  </div>
                  <FormError message={error} />
                  <FormSuccess message={success} />
                  <Button type="submit" disabled={isPending} >
                    {isPending && (
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
                      <Button variant="link" >
                        <a href="/auth/reset" className="font-medium text-primary opacity-[0.70] hover:opacity-[1.0]">
                          Forgot your password?
                        </a>
                      </Button>
                    </div>
                  </div>
                </div>
              </form>
            </Form>
            <div className="relative">
              <div className="absolute inset-0 flex items-center">
                <span className="w-full border-t" />
              </div>
            </div>
            <div className="flex items-center justify-center">
              <Button variant="link" >
                <a href="/auth/signup" className="font-small p-2 text-primary opacity-[0.60] hover:opacity-[1.0]">
                  Don't have an account? Sign Up
                </a>
              </Button>
          </div>
          </div>
        </div>
      </div>
    </div>
  )
}
