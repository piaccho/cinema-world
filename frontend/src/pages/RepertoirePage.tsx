import React, { useEffect, useState } from 'react';
import ApiService from '../ApiService';
import dayjs, { Dayjs } from 'dayjs';
import RepertoireItem from '../components/RepertoireItem';
import { Box, Button, Container, Grid, Typography, CircularProgress, Zoom, Fade, Grow } from '@mui/material';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { DatePicker } from '@mui/x-date-pickers';
import { Showing } from '../mongoSchemas';
import { useNavigate, useParams } from 'react-router-dom';
import { groupShowingsByMovie } from '../util/groupShowings';
import formatDate from '../util/formatDate';
import SentimentVeryDissatisfiedIcon from '@mui/icons-material/SentimentVeryDissatisfied';

const RepertoirePage: React.FC = () => {
    const navigate = useNavigate();
    const { date } = useParams();
    const [value, setValue] = React.useState<Dayjs | null>(dayjs());
    const [showingList, setShowingList] = useState<Showing[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
    const apiService = new ApiService();

    useEffect(() => {
        const fetchShowings = async () => {
            setLoading(true);
            try {
                if (date) {
                    setShowingList(await apiService.getShowingListsByDate(date));
                } else {
                    setError('Invalid or missing date');
                }
            } catch (err) {
                console.log(err);
                setError(`Failed to fetch showings`);
            } finally {
                setLoading(false);
            }
        };

        if (date) {
            fetchShowings();
        } else {
            const today = dayjs().format('YYYY-MM-DD');
            navigate(`/repertoires/date/${encodeURIComponent(today)}`);
        }
    }, [date, navigate]);

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
            sx={{ py: 4, display: 'flex', flexDirection: 'column', alignItems: 'center' }}
            maxWidth="md"
        >
            <Box display='flex' flexDirection='column' alignItems='center' gap={1} mb={6}>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DemoContainer components={['DatePicker']}>
                        <DatePicker
                            value={value || dayjs()}
                            onChange={(newValue) => {
                                setValue(newValue);
                            }}
                        />
                    </DemoContainer>
                </LocalizationProvider>
                <Button
                    sx={{ backgroundColor: "primary.main", color: "primary.contrastText", width: '250px' }}
                    onClick={showRepertoireBtnHandler}
                >
                    Show repertoire
                </Button>
            </Box>
            <Typography variant="h4" mb={6} sx={{ fontWeight: 'bold', color: 'primary.main' }}>
                Repertoire for <Box component="span" sx={{ color: 'secondary.dark' }}>{formatDate(date ?? '')}</Box>
            </Typography>
            {loading ? (
                <CircularProgress color='info' />
            ) : error ? (
                <div>Error: {error}</div>
            ) : showingList.length === 0 ? (
                <Box gap={1} display='flex' flexDirection='row' alignItems='center' my={3} sx={{ color: 'primary.light' }}>
                    <Typography variant="h6" sx={{ fontWeight: 'bold'}}>
                        No showings for this date
                    </Typography>
                    <SentimentVeryDissatisfiedIcon />
                </Box>
            ) : (
                <Grid container spacing={4}>
                    {groupShowingsByMovie(showingList).map((showingGroup, index) => (
                        <Grow in={true} key={index} timeout={1000 + (index * 500)}>
                            <Grid item xs={12} md={12}>
                                <RepertoireItem showings={showingGroup.showings} key={showingGroup.movieId} />
                            </Grid>
                        </Grow>
                    ))}
                </Grid>
            )}
        </Container>
    );
};

export default RepertoirePage;