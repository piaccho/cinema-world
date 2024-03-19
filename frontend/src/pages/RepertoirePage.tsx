import React, { useEffect, useState } from 'react';
import ApiService from '../ApiService';
import dayjs, { Dayjs } from 'dayjs';
import RepertoireItem from '../components/RepertoireItem';
import { Box, Button, Container, Grid, Typography } from '@mui/material';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { DatePicker } from '@mui/x-date-pickers';
import { Showing } from '../types';
import { useNavigate, useParams } from 'react-router-dom';
import { groupShowingsByMovie } from '../util/groupShowings';

const RepertoirePage: React.FC = () => {
    const navigate = useNavigate();
    const { date } = useParams();
    const [value, setValue] = React.useState<Dayjs | null>(dayjs());
    const [showingList, setShowingList] = useState<Showing[]>([]);
    const apiService = new ApiService();

    useEffect(() => {
        const fetchShowings = async () => {
            if (date) {
                setShowingList(await apiService.getShowingListsByDate(date));
            }
        };

        if (date) {
            fetchShowings();
        } else {
            if (value) {
                navigate(`/repertoires/date/${encodeURIComponent(value.format('YYYY-MM-DD'))}`);
            }
        }
    }, []);

    const showRepertoireBtnHandler = () => {
        if (value) {
            navigate(`/repertoires/date/${encodeURIComponent(value.format('YYYY-MM-DD'))}`);
        } else {
            console.log("no value");
        }
    }

    return (
        <Container
            component="main"
            sx={{ py: 4 }} 
            maxWidth="md"
        >
            <Box display='flex' flexDirection='column' alignItems='center' gap={1} mb={6}>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DemoContainer components={['DatePicker']}>
                        <DatePicker
                            value={value || dayjs()}
                            onChange={(newValue) => setValue(newValue)}
                        />
                    </DemoContainer>
                </LocalizationProvider>
                <Button
                    sx={{ backgroundColor: "primary.main", color: "primary.contrastText", width: '250px'}}
                    onClick={showRepertoireBtnHandler}
                >
                    Show repertoire
                </Button>
            </Box>
            <Typography variant="h4" mb={3} sx={{ fontWeight: 'bold', color: 'primary.main' }}>Repertoire: <Box component="span" sx={{ color: 'secondary.dark' }}>{date ? date : ''}</Box></Typography>
            <Grid container spacing={4}>
                {groupShowingsByMovie(showingList).map((showingGroup) => (
                    <RepertoireItem showings={showingGroup.showings} key={showingGroup.movieId}/>
                ))}
            </Grid>
        </Container>
    );
};

export default RepertoirePage;
