import { Separator } from "./ui/separator";


export function NavBar() {
  return (
    <div className="w-full">
      <div className="w-full flex justify-center items-center mx-1 font-bold">
        <h1 className="text-2xl sm:text-3xl p-4">name&apos;s wizard world</h1>
      </div>

      <div className="w-full flex justify-center items-center">
        <Separator className="w-[90%] bg-black" />
      </div>
    </div>
  )
}