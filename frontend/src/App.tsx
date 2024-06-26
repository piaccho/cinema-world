import Layout from './components/Layout';
import HomePage from './pages/HomePage';
import NoMatchPage from './pages/NoMatchPage';
import SignInPage from './pages/SignInPage';
import SignUpPage from './pages/SignUpPage';
import MovieOverviewPage from "./pages/MovieOverviewPage";
import ChooseGenrePage from "./pages/ChooseGenrePage";
import MoviesResultsPage from "./pages/MoviesResultsPage";
import RepertoirePage from "./pages/RepertoirePage";
import TestPage from "./pages/TestPage";

import React from "react";
import { Routes, Route, Navigate, } from "react-router-dom";
import dayjs from 'dayjs';


const App: React.FC = () => {
  // const [mode, setMode] = React.useState<PaletteMode>('light');
  // const colorMode = React.useMemo(
  //   () => ({
  //     // The dark mode switch would invoke this method
  //     toggleColorMode: () => {
  //       setMode((prevMode: PaletteMode) =>
  //         prevMode === 'light' ? 'dark' : 'light',
  //       );
  //     },
  //   }),
  //   [],
  // );
  // // Update the theme only if the mode changes
  // const theme = React.useMemo(() => createTheme(getDesignTokens(mode)), [mode]);

  return (
    <Routes>
      <Route path="/test" element={<TestPage />} />
      <Route path="/" element={<Layout />}>
        <Route index element={<HomePage />} />
        <Route path="sign-in" element={<SignInPage />} />
        <Route path="sign-up" element={<SignUpPage />} />
        <Route path="movies/:id" element={<MovieOverviewPage />} />
        <Route path="movies/genres" element={<ChooseGenrePage />} />
        <Route path="repertoires/date/" element={<Navigate to={`${dayjs().format('YYYY-MM-DD')}`}/>} />
        <Route path="repertoires/date/:date" element={<RepertoirePage />} />
        <Route path="movies/genres/:genreName" element={<MoviesResultsPage />} />
        <Route path="movies/search/:searchQuery" element={<MoviesResultsPage />} />
        <Route path="*" element={<NoMatchPage />} />
      </Route>
    </Routes>
  );
}

export default App;
