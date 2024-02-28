import { MainNav } from "./main-nav"
import { UserNav } from "./user-nav"

export const TopNav = () => {
    return <div className="border-b bg-background">
    <div className="flex h-16 items-center px-4">
        <MainNav className="mx-6" />
        <div className="ml-auto flex items-center space-x-4">
            <UserNav email="name@example.com" />
        </div>
    </div>
</div>
}