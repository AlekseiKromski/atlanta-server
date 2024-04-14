# Description ✏️

The Atlanta server is a piece of dedicated SCV (Supply chain visibility) software that will receive/analyze/distribute data from the IoT applicance

Project development for graduate work at the [Narva Collage of the University of Tartu](https://narva.ut.ee/en)

# Setup 👨‍💻

## Setup `.env` file 📝
As the first step you have to configure `.env` file. Example of this file is `.env.example`

```dotenv
GIN_ADDRESS=:3000
GIN_SECRET=secret
GIN_COOKIE_DOMAIN=localhost
TCP_CONSUMER_ADDRESS=:3017
TCP_CONSUMER_BUF=256
DB_DATABASE=atlanta
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
```
Recommended do not change this variables
`GIN_ADDRESS` & `TCP_CONSUMER_ADDRESS` & `TCP_CONSUMER_BUF`

## Run 🚀

To start the application you need to: 

Build application with docker
```shell
docker compose build --no-cache
```

To run application
```shell
docker compose up --force-recreate -d
```

Application will automatically run db migrations if needed with default roles and user.

## Admin access to application ✅

Default user & password for admin is:
```dotenv
USER: admin
PASSWORD: admin
```

# Documentation 🗒️

All documentation is available here: [`./front-end/public/docs/*`](./front-end/public/docs)

If you have any problem, you can use `help` button on each block with description in application

Recommendation to read:
- [Search](./front-end/public/docs/search.md)
- [Access management](./front-end/public/docs/access_management.md)

