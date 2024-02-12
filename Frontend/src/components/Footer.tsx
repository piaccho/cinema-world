import { Link } from '@mui/material';
import { Box, Container, Grid, Typography } from "@mui/material";
import FacebookIcon from '@mui/icons-material/Facebook';
import InstagramIcon from '@mui/icons-material/Instagram';
import XIcon from '@mui/icons-material/X';
import YoutubeIcon from '@mui/icons-material/YouTube';

export default function Footer() {
    return (
        <Box
            sx={{
                width: "100%",
                height: "300px",
                backgroundColor: "secondary.main",
                paddingTop: "1rem",
                paddingBottom: "1rem",
                bottom: 0,
            }}
        >
            <Container maxWidth="lg">
                <Grid container
                    direction="column"
                    justifyContent="space-between"
                    alignItems="center"
                    style={{ height: '250px' }}>
                    <Grid item container spacing={3} direction="row" justifyContent="space-around">
                        {/* extract to components - FooterColumn */}
                        <Grid item xs={3} container direction="column">
                            <Typography variant="h6" mb={1} sx={{ fontWeight: 'bold' }}>ABOUT US</Typography>
                            <Link variant='body1' href="#" underline="hover">CinemaWorld</Link>
                            <Link variant='body1' href="#" underline="hover">Newsletter</Link>
                            <Link variant='body1' href="#" underline="hover">Contact</Link>
                        </Grid>

                        <Grid item xs={3} container direction="column">
                            <Typography variant="h6" mb={1} sx={{ fontWeight: 'bold' }}>INFORMATIONS</Typography>
                            <Link variant='body1' href="#" underline="hover">Regulations</Link>
                            <Link variant='body1' href="#" underline="hover">Privacy Policy</Link>
                            <Link variant='body1' href="#" underline="hover">Manage cookies</Link>
                            <Link variant='body1' href="#" underline="hover">Cookies Policy</Link>
                        </Grid>

                        <Grid item xs={3} container direction="column">
                            <Typography variant="h6" mb={1} sx={{ fontWeight: 'bold' }}>FOLLOW US</Typography>
                            <Link variant='body1' href="#" underline="hover" style={{display: 'flex', alignItems: 'center' }}><FacebookIcon />Facebook</Link>
                            <Link variant='body1' href="#" underline="hover" style={{display: 'flex', alignItems: 'center' }}><InstagramIcon />Instagram</Link>
                            <Link variant='body1' href="#" underline="hover" style={{display: 'flex', alignItems: 'center' }}><XIcon />Twitter</Link>
                            <Link variant='body1' href="#" underline="hover" style={{display: 'flex', alignItems: 'center' }}><YoutubeIcon />YouTube</Link>
                        </Grid>
                    </Grid>
                    <Grid item container direction="row" justifyContent="center">
                        <Typography color="textSecondary" variant="subtitle1">
                            All rights reserved Cinema World {new Date().getFullYear()} ©
                        </Typography>
                    </Grid>
                </Grid>
            </Container>
        </Box>   
    );
};
