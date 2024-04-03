import React, { useEffect, useState } from 'react';
import { CircularProgress, Container, Typography } from '@mui/material';
import { Movie } from '../mongoSchemas';
import ApiService from '../ApiService';
import Carousel from '../components/Carousel';
import MovieItem from '../components/MovieItem';

const HomePage: React.FC = () => {
    const [popularMovies, setPopularMovies] = useState<Movie[]>([]);
    const [errorPopular, setErrorPopular] = useState<string | null>(null);
    const [loadingPopular, setLoadingPopular] = useState<boolean>(false);

    const [upcomingMovies, setUpcomingMovies] = useState<Movie[]>([]);
    const [errorUpcoming, setErrorUpcoming] = useState<string | null>(null);
    const [loadingUpcoming, setLoadingUpcoming] = useState<boolean>(false);

    const apiService = new ApiService();

    useEffect(() => {
        const fetchPopularMovies = async () => {
            setLoadingPopular(true);
            try {
                const fetchedPopularMovies = await apiService.getPopularMovies();
                if (Array.isArray(fetchedPopularMovies)) {
                    setPopularMovies(fetchedPopularMovies);
                } else {
                    setErrorPopular('Invalid data format');
                }
            } catch (err) {
                setErrorPopular('Failed to fetch popular movies');
            } finally {
                setLoadingPopular(false);
            }
        };
        
        const fetchUpcomingMovies = async () => {
            setLoadingUpcoming(true);
            try {
                const fetchedUpcomingMovies = await apiService.getUpcomingMovies();
                if (Array.isArray(fetchedUpcomingMovies)) {
                    setUpcomingMovies(fetchedUpcomingMovies);
                } else {
                    setErrorUpcoming('Invalid data format');
                }
            } catch (err) {
                setErrorUpcoming('Failed to fetch upcoming movies');
            } finally {
                setLoadingUpcoming(false);
            }
        };
        
        fetchPopularMovies();
        fetchUpcomingMovies();
    }, []);

    return (
        <Container 
            maxWidth="xs" 
            sx={{ padding: 5, display: 'flex', flexDirection: 'column', alignItems: 'center'}}
        >
            <Typography 
                variant="h5" 
                color="initial" 
                mb={2}
                sx={{ fontWeight: 'bold', color: 'primary.light' }}
            >
                Popular Movies
            </Typography>
            {loadingPopular ? (
                <CircularProgress color='info'/>
            ) : errorPopular ? (
                <div>Error: {errorPopular}</div>
            ) : (
                <Carousel
                    elements={popularMovies.map((movie) => (
                        <MovieItem key={`popular-${movie._id}`} movie={movie} />
                    ))}
                />
                )
            }
            <Typography 
                variant="h5" 
                color="initial" 
                mt={2} 
                mb={2}
                sx={{ fontWeight: 'bold', color: 'primary.light' }}
            >Upcoming Movies</Typography>
            {loadingUpcoming ? (
                <CircularProgress color='info' />
            ) : errorUpcoming ? (
                <div>Error: {errorUpcoming}</div>
            ) : (
                <Carousel
                    elements={upcomingMovies.map((movie) => (
                        <MovieItem key={`upcoming-${movie._id}`} movie={movie} />
                    ))}
                />
            )
            }
        </Container>
    );
}

export default HomePage;
