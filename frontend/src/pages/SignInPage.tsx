import * as React from 'react';
import { useForm } from 'react-hook-form';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import ApiService from '../ApiService';
import { Alert, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, useTheme } from '@mui/material';
import { Visibility, VisibilityOff } from '@mui/icons-material';

const signInInfo = "Enter your email address and password to sign in.";

export default function SignInPage() {
    const theme = useTheme();
    const apiService = new ApiService();
    const navigate = useNavigate();
    const { register, handleSubmit, formState: { errors } } = useForm({
        mode: 'onSubmit',
        criteriaMode: 'all',
        shouldFocusError: true,
    });
    const [showPassword, setShowPassword] = React.useState(false);
    const [loginError, setLoginError] = React.useState<string | null>(null);

    const handleClickShowPassword = () => setShowPassword((show) => !show);

    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const onSubmit = async (data: any) => {
        try {
            const response = await apiService.login(data.email, data.password);
            if (response.status === 200) {
                // Handle JWT token
                localStorage.setItem('token', response.data.token);
                navigate('/');
            }
        } catch (error: any) {
            setLoginError(error.message);
        }
    };

    return (
        <Container component="main" maxWidth="xs">
            <Box
                my={5}
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography component="h1" variant="h4" sx={{ color: theme.palette.primary.main, fontWeight: 'bold' }}>
                    Sign in
                </Typography>
                {signInInfo && <Typography component="h1" variant="subtitle1" sx={{ color: theme.palette.primary.light, mt: 2, textAlign: 'center' }}>
                    {signInInfo}
                </Typography>}
                <Box component="form" noValidate onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
                    <TextField
                        {...register("email", { required: true })}
                        margin="normal"
                        required
                        fullWidth
                        id="email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                        autoFocus
                    />
                    <FormControl sx={{ width: '100%' }} variant="outlined">
                        <InputLabel htmlFor="password">Password</InputLabel>
                        <OutlinedInput
                            {...register("password", { required: true })}
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
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Sign In
                    </Button>
                    <Grid mb={3} container>
                        <Grid item>
                            <Link component={RouterLink} to="/sign-up" variant="body2">
                                Don't have an account? Sign Up
                            </Link>
                        </Grid>
                    </Grid>
                    {(errors.email || errors.password || loginError) && <Box mb={5}>
                        <Alert severity="error">
                            Invalid email or password.
                        </Alert>
                    </Box>}
                    
                </Box>
            </Box>
        </Container>
    );
}