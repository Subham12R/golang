# Concurrent Movie Booking Platform

A concurrent movie booking application built with **Go** backend and **Next.js + React** frontend. This project teaches concurrency fundamentals through practical implementation: goroutines, channels, waitgroups, and thread-safe data structures.

---

## Architecture

```
┌─────────────────────────────────────────┐
│       Frontend (Next.js + React)        │
│  ┌──────────┐  ┌──────────────┐  ┌────┐ │
│  │   /      │→ │  Seat Grid   │→ │Pay │ │
│  │  Movies  │  │              │  │ment│ │
│  │          │  │ Bookings     │  │Form│ │
│  └──────────┘  └──────────────┘  └────┘ │
└─────────────────────────────────────────┘
           HTTP API Requests (127.0.0.1:8080)
                        ↓
┌─────────────────────────────────────────┐
│          Backend (Go)                   │
│  ├── Models:                           │
│  │   • Movie, Showtime                │
│  │   • Seat, TheaterGrid              │
│  │   • User, Reservation               │
│  └── Handlers + Store with Sync Primitives│
└─────────────────────────────────────────┘
```

---

## MVP Features

### Backend (Go)
- Movies and showtimes endpoint (`GET /api/movies`)
- Theater seating grid visualization (`GET /api/showtimes/{movieId}`)
- Seat booking with concurrent safety (POST `/api/bookings`)
- User registration and reservations
- Goroutines for parallel seat processing
- Thread-safe data structures using channels and sync.Map

### Frontend (Next.js + React)
- Movie listing page with showtimes
- Interactive seat picker component (8x8 grid)
- Booking confirmation modal
- Mock payment form

---

## Tech Stack

| Layer | Technology | Concurrency Concepts |
|-------|-----------|---------------------|
| **Backend** | Go 1.21+ | Goroutines, channels, WaitGroups, Mutexes, sync.Map |
| **Frontend** | Next.js + React | API integration, state management |
| **Storage** | In-memory (MVP) | N/A |
| **Server** | net/http or Gorilla Mux | Concurrent HTTP handlers |

---

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- Node.js 18+ for Next.js frontend
- Git (optional, for version control)

### Backend Setup
```bash
cd backend
go mod init movie-booking-backend

# Run the server
go run main.go
```

### Frontend Setup
```bash
npx create-next-app@latest frontend
cd frontend
npm run dev
```

---

## Project Structure

```
├── backend/
│   ├── main.go                 # Application entry point with router
│   ├── models/
│   │   ├── movie.go            # Movie and Showtime structs
│   │   ├── seat.go             # Seat and theater grid
│   │   └── user.go             # User and reservation
│   ├── handlers/
│   │   ├── movies.go           # Movie API handlers
│   │   ├── bookings.go         # Booking handlers (goroutine-heavy!)
│   │   └── users.go            # User registration handlers
│   └── store/
│       └── data.go             # In-memory stores with sync primitives
│
├── frontend/                   # Next.js application
│   ├── app/
│   │   ├── movies/
│   │   ├── book/[movieId]/
│   │   └── layout.tsx
│   └── components/
│       └── SeatGrid/           # Interactive seat selection
│
├── README.md
└── docker-compose.yml          # Optional: containerized deployment
```

---

## Learning Goals

Through this project, you'll learn:

### Week 1: Goroutine Fundamentals
- Creating and launching goroutines
- Basic channel patterns (send/receive)
- Buffering vs unbuffered channels
- WaitGroups for synchronization

### Week 2: Concurrency Patterns
- Mutex protection with proper locking
- Lock-free data structures using sync.Map
- Producer-consumer pattern implementation
- Context cancellation for graceful shutdowns

### Week 3: Frontend Integration
- Building concurrent UI responses
- Optimistic updates during async operations
- Error handling across boundaries

---

## Testing Concurrent Scenarios

Try these tests once your backend is running:

```bash
# Test simultaneous seat booking (race condition prevention)
for i in {1..50}; do 
  curl -X POST http://localhost:8080/api/bookings \
    -d '{"movieId":"1","seatRows":4,"seatCols":7}' &
done

# Wait for all requests to complete
wait
```

---

## Contributing

Feel free to submit issues and pull requests! When making changes:

1. Follow Go naming conventions
2. Add comments explaining concurrent operations
3. Test with `go test ./...` when tests exist

---

## License

MIT License - feel free to use this as a learning resource or base for your own projects.

---

## Important Notes

- This MVP uses an in-memory store for simplicity (no database yet)
- All data is lost on server restart
- Production deployment would require a proper database with transactions
- Concurrency safety is enforced manually - no automatic locking
