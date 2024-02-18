"use client";

import * as React from "react"
import { cn } from "@/lib/utils"
import { ViewTemplate } from "@/components/dashboard/template/view"
import { UploadTemplate } from "@/components/dashboard/template/upload"

interface ViewProps extends React.HTMLAttributes<HTMLDivElement> { }

export function Dashboard({ className, ...props }: ViewProps) {
  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <div className="container h-screen flex justify-center items-center">
        <div className="max-w-lg p-8 bg-white rounded-lg shadow-lg">
          <h1 className="text-3xl font-semibold mb-4">Dashboard</h1>
          {/* List of email templates */}
          <div>
            <ViewTemplate />
          </div>
          {/* Upload template button */}
          <div>
            <UploadTemplate />
          </div>
        </div>
      </div>
    </div>
  )
}

