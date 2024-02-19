"use client";

import * as React from "react"
import { useRouter } from 'next/navigation'

import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/registry/new-york/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"

import { buttonVariants } from "@/registry/new-york/ui/button"

interface UploadTemplateProps extends React.HTMLAttributes<HTMLDivElement> { }

export function UploadTemplate({ className, ...props }: UploadTemplateProps) {
    const [isLoading, setIsLoading] = React.useState<boolean>(false)
    const router = useRouter(); // Initialize useRouter

    async function generateUploadLink(event: React.SyntheticEvent) {
        event.preventDefault()
        setIsLoading(true)

        setTimeout(() => {
            setIsLoading(false)
            // Navigate to the dashboard on successful login
            router.push('/dashboard');
        }, 30)
    }

    return (
        <div className={cn("grid gap-6", className)} {...props}>
            <div className="grid gap-3">
                <div className="grid gap-3">
                    <Label className="sr-only" htmlFor="email">
                        Upload Template
                    </Label>
                </div>

                <div>
                    <h2 className="text-lg font-semibold mb-2">Upload Template</h2>
                    <button
                        onClick={generateUploadLink}
                        className={cn(
                            buttonVariants({ variant: "ghost" }),
                            "py-2 px-4 rounded-md shadow-md"
                        )}
                    >
                        Upload Template
                    </button>
                </div>
            </div>
        </div>
    )
}

