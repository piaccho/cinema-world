import * as React from 'react';
import { useForm } from 'react-hook-form';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { Alert, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Stack } from '@mui/material';
import ApiService from '../ApiService';
import IconButton from '@mui/material/IconButton';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import FormControl from '@mui/material/FormControl';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import SentimentSatisfiedAltIcon from '@mui/icons-material/SentimentSatisfiedAlt';
import theme from '../theme';
import { grey } from '@mui/material/colors';

const firstNameRegex = /^[a-zA-ZÀ-ÿ '-]{2,40}$/;
const lastNameRegex = /^([A-ZÀ-ÿ]+|\b[A-ZÀ-ÿ][a-zÀ-ÿ]*){1,3}$/;
const emailRegex = /^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}$/i;
// const passwordRegex = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?!.*\s).{8,16}$/;
const passwordRegex = /^[A-Za-z0-9]{8,}$/; // less strict

const errorInfoMap = {
    'firstName': 'First name must be between 2 and 40 characters long.',
    'lastName': 'Last name must be between 2 and 40 characters long.',
    'email': 'Invalid email address.',
    'password': 'Password must be at least 8 characters long.',
};


export default function SignUpPage() {
    const apiService = new ApiService();
    const navigate = useNavigate();
    const { register, handleSubmit, formState: { errors } } = useForm({
        mode: 'all',
        criteriaMode: 'all',
        shouldFocusError: true,
    });
    const [open, setOpen] = React.useState(false);

    const [showPassword, setShowPassword] = React.useState(false);
    const [registerError, setRegisterError] = React.useState<string | null>(null);

    const handleClickShowPassword = () => setShowPassword((show) => !show);

    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const onSubmit = async (data: any) => {
        try {
            const response = await apiService.register(data.firstName, data.lastName, data.email, data.password);
            if (response.status === 200) {
                setOpen(true);
            } 
        } catch (error: any) {
            setRegisterError(error.message);
        }
    };

    const handleClose = () => {
        setOpen(false);
        navigate('/sign-in');
    };

    return (
        <Container component="main" maxWidth="lg" sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
        }}>
            <Stack
                mt={5}
                mb={3}
                direction="column"
                justifyContent="center"
                alignItems="center"
                maxWidth="30vw"
            >
                <Typography component="h1" variant="h4" sx={{ color: theme.palette.primary.main, fontWeight: 'bold' }}>
                    Sign up
                </Typography>
                <Box component="form" noValidate onSubmit={handleSubmit(onSubmit)} sx={{ mt: 3}}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                {...register("firstName", { required: true, minLength: 2, maxLength: 50, pattern: firstNameRegex})}
                                autoComplete="given-name"
                                name="firstName"
                                required
                                fullWidth
                                id="firstName"
                                label="First Name"
                                autoFocus
                            />
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                {...register("lastName", { required: true, minLength: 2, maxLength: 50, pattern: lastNameRegex})}
                                required
                                fullWidth
                                id="lastName"
                                label="Last Name"
                                name="lastName"
                                autoComplete="family-name"
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                {...register("email", {
                                    required: true, pattern: emailRegex})}
                                required
                                fullWidth
                                id="email"
                                label="Email Address"
                                name="email"
                                autoComplete="email"
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <FormControl sx={{ width: '100%' }} variant="outlined">
                                <InputLabel htmlFor="password">Password</InputLabel>
                                <OutlinedInput
                                    {...register("password", { required: true, pattern: passwordRegex})}
                                    id="password"
                                    type={showPassword ? 'text' : 'password'}
                                    endAdornment={
                                        <InputAdornment position="end">
                                            <IconButton
                                                aria-label="toggle password visibility"
                                                onClick={handleClickShowPassword}
                                                onMouseDown={handleMouseDownPassword}
                                                edge="end"
                                            >
                                                {showPassword ? <VisibilityOff /> : <Visibility />}
                                            </IconButton>
                                        </InputAdornment>
                                    }
                                    label="Password"
                                />
                            </FormControl>
                        </Grid>
                    </Grid>
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Sign Up
                    </Button>
                    <Grid container justifyContent="center">
                        <Grid item>
                            <Link component={RouterLink} to="/sign-in" variant="body2">
                                Already have an account? Sign in
                            </Link>
                        </Grid>
                    </Grid>
                </Box>
            </Stack>
            
            <Dialog
                open={open}
                onClose={handleClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
                sx={{ display: 'flex', justifyContent: 'center', flexDirection: 'column', alignItems: 'center' }}
                
            >
                <DialogTitle id="alert-dialog-title" sx={{ color: theme.palette.primary.main, fontWeight: 'bold', mt: 2}}>
                    Successful Registration
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        <Box gap={1} display='flex' flexDirection='row' alignItems='center' sx={{ color: grey[600] }}>
                            You have successfully registered. You can now login to your account. 
                            <SentimentSatisfiedAltIcon /> 
                        </Box>
                    </DialogContentText>
                </DialogContent>
                <DialogActions
                    sx={{ mb: 3, display: 'flex', justifyContent: 'center', flexDirection: 'row', alignItems: 'center' }}
                >
                    <Button onClick={handleClose} color="primary" variant='contained' autoFocus>
                        Sign In
                    </Button>
                </DialogActions>
            </Dialog>
            {/* VALIDATION INFO */}
            <Box width="30vw" mb={5}>
            {Object.keys(errors).length > 0 ? (
                <Alert severity="error">
                    {Object.entries(errors).map(([key, value]) => (
                        <Box key={key}>* {errorInfoMap[key]}</Box>
                    ))}
                </Alert>
            ) : registerError ? (
                        <Alert severity="error">Registration failed. A user with such data exists.</Alert>
            ): (
                <Alert severity="success">Everything is correct.</Alert>
            )}
            </Box>
            
        </Container>
    );
}