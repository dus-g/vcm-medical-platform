import React from 'react';
import { Link } from 'react-router-dom';
import Layout from '../components/layout/Layout';
import { MEDICAL_SPECIALTIES } from '../utils/constants';

const HomePage: React.FC = () => {
  return (
    <Layout>
      {/* Hero Section */}
      <section className="medical-gradient text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-5xl md:text-6xl font-bold mb-6">
            Advanced Medical Treatments for Complex Conditions
          </h1>
          <p className="text-xl mb-8 max-w-3xl mx-auto">
            VCM Medical Platform connects patients with specialized doctors for cutting-edge treatments 
            in cancer immunotherapy, autoimmune disorders, ophthalmology, and more.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link 
              to="/register" 
              className="bg-white text-primary-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition-colors shadow-lg"
            >
              Start Assessment
            </Link>
            <Link 
              to="/about" 
              className="border-2 border-white text-white px-8 py-4 rounded-lg font-semibold hover:bg-white hover:text-primary-600 transition-colors"
            >
              Learn More
            </Link>
          </div>
        </div>
      </section>

      {/* Medical Specialties */}
      <section className="py-16 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl md:text-4xl font-bold text-center mb-4">Our Medical Specialties</h2>
          <p className="text-xl text-gray-600 text-center mb-12 max-w-3xl mx-auto">
            Cutting-edge treatments across multiple medical disciplines
          </p>
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {MEDICAL_SPECIALTIES.map((specialty) => (
              <div key={specialty.id} className="card hover:shadow-lg transition-all duration-300">
                <div className="text-4xl mb-4">{specialty.icon}</div>
                <h3 className="text-xl font-semibold mb-3">{specialty.title}</h3>
                <p className="text-gray-600 mb-4">{specialty.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 bg-gray-900 text-white">
        <div className="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            Ready to Start Your Treatment Journey?
          </h2>
          <p className="text-xl text-gray-300 mb-8">
            Join thousands of patients who have found effective treatments
          </p>
          <Link 
            to="/register" 
            className="btn-primary text-lg px-8 py-4"
          >
            Get Started Today
          </Link>
        </div>
      </section>
    </Layout>
  );
};

export default HomePage;
