#### **📌 `README.md` ফাইল তৈরি করুন:**
👉 **প্রজেক্টের রুট ডিরেক্টরিতে** `README.md` ফাইল তৈরি করুন এবং নিচের কোড কপি-পেস্ট করুন।

```md
# 📌 Go Fiber + MySQL CRUD API

## 🛠 টেকনোলজি
এই প্রজেক্টে নিম্নলিখিত প্রযুক্তি ব্যবহার করা হয়েছে:
- 🚀 **Golang (Go)**
- ⚡ **Fiber Framework**
- 🛢 **GORM (ORM for Golang)**
- 🐬 **MySQL Database**

---

## 📌 প্রজেক্ট সেটআপ (Installation & Setup)
### 🛠 1️⃣ প্রথমে প্রজেক্ট ক্লোন করুন (অথবা তৈরি করুন)
```sh
git clone https://github.com/Abidanik86/go-fiber-crud-1-practice.git
cd go-fiber-mysql-crud
```
অথবা নতুন **Go Module** তৈরি করুন:
```sh
mkdir go-fiber-mysql-crud
cd go-fiber-mysql-crud
go mod init go-fiber-mysql-crud
```

### 🛠 2️⃣ প্রয়োজনীয় প্যাকেজ ইনস্টল করুন
```sh
go get -u github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

### 🛠 3️⃣ MySQL ডাটাবেজ তৈরি করুন
MySQL এ একটি নতুন ডাটাবেজ তৈরি করুন:
```sql
CREATE DATABASE go_fiber_crud;
```
**📌 MySQL কানেকশন স্ট্রিং (`database/database.go`) পরিবর্তন করুন:**  
```go
dsn := "root:password@tcp(127.0.0.1:3306)/go_fiber_crud?charset=utf8mb4&parseTime=True&loc=Local"
```
> 🔹 **`root:password` পরিবর্তন করে আপনার MySQL ইউজারনেম ও পাসওয়ার্ড দিন।**

---

## 🚀 রান করার নির্দেশনা (Run Instructions)
### 1️⃣ **প্রজেক্ট রান করুন**
```sh
go run main.go
```
আপনার **Go Fiber CRUD API** **`http://localhost:8080`** এ চলবে।

---

## 🌍 API Endpoint List
| Method | Endpoint       | Description         |
|--------|--------------|---------------------|
| **GET** | `/books`     | সমস্ত বই দেখুন      |
| **GET** | `/books/:id` | নির্দিষ্ট বই দেখুন |
| **POST** | `/books`    | নতুন বই যোগ করুন  |
| **PUT** | `/books/:id` | নির্দিষ্ট বই আপডেট করুন |
| **DELETE** | `/books/:id` | নির্দিষ্ট বই মুছুন |

---

## 📌 API টেস্ট (Postman / cURL)
### 🔹 **নতুন বই তৈরি করুন (POST)**
```sh
curl -X POST "http://localhost:8080/books" -H "Content-Type: application/json" -d '{"title":"Golang Basics", "author":"John Doe"}'
```

### 🔹 **সমস্ত বই দেখুন (GET)**
```sh
curl -X GET "http://localhost:8080/books"
```

### 🔹 **নির্দিষ্ট বই দেখুন (GET)**
```sh
curl -X GET "http://localhost:8080/books/1"
```

### 🔹 **বই আপডেট করুন (PUT)**
```sh
curl -X PUT "http://localhost:8080/books/1" -H "Content-Type: application/json" -d '{"title":"Advanced Golang", "author":"Jane Doe"}'
```

### 🔹 **বই ডিলিট করুন (DELETE)**
```sh
curl -X DELETE "http://localhost:8080/books/1"
```

---

## 📌 ফোল্ডার স্ট্রাকচার
```
go-fiber-mysql-crud/
│── database/
│   ├── database.go
│
│── models/
│   ├── book.go
│
│── routes/
│   ├── book_routes.go
│
│── main.go
│── go.mod
│── README.md
```

---

## ⚡ ডাটাবেজ মাইগ্রেশন চালান
```sh
go run main.go
```
> 🔹 এটি MySQL-এ **`books`** টেবিল তৈরি করবে।

---

## 🎯 কন্ট্রিবিউশন
এই প্রজেক্টে কন্ট্রিবিউট করতে চান? **Pull Request** পাঠান, অথবা ইস্যু খুলুন! 🚀

---

## 📜 লাইসেন্স
এই প্রজেক্টটি **MIT License** দ্বারা লাইসেন্সকৃত।

---

## 📞 যোগাযোগ
আপনার যদি কোনো প্রশ্ন থাকে, আমাকে ইমেইল করতে পারেন:  
✉️ **your-abidanik86@gmail.com**
```

