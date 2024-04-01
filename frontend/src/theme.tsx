import { PaletteMode } from '@mui/material';
import { createTheme } from '@mui/material/styles';
import { orange, red, grey, green } from '@mui/material/colors';

const getDesignTokens = (mode: PaletteMode) => ({
  palette: {
    mode,
    ...(mode === 'light'
      ? {
        // palette values for light mode
        primary: {
          main: "#5E2775",
          light: "#835D8F",
          dark: "#3A015C",
          contrastText: "#FFFFFF"
        },
        text: {
          primary: "rgba(0, 0, 0, 0.87)",
          secondary: "rgba(0, 0, 0, 0.6)",
          disabled: "rgba(0, 0, 0, 0.38)",
        },
      }
      : {
        // palette values for dark mode
        primary: {
          main: "#461261",
          light: "#6D3F7F",
          dark: "#2D0843",
          contrastText: "rgba(0, 0, 0, 0.87)"
        },
        text: {
          primary: "#fff",
          secondary: "rgba(255, 255, 255, 0.7)",
          disabled: "rgba(255, 255, 255, 0.5)",
        },
        
      }),
  },
  typography: {
    fontFamily: [
      'Noto Sans',
      'Lora',
      'Arial',
      'sans-serif'
    ].join(','),
  },
});

const theme = createTheme({
  palette: {
    primary: {
      main: "#5E2775",
      light: "#835D8F",
      dark: "#3A015C",
      contrastText: "#FFFFFF"
    },
    secondary: {
      main: "#461261",
      light: "#6D3F7F",
      dark: "#2D0843",
      contrastText: "rgba(0, 0, 0, 0.87)"
    },
    error: {
      main: red.A400,
    },
    warning: {
      main: orange.A400,
    },
    info: {
      main: grey.A400,
    },
    success: {
      main: green.A400,
    }
  },
  typography: {
    fontFamily: [
      'Noto Sans',
      'Lora',
      'Arial',
      'sans-serif'
    ].join(','),
  },
});

export default theme;
