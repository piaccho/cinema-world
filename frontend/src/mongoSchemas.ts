export interface Movie {
    _id: string;
    adult: boolean;
    genres: Genre[];
    image: string;
    length: number;
    originalLanguage: string;
    originalTitle: string;
    overview: string;
    popularity: number;
    releaseDate: Date;
    title: string;
    voteAverage: number;
    voteCount: number;
    reviews: Review[];
}

export interface MovieRef {
    _id: string;
    adult: boolean;
    genres: Genre[];
    originalLanguage: string;
    image: string;
    length: number;
    title: string;
}

export interface Review {
    _id: string;
    userId: string;
    firstName: string;
    content: string;
    rating: number;
}

export interface Genre {
    _id: number;
    name: string;
}

export interface Hall {
    _id: string;
    name: string;
    rows: number;
    seatsPerRow: number;
}

export interface Seat {
    rowNumber: number;
    seatNumber: number;
    isReserved: boolean;
}

export interface Showing {
    _id: string;
    movie: MovieRef;
    hall: Hall;
    startTime: Date;
    endTime: Date;
    seats: Seat[][];
    availableSeats: number;
    bookedSeats: number;
    pricePerTicket: number;
    audioType: 'Dubbing' | 'Subtitles' | 'VoiceOver';
    videoType: '2D' | '3D';
}

export interface ReservedSeat {
    rowNumber: number;
    seatNumber: number;
}

export interface Reservation {
    _id: string;
    showingId: string;
    userId: string;
    movieShowingRef: MovieRef;
    reservedSeats: ReservedSeat[];
    totalPrice: number;
}

export interface ToWatchListItem {
    _id: string;
    movie: MovieRef;
}

export interface User {
    _id: string;
    type: 'Admin' | 'User';
    email: string;
    password: string;
    firstname: string;
    lastname: string;
    toWatch: ToWatchListItem[];
}

