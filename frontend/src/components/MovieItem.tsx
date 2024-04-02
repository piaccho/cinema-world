import * as React from 'react';
import { useNavigate } from 'react-router-dom';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import { CardMedia, useTheme } from '@mui/material';
import { Movie } from '../mongoSchemas';
import formatDate from '../util/formatDate';

const MovieItem: React.FC<{ movie: Movie }> = ({ movie }) => {
    const theme = useTheme();
    const navigate = useNavigate();

    const handleClick = () => {
        navigate(`/movies/${movie._id}`, { state: { movie } });
    };
    
    return (
        <Card key={movie._id} sx={{ width: 240, height: 500, bgcolor: 'white', cursor: 'pointer', borderRadius: '15px' }} onClick={handleClick}>
            <CardMedia
                sx={{ height: 360 }}
                component="img"
                image={movie.image}
                alt={movie.title}
            />
            <CardContent sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'space-between',
            }}>
                <Typography 
                    gutterBottom 
                    variant={movie.title.length >= 18 ? 'body1' : 'h5'} 
                    component="div"
                    sx={{ color: theme.palette.primary.main, fontWeight: 500}}
                >
                    {movie.title}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    Premiere: {formatDate(movie.releaseDate)}
                </Typography>
            </CardContent>
        </Card>
    );
}

export default MovieItem;