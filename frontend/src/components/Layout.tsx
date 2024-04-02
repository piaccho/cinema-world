import { Outlet } from 'react-router-dom';
import Footer from './Footer';
import Header from './Header';
import { Box } from '@mui/material';


const Layout: React.FC = () => {
    return (
        <Box sx={{
            display: 'flex',
            flexDirection: 'column',
            minHeight: '100vh',
        }}>
            <Header />
            <Box sx={{ flex: '1' }}>
                <Outlet />
            </Box>
            <Footer />
        </Box>
    );
};

export default Layout;
