import { ViewTemplate } from "@/components/dashboard/template/playground";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { DocsSidebarNav } from "@/components/dashboard/template/nav";

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

export const SideNav = () => {
  return (
    <div className="max-w-[250px]  p-1 border-slate-200 relative  overflow-hidden flex flex-col flex-1 border-r max-h-full">
      <div className="inset-0 absolute  overflow-y-auto">
        <div>
          <DocsSidebarNav items={items} />
        </div>
        <div>
          <DocsSidebarNav items={items} />
        </div>
        <div>
          <DocsSidebarNav items={items} />
        </div>
        <div>
          <DocsSidebarNav items={items} />
        </div>
      </div>
    </div>
  );
};
