import { Card, CardContent } from "./ui/card";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Separator } from "./ui/separator";

export default function CharacterCard() {
  return(
    <div className="w-full">
      <Dialog>
        <DialogTrigger>
          <Card className="max-w-72 w-96 sm:w-fit cursor-pointer hover:shadow-lg" >
            <CardContent className="flex justify-around items-center gap-4 p-4">
              <div className="w-12 h-12 bg-gray-500 rounded-full">

              </div>
              <div className="w-9/12">
                <h2 className="text-xl font-bold truncate">Harry Potter</h2>
              </div>
            </CardContent>
          </Card>
        </DialogTrigger>
        <DialogContent className="max-w-md max-h-screen md:max-w-2xl lg:max-w-3xl rounded-md overflow-auto">
          <DialogHeader>
            <DialogTitle>
              <div className="w-full flex justify-center items-center pb-2">
                <h1 className="text-3xl font-semibold">Harry Potter</h1>
              </div>
            </DialogTitle>

            <Separator />

          </DialogHeader>

          <div className="w-full flex flex-col gap-4 text-lg">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-justify">
              <div className="space-x-1 flex justify-start items-start">
                <span>species: </span>
                <span>human</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>gender: </span>
                <span>male</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>house: </span>
                <span>Gryffindor</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>dateOfBirth: </span>
                <span>31-07-1980</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>wizard: </span>
                <span>true</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>ancestry: </span>
                <span>half-blood</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>eyeColour: </span>
                <span>green</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>hairColour: </span>
                <span>black</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>patronus: </span>
                <span>stag</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>designation: </span>
                <span>student</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>actor: </span>
                <span>Daniel Radcliffe</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>alive: </span>
                <span>true</span>
              </div>
            </div>

            <div className="flex flex-col justify-start items-start gap-4">
              <div className="space-x-1 flex justify-start items-start ">
                <span>names: </span>
                <span>The Boy Who Lived, The Chosen One, Undesirable No. 1, Potty</span>
              </div>
              <div className="space-x-1 flex justify-start items-start">
                <span>wand: </span>
                <span>holly, phoenix tail feather, 11</span>
              </div>
            </div>
          </div>
        </DialogContent>
      </Dialog>

    </div>
  )
}