package models

type Movie struct {
	ID 			int 	`json:"id"`
	Title		string	`json:"title"`
	Director	string	`json:"director"`
	ReleaseYear	int		`json:"release_year"`
}

type CreateMovieRequest struct{
	Title 		string	`json:"title"`
	Director	string	`json:"director"`
	ReleaseYear	int		`json:"release_year"`
}

type UpdateMovieRequest struct{
	Title		string	`json:"title"`
	Director	string	`json:"director"`
	ReleaseYear	int		`json:"release_year"`
}
