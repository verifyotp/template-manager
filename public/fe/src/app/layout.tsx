import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Inter as FontSans } from "next/font/google";

// const inter = Inter({ subsets: ["latin"] });

import { cn } from "../lib/utils";

export const fontSans = FontSans({
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
      //  className={inter.className}

       className={cn(
        "min-h-screen bg-background font-sans antialiased",
        fontSans.variable
      )}
       >{children}</body>
    </html>
  );
}
