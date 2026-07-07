"use client";

import { useEffect, useState } from "react";
import Image from "next/image";
import Link from "next/link";

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export interface MovieCardProps {
  id: string;
  title: string;
  poster: string;
  screen: string;
  rating: string;
  language: string;
}

interface Showtime {
  id: string;
  time: string;
  screen: string;
}

function getDateTabs() {
  const labels = ["Today", "Tomorrow"];
  return Array.from({ length: 5 }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() + i);
    return {
      date,
      label: labels[i] ?? date.toLocaleDateString("en-US", { weekday: "short" }),
      day: date.toLocaleDateString("en-US", { day: "2-digit", month: "short" }),
    };
  });
}

export function MovieCard({ id, title, poster, screen, rating, language }: MovieCardProps) {
  const [expanded, setExpanded] = useState(false);
  const [activeDate, setActiveDate] = useState(0);
  const [shows, setShows] = useState<Showtime[]>([]);
  const [loading, setLoading] = useState(false);
  const dateTabs = getDateTabs();

  useEffect(() => {
    if (!expanded) return;

    const isoDate = dateTabs[activeDate].date.toISOString();
    let cancelled = false;
    setLoading(true);

    fetch(`${API_URL}/api/movies/${id}/shows?date=${encodeURIComponent(isoDate)}`)
      .then((res) => (res.ok ? res.json() : []))
      .then((data: Showtime[]) => {
        if (!cancelled) setShows(data);
      })
      .finally(() => {
        if (!cancelled) setLoading(false);
      });

    return () => {
      cancelled = true;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [expanded, activeDate, id]);

  return (
    <div className="w-full bg-zinc-200 rounded-md border-2 border-zinc-100/10 overflow-hidden font-sans">
      <div className="flex items-stretch">
        <div className="relative w-32 sm:w-44 h-54 shrink-0 ">
          <Image src={poster} alt={title} fill className="object-cover px-1 py-1 rounded-xl" />
        </div>

        <div className="flex-1 flex flex-col gap-4 px-4 sm:px-2 py-2 bg-zinc-200 m-1 border-2 border-zinc-300 rounded-lg">
          <div className="flex items-start justify-between gap-4">
            <div>
              <h2 className="text-zinc-900 text-xl tracking-tighter">{title}</h2>
              <p className="text-zinc-400 tracking-tighter text-sm">{screen} | {rating} | {language}</p>
            </div>

            <button
              onClick={() => setExpanded((v) => !v)}
              className="shrink-0 py-1 block w-42 h-10 px-6 bg-zinc-100 border-2 border-zinc-100/10 text-zinc-900 cursor-pointer rounded-md font-medium tracking-tighter transition-all active:scale-95"
            >
              {expanded ? "Close" : "Book Ticket"}
            </button>
          </div>

          {expanded && (
            <div className="border-t-2 border-zinc-100/10 pt-4">
              <div className="flex gap-6 border-b-2 border-zinc-800/10 overflow-x-auto">
                {dateTabs.map((tab, i) => (
                  <button
                    key={i}
                    onClick={() => setActiveDate(i)}
                    className={`shrink-0 pb-2 -mb-[2px] text-sm tracking-tighter border-b-2 transition-all ${
                      activeDate === i
                        ? "border-b-zinc-100 border-b text-zinc-800"
                        : "border-transparent text-zinc-400"
                    }`}
                  >
                    {tab.label} <span className="opacity-70">{tab.day}</span>
                  </button>
                ))}
              </div>

              <div className="flex gap-2 flex-wrap mt-3">
                {loading && (
                  <p className="text-sm text-zinc-500 tracking-tighter">Loading times...</p>
                )}
                {!loading && shows.length === 0 && (
                  <p className="text-sm text-zinc-500 tracking-tighter">No shows on this date</p>
                )}
                {!loading &&
                  shows.map((show) => (
                    <Link
                      key={show.id}
                      href={`/movie/${id}/book?showId=${encodeURIComponent(show.id)}&screen=${encodeURIComponent(show.screen)}&time=${encodeURIComponent(show.time)}&language=${encodeURIComponent(language)}&title=${encodeURIComponent(title)}&date=${encodeURIComponent(dateTabs[activeDate].date.toISOString())}`}
                      className="px-4 py-2 rounded-md text-sm bg-zinc-600/10 backdrop-blur-2xl border-2 border-zinc-100/10 text-zinc-600 tracking-tighter"
                    >
                      {show.time}
                    </Link>
                  ))}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
