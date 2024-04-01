import { Box, CircularProgress, Typography } from '@mui/material';
import React from 'react';
import { css } from '@emotion/react';

const Header: React.FC = () => (
    <Box sx={{
        border: '1px dotted grey',
        flex: '0 1 auto',
    }}>
        <Typography variant="body1">
            <span><b>Welcome to Test Page</b><CircularProgress color='info' /></span>
            <br />(sized to content)
        </Typography>
    </Box>
);

const Content: React.FC = () => (
    <Box sx={{
        border: '1px dotted grey',
        display: 'flex',
        flexDirection: 'column',
        flex: '1 1 auto',
        overflowY: 'auto',
    }}>
        {Array.from({ length: 100 }, (_, i) => (
            <Typography variant="body1" key={i}>
                <b>content</b> (fills remaining space)
            </Typography>
        ))}
    </Box>
);

const Footer: React.FC = () => (
    <Box sx={{
        border: '1px dotted grey',
        flex: '0 1 40px',
    }}>
        <Typography variant="body1">
            <b>footer</b> (fixed height)
        </Typography>
    </Box>
);

const LeftSpace: React.FC = () => (
    <Box sx={{
        flex: '1 0 auto',
    }}>
        <Typography variant="body1">
            <b>body</b> 
        </Typography>
    </Box>
);

const TestPage: React.FC = () => {
    return (
        <Box sx={{
            display: 'flex',
            flexDirection: 'column',
            height: '100vh',
            margin: 0,
        }}>
            <Header />
            <LeftSpace />
            {/* <Content /> */}
            <Footer />
        </Box>
    );
};

export default TestPage;