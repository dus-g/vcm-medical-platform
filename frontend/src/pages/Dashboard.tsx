import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuthStore } from '../store/authStore'
import { Button } from '../components/ui/Button'
import { Card } from '../components/ui/Card'
import api from '../utils/api'
import { 
  LogOut, 
  User, 
  Activity, 
  Calendar, 
  ShoppingCart, 
  FileText,
  MessageSquare,
  TrendingUp,
  CheckCircle,
  Settings,
  Bell
} from 'lucide-react'

const Dashboard: React.FC = () => {
  const navigate = useNavigate()
  const { user, logout } = useAuthStore()
  const [dashboardData, setDashboardData] = useState<any>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    fetchDashboard()
  }, [])

  const fetchDashboard = async () => {
    try {
      const response = await api.get('/dashboard')
      setDashboardData(response.data)
    } catch (error) {
      console.error('Failed to fetch dashboard data')
    } finally {
      setIsLoading(false)
    }
  }

  const handleLogout = () => {
    logout()
    navigate('/')
  }

  const getUserTypeLabel = (type: number) => {
    const types: { [key: number]: string } = {
      0: 'Patient',
      1: 'Agent', 
      2: 'Sales Channel',
      3: 'Influencer',
      4: 'Distributor',
      5: 'Doctor',
      10: 'Operator',
      11: 'Admin'
    }
    return types[type] || 'User'
  }

  const quickActions = [
    {
      name: 'Complete Assessment',
      description: 'Fill out medical assessment forms',
      icon: FileText,
      color: 'bg-blue-500',
      action: () => console.log('Navigate to assessments')
    },
    {
      name: 'Book Appointment',
      description: 'Schedule with a doctor',
      icon: Calendar,
      color: 'bg-green-500',
      action: () => console.log('Navigate to appointments')
    },
    {
      name: 'Order Products',
      description: 'Purchase treatment products',
      icon: ShoppingCart,
      color: 'bg-purple-500',
      action: () => console.log('Navigate to orders')
    },
    {
      name: 'Chat Support',
      description: 'Get help from our team',
      icon: MessageSquare,
      color: 'bg-orange-500',
      action: () => console.log('Navigate to chat')
    }
  ]

  const stats = [
    { 
      name: 'Active Assessments', 
      value: dashboardData?.stats?.assessments || '0', 
      icon: Activity,
      color: 'text-blue-600'
    },
    { 
      name: 'Upcoming Appointments', 
      value: dashboardData?.stats?.appointments || '0', 
      icon: Calendar,
      color: 'text-green-600'
    },
    { 
      name: 'Orders Placed', 
      value: dashboardData?.stats?.orders || '0', 
      icon: ShoppingCart,
      color: 'text-purple-600'
    },
    { 
      name: 'Platform Score', 
      value: '95%', 
      icon: TrendingUp,
      color: 'text-orange-600'
    }
  ]

  const recentActivity = [
    { action: 'Account created successfully', date: 'Just now', status: 'completed' },
    { action: 'Profile setup completed', date: '2 minutes ago', status: 'completed' },
    { action: 'Dashboard accessed', date: '5 minutes ago', status: 'completed' },
    { action: 'Medical assessment pending', date: 'Waiting', status: 'pending' }
  ]

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading dashboard...</p>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Navigation */}
      <nav className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-semibold text-gray-900">VCM Medical Dashboard</h1>
            </div>
            <div className="flex items-center space-x-4">
              <button className="p-2 text-gray-400 hover:text-gray-500">
                <Bell className="h-6 w-6" />
              </button>
              <button className="p-2 text-gray-400 hover:text-gray-500">
                <Settings className="h-6 w-6" />
              </button>
              <Button
                onClick={handleLogout}
                variant="outline"
                size="sm"
                className="text-red-600 border-red-300 hover:bg-red-50"
              >
                <LogOut className="h-4 w-4 mr-2" />
                Logout
              </Button>
            </div>
          </div>
        </div>
      </nav>

      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          {/* Welcome Section */}
          <div className="bg-gradient-to-r from-primary-600 to-blue-600 rounded-lg p-6 text-white mb-6">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <div className="w-16 h-16 bg-white/20 rounded-full flex items-center justify-center">
                  <User className="h-8 w-8 text-white" />
                </div>
              </div>
              <div className="ml-4">
                <h1 className="text-2xl font-bold">
                  Welcome back, {user?.name.split(" ")[0]} {user?.name.split(" ")[1] || ""}!
                </h1>
                <p className="text-primary-100 mt-1">
                  {getUserTypeLabel(user?.userType || 0)} • {user?.email}
                </p>
                <p className="text-primary-200 text-sm mt-1">
                  Continue your treatment journey with our advanced medical protocols
                </p>
              </div>
            </div>
          </div>

          {/* Stats Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
            {stats.map((stat) => {
              const Icon = stat.icon
              return (
                <Card key={stat.name}>
                  <div className="flex items-center">
                    <div className="flex-shrink-0">
                      <Icon className={`h-8 w-8 ${stat.color}`} />
                    </div>
                    <div className="ml-4">
                      <p className="text-sm font-medium text-gray-600">{stat.name}</p>
                      <p className="text-2xl font-bold text-gray-900">{stat.value}</p>
                    </div>
                  </div>
                </Card>
              )
            })}
          </div>

          {/* Quick Actions */}
          <div className="mb-6">
            <h2 className="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
              {quickActions.map((action) => {
                const Icon = action.icon
                return (
                  <Card key={action.name} className="hover:shadow-md transition-shadow cursor-pointer" >
                    <div className="text-center">
                      <div className={`inline-flex items-center justify-center w-12 h-12 rounded-lg ${action.color} text-white mb-4`}>
                        <Icon className="h-6 w-6" />
                      </div>
                      <h3 className="text-lg font-medium text-gray-900 mb-2">
                        {action.name}
                      </h3>
                      <p className="text-sm text-gray-600">
                        {action.description}
                      </p>
                    </div>
                  </Card>
                )
              })}
            </div>
          </div>

          {/* Content Grid */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {/* Recent Activity */}
            <Card>
              <h3 className="text-lg font-medium text-gray-900 mb-4">Recent Activity</h3>
              <div className="space-y-4">
                {recentActivity.map((activity, index) => (
                  <div key={index} className="flex items-center space-x-3">
                    <div className={`flex-shrink-0 w-2 h-2 rounded-full ${
                      activity.status === 'completed' ? 'bg-green-500' : 'bg-yellow-500'
                    }`} />
                    <div className="flex-1 min-w-0">
                      <p className="text-sm text-gray-900">{activity.action}</p>
                      <p className="text-xs text-gray-500">{activity.date}</p>
                    </div>
                    {activity.status === 'completed' && (
                      <CheckCircle className="h-4 w-4 text-green-500" />
                    )}
                  </div>
                ))}
              </div>
            </Card>

            {/* Platform Status */}
            <Card>
              <h3 className="text-lg font-medium text-gray-900 mb-4">Platform Status</h3>
              <div className="space-y-4">
                <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                  <div className="flex items-center">
                    <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                    <span className="text-green-800 font-medium">All Systems Operational</span>
                  </div>
                  <p className="text-green-700 text-sm mt-1">
                    VCM Medical Platform is running smoothly
                  </p>
                </div>
                
                <div>
                  <div className="flex justify-between text-sm text-gray-600 mb-1">
                    <span>Account Setup</span>
                    <span>100%</span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div className="bg-green-600 h-2 rounded-full" style={{ width: '100%' }} />
                  </div>
                </div>
                
                <div>
                  <div className="flex justify-between text-sm text-gray-600 mb-1">
                    <span>Profile Completion</span>
                    <span>85%</span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div className="bg-blue-600 h-2 rounded-full" style={{ width: '85%' }} />
                  </div>
                </div>

                <div className="mt-4">
                  <Button size="sm" className="w-full">
                    Complete Profile Setup
                  </Button>
                </div>
              </div>
            </Card>
          </div>

          {/* Success Message */}
          <div className="mt-6 bg-blue-50 border border-blue-200 rounded-lg p-6">
            <h3 className="text-lg font-semibold text-blue-900 mb-2">🎉 Welcome to VCM Medical Platform!</h3>
            <p className="text-blue-800 mb-4">
              Your account has been successfully created and verified. You now have access to our comprehensive 
              medical treatment platform with advanced life-cell based therapies.
            </p>
            <div className="grid md:grid-cols-2 gap-4 text-sm">
              <div>
                <h4 className="font-medium text-blue-900 mb-2">✅ Available Features:</h4>
                <ul className="space-y-1 text-blue-700">
                  <li>• Medical assessment forms</li>
                  <li>• Doctor appointment booking</li>
                  <li>• Real-time chat support</li>
                  <li>• Treatment progress tracking</li>
                </ul>
              </div>
              <div>
                <h4 className="font-medium text-blue-900 mb-2">🚀 Next Steps:</h4>
                <ul className="space-y-1 text-blue-700">
                  <li>• Complete your medical assessment</li>
                  <li>• Upload progress photos</li>
                  <li>• Schedule your first consultation</li>
                  <li>• Explore treatment options</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Dashboard
