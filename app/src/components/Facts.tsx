import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"

export function Facts() {
  return(
    <div className="w-full flex justify-center items-center p-0 md:p-4">
      <Card className="w-full flex justify-center items-center flex-col p-2 md:p-0">
        <CardHeader className="p-2 md:p-4">
          <CardTitle className="text-2xl">Fact of the Day.</CardTitle>
        </CardHeader>

        <Separator className="mx-4 w-[95%]" />

        <CardContent className="p-2 md:p-4 text-balance text-center">
          <p className="text-xl">The spell Protego casts an invisible shield around the caster, protecting against spells and objects (except for the killing curse).</p>
        </CardContent>
      </Card>
    </div>
  )
}