import React, { useEffect, useState } from 'react';
import { Box, Container, Grid, Typography, CircularProgress } from '@mui/material';
import { Movie, MoviesListPageProps } from '../types';
import ApiService from '../ApiService';
import MovieItem from '../components/MovieItem';
import capitalizeLetter from '../util/capitalLetter';
import { useStore } from '../store';
import { useParams } from 'react-router-dom';

interface MovieResultsHeaderProps {
    type: string;
    moviesLength: number;
    searchQuery: string | null;
    genreName: string | null;
}

const MovieResultsHeader: React.FC<MovieResultsHeaderProps> = ({ type, moviesLength, searchQuery, genreName }) => {
    return (
        <Box mb={4} display="flex" justifyContent="space-between" alignItems="center">
            <Typography variant="h4" sx={{ fontWeight: 'bold', color: 'primary.main' }}>
                {
                    type === 'searchQuery' ?
                        <>
                            Search results for: <Box component="span" sx={{ color: 'secondary.dark' }}>{searchQuery}</Box>
                        </>
                        : genreName ?
                            `${capitalizeLetter(genreName)} movies:`
                            : ''
                }
            </Typography>
            <Typography variant="h6" sx={{ fontWeight: 'bold', color: 'primary.light' }}>
                (Total results: {moviesLength})
            </Typography>
        </Box>
    );
}

interface MovieResultsGridProps {
    movies: Movie[];
}

const MovieResultsGrid: React.FC<MovieResultsGridProps> = ({ movies }) => {
    return (
        <Grid container spacing={4}>
            {movies.map((movie, index) => (
                <Grid item key={index} xs={12} sm={5} md={3}>
                    <MovieItem movie={movie} />
                </Grid>
            ))}
        </Grid>
    );
}

const MoviesResultsPage: React.FC<MoviesListPageProps> = ({ type }) => {
    const [movies, setMovies] = useState<Movie[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const apiService = new ApiService();

    // const { genreName, searchQuery } = useParams();
    const genreId = useStore((state) => state.genreId);
    const genreName = useStore((state) => state.genreName);
    const searchQuery = useStore((state) => state.searchQuery);


    useEffect(() => {
        const fetchMovies = async () => {
            console.log(`searchQuery, genreId, genreName: ${searchQuery}, ${genreId}, ${genreName}`);
            setLoading(true);
            try {
                if (type === 'searchQuery' && searchQuery !== null) {
                    setMovies(await apiService.getMoviesBySearchQuery(searchQuery ?? ''));
                } else if (type === 'genre' && genreName !== null) {
                    setMovies(await apiService.getMoviesByGenreName(genreName ?? ''));
                // } else if (type === 'genre' && genreId !== null) {
                // setMovies(await apiService.getMoviesByGenreId(genreId));
                } else {
                    setError('Invalid type or missing search query or genre id');
                }
            } catch (err) {
                setError('Failed to fetch movies');
            } finally {
                setLoading(false);
            }
        };
        fetchMovies();
    }, [searchQuery]);

    return (
        <Container component="main">
            <Container sx={{ py: 8 }} maxWidth="lg">
                {loading ? (
                    <CircularProgress color='info'/>
                ) : error ? (
                    <div>Error: {error}</div>
                ) : (
                    <>
                        <MovieResultsHeader type={type} moviesLength={movies.length} searchQuery={searchQuery} genreName={genreName} />
                        <MovieResultsGrid movies={movies} />
                    </>
                )}
            </Container>
        </Container>
    );
}

export default MoviesResultsPage;
