import React, { useState, FormEvent, CSSProperties } from 'react';

// Define types for better TypeScript support
interface User {
  email: string;
  name: string;
}

interface UserData {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  phone: string;
  dateOfBirth: string;
  gender: string;
  occupation: string;
  medicalCondition: string;
  emergencyContact: string;
  emergencyPhone: string;
}

// Simple routing without react-router-dom to avoid dependency issues
function App() {
  const [currentPage, setCurrentPage] = useState<string>('home');
  const [user, setUser] = useState<User | null>(null);

  const navigate = (page: string) => {
    setCurrentPage(page);
  };

  const mockLogin = (email: string, password: string) => {
    // Mock login
    setUser({ email, name: 'John Doe' });
    navigate('dashboard');
  };

  const mockRegister = (userData: UserData) => {
    // Mock registration
    navigate('verify-otp');
  };

  const mockVerifyOTP = () => {
    setUser({ email: 'user@example.com', name: 'New User' });
    navigate('complete-profile');
  };

  // Define styles with proper TypeScript types
  const styles: { [key: string]: CSSProperties } = {
    container: {
      fontFamily: "'Segoe UI', Tahoma, Geneva, Verdana, sans-serif",
      lineHeight: '1.6',
      margin: 0,
      padding: 0,
      backgroundColor: '#f8f9fa'
    },
    header: {
      backgroundColor: '#2c3e50',
      color: 'white',
      padding: '1rem 2rem',
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
      boxShadow: '0 2px 4px rgba(0,0,0,0.1)'
    },
    logo: {
      fontSize: '1.8rem',
      fontWeight: 'bold',
      margin: 0
    },
    nav: {
      display: 'flex',
      gap: '2rem'
    },
    navButton: {
      background: 'none',
      border: 'none',
      color: 'white',
      cursor: 'pointer',
      fontSize: '1rem',
      padding: '0.5rem 1rem',
      borderRadius: '4px',
      transition: 'background-color 0.3s'
    },
    main: {
      padding: '2rem',
      maxWidth: '1200px',
      margin: '0 auto'
    },
    hero: {
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      color: 'white',
      padding: '4rem 2rem',
      textAlign: 'center' as const,
      borderRadius: '12px',
      marginBottom: '3rem'
    },
    heroTitle: {
      fontSize: '3rem',
      marginBottom: '1rem',
      fontWeight: 'bold'
    },
    heroSubtitle: {
      fontSize: '1.3rem',
      marginBottom: '2rem',
      opacity: 0.9
    },
    card: {
      backgroundColor: 'white',
      borderRadius: '8px',
      padding: '2rem',
      boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
      marginBottom: '2rem'
    },
    cardTitle: {
      fontSize: '1.8rem',
      marginBottom: '1rem',
      color: '#2c3e50'
    },
    form: {
      display: 'flex',
      flexDirection: 'column',
      gap: '1rem'
    },
    input: {
      width: '100%',
      padding: '0.75rem',
      border: '1px solid #ddd',
      borderRadius: '4px',
      fontSize: '1rem',
      marginBottom: '1rem',
      boxSizing: 'border-box' as const
    },
    button: {
      backgroundColor: '#3498db',
      color: 'white',
      border: 'none',
      padding: '0.75rem 1.5rem',
      borderRadius: '4px',
      fontSize: '1rem',
      cursor: 'pointer',
      transition: 'background-color 0.3s',
      marginTop: '1rem'
    },
    buttonSecondary: {
      backgroundColor: '#95a5a6',
      color: 'white',
      border: 'none',
      padding: '0.75rem 1.5rem',
      borderRadius: '4px',
      fontSize: '1rem',
      cursor: 'pointer',
      transition: 'background-color 0.3s',
      marginTop: '0.5rem'
    },
    grid: {
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))',
      gap: '2rem',
      marginTop: '2rem'
    },
    footer: {
      backgroundColor: '#2c3e50',
      color: 'white',
      textAlign: 'center' as const,
      padding: '2rem',
      marginTop: '3rem'
    }
  };

  // Home Page Component
  const HomePage = () => (
    <div style={styles.container}>
      <section style={styles.hero}>
        <h1 style={styles.heroTitle}>VCM Medical Platform</h1>
        <p style={styles.heroSubtitle}>
          Advanced medical treatments with cutting-edge immunotherapy protocols. 
          Connect with our medical experts. Support available 24/7.
        </p>
        <button 
          style={styles.button}
          onClick={() => navigate('register')}
        >
          Get Started Today
        </button>
      </section>

      <div style={styles.grid}>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>🧬 Cancer Immunotherapy</h3>
          <p>CAR-T Cell Therapy, BiTE Antibodies, and Neoantigen-TIL therapy for various cancer types.</p>
        </div>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>🔬 Autoimmune Disorders</h3>
          <p>Precision therapy for Psoriasis, Rheumatoid Arthritis, Lupus, and Hashimoto's Thyroiditis.</p>
        </div>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>👁️ Ophthalmology</h3>
          <p>Advanced treatments for Optic Nerve Atrophy, Glaucoma, Macular Degeneration, and more.</p>
        </div>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>🧠 Neurological Sciences</h3>
          <p>Stroke rehabilitation, Alzheimer's treatment, Autism interventions, and migraine therapy.</p>
        </div>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>🫁 Respiratory Medicine</h3>
          <p>Tuberculosis precision treatment, Pneumonia immunotherapy, and chronic bronchitis management.</p>
        </div>
        <div style={styles.card}>
          <h3 style={styles.cardTitle}>🦠 Infectious Diseases</h3>
          <p>HPV immunotherapy, broad-spectrum antivirals, and antibiotic-resistant infection treatment.</p>
        </div>
      </div>
    </div>
  );

  // Login Page Component
  const LoginPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e: FormEvent) => {
      e.preventDefault();
      mockLogin(email, password);
    };

    return (
      <div style={styles.main}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Welcome Back</h2>
          <form style={styles.form} onSubmit={handleSubmit}>
            <input
              type="email"
              placeholder="Email Address"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              style={styles.input}
              required
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              style={styles.input}
              required
            />
            <button type="submit" style={styles.button}>
              Sign In
            </button>
            <button 
              type="button" 
              style={styles.buttonSecondary}
              onClick={() => navigate('register')}
            >
              Create New Account
            </button>
          </form>
        </div>
      </div>
    );
  };

  // Registration Page Component
  const RegistrationPage = () => {
    const [formData, setFormData] = useState<UserData>({
      firstName: '',
      lastName: '',
      email: '',
      password: '',
      phone: '',
      dateOfBirth: '',
      gender: '',
      occupation: '',
      medicalCondition: '',
      emergencyContact: '',
      emergencyPhone: ''
    });

    const handleSubmit = (e: FormEvent) => {
      e.preventDefault();
      mockRegister(formData);
    };

    const handleInputChange = (field: keyof UserData, value: string) => {
      setFormData(prev => ({
        ...prev,
        [field]: value
      }));
    };

    return (
      <div style={styles.main}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Create Your Account</h2>
          <form style={styles.form} onSubmit={handleSubmit}>
            <input
              type="text"
              placeholder="First Name"
              value={formData.firstName}
              onChange={(e) => handleInputChange('firstName', e.target.value)}
              style={styles.input}
              required
            />
            <input
              type="text"
              placeholder="Last Name"
              value={formData.lastName}
              onChange={(e) => handleInputChange('lastName', e.target.value)}
              style={styles.input}
              required
            />
            <input
              type="email"
              placeholder="Email Address"
              value={formData.email}
              onChange={(e) => handleInputChange('email', e.target.value)}
              style={styles.input}
              required
            />
            <select
              value={formData.gender}
              onChange={(e) => handleInputChange('gender', e.target.value)}
              style={styles.input}
              required
            >
              <option value="">Select Gender</option>
              <option value="male">Male</option>
              <option value="female">Female</option>
              <option value="other">Other</option>
            </select>
            <input
              type="tel"
              placeholder="Phone Number"
              value={formData.phone}
              onChange={(e) => handleInputChange('phone', e.target.value)}
              style={styles.input}
              required
            />
            <input
              type="date"
              placeholder="Date of Birth"
              value={formData.dateOfBirth}
              onChange={(e) => handleInputChange('dateOfBirth', e.target.value)}
              style={styles.input}
              required
            />
            <button type="submit" style={styles.button}>
              Create Account
            </button>
          </form>
        </div>
      </div>
    );
  };

  // OTP Verification Page
  const OTPVerificationPage = () => {
    const [otp, setOtp] = useState('');

    const handleSubmit = (e: FormEvent) => {
      e.preventDefault();
      mockVerifyOTP();
    };

    return (
      <div style={styles.main}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Verify Your Account</h2>
          <p>Please enter the 6-digit verification code sent to your email.</p>
          <form style={styles.form} onSubmit={handleSubmit}>
            <input
              type="text"
              placeholder="000000"
              value={otp}
              onChange={(e) => setOtp(e.target.value)}
              style={{ 
                ...styles.input, 
                textAlign: 'center' as const, 
                fontSize: '2rem', 
                letterSpacing: '0.5rem' 
              }}
              maxLength={6}
              required
            />
            <button type="submit" style={styles.button}>
              Verify Account
            </button>
          </form>
        </div>
      </div>
    );
  };

  // Complete Profile Page
  const CompleteProfilePage = () => {
    const [profileData, setProfileData] = useState({
      occupation: '',
      medicalCondition: '',
      emergencyContact: '',
      emergencyPhone: ''
    });

    const handleSubmit = (e: FormEvent) => {
      e.preventDefault();
      navigate('dashboard');
    };

    const handleInputChange = (field: string, value: string) => {
      setProfileData(prev => ({
        ...prev,
        [field]: value
      }));
    };

    return (
      <div style={styles.main}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Complete Your Profile</h2>
          <form style={styles.form} onSubmit={handleSubmit}>
            <input
              type="text"
              placeholder="Occupation"
              value={profileData.occupation}
              onChange={(e) => handleInputChange('occupation', e.target.value)}
              style={styles.input}
            />
            <input
              type="text"
              placeholder="Primary Medical Condition (if any)"
              value={profileData.medicalCondition}
              onChange={(e) => handleInputChange('medicalCondition', e.target.value)}
              style={styles.input}
            />
            <select
              value={profileData.medicalCondition}
              onChange={(e) => handleInputChange('medicalCondition', e.target.value)}
              style={styles.input}
            >
              <option value="">Select Medical Specialty</option>
              <option value="cancer">Cancer Immunotherapy</option>
              <option value="autoimmune">Autoimmune Disorders</option>
              <option value="ophthalmology">Ophthalmology</option>
              <option value="neurology">Neurological Sciences</option>
              <option value="respiratory">Respiratory Medicine</option>
              <option value="infectious">Infectious Diseases</option>
            </select>
            <input
              type="text"
              placeholder="Emergency Contact Name"
              value={profileData.emergencyContact}
              onChange={(e) => handleInputChange('emergencyContact', e.target.value)}
              style={styles.input}
            />
            <input
              type="tel"
              placeholder="Emergency Contact Phone"
              value={profileData.emergencyPhone}
              onChange={(e) => handleInputChange('emergencyPhone', e.target.value)}
              style={styles.input}
            />
            <button type="submit" style={styles.button}>
              Complete Setup
            </button>
          </form>
        </div>
      </div>
    );
  };

  // Dashboard Page
  const DashboardPage = () => (
    <div style={styles.main}>
      <div style={styles.card}>
        <h2 style={styles.cardTitle}>
          Welcome back, {user?.name || 'User'}!
        </h2>
        <p>Your VCM Medical Platform dashboard is ready.</p>
        
        <div style={styles.grid}>
          <div style={styles.card}>
            <h3>📋 Medical Assessments</h3>
            <p>Complete your comprehensive medical evaluation.</p>
            <button style={styles.button}>Start Assessment</button>
          </div>
          <div style={styles.card}>
            <h3>🩺 Appointments</h3>
            <p>Schedule consultations with our medical experts.</p>
            <button style={styles.button}>Book Appointment</button>
          </div>
          <div style={styles.card}>
            <h3>💬 Chat Support</h3>
            <p>Get instant support from our medical team.</p>
            <button style={styles.button}>Start Chat</button>
          </div>
          <div style={styles.card}>
            <h3>📊 Treatment Progress</h3>
            <p>Track your treatment journey and outcomes.</p>
            <button style={styles.button}>View Progress</button>
          </div>
        </div>
      </div>
    </div>
  );

  // Render the appropriate page based on current route
  const renderPage = () => {
    switch (currentPage) {
      case 'login':
        return <LoginPage />;
      case 'register':
        return <RegistrationPage />;
      case 'verify-otp':
        return <OTPVerificationPage />;
      case 'complete-profile':
        return <CompleteProfilePage />;
      case 'dashboard':
        return <DashboardPage />;
      default:
        return <HomePage />;
    }
  };

  return (
    <div style={styles.container}>
      <header style={styles.header}>
        <h1 style={styles.logo}>VCM Medical</h1>
        <nav style={styles.nav}>
          <button 
            style={styles.navButton}
            onClick={() => navigate('home')}
          >
            Home
          </button>
          {user ? (
            <>
              <button 
                style={styles.navButton}
                onClick={() => navigate('dashboard')}
              >
                Dashboard
              </button>
              <button 
                style={styles.navButton}
                onClick={() => {
                  setUser(null);
                  navigate('home');
                }}
              >
                Logout
              </button>
              <span>Welcome, {user.name}!</span>
            </>
          ) : (
            <>
              <button 
                style={styles.navButton}
                onClick={() => navigate('login')}
              >
                Login
              </button>
              <button 
                style={styles.navButton}
                onClick={() => navigate('register')}
              >
                Register
              </button>
            </>
          )}
        </nav>
      </header>

      <main>
        {renderPage()}
      </main>

      <footer style={styles.footer}>
        <p>&copy; 2025 VCM Medical Platform. Advanced medical treatments for better health outcomes.</p>
      </footer>
    </div>
  );
}

export default App;
