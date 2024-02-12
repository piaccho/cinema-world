import { createTheme } from '@mui/material/styles';
import { orange, red, grey, green } from '@mui/material/colors';



// A custom theme for this app
const theme = createTheme({
  // palette 1
      // primary: {
      //   main: "#2F4460",
      //   light: "#5175A4",
      //   dark: "#141D29",
      //   contrastText: "#FFFFFF"
      // },
      // secondary: {
      //   main: "#BFCDE0",
      //   light: "#FEFCFD",
      //   dark: "#7593BD",
      //   contrastText: "#000000"
      // },
 
  // palette 2
      // primary: {
      //   main: "#634832",
      //   light: "#967259",
      //   dark: "#38220f",
      //   contrastText: "#FFFFFF"
      // },
      // secondary: {
      //   main: "#dbc1ac",
      //   light: "#ece0d1",
      //   dark: "#967259",
      //   contrastText: "#000000"
      // },

  palette: {
    primary: {
      main: "#634832", 
      light: "#967259",
      dark: "#38220f",
      contrastText: "#FFFFFF"
    },
    secondary: {
      main: "#dbc1ac",
      light: "#ece0d1",
      dark: "#967259",
      contrastText: "#000000"
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