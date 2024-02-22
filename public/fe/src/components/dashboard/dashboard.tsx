"use client";

import * as React from "react";
import { cn } from "@/lib/utils";
import { ViewTemplate } from "@/components/dashboard/template/playground";
import { Separator } from "@/components/ui/separator"
import { ApiKeySheet } from "./api-key/sheet";
import { DocsSidebarNav } from "./template/nav";
import {Button } from "@/components/ui/button";

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
    ],
  },
];


export function NavBar() {
  return (
    <div>
      <div className="space-y-1">
        <h4 className="text-sm font-medium leading-none py-[5px]">Dashboard</h4>
        <p className="text-sm text-muted-foreground">
          Manage All your Templates In one Place
        </p>
      </div>
      <Separator className="my-4" />
      <div className="flex h-5 items-center space-x-4 text-sm">
        <Button variant="link">Integrations</Button>
        <Separator orientation="vertical" />
        <Separator orientation="vertical" />
        <Button variant="link">Profile</Button>
        <Separator orientation="vertical" />
        <Button variant="link">Subscription</Button>
        <Separator orientation="vertical" />
       
        <Button variant="link" className="hover:text-red-500">Logout</Button>
      </div>
    </div>
  )
}


interface ViewProps extends React.HTMLAttributes<HTMLDivElement> { }

export function Dashboard({ className, ...props }: ViewProps) {
  return (
    <div className={cn("flex flex-row", className)} {...props}>
      <div className="container items-stretch flex-1">

        <div className="sticky top-3 flex-1 p-[10px] w-full">
          <NavBar />
        </div>

        

        <div className="max-w-[200px] flex-1 border-r border-slate-100 p-6 pt-[25px] h-full">
          <DocsSidebarNav items={items} />
        </div>

      </div>
    </div>
  );
}
