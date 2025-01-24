import { Separator } from "./ui/separator"


export function Footer() {
  return(
    <div className="w-full bottom-0 right-0 left-0 absolute p-2">
      <div className="w-full flex justify-center items-center flex-col gap-2">
      <Separator className="w-[90%]" />
        <p className="text-lg ">Made with love, by NAME</p>
      </div>
    </div>
  )
}