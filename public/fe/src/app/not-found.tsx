import Link from 'next/link'
import { cn } from "@/lib/utils";
import { buttonVariants } from "@/components/ui/button";

export default function NotFound() {
    return (
        <div className="container h-screen flex flex-col bg-white justify-center items-center relative">
            <div className="max-h-100 max-w-100 justify-center">
            <img src="/404.jpeg" alt="404" className="w-96" />
            </div>
            <div className="justify-center flex flex-col items-center relative">
            <h1 className="text-5xl font-bold mb-4">404</h1>
            <p className="text-xl text-gray-700 mb-8">Page not found</p>
            <Link
                href="/auth/login"
                className={cn(
                    buttonVariants({ variant: "default" }),
                    "py-3 px-4 rounded-md shadow-md text-white"
                )}
            >
                Home Page
            </Link>
            </div>
        </div>
  )
}