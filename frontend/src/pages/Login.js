import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios';

function LoginPage() {
  const [form, setForm] = useState({ email: '', password: '' });
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const res = await axios.post(`${process.env.REACT_APP_API_URL}/login`, form);
      localStorage.setItem('token', res.data.token);
      localStorage.setItem('userId', res.data.userId);
      setMessage('Успешный вход!');
      navigate('/map');
    } catch (err) {
      setMessage('Ошибка входа');
    }
  };

  return (
    <div style={{ padding: '2rem' }}>
      <h2>Вход</h2>
      <form onSubmit={handleSubmit}>
        <input name="email" type="email" placeholder="Email" onChange={handleChange} required />
        <br />
        <input name="password" type="password" placeholder="Пароль" onChange={handleChange} required />
        <br />
        <button type="submit">Войти</button>
      </form>
      <p>{message}</p>
      <p>Нет аккаунта? <Link to="/register">Регистрация</Link></p>
    </div>
  );
}

export default LoginPage;
