import * as React from 'react';
import Typography from '@mui/material/Typography';
import { Box, Button, Card, CardContent, Grid } from '@mui/material';
import { Showing } from '../mongoSchemas';
import formatDate from '../util/formatDate';
import getTimeFromDate from '../util/getTimeFromDate';

const MovieOverviewShowing: React.FC<{ showings: Showing[] }> = ({ showings }) => {
    return (
        <Grid item xs={12} md={12}>
            <Card sx={{ display: 'flex' }}>
                <CardContent sx={{ flex: 1 }}>
                    <Typography variant="h6" mb={2} sx={{ fontWeight: 'bold', color: 'primary.main' }}>{formatDate(showings[0].datetime.toString())}</Typography>
                    <Typography variant="body1" mb={3}>
                        {
                            showings[0].type
                            // showings[0].type === 'dub' ? 'Dubbed'
                            //     : showings[0].type === 'sub' ? 'Subbed'
                            //         : 'Voiceover'
                        }
                    </Typography>
                    <Box sx={{ display: 'flex', gap: '10px', flexWrap: 'wrap' }}>
                        {showings.map((show, index) => (
                            <Button key={index} variant={'contained'} disabled={show.freeSeats <= 0 ? true : false}>{getTimeFromDate(show.datetime.toString())}</Button>
                        ))}
                    </Box>
                </CardContent>
            </Card>
        </Grid>
    );
}

export default MovieOverviewShowing;