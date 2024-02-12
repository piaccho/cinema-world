import React from "react";

export interface Movie {
    movieId: number;
    adult: boolean;
    categories: Category[];
    title: string;
    originalTitle: string;
    releaseDate: string;
    image: string;
    overview: string;
    length: number;
    voteAverage: number;
    voteCount: number;
    popularity: number;
    originalLanguage: string;
}

export interface BasicMovie {
    movieId: number;
    categories: Category[];
    title: string;
    image: string;
    length: number;
}

export interface Category {
    id: number;
    name: string;
}

export interface Hall {
    id: string;
    hallNumber: number;
    seatsNumber: number;
}

export interface Showing {
    showingId: string;
    movie: BasicMovie;
    hall: Hall;
    datetime: Date;
    freeSeats: number;
    type: 'sub' | 'dub' | 'voiceOver';
    price: number;
}

export interface Reservation {
    id: string;
    showing: Showing;
    userId: string;
}

export interface User {
    userId: string;
    type: string;
    login: string;
    password: string;
    firstname: string;
    lastname: string;
    to_watch: BasicMovie[];
    reservations: Reservation[];
}

export interface CarouselProps {
    elements: React.ReactElement[];
}

export interface MoviesListPageProps {
    type: 'searchQuery' | 'genre';
}

export interface MovieOverviewCardProps {
    movie: Movie;
    onButtonClick: () => void;
}

export type DateGroupedShowings = {
    date: string;
    showings: Showing[];
};


export type MovieIdGroupedShowings = {
    movieId: string;
    showings: Showing[];
};