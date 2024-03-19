import * as React from 'react';
import Typography from '@mui/material/Typography';
import { Box, Grid } from '@mui/material';
import { Showing } from '../types';
import { groupShowingsByMovie } from '../util/groupShowings';
import { useParams } from 'react-router-dom';
import ApiService from '../ApiService';
import RepertoireItem from './RepertoireItem';

const RepertoireList: React.FC = () => {
    const { date } = useParams();
    const [showingList, setShowingList] = React.useState<Showing[]>([]);
    const apiService = new ApiService();

    React.useEffect(() => {
        const fetchShowings = async () => {
            if (date) {
                setShowingList(await apiService.getShowingListsByDate(date));
            }
        };

        if (date) {
            fetchShowings();
        }
    }, []);

    return (
        <Box>
            <Typography variant="h4" mb={3} sx={{ fontWeight: 'bold', color: 'primary.main' }}>Repertoire: <Box component="span" sx={{ color: 'secondary.dark' }}>{date ? date : ''}</Box></Typography>
            <Grid container spacing={4}>
                {groupShowingsByMovie(showingList).map((showingGroup) => (
                    <RepertoireItem showings={showingGroup.showings} key={showingGroup.movieId}/>
                ))}
            </Grid>
        </Box>
    );
}

export default RepertoireList;