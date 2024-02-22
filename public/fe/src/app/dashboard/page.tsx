import { Metadata } from "next"
import Link from "next/link"

import {Search} from "./_components/search"
import { MainNav } from "./_components/main-nav"
import { UserNav } from "./_components/user-nav"
import TeamSwitcher from "./_components/team-switcher"
import { Dashboard } from "@/components/dashboard/dashboard"

export const metadata: Metadata = {
    title: "Dashboard",
    description: "Manage templates and upload new templates.",
}

export default function DashboardPage() {


    return (
        <div className="hidden flex-col md:flex">
            <div className="border-b sticky y-6">
                <div className="flex h-16 items-center px-4">
                    <TeamSwitcher />
                    <MainNav className="mx-6" />
                    <div className="ml-auto flex items-center space-x-4">
                        <Search />
                        <UserNav />
                    </div>
                </div>
            </div>
            <div>
                <Dashboard />
            </div>
        </div>
    )
}
