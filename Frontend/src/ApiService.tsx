import axios, { AxiosInstance } from 'axios';
import { API } from './config/config';
import { Category, Movie, Showing } from './types';

// Klasa serwisu API
export default class ApiService {
    private api: AxiosInstance;

    constructor() {
        this.api = axios.create({
            baseURL: API.BASE_URL,
            headers: {
                "Content-type": "application/json",
                "Access-Control-Allow-Origin": "*",
            }
        });
    }

    async getPopularMovies(): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(API.GET_POPULAR_MOVIES);
        return response.data;
        // return fetch("src/mocks/popularMovies.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Movie[]) => {
        //         console.log(`Got ${data.length} popular movies`);
        //         return data;
        //     });
    }

    async getUpcomingMovies(): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(API.GET_UPCOMING_MOVIES);
        return response.data;
        // return fetch("src/mocks/upcomingMovies.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Movie[]) => {
        //         console.log(`Got ${data.length} upcoming movies`);
        //         return data;
        //     });
    }

    async getGenres(): Promise<Category[]> {
        const response = await this.api.get<Category[]>(API.GET_GENRES);
        return response.data;
        // return fetch("/src/mocks/genres.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Genre[]) => {
        //         console.log(`Got ${data.length} genres`);
        //         return data;
        //     });
    }

    async getMoviesBySearchQuery(query: string): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(`${API.GET_MOVIES_BY_SEARCH_QUERY}`, {
            withCredentials: false,
            params: {
                query: query
            }
        });
        console.log(`got movies by search query: ${query}`, response.data);
        return response.data;
        // return fetch("/src/mocks/searchQueryMovies.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Movie[]) => {
        //         console.log(`Got ${data.length} genres`);
        //         return data;
        //     });
    } 

    async getMoviesByGenre(genreName: string): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(`${API.GET_MOVIES_BY_GENRE}${genreName}`);
        return response.data;
        // return fetch("/src/mocks/actionMovies.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Movie[]) => {
        //         console.log(`Got ${data.length} genres`);
        //         return data;
        //     });
    }

    async getShowingListsByDate(date: string): Promise<Showing[]> {
        const response = await this.api.get<Showing[]>(`${API.GET_SHOWINGS_BY_DATE}${date}`);
        console.log(`Got repertoire for ${date}`, response.data);
        return response.data;
        // return fetch("/src/mocks/showings.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Showing[]) => {
        //         console.log(`Got repertoire with ${data.length} showings`);
        //         return data;
        //     });
    }

    async getShowingListsByMovieId(movieId: number): Promise<Showing[]> {
        const response = await this.api.get<Showing[]>(`${API.GET_SHOWINGS_BY_MOVIEID}${movieId}`);
        console.log(`Got ${response.data.length} repertoires`, response.data);
        return response.data;
        // return fetch("/src/mocks/showingsOneMovie.json")
        //     .then(response => {
        //         if (!response.ok) {
        //             throw new Error(`HTTP error! status: ${response.status}`);
        //         }
        //         return response.json();
        //     })
        //     .then((data: Showing[]) => {
        //         console.log(`Got ${data.length} repertoires`);
        //         return data;
        //     });
    }

    // async createPost(post: Post): Promise<Post> {
    //     const response = await this.api.post<Post>('/posts', post);
    //     return response.data;
    // }

    // async updatePost(id: number, post: Partial<Post>): Promise<Post> {
    //     const response = await this.api.put<Post>(`/posts/${id}`, post);
    //     return response.data;
    // }

    // async deletePost(id: number): Promise<void> {
    //     await this.api.delete(`/posts/${id}`);
    // }
}
