# API-GOLANG-EXAMPLE

this api was develop to study.

### RUN LOCAL

```bash

# create folder persistent data to postgresql
mkdir postgresql

# start application
docker-compose up --build
```

### Test API Response

```bash

curl --location --request GET 'localhost:8080'
```

### ADD USER

```bash

curl --location --request POST 'localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Romaryo",
    "email": "r@gmail.com"
}'
```

### LIST USERS

```bash

curl --location --request GET 'localhost:8080/users'
```