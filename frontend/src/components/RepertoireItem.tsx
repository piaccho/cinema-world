import * as React from 'react';
import Typography from '@mui/material/Typography';
import { Box, Button, Card, CardContent, CardMedia, Grid } from '@mui/material';
import { Showing } from '../types';
import getTimeFromDate from '../util/getTimeFromDate';

const RepertoireItem: React.FC<{ showings: Showing[] }> = ({ showings }) => {
    return (
        <Grid item xs={12} md={12}>
            <Card sx={{ display: 'flex' }}>
                <CardMedia
                    component="img"
                    sx={{ width: 200, display: { xs: 'none', sm: 'block' } }}
                    image={showings[0].movie.image}
                    alt='thumbnail'
                />
                <CardContent sx={{ flex: 1 }}>
                    <Typography variant="h5" mb={2} sx={{ fontWeight: 'bold', color: 'primary.main' }}>{showings[0].movie.title}</Typography>
                    {/* <Typography variant="h6" mb={2}>{showings[0].movie.genres[0]} | {showings[0].movie.length} min</Typography> */}
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

export default RepertoireItem;