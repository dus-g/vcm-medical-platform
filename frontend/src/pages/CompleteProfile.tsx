import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuthStore } from '../store/authStore';
import api from '../utils/api';

interface Country {
  cdCountry: number;
  countryName: string;
  countryAbbr: string;
}

interface State {
  cdState: number;
  stateName: string;
  stateAbbr: string;
}

interface City {
  cdCity: number;
  cityName: string;
  cityAbbr: string;
}

const CompleteProfile: React.FC = () => {
  const navigate = useNavigate();
  const { completeProfile, isLoading, error, clearError } = useAuthStore();
  
  const [formData, setFormData] = useState({
    phoneNumber: '',
    gender: '',
    dateOfBirth: '',
    countryId: 0,
    stateId: 0,
    cityId: 0,
    streetAddress: '',
    postalCode: '',
  });

  const [countries, setCountries] = useState<Country[]>([]);
  const [states, setStates] = useState<State[]>([]);
  const [cities, setCities] = useState<City[]>([]);
  const [loadingStates, setLoadingStates] = useState(false);
  const [loadingCities, setLoadingCities] = useState(false);

  useEffect(() => {
    // Load countries on component mount
    const loadCountries = async () => {
      try {
        const countriesData = await api.getCountries();
        setCountries(countriesData);
      } catch (err) {
        console.error('Failed to load countries:', err);
      }
    };
    loadCountries();
  }, []);

  useEffect(() => {
    // Load states when country changes
    if (formData.countryId) {
      setLoadingStates(true);
      api.getStates(formData.countryId)
        .then(setStates)
        .catch(err => console.error('Failed to load states:', err))
        .finally(() => setLoadingStates(false));
      
      // Reset state and city when country changes
      setFormData(prev => ({ ...prev, stateId: 0, cityId: 0 }));
      setCities([]);
    }
  }, [formData.countryId]);

  useEffect(() => {
    // Load cities when state changes
    if (formData.countryId && formData.stateId) {
      setLoadingCities(true);
      api.getCities(formData.countryId, formData.stateId)
        .then(setCities)
        .catch(err => console.error('Failed to load cities:', err))
        .finally(() => setLoadingCities(false));
      
      // Reset city when state changes
      setFormData(prev => ({ ...prev, cityId: 0 }));
    }
  }, [formData.countryId, formData.stateId]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: name.includes('Id') ? parseInt(value) || 0 : value
    }));
    if (error) clearError();
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // Validate required fields
    if (!formData.phoneNumber || !formData.gender || !formData.dateOfBirth || 
        !formData.countryId || !formData.stateId || !formData.cityId) {
      alert('Please fill in all required fields');
      return;
    }

    try {
      await completeProfile(formData);
      navigate('/dashboard');
    } catch (err) {
      // Error is handled by the store
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-600 via-purple-600 to-blue-800 flex items-center justify-center p-4">
      <div className="max-w-2xl w-full bg-white rounded-2xl shadow-2xl p-8">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Complete Your Profile</h1>
          <p className="text-gray-600">
            Please provide additional information to complete your registration
          </p>
        </div>

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
            <p className="text-red-600 text-sm">{error}</p>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Phone Number *
              </label>
              <input
                type="tel"
                name="phoneNumber"
                value={formData.phoneNumber}
                onChange={handleChange}
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
                placeholder="Enter phone number"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Gender *
              </label>
              <select
                name="gender"
                value={formData.gender}
                onChange={handleChange}
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
              >
                <option value="">Select Gender</option>
                <option value="Male">Male</option>
                <option value="Female">Female</option>
                <option value="Other">Other</option>
                <option value="Prefer not to say">Prefer not to say</option>
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Date of Birth *
              </label>
              <input
                type="date"
                name="dateOfBirth"
                value={formData.dateOfBirth}
                onChange={handleChange}
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Country *
              </label>
              <select
                name="countryId"
                value={formData.countryId}
                onChange={handleChange}
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
              >
                <option value="">Select Country</option>
                {countries.map((country) => (
                  <option key={country.cdCountry} value={country.cdCountry}>
                    {country.countryName}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                State/Province *
              </label>
              <select
                name="stateId"
                value={formData.stateId}
                onChange={handleChange}
                required
                disabled={!formData.countryId || loadingStates}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors disabled:opacity-50"
              >
                <option value="">
                  {loadingStates ? 'Loading...' : 'Select State/Province'}
                </option>
                {states.map((state) => (
                  <option key={state.cdState} value={state.cdState}>
                    {state.stateName}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                City *
              </label>
              <select
                name="cityId"
                value={formData.cityId}
                onChange={handleChange}
                required
                disabled={!formData.stateId || loadingCities}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors disabled:opacity-50"
              >
                <option value="">
                  {loadingCities ? 'Loading...' : 'Select City'}
                </option>
                {cities.map((city) => (
                  <option key={city.cdCity} value={city.cdCity}>
                    {city.cityName}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Street Address
            </label>
            <input
              type="text"
              name="streetAddress"
              value={formData.streetAddress}
              onChange={handleChange}
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
              placeholder="Enter street address"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Postal Code
              </label>
              <input
                type="text"
                name="postalCode"
                value={formData.postalCode}
                onChange={handleChange}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
                placeholder="Enter postal code"
              />
            </div>
          </div>

          <button
            type="submit"
            disabled={isLoading}
            className="w-full bg-gradient-to-r from-blue-600 to-purple-600 text-white py-3 px-4 rounded-lg font-medium hover:from-blue-700 hover:to-purple-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isLoading ? (
              <div className="flex items-center justify-center">
                <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white mr-2"></div>
                Completing Profile...
              </div>
            ) : (
              'Complete Profile'
            )}
          </button>
        </form>
      </div>
    </div>
  );
};

export default CompleteProfile;
