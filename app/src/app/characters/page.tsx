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

export default function Characters() {
  return(
    <div className="w-full">
      <div className="w-full flex flex-col justify-center items-center">
        <div className="flex justify-between items-center w-full p-4">
          <Button variant="outline" size="icon" className="p-2">
            <ChevronLeft />
          </Button>
          <h1 className="text-2xl font-semibold px-4">Characters from Harry Potter</h1>
          <div></div>
        </div>

        <Separator />

        <div className="p-4 w-full flex justify-center items-center">
          <div className="grid grid-cols-1 md:grid-cols-3 sm:grid-cols-2 gap-6">
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
            <CharacterCard />
          </div>
        </div>

        <div className="w-full mt-2 mb-10">
          <Pagination>
            <PaginationContent>
              <PaginationItem>
                <PaginationPrevious href="#" />
              </PaginationItem>
              <PaginationItem>
                <PaginationLink href="#">1</PaginationLink>
              </PaginationItem>
              <PaginationItem>
                <PaginationLink href="#" isActive>
                  2
                </PaginationLink>
              </PaginationItem>
              <PaginationItem>
                <PaginationLink href="#">3</PaginationLink>
              </PaginationItem>
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
              <PaginationItem>
                <PaginationNext href="#" />
              </PaginationItem>
            </PaginationContent>
          </Pagination>
        </div>
      </div>
    </div>
  )
}