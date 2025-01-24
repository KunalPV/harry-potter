import Link from "next/link";
import { Separator } from "./ui/separator";

export function NavBar() {
  return (
    <div className="w-full">
      <div className="w-full flex justify-center items-center mx-1 font-bold">
        <Link href={"/"}>
        <h1 className="text-5xl sm:text-6xl p-4 font-harry tracking-wide">Potterhead Puzzles</h1>
        </Link>
      </div>

      <div className="w-full flex justify-center items-center">
        <Separator className="w-[90%] bg-black" />
      </div>
    </div>
  )
}