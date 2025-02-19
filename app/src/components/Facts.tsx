"use client"
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { useEffect, useState } from "react";

export function Facts() {
  const [fact, setFact] = useState<string>("Loading today's magical fact...");

  useEffect(() => {
    const fetchFact = async () => {
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/facts`);
        if (!response.ok) {
          throw new Error("Failed to fetch the fact");
        }
        const data = await response.json();
        setFact(data.fact); // Assuming the API returns { fact: "..." }
      } catch (error) {
        console.error("Error fetching fact:", error);
        setFact("Unable to fetch today's fact. Please try again later.");
      }
    };

    fetchFact();
  }, []);

  return(
    <div className="w-full flex justify-center items-center p-0 md:p-4">
      <Card className="w-full flex justify-center items-center flex-col p-2 md:p-0 bg-white/30 backdrop-blur-sm border border-white/5">
        <CardHeader className="p-2 md:p-4">
          <CardTitle className="text-3xl font-bold font-im-fell">Fact of the Day.</CardTitle>
        </CardHeader>

        <Separator className="mx-4 w-[95%] bg-black" />

        <CardContent className="p-2 md:p-4 text-balance text-center">
          <p className="text-2xl font-medieval font-bold">{fact}</p>
        </CardContent>
      </Card>
    </div>
  )
}