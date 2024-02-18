import type { Metadata } from "next";
import "./globals.css";
import { Inter as FontSans } from "next/font/google";
import { Toaster } from "@/components/ui/toaster"
import { cn } from "../lib/utils";

const fontSans = FontSans({
  subsets: ["latin"],
  variable: "--font-sans",
})

export const metadata: Metadata = {
  title: "Template Manager",
  description: "Manage your templates with ease",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
       className={cn(
        "min-h-screen bg-background font-sans antialiased",
        fontSans.variable
      )}  
       >
        <main>{children}</main>
       <Toaster />
       </body>
    </html>
  );
}
