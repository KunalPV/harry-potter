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
        <CardContent className="flex justify-around items-center gap-4 p-6 flex-col">
          <div className="w-full flex justify-between items-center">
            <div>
              <p className="text-2xl font-semibold text-justify">{spell.name}</p>
            </div>
            <div className="w-8 h-8 bg-gray-500 rounded-full" />
          </div>

          <Separator />

          <div className="">
            <h2 className="text-xl text-center">{spell.description}</h2>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}