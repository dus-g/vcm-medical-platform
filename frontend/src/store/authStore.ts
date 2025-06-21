import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import api from '../utils/api';

export interface User {
  cdUser: number;
  email: string;
  firstName: string;
  lastName: string;
  userStatus: string;
  tyUser: number; // User type: 1=Patient, 5=Doctor, etc.
  phoneNumber?: string;
  gender?: string;
  dateOfBirth?: string;
  cdCountry?: number;
  cdState?: number;
  cdCity?: number;
  streetAddress?: string;
  postalCode?: string;
  profileComplete: boolean;
  createdAt: string;
  updatedAt: string;
}

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
  
  // Actions
  login: (email: string, password: string) => Promise<void>;
  register: (data: RegisterData) => Promise<void>;
  verifyOTP: (email: string, otpCode: string) => Promise<void>;
  completeProfile: (data: ProfileData) => Promise<void>;
  logout: () => void;
  clearError: () => void;
  setUser: (user: User) => void;
  setToken: (token: string) => void;
}

interface RegisterData {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  userType: number;
  phoneNumber?: string;
  gender?: string;
  dateOfBirth?: string;
  countryId?: number;
  stateId?: number;
  cityId?: number;
}

interface ProfileData {
  phoneNumber: string;
  gender: string;
  dateOfBirth: string;
  countryId: number;
  stateId: number;
  cityId: number;
  streetAddress?: string;
  postalCode?: string;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      isLoading: false,
      error: null,

      login: async (email: string, password: string) => {
        set({ isLoading: true, error: null });
        try {
          const response = await api.login({ email, password });
          
          if (response.token) {
            localStorage.setItem('vcm-token', response.token);
            set({
              user: response.user,
              token: response.token,
              isAuthenticated: true,
              isLoading: false,
            });
          }
        } catch (error) {
          set({ 
            isLoading: false, 
            error: error instanceof Error ? error.message : 'Login failed' 
          });
          throw error;
        }
      },

      register: async (data: RegisterData) => {
        set({ isLoading: true, error: null });
        try {
          await api.register(data);
          set({ isLoading: false });
        } catch (error) {
          set({ 
            isLoading: false, 
            error: error instanceof Error ? error.message : 'Registration failed' 
          });
          throw error;
        }
      },

      verifyOTP: async (email: string, otpCode: string) => {
        set({ isLoading: true, error: null });
        try {
          const response = await api.verifyOTP({ email, otpCode });
          
          if (response.token) {
            localStorage.setItem('vcm-token', response.token);
            set({
              user: response.user,
              token: response.token,
              isAuthenticated: true,
              isLoading: false,
            });
          }
        } catch (error) {
          set({ 
            isLoading: false, 
            error: error instanceof Error ? error.message : 'OTP verification failed' 
          });
          throw error;
        }
      },

      completeProfile: async (data: ProfileData) => {
        set({ isLoading: true, error: null });
        try {
          const response = await api.completeProfile(data);
          set({
            user: response.user,
            isLoading: false,
          });
        } catch (error) {
          set({ 
            isLoading: false, 
            error: error instanceof Error ? error.message : 'Profile completion failed' 
          });
          throw error;
        }
      },

      logout: () => {
        localStorage.removeItem('vcm-token');
        set({
          user: null,
          token: null,
          isAuthenticated: false,
          error: null,
        });
      },

      clearError: () => {
        set({ error: null });
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
