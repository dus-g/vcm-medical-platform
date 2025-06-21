export interface User {
  cd_user: number;
  email: string;
  first_name: string;
  last_name: string;
  ty_user: number;
  user_status: string;
  phone_number?: string;
  date_of_birth?: string;
  gender?: string;
  created_at: string;
  // Add computed properties for backwards compatibility
  name?: string;
  userType?: number;
}

export interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  loading: boolean;
  login: (email: string, password: string) => Promise<void>;
  register: (data: RegisterData) => Promise<void>;
  verifyOTP: (email: string, otp: string) => Promise<void>;
  logout: () => void;
  setUser: (user: User) => void;
}

export interface RegisterData {
  email: string;
  password: string;
  confirmPassword: string;
  first_name: string;
  last_name: string;
  ty_user: number;
}

export interface LoginData {
  email: string;
  password: string;
}

export interface ApiResponse<T = any> {
  success: boolean;
  data: T;
  message?: string;
  error?: string;
}
