import React from 'react';
import { Typography, Button, Container } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const NoMatchPage: React.FC = () => {
    const navigate = useNavigate();

    const handleGoBack = () => {
        navigate("/");
    };

    return (
        <Container 
            maxWidth="lg"
            component="main"
            sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                marginTop: '5rem',
            }}
            >
                <Typography variant="h4" component="h1" mb={3}>
                    404 - Page Not Found
                </Typography>
                <Typography variant="body1" mb={5}>
                    Oops! The page you are looking for does not exist.
                </Typography>
                <Button variant="contained" color="primary" onClick={handleGoBack}>
                    Go Back
                </Button>
        </Container>
    );
};

export default NoMatchPage;