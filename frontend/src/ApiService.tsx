import axios, { AxiosInstance } from 'axios';
import { API } from './config/config';
import { Genre, Movie, Showing } from './types';

type ApiResponse = {
    status: number;
    message: string;
    data: { [key: string]: any };
};

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

    async getAllGenres(): Promise<Genre[]> {
        const URL = `${API.CRUD_GENRES}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getGenreByName(genreName: string): Promise<Genre[]> {
        const URL = `${API.GET_GENRE_BY_NAME}${genreName}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getMovieByTitle(title: string): Promise<Genre[]> {
        const URL = `${API.GET_MOVIE_BY_TITLE}${title}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getPopularMovies(): Promise<Movie[]> {
        const URL = `${API.GET_POPULAR_MOVIES}${API.DEFAULT_ENTRIES_NUMBER}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getUpcomingMovies(): Promise<Movie[]> {
        const URL = `${API.GET_UPCOMING_MOVIES}${API.DEFAULT_ENTRIES_NUMBER}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getMoviesBySearchQuery(query: string): Promise<Movie[]> {
        const URL = `${API.GET_MOVIES_BY_SEARCH_QUERY}${query}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getMoviesByGenreId(genreId: number): Promise<Movie[]> {
        const URL = `${API.GET_MOVIES_BY_GENRE_ID}${genreId}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getMoviesByGenreName(genreName: string): Promise<Movie[]> {
        const URL = `${API.GET_MOVIES_BY_GENRE_NAME}${genreName}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getShowingListsByDate(date: string): Promise<Showing[]> {
        const URL = `${API.GET_SHOWINGS_BY_DATE}${date}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }

    async getShowingListsByMovieId(movieId: number): Promise<Showing[]> {
        const URL = `${API.GET_SHOWINGS_BY_MOVIEID}${movieId}`;
        const response = await this.api.get<ApiResponse>(URL);
        console.log(`GET ${URL}`, response.data);
        return response.data.data.data || [];
    }
}