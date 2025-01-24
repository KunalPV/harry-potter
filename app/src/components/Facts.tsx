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
        const response = await fetch("http://localhost:8080/api/facts");
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
      <Card className="w-full flex justify-center items-center flex-col p-2 md:p-0">
        <CardHeader className="p-2 md:p-4">
          <CardTitle className="text-2xl">Fact of the Day.</CardTitle>
        </CardHeader>

        <Separator className="mx-4 w-[95%]" />

        <CardContent className="p-2 md:p-4 text-balance text-center">
          <p className="text-xl">{fact}</p>
        </CardContent>
      </Card>
    </div>
  )
}