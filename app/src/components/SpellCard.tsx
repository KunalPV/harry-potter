import { Card, CardContent } from "./ui/card";
import { Separator } from "./ui/separator";


export default function SpellCard() {
  return(
    <div className="w-full">
      <Card className=" sm:w-fit" >
        <CardContent className="flex justify-around items-center gap-4 p-6 flex-col">
          <div className="w-full flex justify-between items-center">
            <div>
              <p className="text-2xl font-semibold text-justify">Capacious Extremis</p>
            </div>
            <div className="w-8 h-8 bg-gray-500 rounded-full" />
          </div>

          <Separator />

          <div className="">
            <h2 className="text-xl text-justify">The Patronus Charm is a powerful projection of hope and happiness that drives away Dementors; a corpeal Patronus takes the the respective animal form of the caster, while a non-corpeal appears as a wisp of light; at 13, Harry Potter was the youngest known witch or wizard to prouduce a corpeal Patronus</h2>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}