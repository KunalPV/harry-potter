"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { ChevronLeft } from "lucide-react";
import { useRouter } from "next/navigation";
import Question from "@/components/Question";
import EndGame from "@/components/EndGame";

type QuestionData = {
  question: string;
  options: string[];
  answer: string;
  difficulty: string;
  type: string;
};

export default function Trivia() {
  const router = useRouter();
  const [question, setQuestion] = useState<QuestionData | null>(null);
  const [lives, setLives] = useState(3);
  const [gameOver, setGameOver] = useState(false);
  const [totalQuestions, setTotalQuestions] = useState(0);

  const fetchQuestion = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/game/question");
      if (!response.ok) throw new Error("Failed to fetch question.");
      const data: QuestionData = await response.json();
      setQuestion(data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleEndGame = () => setGameOver(true);

  const handleIncorrect = () => {
    setLives((prev) => {
      const newLives = prev - 1;
      if (newLives === 0) setGameOver(true);
      return newLives;
    });
  };

  const handleCorrect = () => {
    setTotalQuestions(prev => prev + 1);
    fetchQuestion();
  };

  const resetGame = () => {
    setLives(3);
    setTotalQuestions(0);
    setGameOver(false);
    fetchQuestion();
  };

  useEffect(() => {
    fetchQuestion();
  }, []);

  if (gameOver) {
    return <EndGame resetGame={resetGame} totalQuestions={totalQuestions} />;
  }

  return (
    <div className="w-full p-2 mt-2 mb-6">
      <div className="w-full flex flex-col justify-center items-start gap-4">
        <div className="flex justify-between items-center w-full">
          <Button 
            variant="outline" 
            size="icon" 
            onClick={() => router.back()} 
            className="bg-white/30 backdrop-blur-sm border border-white/5"
          >
            <ChevronLeft />
          </Button>
          <h1 className="text-2xl font-semibold">Trivia Game</h1>
          <div></div>
        </div>

        <Separator />

        <div className="w-full">
          {question && (
            <Question
              question={question}
              lives={lives}
              handleIncorrect={handleIncorrect}
              handleCorrect={handleCorrect}
            />
          )}
        </div>

        <div className="w-full flex justify-center items-center">
          <Button size={"lg"} className="text-xl w-64" onClick={handleEndGame}>
            End Trivia
          </Button>
        </div>
      </div>
    </div>
  );
}