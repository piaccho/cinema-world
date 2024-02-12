import React, { useState } from 'react';
import { styled, alpha } from '@mui/material/styles';
import SearchIcon from '@mui/icons-material/Search';
import InputBase from '@mui/material/InputBase';
import { useNavigate } from 'react-router-dom';
import { Box, IconButton } from '@mui/material';

const Search = styled('div')(({ theme }) => ({
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    '&:hover': {
        backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    marginLeft: 0,
    width: '100%',
    [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: 'auto',
    },
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
    color: 'inherit',
    width: '100%',
    '& .MuiInputBase-input': {
        padding: theme.spacing(1, 1, 1, 0),
        // vertical padding + font size from searchIcon
        paddingLeft: `calc(1em + ${theme.spacing(1)})`,
        transition: theme.transitions.create('width'),
        [theme.breakpoints.up('sm')]: {
            width: '15ch',
            '&:focus': {
                width: '30ch',
            },
        },
    },
}));

const SearchBar: React.FC = () => {
    const [searchQuery, setSearchQuery] = useState('');
    const navigate = useNavigate();

    const handleSearch = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter' && searchQuery !== '') {
            navigate(`/movies/search/${searchQuery}`);
        }
    };
    const handleClick = () => {
        if (searchQuery !== '') {
            navigate(`/movies/search/${searchQuery}`);
        }
    };


    return (
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Search>
                <StyledInputBase
                    placeholder="Search movie…"
                    inputProps={{ 'aria-label': 'search' }}
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    onKeyDown={handleSearch}
                />
            </Search>
            <IconButton onClick={handleClick} size="large" aria-label="search" color="inherit">
                <SearchIcon />
            </IconButton>
        </Box>
    );
};

export default SearchBar;
