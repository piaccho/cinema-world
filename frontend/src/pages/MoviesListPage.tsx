import React, { useEffect, useState } from 'react';
import { Box, Container, Grid, Typography } from '@mui/material';
import { Movie, MoviesListPageProps } from '../types';
import ApiService from '../ApiService';
import MovieItem from '../components/MovieItem';
import { useParams } from 'react-router-dom';
import capitalizeLetter from '../util/capitalLetter';

const MoviesListPage: React.FC<MoviesListPageProps> = ({ type }) => {
    const [movies, setMovies] = useState<Movie[]>([]);
    const apiService = new ApiService();
    const { name, q } = useParams();

    useEffect(() => {
        const fetchMovies = async () => {
            if (type === 'searchQuery' && q !== undefined) {
                setMovies(await apiService.getMoviesBySearchQuery(q));
            } else if (type === 'genre' && name !== undefined) {
                setMovies(await apiService.getMoviesByGenre(name));
            }
        };
        fetchMovies();
    }, []);

    return (
        <Container
            component="main">
            <Container sx={{ py: 8 }} maxWidth="lg">
                <Typography variant="h4" mb={3} sx={{ fontWeight: 'bold', color: 'primary.main' }}>
                    {type === 'searchQuery' ?
                        <>Search results ({movies.length}) for: <Box component="span" sx={{ color: 'secondary.dark' }}>{q}</Box></> :
                        name ?
                            `${capitalizeLetter(name)} movies (${movies.length}):` : ''
                    }
                </Typography>
                <Grid container spacing={4}>
                    {movies.map((movie, index) => (
                        <Grid item key={index} xs={12} sm={5} md={3}>
                            <MovieItem movie={movie} />
                        </Grid>
                    ))}
                </Grid>
            </Container>
        </Container>
    );
};

export default MoviesListPage;
