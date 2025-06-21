import React, { useState, useEffect } from 'react';
import api from '../utils/api';

interface Country {
  cd_country: number;
  country_name: string;
  country_abbr: string;
}

const Test: React.FC = () => {
  const [countries, setCountries] = useState<Country[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCountries = async () => {
      try {
        const countriesData = await api.getCountries();
        setCountries(countriesData);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch countries');
      } finally {
        setLoading(false);
      }
    };

    fetchCountries();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading countries...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="text-red-600 text-xl">❌ Error</div>
          <p className="mt-2 text-gray-600">{error}</p>
          <p className="mt-4 text-sm text-gray-500">
            Make sure your backend is running on http://localhost:8080
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-4xl mx-auto px-4">
        <div className="bg-white rounded-lg shadow-md p-6">
          <h1 className="text-2xl font-bold text-gray-900 mb-6">
            🧪 Backend Integration Test
          </h1>
          
          <div className="mb-6">
            <div className="bg-green-50 border border-green-200 rounded-lg p-4">
              <div className="flex items-center">
                <div className="text-green-600 text-xl mr-2">✅</div>
                <div>
                  <h3 className="text-green-800 font-medium">Backend Connected!</h3>
                  <p className="text-green-700 text-sm">Successfully fetched {countries.length} countries</p>
                </div>
              </div>
            </div>
          </div>

          <div>
            <h2 className="text-lg font-semibold text-gray-900 mb-4">Countries from Backend:</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {countries.map((country) => (
                <div
                  key={country.cd_country}
                  className="bg-gray-50 rounded-lg p-4 border border-gray-200"
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <h3 className="font-medium text-gray-900">{country.country_name}</h3>
                      <p className="text-sm text-gray-600">Code: {country.country_abbr}</p>
                    </div>
                    <div className="text-2xl font-bold text-blue-600">
                      {country.cd_country}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          <div className="mt-8 text-center">
            <p className="text-gray-600">
              🎉 Your VCM Medical Platform backend is working perfectly!
            </p>
            <p className="text-sm text-gray-500 mt-2">
              Frontend ↔️ Backend integration successful
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Test;
