import { Link, useTheme } from '@mui/material';
import { Box, Container, Grid, Typography } from "@mui/material";
import FacebookIcon from '@mui/icons-material/Facebook';
import InstagramIcon from '@mui/icons-material/Instagram';
import XIcon from '@mui/icons-material/X';
import YoutubeIcon from '@mui/icons-material/YouTube';
import { grey } from '@mui/material/colors';

const footerColumns = [
    {
        title: 'ABOUT US',
        links: ['CinemaWorld', 'Newsletter', 'Contact']
    },
    {
        title: 'INFORMATIONS',
        links: ['Regulations', 'Privacy Policy', 'Manage cookies', 'Cookies Policy']
    },
    {
        title: 'FOLLOW US',
        links: [
            { name: 'Facebook', icon: <FacebookIcon /> },
            { name: 'Instagram', icon: <InstagramIcon /> },
            { name: 'Twitter', icon: <XIcon /> },
            { name: 'YouTube', icon: <YoutubeIcon /> }
        ]
    }
];

export default function Footer() {
    const theme = useTheme();

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
                        {footerColumns.map((column, index) => (
                            <Grid key={index} item xs={3} container direction="column">
                                <Typography variant="h6" mb={1} sx={{ fontWeight: 'bold' }}>{column.title}</Typography>
                                {column.links.map((link, linkIndex) => (
                                    typeof link === 'string' ?
                                        <Link key={linkIndex} variant='body1' href="#" underline="hover" style={{ color: theme.palette.primary.contrastText }}>{link}</Link>
                                        :
                                        <Link key={linkIndex} variant='body1' href="#" underline="hover" style={{ display: 'flex', alignItems: 'center', color: theme.palette.primary.contrastText }}>{link.icon}{link.name}</Link>
                                ))}
                            </Grid>
                        ))}
                    </Grid>
                    <Grid item container direction="row" justifyContent="center">
                        <Typography color={grey[700]} variant="subtitle1">
                            All rights reserved Cinema World {new Date().getFullYear()} ©
                        </Typography>
                    </Grid>
                </Grid>
            </Container>
        </Box>   
    );
};
