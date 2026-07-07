import { MovieCard } from "@/components/MovieCard";

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

interface ApiMovie {
  id: string;
  title: string;
  poster: string;
  screen: string;
  certificate: string;
  language: string;
}

async function getMovies(): Promise<ApiMovie[]> {
  const res = await fetch(`${API_URL}/api/movies`, { cache: "no-store" });
  if (!res.ok) return [];
  return res.json();
}

export default async function Home() {
  const movies = await getMovies();

  return (
    <>
      <div className="min-h-screen bg-[#fffff]">
        <div className="max-w-4xl mx-auto  border-l-2 border-r-2 h-screen border-zinc-500/5 px-4 py-4 ">
          <div className="max-w-4xl flex flex-col gap-4 sticky top-0 bg-white pt-4 ">
            <h1 className="text-4xl tracking-tighter font-sans font-medium text-zinc-900">
              Available Movies
            </h1>
            <div className="border-b-2 border-zinc-500/5"></div>
          </div>
          <div className="grid grid-cols-1 gap-2 w-full items-start justify-start mt-4">
            {movies.map((movie) => (
              <MovieCard
                key={movie.id}
                id={movie.id}
                title={movie.title}
                poster={movie.poster}
                screen={movie.screen}
                rating={movie.certificate}
                language={movie.language}
              />
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
