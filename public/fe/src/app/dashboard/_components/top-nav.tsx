import { Search } from "./search"
import { MainNav } from "./main-nav"
import { UserNav } from "./user-nav"
import TeamSwitcher from "./team-switcher"

export const TopNav = () => {
    return <div className="border-b bg-background">
    <div className="flex h-16 items-center px-4">
        <TeamSwitcher />
        <MainNav className="mx-6" />
        <div className="ml-auto flex items-center space-x-4">
            <Search />
            <UserNav />
        </div>
    </div>
</div>
}