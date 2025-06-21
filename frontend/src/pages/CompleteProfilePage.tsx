import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { User, Phone, Calendar, Users, MapPin } from 'lucide-react';
import Layout from '../components/layout/Layout';
import { useAuthStore } from '../store/authStore';

interface ProfileFormData {
  phone_number: string;
  date_of_birth: string;
  gender: string;
  occupation: string;
  emergency_contact_name: string;
  emergency_contact_phone: string;
  address: string;
}

const CompleteProfilePage: React.FC = () => {
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<ProfileFormData>();
  const { user, setUser } = useAuthStore();
  const navigate = useNavigate();

  const onSubmit = async (data: ProfileFormData) => {
    try {
      // Mock profile completion
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      if (user) {
        const updatedUser = { ...user, ...data, profile_complete: true };
        setUser(updatedUser);
      }
      
      navigate('/dashboard');
    } catch (error) {
      console.error('Profile completion failed:', error);
    }
  };

  return (
    <Layout showFooter={false}>
      <div className="min-h-screen bg-gradient-to-br from-primary-50 to-medical-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl w-full">
          <div className="bg-white rounded-2xl shadow-xl p-8">
            <div className="text-center mb-8">
              <div className="w-16 h-16 bg-gradient-to-br from-primary-600 to-primary-700 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg">
                <User className="w-8 h-8 text-white" />
              </div>
              <h2 className="text-3xl font-bold text-gray-900 mb-2">Complete Your Profile</h2>
              <p className="text-gray-600">
                Please provide additional information to complete your medical profile
              </p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Phone Number
                  </label>
                  <div className="relative">
                    <Phone className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <input
                      {...register('phone_number', {
                        required: 'Phone number is required',
                        pattern: {
                          value: /^[\+]?[1-9][\d]{0,15}$/,
                          message: 'Invalid phone number'
                        }
                      })}
                      type="tel"
                      className="input-field pl-10"
                      placeholder="+1 (555) 123-4567"
                    />
                  </div>
                  {errors.phone_number && (
                    <p className="mt-1 text-sm text-red-600">{errors.phone_number.message}</p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Date of Birth
                  </label>
                  <div className="relative">
                    <Calendar className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <input
                      {...register('date_of_birth', {
                        required: 'Date of birth is required'
                      })}
                      type="date"
                      className="input-field pl-10"
                    />
                  </div>
                  {errors.date_of_birth && (
                    <p className="mt-1 text-sm text-red-600">{errors.date_of_birth.message}</p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Gender
                  </label>
                  <div className="relative">
                    <Users className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <select
                      {...register('gender', {
                        required: 'Please select your gender'
                      })}
                      className="input-field pl-10"
                    >
                      <option value="">Select gender</option>
                      <option value="Male">Male</option>
                      <option value="Female">Female</option>
                      <option value="Other">Other</option>
                      <option value="Prefer not to say">Prefer not to say</option>
                    </select>
                  </div>
                  {errors.gender && (
                    <p className="mt-1 text-sm text-red-600">{errors.gender.message}</p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Occupation
                  </label>
                  <div className="relative">
                    <User className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <input
                      {...register('occupation')}
                      type="text"
                      className="input-field pl-10"
                      placeholder="Your occupation"
                    />
                  </div>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Emergency Contact Name
                  </label>
                  <div className="relative">
                    <User className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <input
                      {...register('emergency_contact_name', {
                        required: 'Emergency contact name is required'
                      })}
                      type="text"
                      className="input-field pl-10"
                      placeholder="Emergency contact name"
                    />
                  </div>
                  {errors.emergency_contact_name && (
                    <p className="mt-1 text-sm text-red-600">{errors.emergency_contact_name.message}</p>
                  )}
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Emergency Contact Phone
                  </label>
                  <div className="relative">
                    <Phone className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                    <input
                      {...register('emergency_contact_phone', {
                        required: 'Emergency contact phone is required'
                      })}
                      type="tel"
                      className="input-field pl-10"
                      placeholder="+1 (555) 123-4567"
                    />
                  </div>
                  {errors.emergency_contact_phone && (
                    <p className="mt-1 text-sm text-red-600">{errors.emergency_contact_phone.message}</p>
                  )}
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Address
                </label>
                <div className="relative">
                  <MapPin className="absolute left-3 top-3 w-5 h-5 text-gray-400" />
                  <textarea
                    {...register('address')}
                    rows={3}
                    className="input-field pl-10"
                    placeholder="Your full address"
                  />
                </div>
              </div>

              <button
                type="submit"
                disabled={isSubmitting}
                className="w-full btn-primary disabled:opacity-50"
              >
                {isSubmitting ? 'Completing Profile...' : 'Complete Profile'}
              </button>
            </form>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default CompleteProfilePage;
