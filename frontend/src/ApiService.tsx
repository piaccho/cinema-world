import axios, { AxiosInstance } from 'axios';
import { API } from './config/config';
import { Genre, Movie, Showing } from './types';
// import { getPopularMoviesMock, getUpcomingMoviesMock, getGenresMock, getMoviesBySearchQueryMock, getMoviesByGenreMock, getShowingListsByDateMock, getShowingListsByMovieIdMock } from './util/mockServices';


// Klasa serwisu API
export default class ApiService {
    private api: AxiosInstance;

    constructor() {
        const baseURL = import.meta.env.VITE_API_URI || 'http://localhost:8080/api';
        
        this.api = axios.create({
            baseURL: baseURL,
            headers: {
                "Content-type": "application/json"
            }
        });
    }

    async getPopularMovies(): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(API.GET_POPULAR_MOVIES);
        return response.data;
        // return getPopularMoviesMock();
    }

    async getUpcomingMovies(): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(API.GET_UPCOMING_MOVIES);
        return response.data;
        // return getUpcomingMoviesMock();
    }

    async getGenres(): Promise<Genre[]> {
        const response = await this.api.get<Genre[]>(API.GET_GENRES);
        return response.data;
        // return getGenresMock();
    }

    async getMoviesBySearchQuery(query: string): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(`${API.GET_MOVIES_BY_SEARCH_QUERY}${query}`);
        console.log(`got movies by search query: ${query}`, response.data);
        return response.data;
        // return getMoviesBySearchQueryMock(query);
    } 

    async getMoviesByGenre(genreName: string): Promise<Movie[]> {
        const response = await this.api.get<Movie[]>(`${API.GET_MOVIES_BY_GENRE}${genreName}`);
        return response.data;
        // return getMoviesByGenreMock(genreName);
    }

    async getShowingListsByDate(date: string): Promise<Showing[]> {
        const response = await this.api.get<Showing[]>(`${API.GET_SHOWINGS_BY_DATE}${date}`);
        console.log(`Got repertoire for ${date}`, response.data);
        return response.data;
        // return getShowingListsByDateMock(date);
    }

    async getShowingListsByMovieId(movieId: number): Promise<Showing[]> {
        const response = await this.api.get<Showing[]>(`${API.GET_SHOWINGS_BY_MOVIEID}${movieId}`);
        console.log(`Got ${response.data.length} repertoires`, response.data);
        return response.data;
        // return getShowingListsByMovieIdMock(movieId);
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
