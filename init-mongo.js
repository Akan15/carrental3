// === Создаём пользователя для базы carrental ===
db = db.getSiblingDB("carrental");

db.createUser({
  user: "admin",
  pwd: "secret",
  roles: [
    {
      role: "readWrite",
      db: "carrental"
    }
  ]
});

// Добавляем тестовые машины
db.cars.insertMany([
  {
    id: "1",
    brand: "Toyota",
    model: "Camry"
  },
  {
    id: "2",
    brand: "Tesla",
    model: "Model 3"
  }
]);

// === Создаём базу restaurant и коллекции ===
db = db.getSiblingDB("restaurant");

db.createUser({
  user: "admin",
  pwd: "secret",
  roles: [{ role: "readWrite", db: "restaurant" }],
});

// Коллекции создаются автоматически при вставке, но можно вручную:
db.createCollection("users");
db.createCollection("cars");
db.createCollection("rentals");
