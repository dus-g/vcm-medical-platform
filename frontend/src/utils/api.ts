const API_BASE_URL = typeof window !== 'undefined' && window.location.hostname === 'localhost'
  ? 'http://localhost:8080/api/v1' 
  : '/api/v1';

interface ApiResponse {
  [key: string]: any;
}

class ApiClient {
  private getHeaders(): HeadersInit {
    const token = localStorage.getItem('vcm-token');
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
    };
    
    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }
    
    return headers;
  }

  async request(endpoint: string, options: RequestInit = {}): Promise<ApiResponse> {
    const url = `${API_BASE_URL}${endpoint}`;
    const config: RequestInit = {
      headers: this.getHeaders(),
      ...options,
    };

    const response = await fetch(url, config);
    
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ 
        error: `HTTP error! status: ${response.status}` 
      }));
      throw new Error(errorData.error || errorData.message || 'Network error');
    }

    return response.json();
  }

  // HTTP method helpers
  async get(endpoint: string): Promise<ApiResponse> {
    return this.request(endpoint, { method: 'GET' });
  }

  async post(endpoint: string, data?: any): Promise<ApiResponse> {
    return this.request(endpoint, {
      method: 'POST',
      body: data ? JSON.stringify(data) : undefined,
    });
  }

  async put(endpoint: string, data?: any): Promise<ApiResponse> {
    return this.request(endpoint, {
      method: 'PUT',
      body: data ? JSON.stringify(data) : undefined,
    });
  }

  async delete(endpoint: string): Promise<ApiResponse> {
    return this.request(endpoint, { method: 'DELETE' });
  }

  // Auth methods
  async register(data: {
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
  }) {
    return this.post('/auth/register', data);
  }

  async verifyOTP(data: { email: string; otpCode: string }) {
    return this.post('/auth/verify-otp', data);
  }

  async login(data: { email: string; password: string }) {
    return this.post('/auth/login', data);
  }

  async completeProfile(data: {
    phoneNumber: string;
    gender: string;
    dateOfBirth: string;
    countryId: number;
    stateId: number;
    cityId: number;
    streetAddress?: string;
    postalCode?: string;
  }) {
    return this.post('/auth/complete-profile', data);
  }

  async getProfile() {
    return this.get('/auth/me');
  }

  async updateProfile(data: any) {
    return this.put('/auth/profile', data);
  }

  async resendOTP(data: { email: string }) {
    return this.post('/auth/resend-otp', data);
  }

  // Location methods
  async getCountries() {
    const response = await this.get('/location/countries');
    return response.countries || [];
  }

  async getStates(countryId: number) {
    const response = await this.get(`/location/states/${countryId}`);
    return response.states || [];
  }

  async getCities(countryId: number, stateId: number) {
    const response = await this.get(`/location/cities/${countryId}/${stateId}`);
    return response.cities || [];
  }
}

export const api = new ApiClient();
export default api;
