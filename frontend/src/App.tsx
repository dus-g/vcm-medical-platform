import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PublicHeader from './components/layout/PublicHeader';
import PublicFooter from './components/layout/PublicFooter';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import VerifyOTP from './pages/VerifyOTP';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50 flex flex-col">
        <PublicHeader />
        <main className="flex-grow">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/verify-otp" element={<VerifyOTP />} />
            {/* Placeholder routes for new pages */}
            <Route path="/about" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">About Us</h1><p>Coming soon...</p></div>} />
            <Route path="/therapies" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Medical Therapies</h1><p>Coming soon...</p></div>} />
            <Route path="/therapies/*" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Therapy Details</h1><p>Coming soon...</p></div>} />
            <Route path="/shop" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Shop</h1><p>Coming soon...</p></div>} />
            <Route path="/research" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Research</h1><p>Coming soon...</p></div>} />
            <Route path="/contact" element={<div className="pt-20 p-8"><h1 className="text-2xl font-bold">Contact Us</h1><p>Coming soon...</p></div>} />
          </Routes>
        </main>
        <PublicFooter />
      </div>
    </Router>
  );
}

export default App;
