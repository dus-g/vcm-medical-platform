import React from 'react';
import { Link } from 'react-router-dom';
import { 
  Mail, 
  Phone, 
  MapPin,
  Facebook,
  Twitter,
  Linkedin,
  Instagram,
  ArrowRight,
  Heart,
  Shield,
  Clock,
  Users
} from 'lucide-react';

export const PublicFooter: React.FC = () => {
  const currentYear = new Date().getFullYear();

  const quickLinks = [
    { name: 'Home', path: '/' },
    { name: 'About Us', path: '/about' },
    { name: 'Medical Therapies', path: '/therapies' },
    { name: 'Contact', path: '/contact' },
    { name: 'Login', path: '/login' },
    { name: 'Register', path: '/register' }
  ];

  const medicalSpecialties = [
    { name: 'Cancers (CAR-T, Melanoma)', path: '/therapies/cancers' },
    { name: 'Viral Infections (RNA/DNA)', path: '/therapies/infections' },
    { name: 'Autoimmune Disorders', path: '/therapies/autoimmune' },
    { name: 'Eye Diseases', path: '/therapies/eye-diseases' },
    { name: 'Neurological Disorders', path: '/therapies/neurological' },
    { name: 'Antibiotic-Resistant Infections', path: '/therapies/antibiotic-resistant' }
  ];

  const legalLinks = [
    { name: 'Privacy Policy', path: '/privacy' },
    { name: 'Terms of Service', path: '/terms' },
    { name: 'Medical Disclaimer', path: '/disclaimer' },
    { name: 'Cookie Policy', path: '/cookies' }
  ];

  const socialLinks = [
    { name: 'Facebook', icon: Facebook, url: '#' },
    { name: 'Twitter', icon: Twitter, url: '#' },
    { name: 'LinkedIn', icon: Linkedin, url: '#' },
    { name: 'Instagram', icon: Instagram, url: '#' }
  ];

  const features = [
    {
      icon: <Shield className="w-5 h-5" />,
      text: '95% Treatment Efficacy'
    },
    {
      icon: <Clock className="w-5 h-5" />,
      text: 'World\'s First Antibiotic Resistance Trials'
    },
    {
      icon: <Users className="w-5 h-5" />,
      text: 'Led by Prof. Sergey I. Chernysh'
    },
    {
      icon: <Heart className="w-5 h-5" />,
      text: 'Shanghai Global Operations'
    }
  ];

  return (
    <footer className="bg-gray-900 text-white">
      {/* Main Footer Content */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {/* Company Information */}
          <div className="lg:col-span-1">
            <div className="flex items-center space-x-2 mb-4">
              <div className="h-8 w-8 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-sm">VCM</span>
              </div>
              <span className="text-white font-bold text-lg">VCM Medical</span>
            </div>
            <p className="text-gray-300 mb-6 text-sm leading-relaxed">
              VAMOS BIOTECH (Shanghai) Co., Ltd. - Bio-pharmaceutical company specializing in 
              advanced life-cell based therapies for cancer, viral infections, and autoimmune disorders 
              with proven 95% treatment efficacy.
            </p>
            
            {/* Features */}
            <div className="space-y-3">
              {features.map((feature, index) => (
                <div key={index} className="flex items-center space-x-2 text-sm text-gray-300">
                  <div className="text-blue-400">
                    {feature.icon}
                  </div>
                  <span>{feature.text}</span>
                </div>
              ))}
            </div>
          </div>

          {/* Quick Links */}
          <div>
            <h3 className="text-lg font-semibold mb-6">Quick Links</h3>
            <ul className="space-y-3">
              {quickLinks.map((link) => (
                <li key={link.name}>
                  <Link
                    to={link.path}
                    className="text-gray-300 hover:text-blue-400 transition-colors text-sm flex items-center group"
                  >
                    <ArrowRight className="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" />
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Treatment Specialties */}
          <div>
            <h3 className="text-lg font-semibold mb-6">Treatment Specialties</h3>
            <ul className="space-y-3">
              {medicalSpecialties.map((specialty) => (
                <li key={specialty.name}>
                  <Link
                    to={specialty.path}
                    className="text-gray-300 hover:text-blue-400 transition-colors text-sm flex items-center group"
                  >
                    <ArrowRight className="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" />
                    {specialty.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Contact & Platform Info */}
          <div>
            <h3 className="text-lg font-semibold mb-6">Contact Us</h3>
            <div className="space-y-4">
              <div className="flex items-start space-x-3">
                <MapPin className="w-5 h-5 text-blue-400 mt-0.5 flex-shrink-0" />
                <div className="text-sm text-gray-300">
                  <p>Building #5, Lin Gang Fengxian Industrial Park</p>
                  <p>1800 Xin Yang Road, Feng Xian District</p>
                  <p>Shanghai 201413, P.R. China</p>
                </div>
              </div>
              
              <div className="flex items-center space-x-3">
                <Phone className="w-5 h-5 text-blue-400 flex-shrink-0" />
                <span className="text-sm text-gray-300">+86 (21) 1234-5678</span>
              </div>
              
              <div className="flex items-center space-x-3">
                <Mail className="w-5 h-5 text-blue-400 flex-shrink-0" />
                <span className="text-sm text-gray-300">info@vamosbiotech.com</span>
              </div>
            </div>

            {/* Company Info */}
            <div className="mt-6">
              <h4 className="text-sm font-semibold mb-3 text-gray-200">Company Details</h4>
              <div className="space-y-1 text-xs text-gray-400">
                <div>Registration: 91310000MAH3AQB3D</div>
                <div>Founded: 2014</div>
                <div>Startup with Global Operations</div>
                <div>Led by Prof. Sergey I. Chernysh</div>
              </div>
            </div>
          </div>
        </div>

        {/* Newsletter Signup */}
        <div className="mt-12 pt-8 border-t border-gray-800">
          <div className="max-w-md mx-auto text-center lg:max-w-none lg:text-left lg:flex lg:items-center lg:justify-between">
            <div className="lg:flex-1">
              <h3 className="text-lg font-semibold mb-2">Clinical Updates</h3>
              <p className="text-gray-300 text-sm">
                Get the latest updates on breakthrough treatments, clinical trials, and medical research.
              </p>
            </div>
            <div className="mt-4 lg:mt-0 lg:ml-8">
              <div className="flex flex-col sm:flex-row max-w-md">
                <input
                  type="email"
                  placeholder="Enter your email"
                  className="px-4 py-2 bg-gray-800 border border-gray-700 rounded-l-md sm:rounded-r-none focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-sm flex-1"
                />
                <button className="px-6 py-2 bg-blue-600 hover:bg-blue-700 transition-colors rounded-r-md sm:rounded-l-none text-sm font-medium">
                  Subscribe
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Bottom Footer */}
      <div className="border-t border-gray-800">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
          <div className="flex flex-col md:flex-row md:items-center md:justify-between">
            <div className="text-sm text-gray-400">
              <p>© {currentYear} VAMOS BIOTECH (Shanghai) Co., Ltd. All rights reserved.</p>
              <p className="mt-1">
                Bio-pharmaceutical innovation for advanced life-cell based therapies.
              </p>
            </div>
            
            <div className="mt-4 md:mt-0 flex flex-col sm:flex-row sm:items-center sm:space-x-6">
              {/* Legal Links */}
              <div className="flex flex-wrap items-center space-x-4 text-xs text-gray-400">
                {legalLinks.map((link, index) => (
                  <React.Fragment key={link.name}>
                    <Link
                      to={link.path}
                      className="hover:text-blue-400 transition-colors"
                    >
                      {link.name}
                    </Link>
                    {index < legalLinks.length - 1 && (
                      <span className="text-gray-600">•</span>
                    )}
                  </React.Fragment>
                ))}
              </div>

              {/* Social Links */}
              <div className="flex items-center space-x-4 mt-3 sm:mt-0">
                {socialLinks.map((social) => {
                  const Icon = social.icon;
                  return (
                    <a
                      key={social.name}
                      href={social.url}
                      className="text-gray-400 hover:text-blue-400 transition-colors"
                      aria-label={social.name}
                    >
                      <Icon className="w-5 h-5" />
                    </a>
                  );
                })}
              </div>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default PublicFooter;
