
import { AppBar, Box, Button, Divider, IconButton, Toolbar, Typography, useTheme } from "@mui/material";
import SearchBar from "./SearchBar";
import Logo from '../assets/logo_light.svg';
import { Link } from "react-router-dom";

export default function Header() {
    const theme = useTheme();
    return (
        <>
            <AppBar position="fixed">
                <AppBar position="static">
                    <Toolbar>
                        <IconButton component={Link} to="/" edge="start" color="inherit" aria-label="logo">
                            <img src={Logo} alt="Logo" height="40px" />
                        </IconButton>
                        <Box ml="auto" display="flex" flexDirection="row" sx={{ gap: '10px' }}>
                            <Button component={Link} to="/sign-in" color="inherit">Sign in</Button>
                            <Button component={Link} to="/sign-up" color="inherit">Sign up</Button>
                            <SearchBar />
                        </Box>
                    </Toolbar>
                </AppBar>
                <AppBar position="static" style={{ backgroundColor: theme.palette.secondary.main, color: theme.palette.primary.contrastText }}>
                    <Toolbar>
                        <Box display="flex" justifyContent="center" width="100%" fontSize="2rem" sx={{ gap: "20px" }}>
                            <Divider orientation="vertical" flexItem />
                            <Button component={Link} to="/">
                                <Typography variant="h6" sx={{fontWeight: 'bold'}}>
                                    Offer
                                </Typography>
                            </Button>
                            <Divider orientation="vertical" flexItem />
                            <Button component={Link} to="/genres"><Typography variant="h6" sx={{fontWeight: 'bold'}}>Genres</Typography></Button>
                            <Divider orientation="vertical" flexItem />
                            <Button component={Link} to="/repertoires"><Typography variant="h6" sx={{fontWeight: 'bold'}}>Repertoires</Typography></Button>
                            <Divider orientation="vertical" flexItem />
                        </Box>
                    </Toolbar>
                </AppBar>
            </AppBar>
            {/* SOLVE THAT PROBLEM WITHOUT THIS SPACER */}
            <div style={{ height: '128px' }}></div>
        </>
    );
};
