import React, { useEffect, useState } from 'react';
import { Typography, Button, Container, Box, CircularProgress } from '@mui/material';
import { Genre } from '../types';
import ApiService from '../ApiService';
import { Link } from 'react-router-dom';
import { useStore } from '../store';

const ChooseGenrePage: React.FC = () => {
    const [genres, setGenres] = useState<Genre[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const apiService = new ApiService();

    const setGenreId = useStore((state) => state.setGenreId);
    const setGenreName = useStore((state) => state.setGenreName);


    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            try {
                const fetchedGenres = await apiService.getAllGenres();
                if (Array.isArray(fetchedGenres)) {
                    setGenres(fetchedGenres);
                } else {
                    setError('Invalid data format');
                }
            } catch (err) {
                setError('Failed to fetch genres');
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    return (
        <Container
            component="main"
            sx={{ backgroundColor: 'primary.main', color: 'primary.contrastText' }}
        >
            <Box
                my={5}
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography variant="h4" my={3} sx={{ fontWeight: 'bold' }}>
                    Choose a Genre
                </Typography>
                {loading ? (
                    <CircularProgress color='info'/>
                ) : error ? (
                    <div>Error: {error}</div>
                ) : (
                    genres.map((genre, index) => (
                        <Button
                            key={index}
                            component={Link}
                            to={`/movies/genres/${genre.name}`}
                            onClick={() => { setGenreId(genre._id); setGenreName(genre.name) }}
                            variant="contained"
                            color="primary"
                            sx={{ width: '100%' }}
                        >
                            {genre.name}
                        </Button>
                    ))
                )}
            </Box>
        </Container>
    );
};

export default ChooseGenrePage;