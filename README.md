# 🚗 Carrental Akan

## 📌 Overview
**Carrental** — это микросервисная платформа для аренды автомобилей с поминутной тарификацией и отображением доступных машин на карте. Система построена на gRPC и включает регистрацию пользователей, бронирование машин, завершение аренды, e-mail уведомления и мониторинг.

---

## 🧱 Services

- `user-service` – управление пользователями (регистрация, логин, email)
- `car-service` – CRUD машин, фильтрация, геолокация
- `rental-service` – аренда машины, завершение, расчёт цены
- `api-gateway` – REST вход, gRPC вызовы (WIP)
- `consumer-service` – подписка на события (NATS)

---

## ⚙️ Technologies

- Go (Golang)
- gRPC + Protocol Buffers
- MongoDB
- Docker + Docker Compose
- NATS (message broker)
- Prometheus + Grafana (monitoring)
- Clean Architecture
- SMTP Email (Mail.ru/Gmail)
- grpcurl (тестирование)
- JS (Frontend – карта) [*Bonus*]

---

## ▶️ How to Run

```bash
# Клонировать проект и перейти в папку
git clone https://github.com/yourusername/carrental.git
cd carrental

# Собрать и запустить все микросервисы
docker-compose up --build
