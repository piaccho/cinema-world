import { Routes, Route, } from "react-router-dom";
import Layout from './components/Layout';
import HomePage from './pages/HomePage';
import NoMatchPage from './pages/NoMatchPage';
import SignInPage from './pages/SignInPage';
import SignUpPage from './pages/SignUpPage';
import MovieOverviewPage from "./pages/MovieOverviewPage";
import ChooseGenrePage from "./pages/ChooseGenrePage";
import MoviesListPage from "./pages/MoviesListPage";
import RepertoirePage from "./pages/RepertoirePage";

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<HomePage />} />
        <Route path="sign-in" element={<SignInPage />} />
        <Route path="sign-up" element={<SignUpPage />} />
        <Route path="movies/:id" element={<MovieOverviewPage/>} />
        <Route path="genres" element={<ChooseGenrePage />} />
        <Route path="repertoires" element={<RepertoirePage />} />
        <Route path="repertoires/date/:date" element={<RepertoirePage />} />
        <Route path="movies/genres/:name" element={<MoviesListPage type="genre"/>} />
        <Route path="movies/search/:q" element={<MoviesListPage type="searchQuery"/>} />
        <Route path="*" element={<NoMatchPage />} />
      </Route>
    </Routes>
  );
}

export default App;
