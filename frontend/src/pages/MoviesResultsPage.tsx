import React, { useEffect, useState } from 'react';
import { Box, Container, Grid, Typography, CircularProgress, Zoom } from '@mui/material';
import { Movie } from '../mongoSchemas';
import ApiService from '../ApiService';
import MovieItem from '../components/MovieItem';
import capitalizeLetter from '../util/capitalLetter';
import { useParams } from 'react-router-dom';

interface MovieResultsHeaderProps {
    moviesLength: number;
    searchQuery: string | null;
    genreName: string | null;
}

const MovieResultsHeader: React.FC<MovieResultsHeaderProps> = ({ moviesLength, searchQuery, genreName }) => {
    return (
        <Box mb={4} display="flex" justifyContent="space-between" alignItems="center">
            <Typography variant="h4" sx={{ fontWeight: 'bold', color: 'primary.main' }}>
                {
                    searchQuery ?
                        <>
                            Search results for: <Box component="span" sx={{ color: 'primary.light' }}>"{searchQuery}"</Box>
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
                <Zoom in={true} key={index} timeout={500 + (50 * index)}>
                    <Grid item xs={12} sm={5} md={3}>
                        <MovieItem movie={movie} />
                    </Grid>
                </Zoom>
            ))}
        </Grid>
    );
}

const MoviesResultsPage: React.FC = () => {
    const [movies, setMovies] = useState<Movie[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const apiService = new ApiService();

    const { genreName, searchQuery } = useParams();

    useEffect(() => {
        const fetchMovies = async () => {
            console.log(`searchQuery, genreName: ${searchQuery}, ${genreName}`);
            setLoading(true);
            try {
                if (searchQuery) {
                    setMovies(await apiService.getMoviesBySearchQuery(searchQuery ?? ''));
                } else if (genreName) {
                    setMovies(await apiService.getMoviesByGenreName(genreName ?? ''));
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
                        <MovieResultsHeader moviesLength={movies.length} searchQuery={searchQuery ?? null} genreName={genreName ?? null} />
                        <MovieResultsGrid movies={movies} />
                    </>
                )}
            </Container>
        </Container>
    );
}

export default MoviesResultsPage;
