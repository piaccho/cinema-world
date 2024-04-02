import React, { useEffect, useState } from 'react';
import { Typography, Button, Container, Box, CircularProgress, Grow, useTheme } from '@mui/material';
import { Genre } from '../mongoSchemas';
import ApiService from '../ApiService';
import { Link } from 'react-router-dom';
import { useStore } from '../store';

const ChooseGenrePage: React.FC = () => {
    const theme = useTheme();
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
            sx={{ color: theme.palette.primary.light}}
        >
            <Box
                my={5}
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography variant="h4" mb={4} sx={{ fontWeight: 'bold' }}>
                    Choose a Genre
                </Typography>
                {loading ? (
                    <CircularProgress color='info' />
                ) : error ? (
                    <div>Error: {error}</div>
                ) : (
                    genres.map((genre, index) => (
                        <Grow in={true} key={index} timeout={1000}>
                            <Button
                                component={Link}
                                to={`/movies/genres/${genre.name}`}
                                onClick={() => { setGenreId(genre._id); setGenreName(genre.name) }}
                                variant="contained"
                                color="primary"
                                sx={{ width: '40%', mb:2, height: '50px', fontSize: '1.2rem', fontWeight: 'bold'}}
                            >
                                {genre.name}
                            </Button>
                        </Grow>
                    ))
                )}
            </Box>
        </Container>
    );
};

export default ChooseGenrePage;