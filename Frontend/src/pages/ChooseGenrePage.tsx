import React, { useEffect, useState } from 'react';
import { Typography, Button, Container, Box } from '@mui/material';
import { Category } from '../types';
import ApiService from '../ApiService';
import { Link } from 'react-router-dom';


const ChooseGenrePage: React.FC = () => {
    const [genres, setGenres] = useState<Category[]>([]);
    const apiService = new ApiService();

    useEffect(() => {
        const fetchGenres = async () => {
            const genres = await apiService.getGenres();
            setGenres(genres);
        };
        fetchGenres();
    }, []);

    return (
        <Container 
        
            component="main"
            sx={{ backgroundColor: 'primary.main', color: 'primary.contrastText'}}>
            <Box
                my={5}
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography variant="h4" my={3} sx={{fontWeight: 'bold'}}>
                    Choose a Genre
                </Typography>
                {genres.map((genre, index) => (
                    <Button 
                        key={index}
                        component={Link} 
                        to={`/movies/genres/${genre.name}`} 
                        variant="contained" 
                        color="primary" 
                        sx={{ width: '100%' }}
                    >
                        {genre.name}
                    </Button>
                ))}
            </Box>
        </Container>
    );
};

export default ChooseGenrePage;
