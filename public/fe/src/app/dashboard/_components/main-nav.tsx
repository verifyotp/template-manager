import Link from "next/link"

import { cn } from "@/lib/utils"
import { ApiKeySheet } from "@/components/dashboard/api-key/sheet"

export function MainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <nav
      className={cn("flex items-center space-x-4 lg:space-x-6", className)}
      {...props}
    >
      <Link
        href=""
        className="text-sm font-medium transition-colors hover:text-primary"
      >
        Templates
      </Link>
      <Link
        href=""
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Integrations
      </Link>
      <Link
        href=""
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        <ApiKeySheet >
        <div>
          Settings
        </div>
        </ApiKeySheet>
      </Link>
    </nav>
  )
}