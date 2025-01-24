"use client"

import SpellCard from "@/components/SpellCard";
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

type Spell = {
  id: string;
  name: string;
  description: string;
};

export default function Spells() {
  const [spells, setSpells] = useState<Spell[]>([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const fetchSpells = async () => {
      setLoading(true);
      try {
        const response = await fetch(`http://localhost:8080/api/spells?page=${page}&limit=10`);
        if (!response.ok) {
          throw new Error(`Error fetching spells: ${response.statusText}`);
        }
        const data = await response.json();
        setSpells(data.data);
        setTotalPages(data.totalPages);
      } catch (error) {
        if (error instanceof Error) {
          console.error(error.message);
        } else {
          console.error("An unknown error occurred.");
        }
      } finally {
        setLoading(false);
      }
    };

    fetchSpells();
  }, [page]);

  return(
    <div className="w-full">
      <div className="w-full flex flex-col justify-center items-center">
        <div className="flex justify-between items-center w-full p-4">
          <Button variant="outline" size="icon" onClick={() => router.back()} className="bg-white/30 backdrop-blur-sm border border-white/5">
            <ChevronLeft />
          </Button>
          <h1 className="text-2xl font-semibold px-4">Spell Book</h1>
          <div></div>
        </div>

        <Separator className="bg-black" />

        {loading ? (
          <div className="text-center">Loading...</div>
        ) : (
          <>
            <div className="p-4 w-full flex justify-center items-center">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full">
                {spells.map((spell) => (
                  <SpellCard key={spell.id} spell={spell} />
                ))}
              </div>
            </div>

            <div className="w-full flex justify-center items-center gap-4 mt-2 mb-10">
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
                        className="bg-white/30 backdrop-blur-sm border border-white/5"
                      />
                    ) : (
                      <PaginationPrevious
                        href="#"
                        className="pointer-events-none opacity-50 bg-white/10 backdrop-blur-sm border border-white/5"
                        aria-disabled="true"
                      />
                    )}
                  </PaginationItem>

                  {/* Dynamic Page Numbers */}
                  {Array.from({ length: totalPages }, (_, index) => index + 1)
                    .filter(
                      (pageNumber) =>
                        pageNumber === 1 ||
                        pageNumber === totalPages ||
                        (pageNumber >= page - 2 && pageNumber <= page + 2)
                    )
                    .map((pageNumber, index, filteredPages) => {
                      const prevPage = filteredPages[index - 1];
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
                            className="bg-white/30 backdrop-blur-sm border border-white/5"
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
                        className="bg-white/30 backdrop-blur-sm border border-white/5"
                        onClick={(e) => {
                          e.preventDefault();
                          setPage(page + 1);
                        }}
                      />
                    ) : (
                      <PaginationNext
                        href="#"
                        className="pointer-events-none bg-white/10 backdrop-blur-sm border border-white/5"
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