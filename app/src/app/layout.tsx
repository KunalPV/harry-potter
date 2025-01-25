import type { Metadata } from "next";
import { Geist, Geist_Mono, MedievalSharp, IM_Fell_English } from "next/font/google";
import "./globals.css";
import { NavBar } from "@/components/NavBar";
import MaxWidthWrapper from "@/components/MaxWidthWrapper";
import Footer from "@/components/Footer";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const medieval = MedievalSharp({
  weight: '400',
  subsets: ['latin'],
  variable: '--font-medieval'
});

const imFellEnglish = IM_Fell_English({
  weight: '400',
  subsets: ['latin'],
  variable: '--font-im-fell'
});

export const metadata: Metadata = {
  title: "Potterhead Puzzles",
  description: "🧙♂️⚡ Test your wizarding wisdom with Potterhead Puzzles...",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <meta name="apple-mobile-web-app-title" content="Potterhead Puzzles" />
      </head>
      <body
        className={`${geistSans.variable} ${geistMono.variable} ${medieval.variable} ${imFellEnglish.variable} antialiased flex flex-col min-h-screen`}
      >
        <header>
          <NavBar />
        </header>
        <MaxWidthWrapper className="flex-grow">
          {children}
        </MaxWidthWrapper>
        <footer>
          <Footer />
        </footer>
      </body>
    </html>
  );
}
