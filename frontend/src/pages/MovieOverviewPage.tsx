import React, { useEffect, useRef, useState } from 'react';
import { Box, Container, Grid, Typography } from '@mui/material';
import MovieOverviewCard from '../components/MovieOverviewCard';
import { useLocation } from 'react-router-dom';
import { Showing } from '../types';
import ApiService from '../ApiService';
import MovieOverviewShowing from '../components/MovieOverviewShowing';
import {groupShowingsByDate} from '../util/groupShowings';

const MovieOverviewPage: React.FC = () => {
    const location = useLocation();
    const movie = location.state.movie;
    const [showingList, setShowingList] = useState<Showing[]>([]);
    const apiService = new ApiService();

    const showingRef = useRef<HTMLDivElement | null>(null);

    const handleScrollToShowing = () => {
        if (showingRef.current) {
            showingRef.current.scrollIntoView({ behavior: 'smooth' });
        }
    };
    
    useEffect(() => {
        const fetchShowings = async () => {
            setShowingList(await apiService.getShowingListsByMovieId(movie._id));
        };
        fetchShowings();
    }, []);

    return (
        <Container maxWidth="xl">
            <Grid container spacing={2} display='flex' direction='column' alignItems='center' mb={5}>
                <MovieOverviewCard movie={movie} onButtonClick={handleScrollToShowing} />
                <Box ref={showingRef} sx={{ display: 'flex', flexDirection: 'column', width: '1000px', marginTop: '100px', marginBottom: '20px' }}>
                    <Typography variant="h4" mb={2} sx={{ color: 'primary.main', fontWeight: 'Bold' }}>
                        Repertoires:
                    </Typography>
                    {groupShowingsByDate(showingList).length === 0 ? 
                            <Typography variant="h6" mb={4} sx={{ color: 'primary.dark', fontWeight: 'Bold' }}>
                                No showings available
                            </Typography>
                            : null
                        }
                    <Grid container spacing={4} >   
                        {groupShowingsByDate(showingList).map((showingGroup) => (
                            <MovieOverviewShowing showings={showingGroup.showings} key={showingGroup.date} />
                        ))}
                    </Grid>
                </Box>
            </Grid>
        </Container>
    );
};

export default MovieOverviewPage;


