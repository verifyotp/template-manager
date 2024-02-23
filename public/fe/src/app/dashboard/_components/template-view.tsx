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

export const TemplateView = () => {
  return (
    <div className="relative  overflow-hidden  items-stretch flex-1  p-1">
      <div className=" pt-[25px] justify-center text-center inset-0 absolute  overflow-y-auto">
        Template View
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
