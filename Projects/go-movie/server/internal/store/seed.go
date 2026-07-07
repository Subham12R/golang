package store

import (
	"fmt"
	"go-movie-server/internal/models"
	"time"
)

type demoMovie struct {
	id          string
	title       string
	poster      string
	screen      string
	certificate string
	language    string
	times       []string
}

var demoMovies = []demoMovie{
	{
		id:          "1",
		title:       "Inception",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=Inception",
		screen:      "SCR-1",
		certificate: "UA13+",
		language:    "English",
		times:       []string{"10:00 AM", "1:30 PM", "6:00 PM", "9:15 PM"},
	},
	{
		id:          "2",
		title:       "Interstellar",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=Interstellar",
		screen:      "SCR-2",
		certificate: "U",
		language:    "English",
		times:       []string{"11:00 AM", "2:30 PM", "7:00 PM"},
	},
	{
		id:          "3",
		title:       "The Dark Knight",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=The+Dark+Knight",
		screen:      "SCR-3",
		certificate: "UA16+",
		language:    "English",
		times:       []string{"12:00 PM", "3:30 PM", "8:00 PM", "10:45 PM"},
	},
	{
		id:          "4",
		title:       "Dune: Part Two",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=Dune+Part+Two",
		screen:      "SCR-1",
		certificate: "UA13+",
		language:    "English",
		times:       []string{"1:00 PM", "4:30 PM", "8:15 PM"},
	},
	{
		id:          "5",
		title:       "Oppenheimer",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=Oppenheimer",
		screen:      "SCR-2",
		certificate: "A",
		language:    "English",
		times:       []string{"10:30 AM", "2:00 PM", "6:30 PM", "9:45 PM"},
	},
	{
		id:          "6",
		title:       "Parasite",
		poster:      "https://placehold.co/400x600/1a1a1a/ffffff?text=Parasite",
		screen:      "SCR-3",
		certificate: "UA16+",
		language:    "Korean",
		times:       []string{"11:45 AM", "5:00 PM", "9:00 PM"},
	},
}

const showDays = 5

// SeedDemoData populates the store with a fixed set of demo movies and
// generates showtimes for the next few days from each movie's times list.
func SeedDemoData(s *SeatStore) {
	s.Lock()
	defer s.Unlock()

	nextShowID := 1
	today := time.Now()

	for _, m := range demoMovies {
		s.Movies[m.id] = &models.Movie{
			ID:          m.id,
			Title:       m.title,
			Poster:      m.poster,
			Screen:      m.screen,
			Certificate: m.certificate,
			Language:    m.language,
		}

		var shows []*models.Show
		for day := 0; day < showDays; day++ {
			date := today.AddDate(0, 0, day)
			for _, t := range m.times {
				parsed, err := time.Parse("3:04 PM", t)
				if err != nil {
					continue
				}
				startTime := time.Date(
					date.Year(), date.Month(), date.Day(),
					parsed.Hour(), parsed.Minute(), 0, 0,
					date.Location(),
				)
				shows = append(shows, &models.Show{
					ID:        fmt.Sprintf("SH-%d", nextShowID),
					MovieID:   m.id,
					StartTime: startTime,
					ScreenID:  m.screen,
				})
				nextShowID++
			}
		}
		s.Shows[m.id] = shows
	}
}
