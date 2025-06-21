import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PublicHeader from './components/layout/PublicHeader';
import PublicFooter from './components/layout/PublicFooter';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import VerifyOTP from './pages/VerifyOTP';
import CompleteProfile from './pages/CompleteProfile';
import Dashboard from './pages/Dashboard';
import Test from './pages/Test';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50 flex flex-col">
        <Routes>
          {/* Public routes with header/footer */}
          <Route path="/" element={
            <>
              <PublicHeader />
              <main className="flex-grow">
                <Home />
              </main>
              <PublicFooter />
            </>
          } />
          
          {/* Auth routes without header/footer */}
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/verify-otp" element={<VerifyOTP />} />
          <Route path="/complete-profile" element={<CompleteProfile />} />
          
          {/* Protected routes */}
          <Route path="/dashboard" element={<Dashboard />} />
          
          {/* Test route */}
          <Route path="/test" element={<Test />} />
          
          {/* Placeholder routes */}
          <Route path="/about" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">About Us</h1><p>Coming soon...</p></div>} />
          <Route path="/therapies" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Medical Therapies</h1><p>Coming soon...</p></div>} />
          <Route path="/shop" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Shop</h1><p>Coming soon...</p></div>} />
          <Route path="/research" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Research</h1><p>Coming soon...</p></div>} />
          <Route path="/contact" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Contact Us</h1><p>Coming soon...</p></div>} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
