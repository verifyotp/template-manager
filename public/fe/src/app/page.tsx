
import { Metadata } from "next";
import Link from "next/link";
import { cn } from "@/lib/utils";
import { buttonVariants } from "@/components/ui/button";
import React from "react";

export const metadata: Metadata = {
  title: "Template Manager",
  description: "Manage your templates with ease",
};

const FeatureCard = ({ children }: React.PropsWithChildren<{}>) => (
  <div className="p-4 bg-gray-50 rounded-lg shadow-md mb-4">
    {children}
  </div>
);

export default function LandingPage() {
  return (
    <div className="container h-screen bg-white flex flex-col justify-center items-center relative">
      {/* Animated scrollable text */}
      {/* <Scroll /> */}
      <h1 className="text-5xl font-bold mb-4">Template Manager</h1>
      <p className="text-xl text-gray-700 mb-8">Create and manage your templates for email, sms, push and more.</p>
      <div className="flex space-x-4">
        <Link
          href="/auth/signup"
          className={cn(
            buttonVariants({ variant: "default" }),
            "py-3 px-4 rounded-md shadow-md text-white"
          )}
        >
          Get Started
        </Link>
        <Link
          href="/auth/login"
          className={cn(
            buttonVariants({ variant: "secondary" }),
            "py-3 px-4 rounded-md shadow-md text-gray-800"
          )}
        >
          Login
        </Link>
      </div>
    </div>
  );
}
