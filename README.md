# Final Project API Documentation

Ini adalah dokumentasi API untuk project `Final Project` menggunakan `Golang` + `Gin Gonic`.

## 🛠️ Authentication
- `POST /register` - Register new user
- `POST /login` - Login and get token
- `GET /catalog` - Get all product catalogs (Public)

---

## 👤 Customer Routes (Authorization: Bearer Token `customer`)
| Method | Endpoint | Description |
|:------:|:--------|:------------|
| GET | `/customers/:id` | Get customer by ID |
| POST | `/customers` | Create new customer |
| PUT | `/customers/:id` | Update customer data |
| GET | `/products` | Get all products |
| GET | `/orders/user/:userID` | Get all orders by user ID |
| GET | `/orders/:id` | Track order by order ID |
| POST | `/orders` | Create new order |
| PUT | `/orders/:id/payment` | Process order payment |
| PUT | `/orders/:id/complete` | Complete delivery process |
| GET | `/users/:id` | Get user data by ID |
| PUT | `/users/:id` | Update user data |
| GET | `/categories` | Get all categories |

---

## 🛡️ Admin Routes (Authorization: Bearer Token `admin`)
| Method | Endpoint | Description |
|:------:|:--------|:------------|
| POST | `/categories` | Create new category |
| GET | `/categories/all` | Get all categories (admin view) |
| DELETE | `/categories/:id` | Delete category by ID |
| GET | `/customers` | Get all customers |
| DELETE | `/customers/:id` | Delete customer by ID |
| GET | `/products/all` | Get all products (admin view) |
| POST | `/products` | Create new product |
| PUT | `/products/:id` | Update product by ID |
| DELETE | `/products/:id` | Delete product by ID |
| GET | `/orders/order/:userID` | Get orders by user ID (admin view) |
| GET | `/users` | Get all users |
| DELETE | `/users/:id` | Delete user by ID |
| GET | `/reports/products-by-quantity` | Report: Products sorted by quantity sold |
| GET | `/reports/customer-spendings` | Report: Customer spending data |
| GET | `/reports/products-by-nominal` | Report: Products sorted by nominal sales |

---

## 🔐 Authentication Rules
- **Customer token** dibutuhkan untuk mengakses customer routes.
- **Admin token** dibutuhkan untuk mengakses admin routes.

---

## 📚 Tech Stack
- Golang
- Gin Gonic
- JWT Authentication
- MySQL
- GORM

---

## 🚀 How to Run
```bash
go mod tidy
go run main.go
```

---

Kalau mau, aku juga bisa bantuin bikin tabel lebih lengkap — kayak request body & contoh responsenya.  
Mau sekalian sekalian? 🔥  
(opsional: buat langsung siap publish ke Postman docs juga) 🚀  
Mau lanjut? 🚀💬 
