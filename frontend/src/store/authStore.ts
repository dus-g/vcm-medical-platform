import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { AuthState, User, RegisterData } from '../types';
import api from '../utils/api';

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      loading: false,

      login: async (email: string, password: string) => {
        set({ loading: true });
        try {
          const response = await api.post('/v1/auth/login', { email, password });
          const { user, token } = response.data;
          
          // Add computed properties for backwards compatibility
          const userWithComputed = {
            ...user,
            name: `${user.first_name} ${user.last_name}`.trim(),
            userType: user.ty_user
          };
          
          set({ user: userWithComputed, token, isAuthenticated: true, loading: false });
          localStorage.setItem('token', token);
        } catch (error: any) {
          set({ loading: false });
          throw new Error(error.response?.data?.message || 'Login failed');
        }
      },

      register: async (data: RegisterData) => {
        set({ loading: true });
        try {
          const response = await api.post('/v1/auth/register', data);
          set({ loading: false });
          return response.data;
        } catch (error: any) {
          set({ loading: false });
          throw new Error(error.response?.data?.message || 'Registration failed');
        }
      },

      verifyOTP: async (email: string, otp: string) => {
        set({ loading: true });
        try {
          const response = await api.post('/v1/auth/verify-otp', { email, otp });
          const { user, token } = response.data;
          
          // Add computed properties for backwards compatibility
          const userWithComputed = {
            ...user,
            name: `${user.first_name} ${user.last_name}`.trim(),
            userType: user.ty_user
          };
          
          set({ user: userWithComputed, token, isAuthenticated: true, loading: false });
          localStorage.setItem('token', token);
        } catch (error: any) {
          set({ loading: false });
          throw new Error(error.response?.data?.message || 'OTP verification failed');
        }
      },

      logout: () => {
        set({ user: null, token: null, isAuthenticated: false });
        localStorage.removeItem('token');
      },

      setUser: (user: User) => {
        const userWithComputed = {
          ...user,
          name: `${user.first_name} ${user.last_name}`.trim(),
          userType: user.ty_user
        };
        set({ user: userWithComputed });
      },
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({ 
        user: state.user, 
        token: state.token, 
        isAuthenticated: state.isAuthenticated 
      }),
    }
  )
);
