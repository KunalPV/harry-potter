import { Card, CardContent } from "./ui/card";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Separator } from "./ui/separator";

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

export default function CharacterCard({ character }: { character: Character }) {
  const {
    name,
    species,
    gender,
    house,
    dateOfBirth,
    wizard,
    patronus,
    actor,
    ancestry,
    hairColour,
    eyeColour,
    hogwartsStudent,
    hogwartsStaff,
    alive,
    alternate_names,
    wand,
  } = character;

  const designation = hogwartsStudent
    ? "Student"
    : hogwartsStaff
    ? "Staff"
    : "Unknown";

  const wandDetails = `${wand.wood || "unknown"} wood, ${wand.core || "unknown"} core, ${wand.length || "unknown"} length`;

  return(
    <div className="w-full">
      <Dialog>
        <DialogTrigger className="w-full flex justify-center items-center">
          <Card className="w-80 sm:w-full cursor-pointer hover:shadow-lg" >
            <CardContent className="flex justify-around items-center gap-4 p-4">
              <div className="w-12 h-12 bg-gray-500 rounded-full">

              </div>
              <div className="w-9/12">
                <h2 className="text-xl font-bold truncate">{name}</h2>
              </div>
            </CardContent>
          </Card>
        </DialogTrigger>
        <DialogContent className="max-w-md max-h-screen md:max-w-2xl lg:max-w-3xl rounded-md overflow-auto">
          <DialogHeader>
            <DialogTitle>
              <div className="w-full flex justify-center items-center pb-2">
                <h1 className="text-3xl font-semibold">{name}</h1>
              </div>
            </DialogTitle>

            <Separator />

          </DialogHeader>

          <div className="w-full flex flex-col gap-4 text-lg">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-justify">
              <div className="space-x-1 flex justify-start items-start ">
                <span className="basis-1/2">species: </span>
                <span className="basis-1/2">{species || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">gender: </span>
                <span className="basis-1/2">{gender || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">house: </span>
                <span className="basis-1/2">{house || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">Date-Of-Birth: </span>
                <span className="basis-1/2">{dateOfBirth || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">wizard: </span>
                <span className="basis-1/2">{wizard ? "Yes" : "No / Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">ancestry: </span>
                <span className="basis-1/2">{ancestry || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">eyeColour: </span>
                <span className="basis-1/2">{eyeColour || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">hairColour: </span>
                <span className="basis-1/2">{hairColour || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">patronus: </span>
                <span className="basis-1/2">{patronus || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">designation: </span>
                <span className="basis-1/2">{designation}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">actor: </span>
                <span className="basis-1/2">{actor || "Unknown"}</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span className="basis-1/2">alive: </span>
                <span className="basis-1/2">{alive ? "Yes" : "No / Unknown"}</span>
              </div>
            </div>

            <div className="flex flex-col justify-start items-start gap-4">
              <div className="w-full space-x-1 flex justify-start items-start">
                <span className="basis-3/12">Alternate Names: </span>
                <span className="basis-9/12">
                  {alternate_names?.length
                    ? alternate_names.join(", ")
                    : "None"}
                </span>
              </div>
              <div className="w-full space-x-1 flex justify-start items-start">
                <span className="basis-3/12">wand: </span>
                <span className="basis-9/12">{wandDetails}</span>
              </div>
            </div>
          </div>
        </DialogContent>
      </Dialog>

    </div>
  )
}