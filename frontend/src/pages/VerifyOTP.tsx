import React, { useState } from 'react'
import { Link, useNavigate, useLocation } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { toast } from 'react-hot-toast'
import { Button } from '../components/ui/Button'
import { Card } from '../components/ui/Card'
import { useAuthStore } from '../store/authStore'
import api from '../utils/api'
import { Shield, ArrowLeft } from 'lucide-react'

interface VerifyOTPForm {
  otp: string
}

const VerifyOTP: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()
  const location = useLocation()
  const login = useAuthStore((state) => state.login)
  
  const email = location.state?.email || ''
  const { register, handleSubmit, formState: { errors } } = useForm<VerifyOTPForm>()

  const onSubmit = async (data: VerifyOTPForm) => {
    setIsLoading(true)
    try {
      const response = await api.post('/auth/verify-otp', {
        email: email,
        otp: data.otp
      })
      
      const { token, user } = response.data
      login(user)
      toast.success('Account verified successfully!')
      navigate('/dashboard')
    } catch (error: any) {
      toast.error(error.response?.data?.error || 'OTP verification failed')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <Link to="/register" className="inline-flex items-center text-primary-600 hover:text-primary-700 mb-6">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Register
          </Link>
          <div className="text-center">
            <Shield className="mx-auto h-12 w-12 text-primary-600" />
            <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
              Verify your account
            </h2>
            <p className="mt-2 text-sm text-gray-600">
              Enter the 6-digit code sent to<br />
              <strong>{email}</strong>
            </p>
          </div>
        </div>
        
        <Card>
          <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
            <div>
              <label className="block text-sm font-medium text-gray-700">Verification Code</label>
              <input
                type="text"
                maxLength={6}
                {...register('otp', { 
                  required: 'OTP is required',
                  pattern: {
                    value: /^[0-9]{6}$/,
                    message: 'OTP must be 6 digits'
                  }
                })}
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm text-center text-lg tracking-widest focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                placeholder="123456"
              />
              {errors.otp && (
                <p className="mt-1 text-sm text-red-600">{errors.otp.message}</p>
              )}
            </div>

            <Button
              type="submit"
              size="lg"
              className="w-full"
              disabled={isLoading}
            >
              {isLoading ? 'Verifying...' : 'Verify Account'}
            </Button>
          </form>
          
          <div className="mt-6 p-4 bg-green-50 rounded-lg">
            <p className="text-sm text-green-800">
              <strong>Demo OTP:</strong> 123456<br />
              <em>In production, this would be sent via email/SMS</em>
            </p>
          </div>
          
          <div className="mt-6 text-center">
            <button className="text-sm text-primary-600 hover:text-primary-500">
              Didn't receive the code? Resend
            </button>
          </div>
        </Card>
      </div>
    </div>
  )
}

export default VerifyOTP
