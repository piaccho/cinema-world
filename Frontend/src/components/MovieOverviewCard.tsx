import { Card, CardContent, Typography, Button, Box, CardMedia } from '@mui/material';
import { MovieOverviewCardProps } from '../types';
import StarIcon from '@mui/icons-material/Star';
import TodayIcon from '@mui/icons-material/Today';
import ScheduleIcon from '@mui/icons-material/Schedule';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';



const MovieOverviewCard: React.FC<MovieOverviewCardProps> = ({ movie, onButtonClick }) => {
    return (
        <Card sx={{ display: 'flex', flexDirection: 'column', width: '1000px', marginTop: '20px'}}>
            <CardContent >
                <Box sx={{ display: 'flex', justifyContent: "space-between"}}>
                    <Typography variant="h3">
                        {movie.title}
                    </Typography>
                    <Box sx={{ display: 'flex', gap: '20px'}}>
                        <Button onClick={onButtonClick} variant="contained" color="primary" sx={{ width: '150px' }}>
                            BUY TICKET
                        </Button>
                        <Button variant="contained" color="primary">
                            <StarIcon />
                        </Button>
                    </Box>
                </Box>
            </CardContent>
            <Box sx={{ display: 'flex' }}>
                <CardContent >
                    <Box sx={{ display: 'flex', flexDirection: 'row' }} mb={3}>
                        <Box sx={{ display: 'flex', flexDirection: 'row', alignItems: 'center'}}>
                            <TodayIcon fontSize='large' sx={{mr:'5px'}}/>
                            <Box sx={{ display: 'flex', flexDirection: 'column', mr: '5em' }}>
                                <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                    RELEASE DATE:
                                </Typography>
                                <Typography variant="h6">
                                    {movie.releaseDate}
                                </Typography>
                            </Box>
                        </Box>
                        <Box sx={{ display: 'flex', flexDirection: 'row', alignItems: 'center' }}>
                            <ScheduleIcon fontSize='large' sx={{ mr: '5px' }} />
                            <Box sx={{ display: 'flex', flexDirection: 'column' }}>
                                <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                    DURATION:
                                </Typography>
                                <Typography variant="h6">
                                    {movie.length} min
                                </Typography>
                            </Box>
                        </Box>
                    </Box>
                    <Typography variant="body1" mb={3}>
                        {movie.overview}
                    </Typography>

                    <TableContainer component={Paper}>
                        <Table sx={{ minWidth: 350 }} aria-label="simple table">
                            <TableBody>
                                <TableRow
                                    key={1}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        Original title:
                                    </TableCell>
                                    <TableCell align="right">{movie.originalTitle}</TableCell>
                                </TableRow>
                                <TableRow
                                    key={2}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        Original language:
                                    </TableCell>
                                    <TableCell align="right">
                                        {movie.originalLanguage.toUpperCase()}
                                    </TableCell>
                                </TableRow>
                                <TableRow
                                    key={3}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        Genres:
                                    </TableCell>
                                    <TableCell align="right">
                                        {movie.categories.map(category => category.name).join(', ')}
                                    </TableCell>
                                </TableRow>
                                <TableRow
                                    key={4}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        Age restriction:
                                    </TableCell>
                                    <TableCell align="right">
                                        {movie.adult ? `For adults` : `None`}
                                    </TableCell>
                                </TableRow>
                                <TableRow
                                    key={5}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        Popularity:
                                    </TableCell>
                                    <TableCell align="right" sx={{ display: 'flex', alignItems: 'center', flexDirection: 'row', justifyContent: 'end'}}>
                                        <Box component='span'>
                                            {movie.popularity.toString().replace(".", " ")}
                                        </Box>
                                        <StarIcon sx={{width: '15px', ml: '2px'}}/>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </TableContainer>
                </CardContent>
                <CardMedia
                    component="img"
                    sx={{ width: 400, display: { xs: 'none', sm: 'block' } }}
                    image={movie.image}
                    alt='movie poster'
                />
            </Box>
        </Card>
    );
};

export default MovieOverviewCard;