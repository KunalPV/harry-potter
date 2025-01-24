"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { Heart, LoaderCircle } from "lucide-react";

type QuestionProps = {
  question: {
    question: string;
    options: string[];
    answer: string;
    difficulty: string;
    type: string;
  };
  lives: number;
  handleIncorrect: () => void;
  handleCorrect: () => void; // Changed from fetchNextQuestion
};

export default function Question({
  question,
  lives,
  handleIncorrect,
  handleCorrect, // Updated prop
}: QuestionProps) {
  const [selectedOption, setSelectedOption] = useState<string | null>(null);
  const [showResult, setShowResult] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [message, setMessage] = useState<string | null>(null);

  const handleSubmit = () => {
    if (!selectedOption) return;

    if (selectedOption === question.answer) {
      setShowResult(true);
      setLoading(true);
      setTimeout(() => {
        setLoading(false);
        setShowResult(false);
        setSelectedOption(null);
        handleCorrect(); // Use handleCorrect instead of fetchNextQuestion
      }, 2000);
    } else {
      setMessage("Oops! Wrong answer. Select the right option to continue.");
      setShowResult(true);
      handleIncorrect();
    }
  };

  const handleOptionClick = (option: string) => {
    if (showResult) {
      setMessage(null);
      setShowResult(false);
    }
    setSelectedOption(option);
  };

  return (
    <Card className="bg-white/30 backdrop-blur-sm border border-white/5 pt-1">
      <CardContent className="p-4">
        <div className="w-full flex flex-col justify-start items-center px-4 gap-4">
          <div>
            <p className="text-lg md:text-xl text-justify">{question.question}</p>
          </div>

          <Separator className="bg-black" />

          <div className="flex flex-col justify-start items-start w-full gap-2 text-lg text-justify">
            {question.options.map((option, index) => (
              <p
                key={index}
                onClick={() => handleOptionClick(option)}
                className={`rounded p-2 cursor-pointer w-full ${
                  selectedOption === option
                    ? showResult
                      ? option === question.answer
                        ? "bg-green-500/20"
                        : "bg-red-500/20"
                      : "bg-zinc-700/20"
                    : showResult && option === question.answer
                    ? "bg-green-500/20"
                    : ""
                }`}
              >
                {option}
              </p>
            ))}
          </div>

          <Separator className="bg-black" />

          {message && (
            <div className="w-full text-center text-red-800 italic mt-2">
              {message}
            </div>
          )}

          <div className="w-full flex justify-between items-center px-2">
            <div className="flex gap-1">
              {Array.from({ length: 3 }).map((_, index) => (
                <Heart
                  key={index}
                  className={`w-5 h-5 ${
                    index < lives 
                      ? "text-red-700 fill-red-700" 
                      : "text-gray-200 fill-gray-200"
                  }`}
                />
              ))}
            </div>

            <div>
              <Button
                size={"lg"}
                className="text-lg flex items-center gap-2"
                disabled={!selectedOption || loading}
                onClick={handleSubmit}
              >
                {loading ? <LoaderCircle className="animate-spin w-5 h-5" /> : "Submit"}
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}