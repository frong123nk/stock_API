# stock_API

#### Add .env file 
```
ACCESS_SECRET = YOUR_ACCESS_SECRET_KEY
DB_HOST       = YOUR_DB_HOST
DB_USERNAME   = YOUR_DB_USERNAME
DB_PASSWORD   = YOUR_DB_PASSWORD
DB_NAME       = YOUR_DB_NAME
DB_PORT       = YOUR_DB_PORT

```
#### URL Paths

```
User : 
  POST : /api/v2/login (login)
  POST : /api/v2/register (register)
Product :
  GET    : /api/v2/product (Get all product or Get product of url request parameter)
  POST   : /api/v2/product (Create product from Multipart or Urlencoded Form)
  PUT    : /api/v2/product (Edit product from Multipart or Urlencoded Form)
  DELETE : /api/v2/product (Delete product from Multipart or Urlencoded Form)
Transaction :
  GET  : /api/v2/transaction (Get transaction in user)
  POST : /api/v2/transaction (Create transaction in user)
```

#### Start 

```
go run main.go
```
