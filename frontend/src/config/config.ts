export const API = {

    DEFAULT_ENTRIES_NUMBER: 20,
    // AUTH
    LOGIN: "/auth/signIn",
    REGISTER: "/auth/signUp",

    // GENRES 
    CRUD_GENRES: "/genres/",
    GET_GENRE_BY_NAME: "/genres/name/",

    // HALLS
    CRUD_HALLS: "/halls/",

    // SHOWINGS
    CRUD_SHOWINGS: "/showings/",
    GET_SHOWINGS_BY_DATE: "/showings/date/",
    GET_SHOWINGS_BY_MOVIEID: "/showings/movie/",
    GET_SHOWINGS_BY_HALLID: "/showings/hall/",

    // RESERVATIONS
    CRUD_RESERVATIONS: "/reservations/",
    GET_RESERVATIONS_BY_USERID: "/reservations/user/",

    // MOVIES
    CRUD_MOVIES: "/movies/",
    GET_MOVIES_BY_SEARCH_QUERY: "/movies/search/",
    GET_MOVIES_BY_GENRE_ID: "/movies/genre/",
    GET_MOVIES_BY_GENRE_NAME: "/movies/genre/name/",
    GET_POPULAR_MOVIES: "/movies/popular/",
    GET_UPCOMING_MOVIES: "/movies/upcoming/",
    GET_MOVIES_REFS: "/movies/refs",
    GET_MOVIE_BY_TITLE: "/movies/title/",


    // REVIEWS
    CRUD_REVIEWS: "/movies/reviews/",

    // USERS
    CRUD_USERS: "/users/",
    CRUD_USER_WATCHLIST: "/toWatchList/items/",
    GET_USER_REVIEWS: "/reviews/",
};