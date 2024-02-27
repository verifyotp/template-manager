"use client";

import * as React from "react"
import { useState } from "react";
import { useTransition } from "react";
import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/new-york/input"
import { signUpRequest as signUp } from "@/actions/signup"
import { useRouter } from 'next/navigation'
import { useToast } from "@/components/ui/use-toast"
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
    </div>
  )
}