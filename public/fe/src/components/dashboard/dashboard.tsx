"use client";

import * as React from "react";
import { cn } from "@/lib/utils";
import { ViewTemplate } from "@/components/dashboard/template/playground";
import { Separator } from "@/components/ui/separator"
import { ApiKeySheet } from "../settings/api-key/sheet";
import { DocsSidebarNav } from "./template/nav";
import { Button } from "@/components/ui/button";

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

interface ViewProps extends React.HTMLAttributes<HTMLDivElement> { }

export function Dashboard({ className, ...props }: ViewProps) {
  return (
    <div className={cn("container flex", className)} {...props}>
      <div className="flex w-full items-stretch h-full justify-between">
        <div className="max-w-[200px] border-slate-100  items-stretch flex-1 border-r p-1" >
          <div className=" pt-[25px] ">
            <DocsSidebarNav items={items} />
          </div>
          <div className="pt-[25px]">
            <DocsSidebarNav items={items} />
          </div>
          <div className="pt-[25px]">
            <DocsSidebarNav items={items} />
          </div>
          <div className="pt-[25px]">
            <DocsSidebarNav items={items} />
          </div>
        </div>
        <div className="max-w-[200px] border-slate-100  items-stretch flex-1 border p-1" >
        <div className=" pt-[25px] justify-center text-center ">
          Template View
        </div>
        </div>
        <div className="max-w-[200px] border-slate-100  items-stretch flex-1 border-l p-1" >
          <div className=" pt-[25px]  justify-center text-center ">
            Template Settings
          </div>
        </div>
      </div>
    </div>
  );
}
