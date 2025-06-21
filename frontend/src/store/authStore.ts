import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import api from '../utils/api';

export interface User {
  cdUser: number;
  email: string;
  firstName: string;
  lastName: string;
  userStatus: string;
  tyUser: number; // User type (1=Patient, 5=Doctor, etc.)
  phoneNumber: string;
  gender: string;
  dateOfBirth: string;
  countryId: number;
  stateId: number;
  cityId: number;
  createdAt: string;
  updatedAt: string;
}

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
  register: (data: RegisterData) => Promise<void>;
  verifyOTP: (email: string, otpCode: string) => Promise<void>;
  setUser: (user: User) => void;
  setToken: (token: string) => void;
}

interface RegisterData {
  email: string;
  password: string;
  userType: number;
  firstName: string;
  lastName: string;
  phoneNumber: string;
  gender: string;
  dateOfBirth: string;
  countryId: number;
  stateId: number;
  cityId: number;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      isLoading: false,

      login: async (email: string, password: string) => {
        set({ isLoading: true });
        try {
          const response = await api.login({ email, password });
          
          // Store token in localStorage and state
          localStorage.setItem('vcm-token', response.token);
          
          set({
            user: response.user,
            token: response.token,
            isAuthenticated: true,
            isLoading: false,
          });
        } catch (error) {
          set({ isLoading: false });
          throw error;
        }
      },

      register: async (data: RegisterData) => {
        set({ isLoading: true });
        try {
          await api.register(data);
          set({ isLoading: false });
        } catch (error) {
          set({ isLoading: false });
          throw error;
        }
      },

      verifyOTP: async (email: string, otpCode: string) => {
        set({ isLoading: true });
        try {
          const response = await api.verifyOTP({ email, otpCode });
          
          // Store token and update user
          localStorage.setItem('vcm-token', response.token);
          
          set({
            user: response.user,
            token: response.token,
            isAuthenticated: true,
            isLoading: false,
          });
        } catch (error) {
          set({ isLoading: false });
          throw error;
        }
      },

      logout: () => {
        localStorage.removeItem('vcm-token');
        set({
          user: null,
          token: null,
          isAuthenticated: false,
        });
      },

      setUser: (user: User) => {
        set({ user, isAuthenticated: true });
      },

      setToken: (token: string) => {
        localStorage.setItem('vcm-token', token);
        set({ token });
      },
    }),
    {
      name: 'vcm-auth-storage',
      partialize: (state) => ({
        user: state.user,
        token: state.token,
        isAuthenticated: state.isAuthenticated,
      }),
    }
  )
);
