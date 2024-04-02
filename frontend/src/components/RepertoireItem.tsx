import * as React from 'react';
import Typography from '@mui/material/Typography';
import { Box, Button, Card, CardContent, CardMedia, Grid, Stack } from '@mui/material';
import { Showing } from '../mongoSchemas';
import getTimeFromDate from '../util/getTimeFromDate';
import { green, grey, red } from '@mui/material/colors';
import getFlagEmoji from '../util/getFlagEmoji';
import ExplicitIcon from '@mui/icons-material/Explicit';
import FamilyRestroomIcon from '@mui/icons-material/FamilyRestroom';

// interface Seat {
//     rowNumber: number;
//     seatNumber: number;
//     isReserved: boolean;
// }

// interface Showing {
//     _id: string;
//     movie: MovieRef;
//     hall: Hall;
//     startTime: Date;
//     endTime: Date;
//     seats: Seat[][];
//     availableSeats: number;
//     bookedSeats: number;
//     pricePerTicket: number;
//     audioType: 'Dubbing' | 'Subtitles' | 'VoiceOver';
//     videoType: '2D' | '3D';
// }

// export interface MovieRef {
//     _id: string;
//     adult: boolean;
//     genres: Genre[];
//     originalLanguage: string;
//     image: string;
//     length: number;
//     title: string;
// }


const RepertoireItem: React.FC<{ showings: Showing[] }> = ({ showings }) => {
    return (
        <Card sx={{ display: 'flex', borderRadius: '15px'}}>
            <CardMedia
                component="img"
                sx={{ width: 200, display: { xs: 'none', sm: 'block' } }}
                image={showings[0].movie.image || 'https://via.placeholder.com/200'}
                alt={`${showings[0].movie.title} - thumbnail`}
            />
            <CardContent sx={{ flex: 1 }}>
                <Stack direction="row" justifyContent="space-between">
                    <Stack direction='column' flex={1}>
                        <Stack
                            direction='row'
                            gap={1}
                            alignItems='center'
                        >
                            <Typography
                                variant="h5"
                                sx={{ fontWeight: 'bold', color: 'primary.main' }}
                            >
                                {showings[0].movie.title}
                            </Typography>
                            <Typography variant="h6" sx={{ color: 'primary.light', fontWeight: 'bold' }}>
                                {getFlagEmoji(showings[0].movie.originalLanguage)}
                            </Typography>
                        </Stack>
                        <Typography variant="body1" mb={2} sx={{ color: grey[600] }}>
                            {showings[0].movie.genres.map((genre, index) => (<span>{genre.name}{index !== showings[0].movie.genres.length - 1 ? (<>, </>): (<></>)}</span>))}
                        </Typography>
                        {showings[0].movie.adult ? (
                            <Button sx={{ px: 2, mb: 2, backgroundColor: red[300], color: 'white', pointerEvents: 'none', alignSelf: 'flex-start' }} startIcon={<ExplicitIcon />}>Adult content</Button>
                        ) : (
                            <Button sx={{ px: 2, mb: 2, backgroundColor: green[400], color: 'white', pointerEvents: 'none', alignSelf: 'flex-start' }} startIcon={<FamilyRestroomIcon />}>Family-friendly</Button>
                        )
                        }
                        <Typography variant="body1" mb={2} sx={{ color: grey[900] }}>
                            Duration: {showings[0].movie.length} min
                        </Typography>
                    </Stack>
                    <Stack gap='10px' direction='column' flex={1}>
                        {showings.map((showing, index) => (
                            <Stack direction='row' justifyContent='space-between' key={index}>
                                <Button
                                    variant={'contained'}
                                    disabled={showing.availableSeats <= 0 ? true : false}
                                    sx={{ px: 1, flex: 1, borderRadius: 0, pointerEvents: 'none', backgroundColor: 'primary.light', borderTopLeftRadius: '7px', borderBottomLeftRadius: '7px' }}
                                >
                                    {showing.videoType} | {showing.audioType} | {showing.hall.name} | {showing.pricePerTicket} €
                                </Button>
                                <Button
                                    variant={'contained'}
                                    disabled={showing.availableSeats <= 0 ? true : false}
                                    sx={{ borderRadius: 0, borderTopRightRadius: '7px', borderBottomRightRadius: '7px' }}
                                >
                                    <Stack direction='row' justifyContent='space-around' alignItems='center'>
                                        <Typography variant="body1" sx={{ color: 'primary.contrastText', fontWeight: 'bold' }}>
                                            {getTimeFromDate(showing.startTime)}
                                        </Typography>
                                    </Stack>
                                </Button>
                            </Stack>
                        ))}
                    </Stack>    
                </Stack>
            </CardContent>
        </Card>
    );
}

export default RepertoireItem;