"use client"

import { Facts } from "./Facts";
import { Button } from "@/components/ui/button"
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export function Hero() {
  const router = useRouter();
  const [hpMessage, setHpMessage] = useState("");

  useEffect(() => {
    const messages = [
      "The Sorting Hat wonders what you'll learn today!",
      "Ten points to your house for returning!",
      "Dobby is pleased to see you again!",
      "The Hogwarts Express awaits your command...",
      "Constant vigilance! New challenges approach...",
      "Nitwit! Blubber! Oddment! Tweak! (Translation: Let's begin!)"
    ];
    setHpMessage(messages[Math.floor(Math.random() * messages.length)]);
  }, []);

  return(
    <div className="w-full px-4 md:px-20 py-5 md:py-14 flex flex-col justify-center items-start font-im-fell">
      <div className="w-full flex flex-col gap-6">
        <p className="font-bold text-4xl md:text-5xl text-center md:text-start pl-0 md:pl-5">
          Greetings, <span className="text-5xl md:text-6xl tracking-wide animate-pulse">Young Wizard!</span>
        </p>
        <p className="font-semibold text-3xl md:text-4xl text-center md:text-start pl-0 md:pl-5 animate-bounce">
          {hpMessage}
        </p>
      </div>
      <div className="w-full mt-8 md:mt-10">
        <Facts />
      </div>
      <div className="w-full flex justify-center items-center mt-10 md:mt-5 font-magic">
        <div className="w-full flex flex-col md:flex-row justify-center items-center gap-4 md:gap-10">
          <Button 
            size={"lg"} 
            className="text-3xl w-64 text-white/90"
            onClick={() => router.push("/trivia")}
          >
            Trivia
          </Button>
          <Button 
            size={"lg"} 
            className="text-3xl w-64 text-white/90"
            onClick={() => router.push("/characters")}
          >
            Characters
          </Button>
          <Button 
            size={"lg"} 
            className="text-3xl w-64 text-white/90"
            onClick={() => router.push("/spells")}
          >
            Spells
          </Button>
        </div>
      </div>
    </div>
  )
}