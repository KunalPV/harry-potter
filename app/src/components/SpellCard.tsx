import Image from "next/image";
import { Card, CardContent } from "./ui/card";
import { Separator } from "./ui/separator";

type Spell = {
  id: string;
  name: string;
  description: string;
};

export default function SpellCard({ spell }: { spell: Spell }) {
  return(
    <div className="w-full">
      <Card className="w-full min-h-full bg-white/50 backdrop-blur-sm border border-white/5" >
        <CardContent className="flex justify-around items-center gap-2 px-6 py-2 flex-col font-bold">
          <div className="w-full flex justify-between items-center">
            <div>
              <p className="text-3xl text-justify font-im-fell animate-bounce">{spell.name}</p>
            </div>

            <div className="overflow-hidden">
              <Image 
                src="/spell_logo.png"
                alt="Spell logo"
                width={36}  // w-8 = 32px (8 * 4)
                height={36}
                className="w-full h-full object-contain p-1"
              />
            </div>
          </div>

          <Separator className="bg-black" />

          <div className="my-2">
            <h2 className="text-2xl text-center font-medieval">{spell.description}</h2>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}