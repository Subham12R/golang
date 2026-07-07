"use client";

import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "next/navigation";
import Link from "next/link";

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
const DEMO_USER_ID = 1;

const SCREEN_ROWS = 8;
const SCREEN_COLS = 10;
const PREMIER_ROWS = 2;
const SILVER_PRICE = 149;
const PREMIER_PRICE = 199;

type SeatStatus = "available" | "bestseller" | "sold";

interface Seat {
  id: string;
  status: SeatStatus;
}

interface Section {
  name: string;
  price: number;
  rows: Seat[][];
}

function buildSections(availableIds: Set<string>): Section[] {
  const allRows: Seat[][] = [];
  for (let r = 0; r < SCREEN_ROWS; r++) {
    const rowLetter = String.fromCharCode(65 + r);
    const row: Seat[] = [];
    for (let c = SCREEN_COLS; c >= 1; c--) {
      const id = `${rowLetter}${c}`;
      const isAvailable = availableIds.has(id);
      const seed = (r * 7 + c * 3) % 11;
      const isBestseller = isAvailable && (seed === 3 || seed === 4);
      row.push({
        id,
        status: !isAvailable ? "sold" : isBestseller ? "bestseller" : "available",
      });
    }
    allRows.push(row);
  }

  return [
    { name: "SILVER", price: SILVER_PRICE, rows: allRows.slice(0, SCREEN_ROWS - PREMIER_ROWS) },
    { name: "PREMIER", price: PREMIER_PRICE, rows: allRows.slice(SCREEN_ROWS - PREMIER_ROWS) },
  ];
}

function seatClass(seat: Seat, isSelected: boolean) {
  if (seat.status === "sold") {
    return "border-zinc-200 bg-zinc-500/20 text-zinc-500 cursor-not-allowed";
  }
  if (isSelected) {
    return "border-emerald-600 bg-emerald-500 text-white";
  }
  if (seat.status === "bestseller") {
    return "border-amber-400 text-amber-400 hover:bg-amber-400/10";
  }
  return "border-emerald-400 text-emerald-500 hover:bg-emerald-500/10";
}

export default function BookingPage() {
  const searchParams = useSearchParams();

  const title = searchParams.get("title") ?? "";
  const language = searchParams.get("language") ?? "";
  const screen = searchParams.get("screen") ?? "";
  const time = searchParams.get("time") ?? "";
  const showId = searchParams.get("showId") ?? "";
  const dateParam = searchParams.get("date");
  const date = dateParam ? new Date(dateParam) : new Date();

  const [availableIds, setAvailableIds] = useState<Set<string>>(new Set());
  const [selected, setSelected] = useState<Set<string>>(new Set());
  const [payStatus, setPayStatus] = useState<"idle" | "submitting" | "success" | "error">("idle");
  const [message, setMessage] = useState("");

  const loadAvailableSeats = () => {
    if (!screen) return;
    fetch(`${API_URL}/api/seats?screen_id=${encodeURIComponent(screen)}`)
      .then((res) => (res.ok ? res.json() : []))
      .then((seats: { seat_id: string }[]) => {
        setAvailableIds(new Set(seats.map((s) => s.seat_id)));
      });
  };

  useEffect(loadAvailableSeats, [screen]);

  const sections = useMemo(() => buildSections(availableIds), [availableIds]);

  const toggleSeat = (seat: Seat) => {
    if (seat.status === "sold") return;
    setSelected((prev) => {
      const next = new Set(prev);
      if (next.has(seat.id)) next.delete(seat.id);
      else next.add(seat.id);
      return next;
    });
  };

  const total = sections.reduce((sum, section) => {
    const count = section.rows.flat().filter((seat) => selected.has(seat.id)).length;
    return sum + count * section.price;
  }, 0);

  const handlePay = async () => {
    if (selected.size === 0 || !showId) return;
    setPayStatus("submitting");
    setMessage("");

    try {
      const bookings = await Promise.all(
        Array.from(selected).map((seatId) =>
          fetch(`${API_URL}/api/bookings`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              screen_id: screen,
              seat_id: seatId,
              show_id: showId,
              user_id: DEMO_USER_ID,
            }),
          }).then(async (res) => {
            if (!res.ok) throw new Error(`Could not book seat ${seatId}`);
            return res.json();
          })
        )
      );

      setPayStatus("success");
      setMessage(`Booking confirmed: ${bookings.map((b) => b.id).join(", ")}`);
      setSelected(new Set());
      loadAvailableSeats();
    } catch (err) {
      setPayStatus("error");
      setMessage(err instanceof Error ? err.message : "Booking failed");
      loadAvailableSeats();
    }
  };

  return (
    <div className="min-h-screen bg-[#fffff] text-zinc-900 flex flex-col font-sans">
      <div className="max-w-4xl mx-auto w-full border-l-2 border-r-2 border-zinc-500/5 px-4 py-4 flex-1">
        <Link
          href="/"
          className="inline-block mb-4 text-sm text-zinc-500 tracking-tighter hover:text-zinc-900 transition-all"
        >
        Back to home
        </Link>
        <h1 className="text-xl tracking-tighter font-medium">
          {title}
          {language && ` - (${language})`}
        </h1>
        <p className="text-zinc-900 text-sm tracking-tighter mt-1">
          {screen} | {date.toLocaleDateString("en-US", { weekday: "short", day: "2-digit", month: "long", year: "numeric" })} | {time}
        </p>

        <div className="border-b-2 border-zinc-500/5 mt-4" />

        <div className="mt-8 flex flex-col gap-10 items-center">
          {sections.map((section) => (
            <div key={section.name} className="flex flex-col items-center gap-3 w-full">
              <p className="text-sm text-zinc-400 tracking-tighter">
                &#8377;{section.price} {section.name}
              </p>
              <div className="border-b border-zinc-500/10 w-full" />
              <div className="flex flex-col gap-2 mt-1">
                {section.rows.map((row, i) => (
                  <div key={i} className="flex gap-1.5 justify-center">
                    {row.map((seat) => (
                      <button
                        key={seat.id}
                        disabled={seat.status === "sold"}
                        onClick={() => toggleSeat(seat)}
                        className={`w-7 h-7 text-[11px] rounded-md border-2 flex items-center justify-center tracking-tighter transition-all ${seatClass(
                          seat,
                          selected.has(seat.id)
                        )}`}
                      >
                        {seat.id.slice(1)}
                      </button>
                    ))}
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>

        <div className="flex flex-col items-center mt-12 gap-2">
          <div className="w-150 h-3 bg-blue-200 rounded-sm" />
          <p className="text-md text-zinc-500 tracking-tighter font-medium font-sans">screen this way</p>
        </div>

        <div className="flex items-center justify-center gap-6 mt-10 text-xs text-zinc-400 flex-wrap">
          <div className="flex items-center gap-1.5">
            <span className="w-3.5 h-3.5 rounded border-2 border-amber-400" /> Bestseller
          </div>
          <div className="flex items-center gap-1.5">
            <span className="w-3.5 h-3.5 rounded border-2 border-emerald-500" /> Available
          </div>
          <div className="flex items-center gap-1.5">
            <span className="w-3.5 h-3.5 rounded bg-emerald-500" /> Selected
          </div>
          <div className="flex items-center gap-1.5">
            <span className="w-3.5 h-3.5 rounded bg-zinc-700" /> Sold
          </div>
        </div>

        {message && (
          <p
            className={`text-center text-sm mt-6 tracking-tighter ${
              payStatus === "error" ? "text-red-500" : "text-emerald-600"
            }`}
          >
            {message}
          </p>
        )}
      </div>

      {selected.size > 0 && (
        <div className="sticky bottom-0 max-w-4xl mx-auto w-full border-l-2 border-r-2 border-zinc-100/5 px-4 py-3 bg-[#fffff]">
          <div className="max-w-4xl mx-auto flex items-center justify-between">
            <p className="text-sm text-zinc-400 tracking-tighter">
              {selected.size} seat{selected.size > 1 ? "s" : ""} selected
            </p>
            <button
              onClick={handlePay}
              disabled={payStatus === "submitting" || !showId}
              className="py-1 px-6 rounded-md bg-blue-500 text-white font-medium tracking-tighter active:scale-95 transition-all disabled:opacity-50"
            >
              {payStatus === "submitting" ? "Booking..." : `Pay ₹${total}`}
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
