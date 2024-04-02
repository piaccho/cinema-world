import { Genre, Movie, Showing } from './mongoSchemas';

export async function getPopularMoviesMock(): Promise<Movie[]> {
    return fetch("src/mocks/popularMovies.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getUpcomingMoviesMock(): Promise<Movie[]> {
    return fetch("src/mocks/upcomingMovies.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getGenresMock(): Promise<Genre[]> {
    return fetch("/src/mocks/genres.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getMoviesBySearchQueryMock(query: string): Promise<Movie[]> {
    return fetch("/src/mocks/searchQueryMovies.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getMoviesByGenreMock(genreName: string): Promise<Movie[]> {
    return fetch("/src/mocks/actionMovies.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getShowingListsByDateMock(date: string): Promise<Showing[]> {
    return fetch("/src/mocks/showings.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}

export async function getShowingListsByMovieIdMock(movieId: number): Promise<Showing[]> {
    return fetch("/src/mocks/showingsOneMovie.json")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        });
}