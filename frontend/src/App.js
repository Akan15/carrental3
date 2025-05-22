import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, Link } from 'react-router-dom';
import MapPage from './pages/MapPage';
import LoginPage from './pages/Login';
import RegisterPage from './pages/Register'; 


function App() {
  return (
    <Router>
      <div>
        <nav style={{ display: 'flex', justifyContent: 'flex-end', padding: '1rem' }}>
          <Link to="/login">Войти</Link>
        </nav>

        <Routes>
          {/* При открытии корня — редирект на /map */}
          <Route path="/" element={<Navigate to="/map" />} />
          <Route path="/map" element={<MapPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
