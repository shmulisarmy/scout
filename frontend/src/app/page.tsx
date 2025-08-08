import Image from "next/image";
import { notificationMessages } from "../../apiglue/svelte_sync";
import Counter from "./counter";
import Todos from "./todos";

export default function Home() {
  return (
    <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
      <Counter/>
      <Todos/>
    </div>
  );
}
