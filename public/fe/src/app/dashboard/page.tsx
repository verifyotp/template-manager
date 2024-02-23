import { Metadata } from "next";

import { TopNav } from "./_components/top-nav";
import { SideNav } from "./_components/side-nav";
import { TemplateView } from "./_components/template-view";
import { TemplateSettings } from "./_components/template-settings";

export const metadata: Metadata = {
  title: "Dashboard",
  description: "Manage templates and upload new templates.",
};

export default function DashboardPage() {
  return (
    <div className="flex flex-col items-stretch  h-full">
      <TopNav />
        <div className="flex flex-1 items-stretch justify-between">
            <SideNav />
            <TemplateView />
            <TemplateSettings />
        </div>
    </div>
  );
}
