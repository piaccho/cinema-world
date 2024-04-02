import { Showing } from "../mongoSchemas";

type MovieIdGroupedShowings = {
    movieId: string;
    showings: Showing[];
};

type DateGroupedShowings = {
    date: string;
    showings: Showing[];
};

export function groupShowingsByDate(showings: Showing[]): DateGroupedShowings[] {
    const grouped: { [key: string]: Showing[] } = {};

    for (const showing of showings) {
        const date = showing.startTime.toString().split('T')[0];
        if (!grouped[date]) {
            grouped[date] = [];
        }
        grouped[date].push(showing);
    }

    return Object.keys(grouped).map(date => ({
        date,
        showings: grouped[date],
    }));
}

export function groupShowingsByMovie(showings: Showing[]): MovieIdGroupedShowings[] {
    const grouped: { [key: string]: Showing[] } = {};

    for (const showing of showings) {
        const movieId = showing.movie._id;
        if (!grouped[movieId]) {
            grouped[movieId] = [];
        }
        grouped[movieId].push(showing);
    }

    return Object.keys(grouped).map(movieId => ({
        movieId,
        showings: grouped[movieId],
    }));
}