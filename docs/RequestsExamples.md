
***

### **localhost:8080/user**

#### **POST** - Создание пользователя

```bash
  curl --location 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --data '{
        "username": "USERNAME",
        "password": "PASSWORD"
    }'
```

#### **DELETE** - Удаление пользователя

```bash
  curl --location --request DELETE 'http://localhost:8080/user?username=USERNAME
```

***

### **localhost:8080/booking**

#### **GET** - Получение всех бронирований пользователя

```bash
  curl --location 'localhost:8080/booking?username=USERNAME'
```

#### **POST** - Создание бронирования

```bash
  curl --location 'localhost:8080/booking' \
    --header 'Content-Type: application/json' \
    --data '{
        "username": "USERNAME",
        "delta": 2 
    }'
```

#### **DELETE** - Удаление бронирования

Принимает на вход uuid бронирования, который генерируется при создании бронирования. Для получения uuid можно воспользоваться запросом на получение всех бронирований пользователя.

```bash
  curl --location --request DELETE 'localhost:8080/booking?id=UUID'
```