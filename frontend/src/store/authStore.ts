import { create } from 'zustand';

export interface User {
  id: string;
  name: string;
  email: string;
  userType: number; // 1=Patient, 5=Doctor, etc.
  phone?: string;
  avatar?: string;
}

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  login: (user: User) => void;
  logout: () => void;
  updateUser: (userData: Partial<User>) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  isAuthenticated: false,
  login: (user) => set({ user, isAuthenticated: true }),
  logout: () => set({ user: null, isAuthenticated: false }),
  updateUser: (userData) => set((state) => ({
    user: state.user ? { ...state.user, ...userData } : null
  })),
}));
