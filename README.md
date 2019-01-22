# boilerplate-gorilla-mux-for-rest-api-with-jwt [![License: MIT][license-image]][license-url]

Boilerplate of Gorilla Mux for REST API with JWT

## Prerequisite

- Git installed
- Docker installed

## Let's begin developing

```bash
git clone https://github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt.git
cd boilerplate-gorilla-mux-for-rest-api-with-jwt
docker-compose up
```

## API endpoints

|  Method  |       URI        |             Data               |
| -------- | ---------------- | ------------------------------ |
|   POST   |  /signup         | email,password,name,age,gender |
|   POST   |  /login          | email,password                 |
|   POST   |  /api/users      | email,password,name,age,gender |
|   GET    |  /api/users      |               -                |
|   GET    |  /api/users/{id} |               -                |
|   PUT    |  /api/users/{id} | email,password,name,age,gender |
|  DELETE  |  /api/users/{id} |               -                |

[license-url]: https://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/badge/License-MIT-yellow.svg
