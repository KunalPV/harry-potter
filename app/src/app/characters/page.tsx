"use client"

import CharacterCard from "@/components/CharacterCard";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { ChevronLeft } from "lucide-react";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type Character = {
  id: string;
  name: string;
  species: string;
  gender: string;
  house: string;
  dateOfBirth: string;
  patronus?: string;
  actor: string;
  ancestry?: string;
  eyeColour?: string;
  hairColour?: string;
  hogwartsStudent: boolean;
  hogwartsStaff: boolean;
  alive: boolean;
  wizard: boolean;
  alternate_names?: string[];
  wand: {
    wood?: string;
    core?: string;
    length?: number;
  };
};

export default function Characters() {
  const [characters, setCharacters] = useState<Character[]>([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const fetchCharacters = async () => {
      setLoading(true);
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/characters?page=${page}&limit=12`);
        if (!response.ok) {
          throw new Error(`Error fetching characters: ${response.statusText}`);
        }
        const data = await response.json();
        setCharacters(data.data); // Array of characters
        setTotalPages(data.totalPages); // Total number of pages
      } catch (error) {
        if (error instanceof Error) {
          console.error(error.message);
        } else {
          console.error("An unknown error occurred.");
        }
      } finally {
        setLoading(false); // Stop loading
      }
    };

    fetchCharacters();
  }, [page]);

  return(
    <div className="w-full">
      <div className="w-full flex flex-col justify-center items-center">
        <div className="flex justify-between items-center w-full py-4">
          <Button variant="outline" size="icon" className="p-2 bg-white/30 backdrop-blur-sm border border-white/5" onClick={() => router.back()}>
            <ChevronLeft />
          </Button>
          <h1 className="text-5xl font-bold font-harry tracking-wide px-2">Character Book</h1>
          <div></div>
        </div>

        <Separator className="bg-black" />

          {loading ? (
            <div className="text-center">Loading...</div>
          ) : (
            <>
            <div className="p-4 w-full flex justify-center items-center">
              <div className="w-full grid grid-cols-1 md:grid-cols-3 sm:grid-cols-2 gap-6">
                {characters.map((character) => (
                  <CharacterCard key={character.id} character={character} />
                ))}
              </div>
            </div>

            <div className="w-full mt-2 mb-10">
              <Pagination>
                <PaginationContent>
                  {/* Previous Button */}
                  <PaginationItem>
                    {page > 1 ? (
                      <PaginationPrevious
                        href={`?page=${page - 1}`}
                        onClick={(e) => {
                          e.preventDefault();
                          setPage(page - 1);
                        }}
                        className="bg-white/30 backdrop-blur-sm border border-white/5 font-medieval text-lg font-semibold"
                      />
                    ) : (
                      <PaginationPrevious
                        href="#"
                        className="pointer-events-none bg-white/10 backdrop-blur-sm border border-white/5 font-medieval text-lg font-semibold"
                        aria-disabled="true"
                      />
                    )}
                  </PaginationItem>

                  {/* Dynamic Page Numbers */}
                  {Array.from({ length: totalPages }, (_, index) => index + 1)
                    .filter((pageNumber) => {
                      // Show the first, last, current, and adjacent pages, plus ellipses
                      return (
                        pageNumber === 1 ||
                        pageNumber === totalPages ||
                        (pageNumber >= page - 2 && pageNumber <= page + 2)
                      );
                    })
                    .map((pageNumber, index, filteredPages) => {
                      const prevPage = filteredPages[index - 1];

                      // Add ellipsis where there is a gap in the page sequence
                      if (prevPage && pageNumber - prevPage > 1) {
                        return (
                          <PaginationItem key={`ellipsis-${pageNumber}`}>
                            <PaginationEllipsis />
                          </PaginationItem>
                        );
                      }

                      return (
                        <PaginationItem key={pageNumber}>
                          <PaginationLink
                            href={`?page=${pageNumber}`}
                            isActive={pageNumber === page}
                            className="bg-white/30 backdrop-blur-sm border border-white/5 font-medieval text-lg font-semibold"
                            onClick={(e) => {
                              e.preventDefault();
                              setPage(pageNumber);
                            }}
                          >
                            {pageNumber}
                          </PaginationLink>
                        </PaginationItem>
                      );
                    })}

                    {/* Next Button */}
                    <PaginationItem>
                      {page < totalPages ? (
                        <PaginationNext
                          href={`?page=${page + 1}`}
                          onClick={(e) => {
                            e.preventDefault();
                            setPage(page + 1);
                          }}
                          className="bg-white/30 backdrop-blur-sm border border-white/5 font-medieval text-lg font-semibold"
                        />
                      ) : (
                        <PaginationNext
                          href="#"
                          className="pointer-events-none bg-white/10 backdrop-blur-sm border border-white/5 font-medieval text-lg font-semibold"
                          aria-disabled="true"
                        />
                      )}
                    </PaginationItem>
                  </PaginationContent>
                </Pagination>
              </div>
            </>
          )}

        
      </div>
    </div>
  )
}