import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useAuthStore } from '../../store/authStore';
import { 
  Menu, 
  X, 
  User,
  Globe,
  DollarSign,
  ChevronDown,
  LogIn,
  UserPlus,
  ShoppingCart,
  Activity,
  ChevronRight
} from 'lucide-react';

export const PublicHeader: React.FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isTherapiesOpen, setIsTherapiesOpen] = useState(false);
  const [isProfileMenuOpen, setIsProfileMenuOpen] = useState(false);
  const [mobileTherapiesOpen, setMobileTherapiesOpen] = useState(false);
  
  const [selectedLanguage, setSelectedLanguage] = useState('English');
  const [selectedCurrency, setSelectedCurrency] = useState('USD');
  const [cartItems] = useState(2);
  
  const { user } = useAuthStore();
  const location = useLocation();

  const languages = [
    { code: 'en', name: 'English', flag: '🇺🇸' },
    { code: 'zh', name: '中文', flag: '🇨🇳' }
  ];

  const currencies = [
    { code: 'USD', name: 'US Dollar', symbol: '$' },
    { code: 'CNY', name: 'Chinese Yuan', symbol: '¥' }
  ];

  const therapyCategories = [
    {
      name: 'Autoimmune Disorders',
      items: [
        { name: 'Psoriasis Vulgaris', path: '/therapies/psoriasis' },
        { name: 'Rheumatoid Arthritis', path: '/therapies/rheumatoid-arthritis' },
        { name: 'Lupus', path: '/therapies/lupus' },
        { name: 'Hashimoto\'s Disease', path: '/therapies/hashimotos' }
      ]
    },
    {
      name: 'Cancers',
      items: [
        { name: 'B-cell Leukemia & Lymphoma (CAR-T)', path: '/therapies/car-t' },
        { name: 'HER2-CAR-T for Solid Tumors', path: '/therapies/her2-car-t' },
        { name: 'Colorectal & Pancreatic Cancer', path: '/therapies/colorectal' },
        { name: 'Melanoma Therapy', path: '/therapies/melanoma' }
      ]
    },
    {
      name: 'Eye Diseases',
      items: [
        { name: 'Optic Nerve Atrophy', path: '/therapies/optic-nerve' },
        { name: 'Glaucoma', path: '/therapies/glaucoma' },
        { name: 'Macular Degeneration', path: '/therapies/macular-degeneration' },
        { name: 'Amblyopia', path: '/therapies/amblyopia' },
        { name: 'Strabismus', path: '/therapies/strabismus' }
      ]
    },
    {
      name: 'Gastrointestinal',
      items: [
        { name: 'Gastritis', path: '/therapies/gastritis' },
        { name: 'Stomach Ulcers', path: '/therapies/stomach-ulcers' },
        { name: 'Helicobacter Pylori', path: '/therapies/h-pylori' },
        { name: 'Leaky Gut Syndrome', path: '/therapies/leaky-gut' }
      ]
    },
    {
      name: 'Respiratory',
      items: [
        { name: 'Tuberculosis', path: '/therapies/tuberculosis' },
        { name: 'Pneumonia', path: '/therapies/pneumonia' },
        { name: 'Chronic Bronchitis', path: '/therapies/bronchitis' }
      ]
    },
    {
      name: 'Neurological',
      items: [
        { name: 'Stroke Rehabilitation', path: '/therapies/stroke' },
        { name: 'Alzheimer\'s Disease', path: '/therapies/alzheimers' },
        { name: 'Autism Spectrum Disorders', path: '/therapies/autism' },
        { name: 'Migraines', path: '/therapies/migraines' },
        { name: 'Insomnia', path: '/therapies/insomnia' }
      ]
    }
  ];

  const publicNavItems = [
    { name: 'Home', path: '/' },
    { name: 'Therapies', hasDropdown: true },
    { name: 'About', path: '/about' },
    { name: 'Shop', path: '/shop' },
    { name: 'Research', path: '/research' },
    { name: 'Contact', path: '/contact' }
  ];

  const getDashboardPath = () => {
    if (!user) return '/dashboard';
    return user.userType === 5 ? '/doctor/dashboard' : '/dashboard';
  };

  const isActivePath = (path: string) => {
    if (path === '/') return location.pathname === '/';
    return location.pathname.startsWith(path);
  };

  const handleLanguageChange = (language: any) => {
    setSelectedLanguage(language.name);
  };

  const handleCurrencyChange = (currency: any) => {
    setSelectedCurrency(currency.code);
  };

  const handleMobileNavigation = () => {
    setIsMenuOpen(false);
    setMobileTherapiesOpen(false);
  };

  const getUserInitials = () => {
    if (user?.name) {
      const names = user.name.split(' ');
      if (names.length >= 2) {
        return names[0][0] + names[1][0];
      }
      return names[0][0];
    }
    return 'U';
  };

  return (
    <header className="fixed w-full top-0 z-50 bg-white/95 backdrop-blur-md border-b border-gray-100 shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-20">
          
          {/* Logo - Far Left */}
          <div className="flex items-center flex-shrink-0">
            <Link to="/" className="flex items-center group">
              <div className="h-16 w-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center transition-transform duration-200 group-hover:scale-105">
                <span className="text-white font-bold text-xl">VCM</span>
              </div>
            </Link>
          </div>

          {/* Desktop Navigation - Center */}
          <nav className="hidden lg:flex items-center justify-center flex-1 px-8">
            <div className="flex items-center space-x-8">
              {publicNavItems.map((item) => (
                <div key={item.name} className="relative group">
                  {item.hasDropdown ? (
                    <>
                      <button
                        className="flex items-center text-gray-700 hover:text-blue-600 font-medium transition-colors duration-200 text-base py-2 px-3 rounded-lg hover:bg-gray-50"
                        onMouseEnter={() => setIsTherapiesOpen(true)}
                        onMouseLeave={() => setIsTherapiesOpen(false)}
                      >
                        {item.name}
                        <ChevronDown className="w-4 h-4 ml-1" />
                      </button>
                      
                      {/* Therapies Mega Menu */}
                      {isTherapiesOpen && (
                        <div 
                          className="absolute left-1/2 transform -translate-x-1/2 top-full mt-2 w-[800px] bg-white rounded-2xl shadow-2xl border border-gray-100 z-50"
                          onMouseEnter={() => setIsTherapiesOpen(true)}
                          onMouseLeave={() => setIsTherapiesOpen(false)}
                        >
                          <div className="p-6">
                            {/* Header */}
                            <div className="mb-6 pb-4 border-b border-gray-100">
                              <Link 
                                to="/therapies"
                                className="inline-flex items-center text-lg font-bold text-blue-600 hover:text-blue-700 transition-colors"
                              >
                                All Medical Therapies
                                <ChevronRight className="w-4 h-4 ml-1" />
                              </Link>
                              <p className="text-sm text-gray-600 mt-1">Comprehensive treatment protocols across 6 medical specialties</p>
                            </div>
                            
                            {/* Categories Grid */}
                            <div className="grid grid-cols-3 gap-6">
                              {therapyCategories.map((category, index) => (
                                <div key={index} className="space-y-3">
                                  <h4 className="text-sm font-semibold text-gray-900 uppercase tracking-wider">
                                    {category.name}
                                  </h4>
                                  <ul className="space-y-2">
                                    {category.items.map((item, itemIndex) => (
                                      <li key={itemIndex}>
                                        <Link
                                          to={item.path}
                                          className="block text-sm text-gray-600 hover:text-blue-600 hover:bg-blue-50 px-3 py-2 rounded-lg transition-colors duration-150"
                                        >
                                          {item.name}
                                        </Link>
                                      </li>
                                    ))}
                                  </ul>
                                </div>
                              ))}
                            </div>
                          </div>
                        </div>
                      )}
                    </>
                  ) : (
                    <Link
                      to={item.path!}
                      className={`text-base font-medium transition-colors duration-200 px-3 py-2 rounded-lg ${
                        isActivePath(item.path!)
                          ? 'text-blue-600 bg-blue-50'
                          : 'text-gray-700 hover:text-blue-600 hover:bg-gray-50'
                      }`}
                    >
                      {item.name}
                    </Link>
                  )}
                </div>
              ))}
            </div>
          </nav>

          {/* Right Side Actions */}
          <div className="hidden lg:flex items-center space-x-3 flex-shrink-0">
            
            {/* Cart */}
            <Link 
              to="/shop"
              className="relative p-3 text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors"
            >
              <ShoppingCart className="w-5 h-5" />
              {cartItems > 0 && (
                <span className="absolute -top-1 -right-1 bg-gradient-to-r from-blue-500 to-blue-600 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center font-bold shadow-lg">
                  {cartItems}
                </span>
              )}
            </Link>

            {/* Unified Profile/Menu Dropdown */}
            <div className="relative">
              <button
                onClick={() => setIsProfileMenuOpen(!isProfileMenuOpen)}
                className="flex items-center space-x-2 p-2 hover:bg-gray-50 rounded-lg transition-colors"
              >
                {user ? (
                  <div className="w-9 h-9 bg-gradient-to-r from-blue-500 to-blue-600 rounded-full flex items-center justify-center shadow-lg">
                    <span className="text-white text-sm font-bold">{getUserInitials()}</span>
                  </div>
                ) : (
                  <div className="w-9 h-9 bg-gray-100 rounded-full flex items-center justify-center">
                    <User className="w-5 h-5 text-gray-600" />
                  </div>
                )}
              </button>

              {isProfileMenuOpen && (
                <div className="absolute right-0 mt-2 w-80 bg-white rounded-xl shadow-2xl border border-gray-100 py-2 z-50">
                  
                  {user ? (
                    <>
                      {/* Logged In User Header */}
                      <div className="px-4 py-4 border-b border-gray-100 bg-gradient-to-r from-blue-50 to-blue-100">
                        <div className="flex items-center space-x-3">
                          <div className="w-12 h-12 bg-gradient-to-r from-blue-500 to-blue-600 rounded-full flex items-center justify-center">
                            <span className="text-white font-bold text-lg">{getUserInitials()}</span>
                          </div>
                          <div className="flex-1">
                            <h3 className="font-semibold text-gray-900">{user?.name || 'User'}</h3>
                            <p className="text-sm text-gray-600">{user?.email}</p>
                            <span className="inline-block px-2 py-1 bg-blue-100 text-blue-700 text-xs font-medium rounded-full mt-1">
                              {user.userType === 5 ? 'Doctor' : 'Patient'}
                            </span>
                          </div>
                        </div>
                        <Link
                          to={user.userType === 5 ? "/doctor/profile" : "/profile"}
                          onClick={() => setIsProfileMenuOpen(false)}
                          className="text-blue-600 hover:text-blue-700 text-sm font-medium mt-2 inline-block"
                        >
                          Profile Management
                        </Link>
                      </div>

                      {/* User Menu Items */}
                      <div className="py-2">
                        <Link
                          to="/patient-portal"
                          onClick={() => setIsProfileMenuOpen(false)}
                          className="flex items-center px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 transition-colors"
                        >
                          <Activity className="w-4 h-4 mr-3 text-gray-400" />
                          <span className="font-medium">Patient Portal</span>
                        </Link>

                        <Link
                          to="/orders"
                          onClick={() => setIsProfileMenuOpen(false)}
                          className="flex items-center px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 transition-colors"
                        >
                          <ShoppingCart className="w-4 h-4 mr-3 text-gray-400" />
                          <span className="font-medium">Orders</span>
                        </Link>
                      </div>
                    </>
                  ) : (
                    <>
                      {/* Not Logged In Header */}
                      <div className="px-4 py-4 border-b border-gray-100">
                        <h3 className="font-semibold text-gray-900 mb-2">Welcome to Our Platform</h3>
                        <div className="space-y-2">
                          <Link
                            to="/login"
                            onClick={() => setIsProfileMenuOpen(false)}
                            className="flex items-center justify-center w-full px-4 py-2 text-sm font-medium text-gray-700 border border-gray-300 hover:border-blue-300 hover:text-blue-600 rounded-lg transition-colors"
                          >
                            <LogIn className="w-4 h-4 mr-2" />
                            Sign In
                          </Link>
                          <Link
                            to="/register"
                            onClick={() => setIsProfileMenuOpen(false)}
                            className="flex items-center justify-center w-full px-4 py-2 text-sm font-medium text-white bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 rounded-lg transition-colors shadow-lg"
                          >
                            <UserPlus className="w-4 h-4 mr-2" />
                            Get Started
                          </Link>
                        </div>
                      </div>
                    </>
                  )}

                  {/* Language & Currency Settings (Always Visible) */}
                  <div className="py-2 border-t border-gray-100">
                    <div className="px-4 py-2 text-xs font-semibold text-gray-500 uppercase tracking-wider">
                      Preferences
                    </div>
                    
                    {/* Language Selector */}
                    <div className="px-4 py-3">
                      <div className="flex items-center justify-between mb-2">
                        <div className="flex items-center">
                          <Globe className="w-4 h-4 mr-3 text-gray-400" />
                          <span className="text-sm font-medium text-gray-700">Language: {selectedLanguage}</span>
                        </div>
                        <ChevronRight className="w-4 h-4 text-gray-400" />
                      </div>
                      <div className="grid grid-cols-2 gap-2">
                        {languages.map((language) => (
                          <button
                            key={language.code}
                            onClick={() => handleLanguageChange(language)}
                            className={`flex items-center space-x-2 px-3 py-2 text-sm rounded-lg transition-colors ${
                              selectedLanguage === language.name
                                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                                : 'bg-gray-50 text-gray-600 hover:bg-gray-100'
                            }`}
                          >
                            <span>{language.flag}</span>
                            <span className="font-medium">{language.name}</span>
                          </button>
                        ))}
                      </div>
                    </div>

                    {/* Currency Selector */}
                    <div className="px-4 py-3">
                      <div className="flex items-center justify-between mb-2">
                        <div className="flex items-center">
                          <DollarSign className="w-4 h-4 mr-3 text-gray-400" />
                          <span className="text-sm font-medium text-gray-700">Currency: {selectedCurrency}</span>
                        </div>
                        <ChevronRight className="w-4 h-4 text-gray-400" />
                      </div>
                      <div className="grid grid-cols-1 gap-2">
                        {currencies.map((currency) => (
                          <button
                            key={currency.code}
                            onClick={() => handleCurrencyChange(currency)}
                            className={`flex items-center justify-between px-3 py-2 text-sm rounded-lg transition-colors ${
                              selectedCurrency === currency.code
                                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                                : 'bg-gray-50 text-gray-600 hover:bg-gray-100'
                            }`}
                          >
                            <span className="font-medium">{currency.name}</span>
                            <span className="text-lg font-bold">{currency.symbol}</span>
                          </button>
                        ))}
                      </div>
                    </div>

                    {/* Additional Options */}
                    <div className="border-t border-gray-100 pt-2">
                      {user && (
                        <button
                          onClick={() => setIsProfileMenuOpen(false)}
                          className="flex items-center w-full px-4 py-3 text-sm text-red-600 hover:bg-red-50 transition-colors"
                        >
                          <LogIn className="w-4 h-4 mr-3 rotate-180" />
                          <span className="font-medium">Sign Out</span>
                        </button>
                      )}
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>

          {/* Mobile menu button */}
          <div className="lg:hidden">
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="p-2 rounded-lg text-gray-700 hover:text-blue-600 hover:bg-gray-50 transition-colors"
            >
              {isMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
            </button>
          </div>
        </div>

        {/* Mobile Navigation */}
        {isMenuOpen && (
          <div className="lg:hidden bg-white border-t border-gray-100 shadow-lg">
            <div className="px-4 py-4 space-y-1 max-h-96 overflow-y-auto">
              
              {/* Mobile Nav Items */}
              <Link
                to="/"
                onClick={() => handleMobileNavigation()}
                className="block px-4 py-3 text-base font-medium text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors"
              >
                Home
              </Link>

              {/* Mobile Therapies Dropdown */}
              <div>
                <button
                  onClick={() => setMobileTherapiesOpen(!mobileTherapiesOpen)}
                  className="flex items-center justify-between w-full px-4 py-3 text-base font-medium text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors"
                >
                  <span>Therapies</span>
                  <ChevronDown className={`w-4 h-4 transition-transform duration-200 ${mobileTherapiesOpen ? 'rotate-180' : ''}`} />
                </button>
                
                {mobileTherapiesOpen && (
                  <div className="ml-4 mt-2 space-y-2 border-l-2 border-blue-200 pl-4">
                    <Link
                      to="/therapies"
                      onClick={() => handleMobileNavigation()}
                      className="block px-4 py-2 text-sm font-semibold text-blue-600 hover:text-blue-700 hover:bg-blue-50 rounded-lg transition-colors"
                    >
                      All Therapies
                    </Link>
                    
                    {therapyCategories.slice(0, 3).map((category, index) => (
                      <div key={index} className="space-y-1">
                        <div className="text-xs font-semibold text-gray-500 uppercase tracking-wider px-4 py-2 bg-gray-50 rounded">
                          {category.name}
                        </div>
                        {category.items.slice(0, 2).map((item, itemIndex) => (
                          <Link
                            key={itemIndex}
                            to={item.path}
                            onClick={() => handleMobileNavigation()}
                            className="block px-4 py-2 text-sm text-gray-600 hover:text-blue-600 hover:bg-gray-50 rounded transition-colors"
                          >
                            {item.name}
                          </Link>
                        ))}
                      </div>
                    ))}
                  </div>
                )}
              </div>

              {/* Other Nav Items */}
              {['About', 'Shop', 'Research', 'Contact'].map((item) => (
                <Link
                  key={item}
                  to={`/${item.toLowerCase()}`}
                  onClick={() => handleMobileNavigation()}
                  className="block px-4 py-3 text-base font-medium text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors"
                >
                  {item}
                </Link>
              ))}

              {/* Mobile User Section */}
              <div className="pt-4 border-t border-gray-200 mt-4">
                {user ? (
                  <div className="space-y-2">
                    <div className="flex items-center space-x-3 px-4 py-3 bg-blue-50 rounded-lg">
                      <div className="w-10 h-10 bg-gradient-to-r from-blue-500 to-blue-600 rounded-full flex items-center justify-center">
                        <span className="text-white text-sm font-bold">{getUserInitials()}</span>
                      </div>
                      <div>
                        <div className="text-sm font-medium text-gray-900">{user?.name || 'User'}</div>
                        <div className="text-xs text-gray-600">{user.userType === 5 ? 'Doctor' : 'Patient'}</div>
                      </div>
                    </div>
                    
                    <Link
                      to={getDashboardPath()}
                      onClick={() => setIsMenuOpen(false)}
                      className="flex items-center px-4 py-3 text-base font-medium text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors"
                    >
                      <Activity className="w-5 h-5 mr-3" />
                      Dashboard
                    </Link>
                  </div>
                ) : (
                  <div className="space-y-2">
                    <Link
                      to="/login"
                      onClick={() => setIsMenuOpen(false)}
                      className="block w-full text-center px-4 py-3 text-base font-medium text-gray-700 hover:text-blue-600 border border-gray-300 hover:border-blue-300 rounded-lg transition-colors"
                    >
                      Sign In
                    </Link>
                    <Link
                      to="/register"
                      onClick={() => setIsMenuOpen(false)}
                      className="block w-full text-center px-4 py-3 text-base font-medium text-white bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 rounded-lg transition-colors shadow-lg"
                    >
                      Get Started
                    </Link>
                  </div>
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </header>
  );
};

export default PublicHeader;
