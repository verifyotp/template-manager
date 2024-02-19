import { Metadata } from "next"
import Link from "next/link"


import { Dashboard } from "@/components/dashboard/dashboard"

export const metadata: Metadata = {
    title: "Dashboard",
    description: "Manage templates and upload new templates.",
}

export default function DashboardPage() {


    return (
        <div className="container h-screen flex justify-center items-center">
            <Dashboard />
        </div>
    )
}
