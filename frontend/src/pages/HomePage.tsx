import React, { useEffect, useState } from 'react';
import { Container, Typography } from '@mui/material';
import { Movie } from '../types';
import ApiService from '../ApiService';
import Carousel from '../components/Carousel';
import MovieItem from '../components/MovieItem';

const HomePage: React.FC = () => {
    const [popularMovies, setPopularMovies] = useState<Movie[]>([]);
    const [upcomingMovies, setUpcomingMovies] = useState<Movie[]>([]);
    const apiService = new ApiService();

    useEffect(() => {
        const fetchPopularMovies = async () => {
            const movies = await apiService.getPopularMovies();
            setPopularMovies(movies);
        };
        
        const fetchUpcomingMovies = async () => {
            const movies = await apiService.getUpcomingMovies();
            setUpcomingMovies(movies);
        };
        
        fetchPopularMovies();
        fetchUpcomingMovies();
    }, []);

    return (
        <Container maxWidth="xs" sx={{ padding: 5, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
            {/* extract to components - CarouselSection */}
            <Typography 
                variant="h5" 
                color="initial" 
                mb={2}
                sx={{ fontWeight: 'bold', color: 'primary.light' }}
            >
                Popular Movies
            </Typography>
            <Carousel
                elements={popularMovies.map((movie) => (
                    <MovieItem key={`popular-${movie._id}`} movie={movie} />
                ))} 
            />
            <Typography 
                variant="h5" 
                color="initial" 
                mt={2} 
                mb={2}
                sx={{ fontWeight: 'bold', color: 'primary.light' }}
            >Upcoming Movies</Typography>
            <Carousel
                elements={upcomingMovies.map((movie) => (
                    <MovieItem key={`upcoming-${movie._id}`} movie={movie} />
                ))} 
            />
        </Container>
    );
}

export default HomePage;
