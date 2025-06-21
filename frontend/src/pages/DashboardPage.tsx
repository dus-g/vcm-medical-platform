import React from 'react';
import { Link } from 'react-router-dom';
import Layout from '../components/layout/Layout';
import { useAuthStore } from '../store/authStore';
import { 
  FileText, 
  Calendar, 
  Activity, 
  MessageSquare, 
  Image, 
  Settings
} from 'lucide-react';

const DashboardPage: React.FC = () => {
  const { user } = useAuthStore();

  const quickActions = [
    {
      title: "Medical Assessments",
      description: "Complete your medical assessment forms",
      icon: <FileText className="w-6 h-6 text-primary-600" />,
      href: "/assessments",
    },
    {
      title: "Appointments",
      description: "Schedule and manage your appointments",
      icon: <Calendar className="w-6 h-6 text-blue-600" />,
      href: "/appointments",
    },
    {
      title: "Treatment Plans",
      description: "Track your treatment progress",
      icon: <Activity className="w-6 h-6 text-green-600" />,
      href: "/treatments",
    },
    {
      title: "Chat with Doctors",
      description: "Real-time communication with medical professionals",
      icon: <MessageSquare className="w-6 h-6 text-purple-600" />,
      href: "/chat",
    },
    {
      title: "Medical Images",
      description: "Upload and manage medical images",
      icon: <Image className="w-6 h-6 text-orange-600" />,
      href: "/medical-images",
    },
    {
      title: "Profile Settings",
      description: "Update your personal information",
      icon: <Settings className="w-6 h-6 text-gray-600" />,
      href: "/profile",
    }
  ];

  return (
    <Layout>
      <div className="min-h-screen bg-gray-50">
        <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
          <div className="px-4 py-6 sm:px-0">
            <div className="mb-8">
              <h1 className="text-3xl font-bold text-gray-900">
                Welcome back, {user?.first_name || 'User'}!
              </h1>
              <p className="mt-2 text-gray-600">
                Manage your medical journey and connect with healthcare professionals
              </p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
              {quickActions.map((action, index) => (
                <div key={index} className="card group hover:shadow-lg transition-all duration-200">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center mb-3">
                        {action.icon}
                      </div>
                      <h3 className="text-lg font-semibold text-gray-900 mb-2">
                        {action.title}
                      </h3>
                      <p className="text-gray-600 text-sm mb-4">
                        {action.description}
                      </p>
                      <Link 
                        to={action.href} 
                        className="text-primary-600 hover:text-primary-700 font-medium text-sm group-hover:underline"
                      >
                        Get Started →
                      </Link>
                    </div>
                  </div>
                </div>
              ))}
            </div>

            <div className="bg-gradient-to-r from-primary-600 to-primary-700 rounded-xl p-6 text-white">
              <div className="flex items-center justify-between">
                <div>
                  <h3 className="text-xl font-semibold mb-2">Ready to start your medical assessment?</h3>
                  <p className="text-primary-100">
                    Complete your comprehensive medical evaluation to get personalized treatment recommendations.
                  </p>
                </div>
                <Link 
                  to="/assessments" 
                  className="bg-white text-primary-600 px-6 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors whitespace-nowrap ml-6"
                >
                  Start Assessment
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default DashboardPage;
