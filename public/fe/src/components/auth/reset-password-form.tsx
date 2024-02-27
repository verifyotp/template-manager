"use client";

import * as React from "react"
import { useState } from "react";
import { cn } from "@/lib/utils"
import { Icons } from "@/components/ui/icons"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/new-york/input"
import { Label } from "@/components/ui/new-york/label"
import { useTransition } from "react";
import { useRouter } from 'next/navigation'
import { useForm } from "react-hook-form";
import { initiatePasswordReset } from "@/actions/reset-password"
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormError } from "@/components/form-error";
import { FormSuccess } from "@/components/form-success";
import { ResetSchema } from "@/schemas";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";




interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UserResetPasswordForm({ className, ...props }: UserAuthFormProps) {

  const router = useRouter(); // Initialize useRouter
  const [isPending, startTransition] = useTransition();
  const [error, setError] = useState<string | undefined>("");
  const [success, setSuccess] = useState<string | undefined>("");

  const form = useForm<z.infer<typeof ResetSchema>>({
    resolver: zodResolver(ResetSchema),
    defaultValues: {
      email: "",
    },
  });


  const onSubmit = (values: z.infer<typeof ResetSchema>) => {
    setError("");
    setSuccess("");

    startTransition(() => {
      initiatePasswordReset(values.email)
        .then((data) => {
          if (data?.status === false) {
            form.reset();
            setError(data.message);
          }

          if (data?.status) {
            form.reset();
            setSuccess(data.message);
            // delay redirect to login page
            setTimeout(() => {
              router.push("/auth/login");
            }, 6000);
          }
        })
        .catch(() => setError("Something went wrong"));
    });
  };

  return (
    <div className={cn("grid gap-6", className)} {...props}>
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
              Reset Password
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