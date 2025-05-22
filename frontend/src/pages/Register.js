import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

function RegisterPage() {
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      await axios.post(`${process.env.REACT_APP_API_URL}/register`, form);
      setMessage('Регистрация успешна!');
      navigate('/login');
    } catch (err) {
      setMessage('Ошибка регистрации');
    }
  };

  return (
    <div style={{ padding: '2rem' }}>
      <h2>Регистрация</h2>
      <form onSubmit={handleSubmit}>
        <input name="name" placeholder="Имя" onChange={handleChange} required />
        <br />
        <input name="email" type="email" placeholder="Email" onChange={handleChange} required />
        <br />
        <input name="password" type="password" placeholder="Пароль" onChange={handleChange} required />
        <br />
        <button type="submit">Зарегистрироваться</button>
      </form>
      <p>{message}</p>
    </div>
  );
}

export default RegisterPage;
