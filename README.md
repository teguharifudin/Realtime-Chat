![](https://www.teguharief.com/img/teguh-arief.png)

# Realtime Chat

This is a realtime chat built with Go (JWT auth, Websocket), Next.js, and TypeScript using PostgreSQL.

## Installation

Install the app on terminal

```
git clone git@github.com:teguharifudin/Realtime-Chat.git
```
```
cd Realtime-Chat
```

### Server
```
cd server
```
```
go run main.go
```

### Client
```
cd client
```
```
npm run dev
```

## Usage

- Create PostgreSQL database and Table "users"
```
macbookpro=# CREATE TABLE users (
```
```
macbookpro(# id SERIAL PRIMARY KEY,
```
```
macbookpro(# username VARCHAR(255) NOT NULL,
```
```
macbookpro(# email VARCHAR(255) NOT NULL,
```
```
macbookpro(# password VARCHAR(255) NOT NULL
```
```
macbookpro(# );
```

- Create account in Postman

POST http://127.0.0.1:8080/signup
```
{
    "username":"admin",
    "email":"admin@gmail.com",
    "password":"123456"
}
```
```
{
    "username":"test",
    "email":"test@gmail.com",
    "password":"123456"
}
```

- And run the app on browser at http://localhost:3000/

## Contributing

Please use the [issue tracker](https://github.com/teguharifudin/Realtime-Chat/issues) to report any bugs or file feature requests.
