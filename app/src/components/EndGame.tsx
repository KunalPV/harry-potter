"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import { Card } from "./ui/card";

type EndGameProps = {
  resetGame: () => void;
  totalQuestions: number;
};

export default function EndGame({ resetGame, totalQuestions }: EndGameProps) {
  const router = useRouter();
  const messages = [
    "Fantastic work, wizard! Dumbledore would be proud of you!",
    "Brilliant! Even Hermione would be impressed.",
    "Merlin's beard! That's more correct than Lockhart's published lies!",
    "Accio genius! You've caught the Golden Snitch of knowledge!",
    "Ten points to your house! The Sorting Hat would struggle to place such brilliance!",
    "You've mastered the Patronus Charm of trivia - expect an owl from the Ministry!",
    "Fantastic! Nearly Headless Nick would give you a standing ovation (if he could!)",
    "Brilliant! You've brewed a perfect Polyjuice Potion of knowledge!",
    "You've found all the Horcruxes of wisdom! Dobby would gift you socks!",
    "Prefect material! McGonagall would award 50 house points!"
  ];
  const randomMessage = messages[Math.floor(Math.random() * messages.length)];

  return (
    <div className="w-full mt-4 flex justify-center items-center">
      <Card className="min-w-full flex flex-col justify-center items-center gap-4 p-4 bg-white/30 backdrop-blur-sm border border-white/5 font-im-fell">
        <h1 className="text-4xl font-bold">Game Over</h1>
        <p className="text-2xl font-semibold">You attempted {totalQuestions} questions this time.</p>
        <p className="text-2xl font-semibold text-center">{randomMessage}</p>
        <div className="flex gap-4 mt-4 flex-col sm:flex-row font-magic">
          <Button size="lg" onClick={resetGame} className="text-3xl text-white/90">
            Restart Trivia
          </Button>
          <Button size="lg" onClick={() => router.push("/")} className="text-3xl text-white/90">
            Return to Home
          </Button>
        </div>
      </Card>
    </div>
  );
}