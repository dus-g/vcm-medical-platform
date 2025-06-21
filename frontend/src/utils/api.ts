const API_BASE_URL = process.env.NODE_ENV === 'production' 
  ? '/api/v1' 
  : 'http://localhost:8080/api/v1';

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

  async request(endpoint: string, options: RequestInit = {}) {
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

  // Auth methods matching your backend
  async register(data: { 
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
  }) {
    return this.request('/auth/register', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async verifyOTP(data: { email: string; otpCode: string }) {
    return this.request('/auth/verify-otp', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async login(data: { email: string; password: string }) {
    return this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getProfile() {
    return this.request('/auth/me');
  }

  // Location methods matching your backend
  async getCountries() {
    const response = await this.request('/location/countries');
    return response.countries; // Your backend returns { countries: [...] }
  }

  async getStates(countryId: number) {
    const response = await this.request(`/location/states/${countryId}`);
    return response.states;
  }

  async getCities(countryId: number, stateId: number) {
    const response = await this.request(`/location/cities/${countryId}/${stateId}`);
    return response.cities;
  }

  async resendOTP(data: { email: string }) {
    return this.request('/auth/resend-otp', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }
}

export const api = new ApiClient();
export default api;
