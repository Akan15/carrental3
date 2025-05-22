import React, { useEffect, useState } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import axios from 'axios';

delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
  iconUrl: require('leaflet/dist/images/marker-icon.png'),
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
});

function MapPage() {
  const [cars, setCars] = useState([]);
  const [message, setMessage] = useState('');

  const userId = "682e3fc6585a73b14b43faf1"; // ← Akan ID
 // допустим, сохранён после логина

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_API_URL}/cars/available`)
      .then(res => {
        console.log('cars/available:', res.data);
        setCars(res.data); // при необходимости: setCars(res.data.cars)
      })
      .catch(err => console.error('Ошибка при загрузке машин:', err));
  }, []);

  const handleRent = async (carId) => {
  try {
    const res = await axios.post(`${process.env.REACT_APP_API_URL}/rentals`, {
      userId: "682e3fc6585a73b14b43faf1", // ← временно жёстко
      carId,
      type: 'minute'
    });
    console.log('Rental created:', res.data);
    setMessage('Успешно арендовано!');
    } catch (err) {
        console.error('Ошибка при аренде:', err);
        setMessage('Ошибка при аренде');
    }
    };


  return (
    <div>
      <h2>Available Cars</h2>
      {message && <p>{message}</p>}

      <MapContainer
        center={[51.1605, 71.4704]}
        zoom={13}
        style={{ height: '700px', width: '100%' }}
      >
        <TileLayer
          attribution='&copy; OpenStreetMap contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />

        {Array.isArray(cars) && cars.map(car => {
          if (!car.latitude || !car.longitude) return null;

          return (
            <Marker key={car.id} position={[car.latitude, car.longitude]}>
              <Popup>
                <strong>{car.model}</strong><br />
                {car.city}<br />
                <button onClick={() => handleRent(car.id)}>Арендовать</button>
              </Popup>
            </Marker>
          );
        })}
      </MapContainer>
    </div>
  );
}

export default MapPage;
