import { create } from 'zustand';

type Store = {
  // MovieList
  genreId: number | null;
  setGenreId: (id: number) => void;
  genreName: string | null;
  setGenreName: (name: string) => void;
  searchQuery: string | null;
  setSearchQuery: (query: string) => void;
  // Repertoire
  date: string | null;
  setDate: (date: string) => void;
};

export const useStore = create<Store>((set) => ({
  // MovieList
  genreId: null,
  setGenreId: (id: number) => set({ genreId: id }),
  genreName: null,
  setGenreName: (name: string) => set({ genreName: name }),
  searchQuery: null,
  setSearchQuery: (query: string) => set({ searchQuery: query }),
  // Repertoire
  date: null,
  setDate: (date: string) => set({ date: date }),
}));