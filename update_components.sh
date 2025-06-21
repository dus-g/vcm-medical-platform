#!/bin/bash

echo "🔧 Fixing TypeScript errors and updating components..."

# Fix Home.tsx button component
cat > frontend/src/pages/Home.tsx << 'HOMEFILE'
import React from 'react';
import { Link } from 'react-router-dom';
import { 
  CheckCircle, 
  Users, 
  FileText, 
  Stethoscope, 
  ArrowRight,
  Zap,
  Shield,
  Globe,
  Award,
  Target,
  Activity,
  Heart,
  Brain,
  Eye,
  Microscope,
  ShoppingCart
} from 'lucide-react';

// Simple Button component with proper typing
const Button: React.FC<{
  children: React.ReactNode;
  size?: 'sm' | 'md' | 'lg';
  variant?: 'primary' | 'secondary';
  className?: string;
  onClick?: () => void;
}> = ({ children, size = 'md', variant = 'primary', className = '', onClick }) => {
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded-lg transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2';
  
  const sizeClasses = {
    sm: 'px-3 py-2 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg'
  };
  
  const variantClasses = {
    primary: 'bg-blue-600 hover:bg-blue-700 text-white focus:ring-blue-500',
    secondary: 'bg-gray-200 hover:bg-gray-300 text-gray-900 focus:ring-gray-500'
  };
  
  return (
    <button 
      className={`${baseClasses} ${sizeClasses[size]} ${variantClasses[variant]} ${className}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

// Simple Card component
const Card: React.FC<{
  children: React.ReactNode;
  className?: string;
}> = ({ children, className = '' }) => {
  return (
    <div className={`bg-white rounded-lg shadow-md border border-gray-200 p-6 ${className}`}>
      {children}
    </div>
  );
};

export const Home: React.FC = () => {
  const companyHighlights = [
    {
      icon: <Award className="w-6 h-6 text-blue-600" />,
      title: "95% Treatment Efficacy",
      description: "Breakthrough results in melanoma and cancer treatment"
    },
    {
      icon: <Microscope className="w-6 h-6 text-blue-600" />,
      title: "World's First",
      description: "Clinical trials for antibiotic-resistant infections"
    },
    {
      icon: <Globe className="w-6 h-6 text-blue-600" />,
      title: "24/7 Platform Access",
      description: "Global operations with Shanghai headquarters"
    },
    {
      icon: <Zap className="w-6 h-6 text-blue-600" />,
      title: "Advanced Life-Cell Therapies",
      description: "Cutting-edge medical and naturopathic treatments"
    }
  ];

  const processSteps = [
    {
      title: 'Register & Consultation',
      description: 'Sign up and schedule your initial medical consultation with our specialized doctors',
      icon: <Users className="w-12 h-12 text-blue-600" />,
      step: 1,
      details: 'Create account, complete profile, book appointment'
    },
    {
      title: 'Medical Assessment',
      description: 'Complete comprehensive 50+ field assessment forms for your specific condition',
      icon: <FileText className="w-12 h-12 text-blue-600" />,
      step: 2,
      details: 'Psoriasis or Eye Disease assessment with medical history'
    },
    {
      title: 'Treatment Protocol',
      description: 'Receive personalized treatment plan with specific product recommendations',
      icon: <Target className="w-12 h-12 text-blue-600" />,
      step: 3,
      details: 'Custom protocol with duration and schedule'
    },
    {
      title: 'Purchase Products',
      description: 'Buy the exact products our medical team prescribes for your treatment',
      icon: <ShoppingCart className="w-12 h-12 text-blue-600" />,
      step: 4,
      details: 'Secure payment, fast shipping, product tracking'
    },
    {
      title: 'Ongoing Monitoring',
      description: 'Upload before/during/after photos and track your progress with our doctors',
      icon: <Activity className="w-12 h-12 text-blue-600" />,
      step: 5,
      details: 'Photo tracking, progress reports, doctor consultations'
    }
  ];

  const treatmentCategories = [
    {
      name: 'Cancers',
      icon: <Shield className="w-8 h-8 text-red-500" />,
      conditions: [
        'B-cell leukemia, lymphoma, and myeloma (CAR-T)',
        'Stomach, ovarian, gastrointestinal, bladder cancer',
        'Colorectal, pancreas, prostate cancer',
        'Melanoma therapy'
      ]
    },
    {
      name: 'Autoimmune Disorders',
      icon: <Heart className="w-8 h-8 text-blue-500" />,
      conditions: [
        'Rheumatoid arthritis',
        'Psoriasis Vulgaris',
        'Lupus',
        'Hashimoto disease',
        'Type 2 diabetes'
      ]
    },
    {
      name: 'Eye Diseases',
      icon: <Eye className="w-8 h-8 text-cyan-500" />,
      conditions: [
        'Macular degeneration',
        'Glaucoma',
        'Blindness from glaucoma',
        'Lazy eye (Amblyopia)'
      ]
    },
    {
      name: 'Gastrointestinal & Respiratory',
      icon: <Stethoscope className="w-8 h-8 text-green-500" />,
      conditions: [
        'Gastritis & Stomach ulcers',
        'Helicobacter Pylori',
        'Tuberculosis',
        'Pneumonia & Bronchitis'
      ]
    }
  ];

  return (
    <div className="min-h-screen bg-white">
      {/* Hero Section */}
      <section className="pt-28 pb-16 px-6 bg-gradient-to-br from-blue-50 via-blue-50 to-cyan-50">
        <div className="max-w-7xl mx-auto text-center">
          <div className="inline-flex items-center px-4 py-2 bg-gradient-to-r from-blue-100 to-blue-100 border border-blue-200 rounded-full text-blue-700 text-sm font-medium mb-6">
            <div className="w-2 h-2 bg-blue-500 rounded-full mr-2 animate-pulse"></div>
            VAMOS BIOTECH - Bio-Pharmaceutical Innovation
          </div>

          <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
            Advanced Medical
            <span className="block text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-blue-600">
              Treatment Platform
            </span>
          </h1>

          <p className="text-xl text-gray-600 mb-8 max-w-4xl mx-auto">
            Breakthrough life-cell based therapies for cancer, viral infections, autoimmune disorders, 
            and antibiotic-resistant bacterial infections with proven 95% efficacy rates.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
            <Link to="/register">
              <Button size="lg" className="min-w-[200px] bg-gradient-to-r from-blue-600 to-blue-600 hover:from-blue-700 hover:to-blue-700 shadow-lg">
                Start Treatment
              </Button>
            </Link>
            <Link to="/therapies">
              <Button variant="secondary" size="lg" className="min-w-[200px] border-blue-300 text-blue-700 hover:bg-blue-50">
                Explore Therapies
              </Button>
            </Link>
          </div>

          {/* Company Highlights */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 max-w-6xl mx-auto">
            {companyHighlights.map((highlight, index) => (
              <div key={index} className="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                <div className="flex items-center justify-center w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto">
                  {highlight.icon}
                </div>
                <h3 className="text-lg font-bold text-gray-900 mb-2">{highlight.title}</h3>
                <p className="text-sm text-gray-600">{highlight.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Treatment Process */}
      <section className="py-16 px-6 bg-white">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Your Treatment Journey
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Simple steps to access our breakthrough medical treatments
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
            {processSteps.map((step, index) => (
              <div key={index} className="relative">
                <div className="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 h-full text-center">
                  <div className="mb-4">{step.icon}</div>
                  <div className="inline-flex items-center justify-center w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3">
                    {step.step}
                  </div>
                  <h3 className="text-lg font-bold text-gray-900 mb-2">{step.title}</h3>
                  <p className="text-gray-600 text-sm mb-3">{step.description}</p>
                  <div className="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                    {step.details}
                  </div>
                </div>
                {index < processSteps.length - 1 && (
                  <div className="hidden lg:block absolute top-1/2 -right-3 transform -translate-y-1/2">
                    <ArrowRight className="w-5 h-5 text-blue-300" />
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Treatment Categories */}
      <section className="py-16 px-6 bg-gray-50">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Medical Conditions We Treat
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Comprehensive treatment protocols across major medical specialties
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {treatmentCategories.map((category, index) => (
              <Card key={index} className="hover:shadow-xl transition-all duration-300 border-l-4 border-l-blue-500">
                <div className="flex items-center mb-4">
                  {category.icon}
                  <h3 className="text-xl font-bold text-gray-900 ml-3">{category.name}</h3>
                </div>
                <ul className="space-y-2">
                  {category.conditions.map((condition, conditionIndex) => (
                    <li key={conditionIndex} className="flex items-start">
                      <CheckCircle className="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" />
                      <span className="text-gray-700 text-sm">{condition}</span>
                    </li>
                  ))}
                </ul>
                <div className="mt-4">
                  <Button variant="secondary" size="sm" className="w-full">
                    Learn More
                  </Button>
                </div>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Call to Action */}
      <section className="py-16 px-6 bg-gradient-to-r from-blue-600 to-blue-700">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="text-3xl md:text-4xl font-bold text-white mb-6">
            Ready to Begin Your Treatment?
          </h2>
          <p className="text-xl text-blue-100 mb-8">
            Join thousands of patients who have experienced breakthrough results with our advanced therapies.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link to="/register">
              <Button size="lg" className="bg-white text-blue-600 hover:bg-gray-100 min-w-[200px] shadow-lg">
                Start Your Journey
              </Button>
            </Link>
            <Link to="/contact">
              <Button variant="secondary" size="lg" className="border-white text-white hover:bg-white hover:text-blue-600 min-w-[200px]">
                Contact Our Team
              </Button>
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Home;
HOMEFILE

echo "✅ Updated Home.tsx"
