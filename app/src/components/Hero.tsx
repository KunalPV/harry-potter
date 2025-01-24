import { Facts } from "./Facts";
import { Button } from "@/components/ui/button"


export function Hero() {
  return(
    <div className="w-full px-4 md:px-20 py-5 md:py-14 flex flex-col justify-center items-start">
      <div className="font-semibold w-full">
        <p className="text-3xl text-center md:text-start pl-0 md:pl-5">Welcome Back, <span className="text-4xl">name!</span></p>
      </div>
      <div className="w-full mt-8 md:mt-10">
        <Facts />
      </div>
      <div className="w-full flex justify-center items-center mt-10 md:mt-5">
        <div className="w-full flex flex-col md:flex-row justify-center items-center gap-4 md:gap-10">
          <Button size={"lg"} className="text-xl w-64">Trivia</Button>
          <Button size={"lg"} className="text-xl w-64">Characters</Button>
          <Button size={"lg"} className="text-xl w-64">Spells</Button>
        </div>
      </div>
    </div>
  )
}