import React from 'react';

function App() {
  return (
    <div style={{ 
      minHeight: '100vh', 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'center',
      fontFamily: 'Arial, sans-serif',
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      color: 'white'
    }}>
      <div style={{ textAlign: 'center' }}>
        <div style={{
          width: '80px',
          height: '80px',
          background: 'rgba(255,255,255,0.2)',
          borderRadius: '20px',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          margin: '0 auto 20px',
          fontSize: '32px',
          fontWeight: 'bold'
        }}>
          V
        </div>
        <h1 style={{ fontSize: '48px', margin: '0 0 20px' }}>
          VCM Medical Platform
        </h1>
        <p style={{ fontSize: '20px', margin: '0 0 30px' }}>
          Advanced Medical Treatments for Complex Conditions
        </p>
        <button style={{
          background: 'white',
          color: '#667eea',
          border: 'none',
          padding: '15px 30px',
          fontSize: '18px',
          borderRadius: '10px',
          cursor: 'pointer',
          fontWeight: 'bold'
        }}>
          Coming Soon
        </button>
        <div style={{ marginTop: '20px', fontSize: '14px', opacity: '0.8' }}>
          Platform Status: ✅ Online
        </div>
      </div>
    </div>
  );
}

export default App;
