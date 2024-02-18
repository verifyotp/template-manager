"use client";

import * as React from "react"
import { useRouter } from 'next/navigation'

import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Button } from "@/registry/new-york/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/registry/new-york/ui/label"

import { buttonVariants } from "@/registry/new-york/ui/button"

interface ViewApiKeyProps extends React.HTMLAttributes<HTMLDivElement> { }

export function ViewApiKey({ className, ...props }: ViewApiKeyProps) {
    const [isLoading, setIsLoading] = React.useState<boolean>(false)
    const router = useRouter(); // Initialize useRouter

    async function onSubmit(event: React.SyntheticEvent) {
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
        <form onSubmit={onSubmit}>
          <div className="grid gap-3">
            <div className="grid gap-3">
              <Label className="sr-only" htmlFor="email">
                View API Keys
              </Label>
             </div>
  
             <div className="mb-4">
            <h2 className="text-lg font-semibold mb-2">Email Templates</h2>
            {/* Replace with logic to display email templates */}
            <ul>
              <li>key 1</li>
              <li>key 2</li>
              {/* Add more templates as needed */}
            </ul>
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

