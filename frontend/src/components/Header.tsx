
import { AppBar, Box, Button, Divider, IconButton, Toolbar, Typography, useTheme } from "@mui/material";
import SearchBar from "./SearchBar";
import Logo from '../assets/logo_light.svg';
import { Link } from "react-router-dom";
import React from "react";

export const LogoButton = () => (
    <IconButton component={Link} to="/" edge="start" color="inherit" aria-label="logo" sx={{ ml: 1 }}>
        <img src={Logo} alt="Logo" height="40px" />
    </IconButton>
);

const authOptions = { 'Sign in': '/sign-in', 'Sign up': '/sign-up' }

export const AuthButtons = () => (
    <Box ml="auto" display="flex" flexDirection="row" alignItems="center" sx={{ gap: '10px' }}>
        {Object.entries(authOptions).map(([key, value]) => (
            <Button
                key={key}
                component={Link}
                to={value}
                color="inherit"
                variant="outlined"
                size="medium"
                sx={{ height: '3em', borderRadius: '10px' }}
            >
                {key}
            </Button>
            ))
        }
    </Box>
);

const navOptions = {
    'Offer': '/', 'Genres': '/movies/genres', 'Repertoires': `/repertoires/date`
};

export const Navigation = () => {
    const theme = useTheme();
    return (
        <Box display="flex" justifyContent="center" width="100%" fontSize="2rem" sx={{ gap: "20px" }}>
            {Object.entries(navOptions).map(([key, value]) => (
                <React.Fragment key={key}>
                    <Divider orientation="vertical" flexItem />
                    <Button component={Link} to={value}>
                        <Typography variant="h6" sx={{ fontWeight: 'bold' }} style={{ color: theme.palette.primary.contrastText }}>
                            {key}
                        </Typography>
                    </Button>
                </React.Fragment>
            ))}
        </Box>
    );
};



export default function Header() {
    const theme = useTheme();
    return (
        <>
            <AppBar position="fixed">
                <AppBar position="static">
                    <Toolbar>
                        <LogoButton />
                        <AuthButtons />
                        <SearchBar />
                    </Toolbar>
                </AppBar>
                <AppBar position="static" style={{ backgroundColor: theme.palette.secondary.main, color: theme.palette.primary.contrastText }}>
                    <Toolbar>
                        <Navigation />
                    </Toolbar>
                </AppBar>
            </AppBar>
            {/* SOLVE THAT PROBLEM WITHOUT THIS SPACER */}
            <div style={{ height: '128px' }}></div>
        </>
    );
};
