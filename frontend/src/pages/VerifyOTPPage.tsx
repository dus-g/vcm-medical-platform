import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { Shield } from 'lucide-react';
import Layout from '../components/layout/Layout';
import { useAuthStore } from '../store/authStore';

interface OTPFormData {
  otp: string;
}

const VerifyOTPPage: React.FC = () => {
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<OTPFormData>();
  const { verifyOTP } = useAuthStore();
  const navigate = useNavigate();
  const location = useLocation();
  const email = (location.state as any)?.email || '';

  const onSubmit = async (data: OTPFormData) => {
    try {
      await verifyOTP(email, data.otp);
      navigate('/complete-profile');
    } catch (error) {
      console.error('OTP verification failed:', error);
    }
  };

  return (
    <Layout showFooter={false}>
      <div className="min-h-screen bg-gradient-to-br from-primary-50 to-medical-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full">
          <div className="bg-white rounded-2xl shadow-xl p-8">
            <div className="text-center mb-8">
              <div className="w-16 h-16 bg-gradient-to-br from-primary-600 to-primary-700 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg">
                <Shield className="w-8 h-8 text-white" />
              </div>
              <h2 className="text-3xl font-bold text-gray-900 mb-2">Verify your email</h2>
              <p className="text-gray-600">
                We have sent a verification code to<br />
                <span className="font-medium text-gray-900">{email}</span>
              </p>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Verification Code
                </label>
                <input
                  {...register('otp', {
                    required: 'OTP is required',
                    pattern: {
                      value: /^\d{6}$/,
                      message: 'OTP must be 6 digits'
                    }
                  })}
                  type="text"
                  maxLength={6}
                  className="input-field text-center text-2xl tracking-widest font-mono"
                  placeholder="000000"
                />
                {errors.otp && (
                  <p className="mt-1 text-sm text-red-600">{errors.otp.message}</p>
                )}
              </div>

              <button
                type="submit"
                disabled={isSubmitting}
                className="w-full btn-primary disabled:opacity-50"
              >
                {isSubmitting ? 'Verifying...' : 'Verify Email'}
              </button>
            </form>

            <div className="mt-6 text-center">
              <span className="text-gray-600">Did not receive the code? </span>
              <button className="text-primary-600 hover:text-primary-500 font-medium">
                Resend
              </button>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default VerifyOTPPage;
