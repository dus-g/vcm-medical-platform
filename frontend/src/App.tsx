import React, { useState } from 'react';

// Simple routing without react-router-dom to avoid dependency issues
function App() {
  const [currentPage, setCurrentPage] = useState('home');
  const [user, setUser] = useState(null);

  const navigate = (page) => {
    setCurrentPage(page);
  };

  const mockLogin = (email, password) => {
    // Mock login
    setUser({ email, name: 'John Doe' });
    navigate('dashboard');
  };

  const mockRegister = (userData) => {
    // Mock registration
    navigate('verify-otp');
  };

  const mockVerifyOTP = () => {
    setUser({ email: 'user@example.com', name: 'New User' });
    navigate('complete-profile');
  };

  const mockCompleteProfile = () => {
    navigate('dashboard');
  };

  const logout = () => {
    setUser(null);
    navigate('home');
  };

  // Common styles
  const styles = {
    container: {
      minHeight: '100vh',
      fontFamily: 'Arial, sans-serif',
    },
    header: {
      background: 'white',
      padding: '1rem 0',
      borderBottom: '1px solid #e5e7eb',
      boxShadow: '0 1px 3px rgba(0,0,0,0.1)',
    },
    headerContent: {
      maxWidth: '1200px',
      margin: '0 auto',
      padding: '0 1rem',
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
    },
    logo: {
      display: 'flex',
      alignItems: 'center',
      gap: '0.5rem',
    },
    logoIcon: {
      width: '40px',
      height: '40px',
      background: 'linear-gradient(135deg, #3b82f6, #1d4ed8)',
      borderRadius: '8px',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      color: 'white',
      fontWeight: 'bold',
      fontSize: '20px',
    },
    nav: {
      display: 'flex',
      gap: '2rem',
    },
    navLink: {
      color: '#374151',
      textDecoration: 'none',
      fontWeight: '500',
      cursor: 'pointer',
      padding: '0.5rem 0',
    },
    button: {
      background: '#3b82f6',
      color: 'white',
      border: 'none',
      padding: '0.75rem 1.5rem',
      borderRadius: '8px',
      fontWeight: '600',
      cursor: 'pointer',
      fontSize: '14px',
    },
    buttonSecondary: {
      background: 'white',
      color: '#374151',
      border: '1px solid #d1d5db',
      padding: '0.75rem 1.5rem',
      borderRadius: '8px',
      fontWeight: '600',
      cursor: 'pointer',
      fontSize: '14px',
    },
    hero: {
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      color: 'white',
      padding: '5rem 0',
      textAlign: 'center',
    },
    heroContent: {
      maxWidth: '1200px',
      margin: '0 auto',
      padding: '0 1rem',
    },
    section: {
      padding: '4rem 0',
    },
    sectionContent: {
      maxWidth: '1200px',
      margin: '0 auto',
      padding: '0 1rem',
    },
    card: {
      background: 'white',
      borderRadius: '12px',
      padding: '2rem',
      boxShadow: '0 4px 6px rgba(0,0,0,0.05)',
      border: '1px solid #e5e7eb',
    },
    formContainer: {
      minHeight: '100vh',
      background: 'linear-gradient(135deg, #f3f4f6, #e5e7eb)',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      padding: '2rem',
    },
    form: {
      background: 'white',
      borderRadius: '16px',
      padding: '2rem',
      boxShadow: '0 10px 25px rgba(0,0,0,0.1)',
      width: '100%',
      maxWidth: '400px',
    },
    input: {
      width: '100%',
      padding: '0.75rem',
      border: '1px solid #d1d5db',
      borderRadius: '8px',
      fontSize: '16px',
      marginBottom: '1rem',
      boxSizing: 'border-box',
    },
    grid: {
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))',
      gap: '2rem',
      marginTop: '2rem',
    },
  };

  // Header Component
  const Header = () => (
    <header style={styles.header}>
      <div style={styles.headerContent}>
        <div style={styles.logo}>
          <div style={styles.logoIcon}>V</div>
          <span style={{ fontSize: '20px', fontWeight: 'bold', color: '#111827' }}>
            VCM Medical
          </span>
        </div>
        <nav style={styles.nav}>
          <span style={styles.navLink} onClick={() => navigate('home')}>Home</span>
          <span style={styles.navLink} onClick={() => navigate('about')}>About</span>
          <span style={styles.navLink} onClick={() => navigate('contact')}>Contact</span>
        </nav>
        <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
          {user ? (
            <>
              <span>Welcome, {user.name}!</span>
              <button style={styles.button} onClick={() => navigate('dashboard')}>
                Dashboard
              </button>
              <button style={styles.buttonSecondary} onClick={logout}>
                Logout
              </button>
            </>
          ) : (
            <>
              <button style={styles.buttonSecondary} onClick={() => navigate('login')}>
                Sign In
              </button>
              <button style={styles.button} onClick={() => navigate('register')}>
                Get Started
              </button>
            </>
          )}
        </div>
      </div>
    </header>
  );

  // Home Page
  const HomePage = () => (
    <div style={styles.container}>
      <Header />
      
      {/* Hero Section */}
      <section style={styles.hero}>
        <div style={styles.heroContent}>
          <h1 style={{ fontSize: '3.5rem', fontWeight: 'bold', marginBottom: '1.5rem', margin: 0 }}>
            Advanced Medical Treatments
          </h1>
          <p style={{ fontSize: '1.25rem', marginBottom: '2rem', opacity: 0.9 }}>
            VCM Medical Platform connects patients with specialized doctors for cutting-edge treatments
          </p>
          <div style={{ display: 'flex', gap: '1rem', justifyContent: 'center', flexWrap: 'wrap' }}>
            <button 
              style={{ ...styles.button, background: 'white', color: '#667eea', fontSize: '18px', padding: '1rem 2rem' }}
              onClick={() => navigate('register')}
            >
              Start Assessment
            </button>
            <button 
              style={{ ...styles.button, background: 'transparent', border: '2px solid white', fontSize: '18px', padding: '1rem 2rem' }}
              onClick={() => navigate('about')}
            >
              Learn More
            </button>
          </div>
        </div>
      </section>

      {/* Specialties */}
      <section style={{ ...styles.section, background: '#f9fafb' }}>
        <div style={styles.sectionContent}>
          <h2 style={{ fontSize: '2.5rem', fontWeight: 'bold', textAlign: 'center', marginBottom: '3rem' }}>
            Our Medical Specialties
          </h2>
          <div style={styles.grid}>
            {[
              { icon: '🔬', title: 'Cancer Immunotherapy', desc: 'CAR-T Cell Therapy and advanced treatments' },
              { icon: '🛡️', title: 'Autoimmune Disorders', desc: 'Precision therapy for complex conditions' },
              { icon: '👁️', title: 'Ophthalmology', desc: 'Advanced eye disease treatments' },
              { icon: '🧠', title: 'Neurology', desc: 'Stroke rehabilitation and brain health' },
              { icon: '🫁', title: 'Respiratory', desc: 'Lung disease management' },
              { icon: '🦠', title: 'Infectious Diseases', desc: 'Advanced infection treatments' },
            ].map((specialty, index) => (
              <div key={index} style={styles.card}>
                <div style={{ fontSize: '3rem', marginBottom: '1rem' }}>{specialty.icon}</div>
                <h3 style={{ fontSize: '1.5rem', fontWeight: 'bold', marginBottom: '1rem' }}>
                  {specialty.title}
                </h3>
                <p style={{ color: '#6b7280' }}>{specialty.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section style={{ ...styles.section, background: '#111827', color: 'white' }}>
        <div style={{ ...styles.sectionContent, textAlign: 'center' }}>
          <h2 style={{ fontSize: '2.5rem', fontWeight: 'bold', marginBottom: '1rem' }}>
            Ready to Start Your Treatment Journey?
          </h2>
          <p style={{ fontSize: '1.25rem', marginBottom: '2rem', opacity: 0.8 }}>
            Join thousands of patients who have found effective treatments
          </p>
          <button 
            style={{ ...styles.button, fontSize: '18px', padding: '1rem 2rem' }}
            onClick={() => navigate('register')}
          >
            Get Started Today
          </button>
        </div>
      </section>
    </div>
  );

  // Login Page
  const LoginPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
      e.preventDefault();
      mockLogin(email, password);
    };

    return (
      <div style={styles.formContainer}>
        <div style={styles.form}>
          <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
            <div style={{ ...styles.logoIcon, margin: '0 auto 1rem' }}>V</div>
            <h2 style={{ fontSize: '2rem', fontWeight: 'bold', marginBottom: '0.5rem' }}>Welcome back</h2>
            <p style={{ color: '#6b7280' }}>Sign in to your VCM Medical account</p>
          </div>
          
          <form onSubmit={handleSubmit}>
            <input
              style={styles.input}
              type="email"
              placeholder="Email address"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
            <input
              style={styles.input}
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            <button type="submit" style={{ ...styles.button, width: '100%', marginBottom: '1rem' }}>
              Sign In
            </button>
          </form>
          
          <div style={{ textAlign: 'center' }}>
            <span style={{ color: '#6b7280' }}>Don't have an account? </span>
            <span style={{ color: '#3b82f6', cursor: 'pointer' }} onClick={() => navigate('register')}>
              Sign up
            </span>
          </div>
        </div>
      </div>
    );
  };

  // Register Page
  const RegisterPage = () => {
    const [formData, setFormData] = useState({
      firstName: '',
      lastName: '',
      email: '',
      password: '',
      confirmPassword: '',
      userType: '0'
    });

    const handleSubmit = (e) => {
      e.preventDefault();
      if (formData.password !== formData.confirmPassword) {
        alert('Passwords do not match');
        return;
      }
      mockRegister(formData);
    };

    return (
      <div style={styles.formContainer}>
        <div style={styles.form}>
          <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
            <div style={{ ...styles.logoIcon, margin: '0 auto 1rem' }}>V</div>
            <h2 style={{ fontSize: '2rem', fontWeight: 'bold', marginBottom: '0.5rem' }}>Create account</h2>
            <p style={{ color: '#6b7280' }}>Join VCM Medical Platform today</p>
          </div>
          
          <form onSubmit={handleSubmit}>
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
              <input
                style={styles.input}
                type="text"
                placeholder="First name"
                value={formData.firstName}
                onChange={(e) => setFormData({...formData, firstName: e.target.value})}
                required
              />
              <input
                style={styles.input}
                type="text"
                placeholder="Last name"
                value={formData.lastName}
                onChange={(e) => setFormData({...formData, lastName: e.target.value})}
                required
              />
            </div>
            <input
              style={styles.input}
              type="email"
              placeholder="Email address"
              value={formData.email}
              onChange={(e) => setFormData({...formData, email: e.target.value})}
              required
            />
            <select
              style={styles.input}
              value={formData.userType}
              onChange={(e) => setFormData({...formData, userType: e.target.value})}
            >
              <option value="0">Patient</option>
              <option value="5">Doctor</option>
              <option value="1">Agent</option>
            </select>
            <input
              style={styles.input}
              type="password"
              placeholder="Password"
              value={formData.password}
              onChange={(e) => setFormData({...formData, password: e.target.value})}
              required
            />
            <input
              style={styles.input}
              type="password"
              placeholder="Confirm password"
              value={formData.confirmPassword}
              onChange={(e) => setFormData({...formData, confirmPassword: e.target.value})}
              required
            />
            <button type="submit" style={{ ...styles.button, width: '100%', marginBottom: '1rem' }}>
              Create Account
            </button>
          </form>
          
          <div style={{ textAlign: 'center' }}>
            <span style={{ color: '#6b7280' }}>Already have an account? </span>
            <span style={{ color: '#3b82f6', cursor: 'pointer' }} onClick={() => navigate('login')}>
              Sign in
            </span>
          </div>
        </div>
      </div>
    );
  };

  // Verify OTP Page
  const VerifyOTPPage = () => {
    const [otp, setOtp] = useState('');

    const handleSubmit = (e) => {
      e.preventDefault();
      mockVerifyOTP();
    };

    return (
      <div style={styles.formContainer}>
        <div style={styles.form}>
          <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
            <div style={{ ...styles.logoIcon, margin: '0 auto 1rem' }}>✉️</div>
            <h2 style={{ fontSize: '2rem', fontWeight: 'bold', marginBottom: '0.5rem' }}>Verify your email</h2>
            <p style={{ color: '#6b7280' }}>We've sent a verification code to your email</p>
          </div>
          
          <form onSubmit={handleSubmit}>
            <input
              style={{ ...styles.input, textAlign: 'center', fontSize: '2rem', letterSpacing: '0.5rem' }}
              type="text"
              placeholder="000000"
              maxLength="6"
              value={otp}
              onChange={(e) => setOtp(e.target.value.replace(/[^0-9]/g, ''))}
              required
            />
            <button type="submit" style={{ ...styles.button, width: '100%', marginBottom: '1rem' }}>
              Verify Email
            </button>
          </form>
          
          <div style={{ textAlign: 'center' }}>
            <span style={{ color: '#6b7280' }}>Didn't receive the code? </span>
            <span style={{ color: '#3b82f6', cursor: 'pointer' }}>
              Resend
            </span>
          </div>
        </div>
      </div>
    );
  };

  // Complete Profile Page
  const CompleteProfilePage = () => {
    const [profileData, setProfileData] = useState({
      phone: '',
      dateOfBirth: '',
      gender: '',
      emergencyContact: '',
      emergencyPhone: ''
    });

    const handleSubmit = (e) => {
      e.preventDefault();
      mockCompleteProfile();
    };

    return (
      <div style={styles.formContainer}>
        <div style={{ ...styles.form, maxWidth: '600px' }}>
          <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
            <div style={{ ...styles.logoIcon, margin: '0 auto 1rem' }}>👤</div>
            <h2 style={{ fontSize: '2rem', fontWeight: 'bold', marginBottom: '0.5rem' }}>Complete Your Profile</h2>
            <p style={{ color: '#6b7280' }}>Please provide additional information</p>
          </div>
          
          <form onSubmit={handleSubmit}>
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
              <input
                style={styles.input}
                type="tel"
                placeholder="Phone number"
                value={profileData.phone}
                onChange={(e) => setProfileData({...profileData, phone: e.target.value})}
                required
              />
              <input
                style={styles.input}
                type="date"
                value={profileData.dateOfBirth}
                onChange={(e) => setProfileData({...profileData, dateOfBirth: e.target.value})}
                required
              />
            </div>
            <select
              style={styles.input}
              value={profileData.gender}
              onChange={(e) => setProfileData({...profileData, gender: e.target.value})}
              required
            >
              <option value="">Select gender</option>
              <option value="male">Male</option>
              <option value="female">Female</option>
              <option value="other">Other</option>
            </select>
            <input
              style={styles.input}
              type="text"
              placeholder="Emergency contact name"
              value={profileData.emergencyContact}
              onChange={(e) => setProfileData({...profileData, emergencyContact: e.target.value})}
              required
            />
            <input
              style={styles.input}
              type="tel"
              placeholder="Emergency contact phone"
              value={profileData.emergencyPhone}
              onChange={(e) => setProfileData({...profileData, emergencyPhone: e.target.value})}
              required
            />
            <button type="submit" style={{ ...styles.button, width: '100%', marginBottom: '1rem' }}>
              Complete Profile
            </button>
          </form>
        </div>
      </div>
    );
  };

  // Dashboard Page
  const DashboardPage = () => (
    <div style={styles.container}>
      <Header />
      <div style={{ ...styles.section, background: '#f9fafb', minHeight: '80vh' }}>
        <div style={styles.sectionContent}>
          <h1 style={{ fontSize: '3rem', fontWeight: 'bold', marginBottom: '2rem' }}>
            Welcome back, {user?.name || 'User'}!
          </h1>
          
          <div style={styles.grid}>
            {[
              { icon: '📋', title: 'Medical Assessments', desc: 'Complete your medical assessment forms' },
              { icon: '📅', title: 'Appointments', desc: 'Schedule and manage appointments' },
              { icon: '📈', title: 'Treatment Plans', desc: 'Track your treatment progress' },
              { icon: '💬', title: 'Chat with Doctors', desc: 'Real-time communication' },
              { icon: '🖼️', title: 'Medical Images', desc: 'Upload and manage medical images' },
              { icon: '⚙️', title: 'Profile Settings', desc: 'Update your information' },
            ].map((action, index) => (
              <div key={index} style={{ ...styles.card, cursor: 'pointer' }}>
                <div style={{ fontSize: '3rem', marginBottom: '1rem' }}>{action.icon}</div>
                <h3 style={{ fontSize: '1.5rem', fontWeight: 'bold', marginBottom: '1rem' }}>
                  {action.title}
                </h3>
                <p style={{ color: '#6b7280', marginBottom: '1rem' }}>{action.desc}</p>
                <span style={{ color: '#3b82f6', fontWeight: '600' }}>Get Started →</span>
              </div>
            ))}
          </div>

          <div style={{ ...styles.card, marginTop: '2rem', background: 'linear-gradient(135deg, #3b82f6, #1d4ed8)', color: 'white' }}>
            <h3 style={{ fontSize: '1.5rem', fontWeight: 'bold', marginBottom: '1rem' }}>
              Ready to start your medical assessment?
            </h3>
            <p style={{ marginBottom: '1.5rem', opacity: 0.9 }}>
              Complete your comprehensive medical evaluation to get personalized treatment recommendations.
            </p>
            <button style={{ ...styles.button, background: 'white', color: '#3b82f6' }}>
              Start Assessment
            </button>
          </div>
        </div>
      </div>
    </div>
  );

  // Route rendering
  const renderPage = () => {
    switch (currentPage) {
      case 'home': return <HomePage />;
      case 'login': return <LoginPage />;
      case 'register': return <RegisterPage />;
      case 'verify-otp': return <VerifyOTPPage />;
      case 'complete-profile': return <CompleteProfilePage />;
      case 'dashboard': return <DashboardPage />;
      case 'about': return (
        <div style={styles.container}>
          <Header />
          <div style={{ ...styles.section, textAlign: 'center' }}>
            <div style={styles.sectionContent}>
              <h1 style={{ fontSize: '3rem', fontWeight: 'bold', marginBottom: '2rem' }}>About VCM Medical</h1>
              <p style={{ fontSize: '1.25rem', color: '#6b7280' }}>
                Advanced medical treatments for complex conditions. Coming soon!
              </p>
            </div>
          </div>
        </div>
      );
      case 'contact': return (
        <div style={styles.container}>
          <Header />
          <div style={{ ...styles.section, textAlign: 'center' }}>
            <div style={styles.sectionContent}>
              <h1 style={{ fontSize: '3rem', fontWeight: 'bold', marginBottom: '2rem' }}>Contact Us</h1>
              <p style={{ fontSize: '1.25rem', color: '#6b7280' }}>
                Get in touch with our medical experts. Support available 24/7.
              </p>
            </div>
          </div>
        </div>
      );
      default: return <HomePage />;
    }
  };

  return renderPage();
}

export default App;
