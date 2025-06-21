import React from 'react'
import { Link } from 'react-router-dom'
import { Button } from '../components/ui/Button'
import { Card } from '../components/ui/Card'
import { 
  Users, 
  FileText, 
  Calendar, 
  Zap,
  Globe,
  Award,
  Microscope,
} from 'lucide-react'

const Home: React.FC = () => {
  const companyHighlights = [
    {
      icon: <Award className="w-6 h-6 text-primary-600" />,
      title: "95% Treatment Efficacy",
      description: "Breakthrough results in melanoma and cancer treatment"
    },
    {
      icon: <Microscope className="w-6 h-6 text-primary-600" />,
      title: "World's First",
      description: "Clinical trials for antibiotic-resistant infections"
    },
    {
      icon: <Globe className="w-6 h-6 text-primary-600" />,
      title: "24/7 Platform Access",
      description: "Global operations with Shanghai headquarters"
    },
    {
      icon: <Zap className="w-6 h-6 text-primary-600" />,
      title: "Advanced Life-Cell Therapies",
      description: "Cutting-edge medical and naturopathic treatments"
    }
  ]

  return (
    <div className="min-h-screen bg-white">
      {/* Navigation */}
      <nav className="bg-white shadow-sm border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-bold text-primary-600">VCM Medical Platform</h1>
            </div>
            <div className="flex items-center space-x-4">
              <Link to="/login" className="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium">
                Sign In
              </Link>
              <Link to="/register">
                <Button size="sm">Get Started</Button>
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="pt-16 pb-20 px-6 bg-gradient-to-br from-primary-50 via-blue-50 to-cyan-50">
        <div className="max-w-7xl mx-auto text-center">
          <div className="inline-flex items-center px-4 py-2 bg-gradient-to-r from-primary-100 to-blue-100 border border-primary-200 rounded-full text-primary-700 text-sm font-medium mb-6">
            <div className="w-2 h-2 bg-primary-500 rounded-full mr-2 animate-pulse"></div>
            VAMOS BIOTECH - Bio-Pharmaceutical Innovation
          </div>

          <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
            Advanced Medical
            <span className="block text-transparent bg-clip-text bg-gradient-to-r from-primary-600 to-blue-600">
              Treatment Platform
            </span>
          </h1>

          <p className="text-xl text-gray-600 mb-8 max-w-4xl mx-auto">
            Breakthrough life-cell based therapies for cancer, viral infections, autoimmune disorders, 
            and antibiotic-resistant bacterial infections with proven 95% efficacy rates.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
            <Link to="/register">
              <Button size="lg" className="min-w-[200px] bg-gradient-to-r from-primary-600 to-blue-600 hover:from-primary-700 hover:to-blue-700 shadow-lg">
                Start Treatment
              </Button>
            </Link>
            <Link to="/login">
              <Button variant="secondary" size="lg" className="min-w-[200px] border-primary-300 text-primary-700 hover:bg-primary-50">
                Access Dashboard
              </Button>
            </Link>
          </div>

          {/* Company Highlights */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 max-w-6xl mx-auto">
            {companyHighlights.map((highlight, index) => (
              <div key={index} className="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                <div className="flex items-center justify-center w-12 h-12 bg-primary-100 rounded-lg mb-4 mx-auto">
                  {highlight.icon}
                </div>
                <h3 className="text-lg font-bold text-gray-900 mb-2">{highlight.title}</h3>
                <p className="text-sm text-gray-600">{highlight.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Platform Success */}
      <section className="py-16 px-6 bg-white">
        <div className="max-w-4xl mx-auto text-center">
          <div className="bg-gradient-to-r from-green-50 to-blue-50 rounded-2xl p-8 border border-green-200">
            <div className="text-6xl mb-6">🎉</div>
            <h2 className="text-3xl font-bold text-gray-900 mb-4">
              Platform Successfully Deployed!
            </h2>
            <p className="text-lg text-gray-600 mb-6">
              Your VCM Medical Platform is now fully operational with all systems ready for global medical innovation.
            </p>
            
            <div className="grid md:grid-cols-2 gap-6 mb-8">
              <div className="bg-white rounded-lg p-6 shadow-sm">
                <h3 className="text-lg font-semibold mb-3 text-blue-600">✅ Core Features Ready</h3>
                <ul className="text-sm text-gray-700 space-y-1">
                  <li>• Multi-user Authentication (8 User Types)</li>
                  <li>• Medical Assessment Forms</li>
                  <li>• Appointment Booking System</li>
                  <li>• Real-time Chat Support</li>
                  <li>• Order Management</li>
                  <li>• Progress Tracking</li>
                </ul>
              </div>
              
              <div className="bg-white rounded-lg p-6 shadow-sm">
                <h3 className="text-lg font-semibold mb-3 text-green-600">🚀 Technology Stack</h3>
                <ul className="text-sm text-gray-700 space-y-1">
                  <li>• React + TypeScript Frontend</li>
                  <li>• Go + Fiber Backend</li>
                  <li>• PostgreSQL Database</li>
                  <li>• JWT Authentication</li>
                  <li>• Railway Cloud Hosting</li>
                  <li>• Responsive Design</li>
                </ul>
              </div>
            </div>
            
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link to="/register">
                <Button size="lg" className="min-w-[200px]">
                  Create Account
                </Button>
              </Link>
              <Link to="/login">
                <Button variant="outline" size="lg" className="min-w-[200px]">
                  Sign In
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* Features Grid */}
      <section className="py-16 px-6 bg-gray-50">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">
              Complete Medical Platform
            </h2>
            <p className="text-xl text-gray-600">
              Everything you need for advanced medical treatment management
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card className="text-center hover:shadow-lg transition-shadow">
              <FileText className="w-12 h-12 text-primary-600 mx-auto mb-4" />
              <h3 className="text-lg font-semibold mb-2">Assessment Forms</h3>
              <p className="text-gray-600">Comprehensive 50+ field medical assessments for psoriasis and eye diseases</p>
            </Card>

            <Card className="text-center hover:shadow-lg transition-shadow">
              <Calendar className="w-12 h-12 text-primary-600 mx-auto mb-4" />
              <h3 className="text-lg font-semibold mb-2">Appointment System</h3>
              <p className="text-gray-600">Schedule consultations with specialized doctors and track your appointments</p>
            </Card>

            <Card className="text-center hover:shadow-lg transition-shadow">
              <Users className="w-12 h-12 text-primary-600 mx-auto mb-4" />
              <h3 className="text-lg font-semibold mb-2">Multi-User Support</h3>
              <p className="text-gray-600">Patients, doctors, agents, distributors, and administrators - all in one platform</p>
            </Card>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-white py-12">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div>
              <h3 className="text-lg font-semibold mb-4">VCM Medical Platform</h3>
              <p className="text-gray-300 text-sm">
                Advanced medical treatment platform with breakthrough life-cell based therapies.
              </p>
            </div>
            
            <div>
              <h4 className="text-lg font-semibold mb-4">Quick Links</h4>
              <ul className="space-y-2 text-sm">
                <li><Link to="/register" className="text-gray-300 hover:text-white">Register</Link></li>
                <li><Link to="/login" className="text-gray-300 hover:text-white">Login</Link></li>
                <li><Link to="/dashboard" className="text-gray-300 hover:text-white">Dashboard</Link></li>
              </ul>
            </div>
            
            <div>
              <h4 className="text-lg font-semibold mb-4">Contact</h4>
              <p className="text-gray-300 text-sm">
                VAMOS BIOTECH (Shanghai) Co., Ltd.<br />
                Building #5, Lin Gang Fengxian Industrial Park<br />
                Shanghai 201413, P.R. China
              </p>
            </div>
          </div>
          
          <div className="border-t border-gray-800 mt-8 pt-8 text-center">
            <p className="text-gray-400 text-sm">
              © 2024 VCM Medical Platform. All rights reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default Home
