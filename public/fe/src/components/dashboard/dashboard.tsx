"use client";

import * as React from "react"
import { cn } from "@/lib/utils"
import { ViewTemplate } from "@/components/dashboard/template/view"
import { UploadTemplate } from "@/components/dashboard/template/upload"
import { ApiKeySheet } from "./api-key/sheet";
import { DocsSidebarNav } from "./template/nav";


const items = [
  {
    title: "Email Templates",
    items: [
      {
        items: [],
        title: "View Templates",
        href: "/dashboard/template/view",
      },
      {
        items: [],
        title: "Upload Templates",
        href: "/dashboard/template/upload",
      },
      {
        items: [],
        title: "Upload Templates",
        href: "/dashboard/template/upload",
      },
      {
        items: [],
        title: "Upload Templates",
        href: "/dashboard/template/upload",
      },
      {
        items: [],
        title: "Upload Templates",
        href: "/dashboard/template/upload",
      },
      {
        items: [],
        title: "Upload Templates",
        href: "/dashboard/template/upload",
      },
    ]
  },
]

interface ViewProps extends React.HTMLAttributes<HTMLDivElement> { }

export function Dashboard({ className, ...props }: ViewProps) {
  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <div className="container v-screen">
        <div>
          <DocsSidebarNav items={items} />
        </div>
        <div>
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
            {/* Create API Key */}
            <div>
              <ApiKeySheet />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

