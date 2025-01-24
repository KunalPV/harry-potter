"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

type EndGameProps = {
  resetGame: () => void;
  totalQuestions: number;
};

export default function EndGame({
  resetGame,
  totalQuestions,
}: EndGameProps) {
  const router = useRouter();
  const messages = [
    "Fantastic work, wizard! Dumbledore would be proud of you!",
    "Brilliant! Even Hermione would be impressed.",
    "Amazing! You're ready for the Triwizard Tournament!",
  ];
  const randomMessage = messages[Math.floor(Math.random() * messages.length)];

  return (
    <div className="w-full flex flex-col justify-center items-center gap-4 p-4">
      <h1 className="text-3xl font-semibold">Game Over</h1>
      <p className="text-lg">
      You attempted {totalQuestions} questions this time.
      </p>
      <p className="text-xl">{randomMessage}</p>
      <div className="flex gap-4 mt-4 flex-col sm:flex-row">
        <Button size="lg" onClick={resetGame} className="text-lg">
          Restart Trivia
        </Button>
        <Button size="lg" onClick={() => router.push("/")} className="text-lg">
          Return to Home
        </Button>
      </div>
    </div>
  );
}
