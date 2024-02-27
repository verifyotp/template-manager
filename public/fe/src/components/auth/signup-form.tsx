"use client";

import * as React from "react"
import { useState } from "react";
import { useTransition } from "react";
import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/new-york/input"
import { signUpRequest as signUp } from "@/actions/signup"
import Link from "next/link"
import { buttonVariants } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { SignupSchema } from "@/schemas";
import { FormError } from "@/components/form-error";
import { FormSuccess } from "@/components/form-success";


interface SignupFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function SignupForm({ className, ...props }: SignupFormProps) {

  const [isPending, startTransition] = useTransition();
  const [error, setError] = useState<string | undefined>("");
  const [success, setSuccess] = useState<string | undefined>("");

  const form = useForm<z.infer<typeof SignupSchema>>({
    resolver: zodResolver(SignupSchema),
    defaultValues: {
      email: "",
    },
  });


  const onSubmit = (values: z.infer<typeof SignupSchema>) => {
    setError("");
    setSuccess("");

    startTransition(() => {
      signUp(values.email)
        .then((data) => {
          if (data?.status === false) {
            form.reset();
            setError(data.message);
          }

          if (data?.status) {
            form.reset();
            setSuccess(data.message);
          }
        })
        .catch(() => setError("Something went wrong"));
    });
  };

  return (
    <div className="container relative h-[800px] flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-1 lg:px-0">

      <div className="lg:p-auto pt-10 ">
        <div className="mx-auto flex w-full flex-col justify-center space-y-8 sm:w-[350px]">
          <div className="flex flex-col space-y-2 text-center">
            <h1 className="text-2xl font-semibold tracking-tight">
              Create an account
            </h1>
            <p className="text-sm text-muted-foreground">
              Enter your email below to create your account
            </p>
          </div>
          <div className={cn("grid gap-9", className)} {...props}>
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
              >
                <div className="grid gap-5">
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
                  <FormError message={error} />
                  <FormSuccess message={success} />
                  <Button disabled={isPending}>
                    {isPending && (
                      <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
                    )}
                    Sign Up
                  </Button>
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
                <a href="/auth/login" className="font-small p-2 text-primary opacity-[0.60] hover:opacity-[1.0]">
                  Have an account? Login
                </a>
              </Button>
            </div>
          </div>
          <p className="px-8 text-center text-sm text-muted-foreground">
            By clicking sign up, you agree to our{" "}
            <Link
              href="/terms"
              className="underline underline-offset-4 hover:text-primary"
            >
              Terms of Service
            </Link>{" "}
            and{" "}
            <Link
              href="/privacy"
              className="underline underline-offset-4 hover:text-primary"
            >
              Privacy Policy
            </Link>
            .
          </p>
        </div>
      </div>
    </div>
  )
}