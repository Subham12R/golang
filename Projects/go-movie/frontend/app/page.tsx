import Image from "next/image";

export default function Home() {
  return (
    <>
      <div className="min-h-screen bg-white items-center  flex flex-col leading-tight justify-center">
        <h1 className="text-4xl tracking-tighter font-serif font-medium text-zinc-900 ">Movie Booking Service</h1>
        <span className="text-md tracking-tighter font-sans font-medium text-zinc-400">
          This is a movie booking service made with Next.js and GoLang by <a href="https://subham12r.me" className="underline text-zinc-600 font-serif tracking-tighter">Subham Karmakar</a>.  
        </span>
      </div>
    </>
  );
}
