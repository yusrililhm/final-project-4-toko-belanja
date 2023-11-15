# Final Project 4 - Toko Belanja

For this final project, we were tasked with developing an application called Toko Belanja. 

# Group 5 - GLNG-KS-007
- Januwar By Khaqi
- Yasid Al Mubarok
- Yusril Ilham Kholid

# Demo
- [API](https://final-project-4-toko-belanja-production.up.railway.app)
- [Swagger](https://final-project-4-toko-belanja-production.up.railway.app/swagger/index.html)

# Admin
- Email : admin@tokobelanja.com
- Password : rahasia

# Tech Stack
- [Go](https://go.dev/)
- [Gin-gonic](https://gin-gonic.com/)
- [Govalidator](https://github.com/asaskevich/govalidator)
- [Jwt-go](https://github.com/golang-jwt/jwt)
- [Crypto](https://pkg.go.dev/crypto)
- [Swagger Documentation](https://github.com/swaggo)
- [Postgres Driver](https://pkg.go.dev/github.com/lib/pq)
- [PostgreSQL](https://www.postgresql.org/)
- [Godotenv](https://github.com/joho/godotenv)

# Schema
| Domain       | Method   | Endpoint                        | Middleware                     | Description            |
|--------------|----------|---------------------------------|--------------------------------|------------------------|
| Users        | POST     | /users/register                 | -                              | User register          |
| Users        | POST     | /users/login                    | -                              | User login             |
| Users        | PATCH    | /users/topup                    | Authentication                 | User topup             |
| Categories   | POST     | /categories                     | Authentication & Authorization | Add category           |
| Categories   | GET      | /categories                     | Authentication & Authorization | Get Categories         |
| Categories   | PATCH    | /categories/:categoryId         | Authentication & Authorization | Update Category        |
| Categories   | DELETE   | /categories/:categoryId         | Authentication & Authorization | Delete Category        |
| Products     | POST     | /products                       | Authentication & Authorization | Add Product            |
| Products     | GET      | /products                       | Authentication                 | Get Products           |
| Products     | PUT      | /products/:productId            | Authentication & Authorization | Update Product         |
| Products     | DELETE   | /products/:productId            | Authentication & Authorization | Delete Product         |
| Transactions | POST     | /transactions                   | Authentication                 | Add Transactions       |
| Transactions | GET      | /transactions/my-transactions   | Authentication                 | Get My Transactionss   |
| Transactions | GET      | /transactions/user-transactions | Authentication & Authorization | Get Users Transactions |
