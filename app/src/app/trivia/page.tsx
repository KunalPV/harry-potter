"use client"

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { ChevronLeft, Heart } from "lucide-react";
import { useRouter } from "next/navigation";

export default function Trivia() {
  const router = useRouter();

  return(
    <div className="w-full p-2 mt-2 mb-6">
      <div className="w-full flex flex-col justify-center items-start gap-4">
        <div className="flex justify-between items-center w-full">
          <Button variant="outline" size="icon" onClick={() => router.back()}>
            <ChevronLeft />
          </Button>
          <h1 className="text-2xl font-semibold">Trivia Game</h1>
          <div></div>
        </div>

        <Separator />

        <div className="w-full">
          <Card className="">
            <CardContent className="p-4">
              <div className="w-full flex flex-col justify-start items-center px-4 gap-4">
                <div>
                  <p className="text-lg text-justify">Lorem ipsum dolor sit amet consectetur adipisicing elit. Placeat repellat quam rem culpa distinctio enim cupiditate neque consectetur deleniti ipsum? Lorem ipsum dolor sit amet consectetur adipisicing elit. Iste fugit vitae dolore. Deleniti ipsam voluptate quo iste quas laudantium doloremque itaque est perspiciatis. </p>
                </div>

                <Separator />

                <div className="flex flex-col justify-start items-start w-full gap-2 text-lg text-justify">
                  <p className="bg-green-100 rounded p-2 cursor-pointer">Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ex incidunt consequuntur soluta praesentium nobis quis ut nihil, cupiditate deleniti dignissimos impedit dicta mollitia harum, iure corporis aspernatur quidem accusamus! Perspiciatis?</p>
                  <p className="bg-red-100 rounded p-2 cursor-pointer">Lorem ipsum dolor sit amet consectetur adipisicing elit. Aliquid perferendis earum, nam asperiores neque tempora illum laboriosam officia minima recusandae!</p>
                  <p className="bg-slate-50 rounded p-2 cursor-pointer">Option 3</p>
                  <p className="bg-slate-100 rounded p-2 cursor-pointer">Option 4</p>
                </div>

                <Separator />

                <div className="w-full flex justify-between items-center px-2">
                  <div className="flex gap-1">
                    <Heart className="w-5 h-5" />
                    <Heart className="w-5 h-5" />
                    <Heart className="w-5 h-5" />
                  </div>

                  <div>
                    <Button size={"lg"} className="text-lg">Submit</Button>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="w-full flex justify-center items-center">
        <Button size={"lg"} className="text-xl w-64">End Trivia</Button>
        </div>
      </div>
    </div>
  )
}