import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuthStore } from '../../store/authStore';
import {
  User,
  Menu,
  X,
  ChevronDown,
  LogOut,
  UserCircle,
  Calendar,
  FileText,
  Settings,
  HelpCircle,
  Shield,
  CreditCard,
  Image,
  MoreHorizontal,
} from 'lucide-react';

const PublicHeader: React.FC = () => {
  const navigate = useNavigate();
  const { user, logout, isAuthenticated } = useAuthStore();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isProfileMenuOpen, setIsProfileMenuOpen] = useState(false);
  const [mobileTherapiesOpen, setMobileTherapiesOpen] = useState(false);

  const therapyCategories = [
    {
      name: 'Cancer Therapies',
      items: [
        'Stomach, ovarian, gastrointestinal, bladder cancer',
        'Colorectal, pancreas, prostate cancer',
        'Melanoma therapy'
      ]
    },
    {
      name: 'Infectious Diseases',
      items: [
        'Gastritis & Stomach ulcers',
        'Helicobacter Pylori',
        'Tuberculosis',
        'Pneumonia & Bronchitis'
      ]
    }
  ];

  const handleLogout = () => {
    logout();
    navigate('/');
    setIsProfileMenuOpen(false);
  };

  const getDisplayName = () => {
    if (user?.firstName && user?.lastName) {
      return `${user.firstName} ${user.lastName}`;
    }
    return user?.firstName || 'User';
  };

  const getUserRole = () => {
    return user?.tyUser === 5 ? 'Doctor' : 'Patient';
  };

  const getProfileLink = () => {
    return user?.tyUser === 5 ? '/doctor/profile' : '/profile';
  };

  const getDashboardLink = () => {
    return user?.tyUser === 5 ? '/doctor/dashboard' : '/dashboard';
  };

  const handleMobileNavigation = () => {
    setIsMenuOpen(false);
  };

  return (
    <header className="bg-white shadow-lg relative z-50">
      <nav className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <div className="flex-shrink-0">
            <Link to="/" className="flex items-center">
              <span className="text-2xl font-bold text-blue-600">VCM</span>
              <span className="ml-2 text-xl font-semibold text-gray-900">Medical</span>
            </Link>
          </div>

          {/* Desktop Navigation */}
          <div className="hidden md:block">
            <div className="ml-10 flex items-baseline space-x-4">
              <Link
                to="/"
                className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Home
              </Link>
              <Link
                to="/about"
                className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                About
              </Link>
              
              {/* Therapies Dropdown */}
              <div className="relative group">
                <button className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors flex items-center">
                  Therapies
                  <ChevronDown className="ml-1 h-4 w-4" />
                </button>
                
                <div className="absolute left-0 mt-2 w-96 bg-white rounded-md shadow-lg opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
                  <div className="py-2">
                    {therapyCategories.map((category, index) => (
                      <div key={index} className="px-4 py-2">
                        <h3 className="font-semibold text-gray-900 mb-2">{category.name}</h3>
                        {category.items.map((item, itemIndex) => (
                          <Link
                            key={itemIndex}
                            to={`/therapies/${item.toLowerCase().replace(/\s+/g, '-')}`}
                            className="block px-2 py-1 text-sm text-gray-600 hover:text-blue-600 hover:bg-gray-50 rounded transition-colors"
                          >
                            {item}
                          </Link>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
              </div>

              <Link
                to="/shop"
                className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Shop
              </Link>
              <Link
                to="/research"
                className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Research
              </Link>
              <Link
                to="/contact"
                className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Contact
              </Link>
            </div>
          </div>

          {/* User Menu / Auth Buttons */}
          <div className="hidden md:flex items-center space-x-4">
            {isAuthenticated && user ? (
              <div className="relative">
                <button
                  onClick={() => setIsProfileMenuOpen(!isProfileMenuOpen)}
                  className="flex items-center space-x-3 text-gray-700 hover:text-blue-600 transition-colors"
                >
                  <UserCircle className="h-8 w-8" />
                  <div className="text-left">
                    <div className="text-sm font-medium text-gray-900">{getDisplayName()}</div>
                    <div className="text-xs text-gray-600">{getUserRole()}</div>
                  </div>
                  <ChevronDown className="h-4 w-4" />
                </button>

                {isProfileMenuOpen && (
                  <div className="absolute right-0 mt-2 w-64 bg-white rounded-lg shadow-lg border border-gray-200 py-2 z-50">
                    <div className="px-4 py-3 border-b border-gray-100">
                      <div className="flex items-center space-x-3">
                        <UserCircle className="h-10 w-10 text-blue-600" />
                        <div>
                          <h3 className="font-semibold text-gray-900">{getDisplayName()}</h3>
                          <p className="text-sm text-gray-600">{user.email}</p>
                          <span className="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                            {getUserRole()}
                          </span>
                        </div>
                      </div>
                    </div>
                    
                    <div className="py-2">
                      <Link
                        to={getDashboardLink()}
                        className="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                        onClick={() => setIsProfileMenuOpen(false)}
                      >
                        <User className="mr-3 h-4 w-4" />
                        Dashboard
                      </Link>
                      <Link
                        to={getProfileLink()}
                        className="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                        onClick={() => setIsProfileMenuOpen(false)}
                      >
                        <UserCircle className="mr-3 h-4 w-4" />
                        Profile
                      </Link>
                      <Link
                        to="/appointments"
                        className="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                        onClick={() => setIsProfileMenuOpen(false)}
                      >
                        <Calendar className="mr-3 h-4 w-4" />
                        Appointments
                      </Link>
                      <Link
                        to="/medical-records"
                        className="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                        onClick={() => setIsProfileMenuOpen(false)}
                      >
                        <FileText className="mr-3 h-4 w-4" />
                        Medical Records
                      </Link>
                    </div>
                    
                    <div className="border-t border-gray-100 py-2">
                      <Link
                        to="/settings"
                        className="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                        onClick={() => setIsProfileMenuOpen(false)}
                      >
                        <Settings className="mr-3 h-4 w-4" />
                        Settings
                      </Link>
                      <button
                        onClick={handleLogout}
                        className="w-full flex items-center px-4 py-2 text-sm text-red-600 hover:bg-red-50"
                      >
                        <LogOut className="mr-3 h-4 w-4" />
                        Logout
                      </button>
                    </div>
                  </div>
                )}
              </div>
            ) : (
              <div className="flex items-center space-x-3">
                <Link
                  to="/login"
                  className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors"
                >
                  Login
                </Link>
                <Link
                  to="/register"
                  className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm font-medium transition-colors"
                >
                  Get Started
                </Link>
              </div>
            )}
          </div>

          {/* Mobile menu button */}
          <div className="md:hidden">
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="text-gray-700 hover:text-blue-600 transition-colors"
            >
              {isMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
            </button>
          </div>
        </div>

        {/* Mobile Navigation Menu */}
        {isMenuOpen && (
          <div className="md:hidden">
            <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3 bg-white border-t border-gray-200">
              {/* Mobile User Info */}
              {isAuthenticated && user && (
                <div className="px-3 py-4 border-b border-gray-200">
                  <div className="flex items-center space-x-3">
                    <UserCircle className="h-10 w-10 text-blue-600" />
                    <div>
                      <div className="text-sm font-medium text-gray-900">{getDisplayName()}</div>
                      <div className="text-xs text-gray-600">{getUserRole()}</div>
                    </div>
                  </div>
                </div>
              )}

              {/* Mobile Navigation Links */}
              <Link
                to="/"
                className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                onClick={handleMobileNavigation}
              >
                Home
              </Link>
              <Link
                to="/about"
                className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                onClick={handleMobileNavigation}
              >
                About
              </Link>
              
              {/* Mobile Therapies */}
              <div>
                <button
                  onClick={() => setMobileTherapiesOpen(!mobileTherapiesOpen)}
                  className="w-full text-left text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-base font-medium transition-colors flex items-center justify-between"
                >
                  Therapies
                  <ChevronDown className={`h-4 w-4 transform transition-transform ${mobileTherapiesOpen ? 'rotate-180' : ''}`} />
                </button>
                
                {mobileTherapiesOpen && (
                  <div className="pl-6 space-y-1">
                    {therapyCategories.map((category, index) => (
                      <div key={index} className="py-2">
                        <div className="font-medium text-gray-900 text-sm mb-1">{category.name}</div>
                        {category.items.map((item, itemIndex) => (
                          <Link
                            key={itemIndex}
                            to={`/therapies/${item.toLowerCase().replace(/\s+/g, '-')}`}
                            className="block px-2 py-1 text-sm text-gray-600 hover:text-blue-600 transition-colors"
                            onClick={handleMobileNavigation}
                          >
                            {item}
                          </Link>
                        ))}
                      </div>
                    ))}
                  </div>
                )}
              </div>

              {/* Other Mobile Links */}
              {['Shop', 'Research', 'Contact'].map((item) => (
                <Link
                  key={item}
                  to={`/${item.toLowerCase()}`}
                  className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                  onClick={handleMobileNavigation}
                >
                  {item}
                </Link>
              ))}

              {/* Mobile Auth/User Actions */}
              {isAuthenticated && user ? (
                <div className="border-t border-gray-200 pt-4 space-y-1">
                  <Link
                    to={getDashboardLink()}
                    className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                    onClick={handleMobileNavigation}
                  >
                    Dashboard
                  </Link>
                  <Link
                    to={getProfileLink()}
                    className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                    onClick={handleMobileNavigation}
                  >
                    Profile
                  </Link>
                  <button
                    onClick={handleLogout}
                    className="w-full text-left text-red-600 hover:text-red-700 px-3 py-2 rounded-md text-base font-medium transition-colors"
                  >
                    Logout
                  </button>
                </div>
              ) : (
                <div className="border-t border-gray-200 pt-4 space-y-1">
                  <Link
                    to="/login"
                    className="text-gray-900 hover:text-blue-600 block px-3 py-2 rounded-md text-base font-medium transition-colors"
                    onClick={handleMobileNavigation}
                  >
                    Login
                  </Link>
                  <Link
                    to="/register"
                    className="bg-blue-600 hover:bg-blue-700 text-white block px-3 py-2 rounded-md text-base font-medium transition-colors"
                    onClick={handleMobileNavigation}
                  >
                    Get Started
                  </Link>
                </div>
              )}
            </div>
          </div>
        )}
      </nav>
    </header>
  );
};

export default PublicHeader;
