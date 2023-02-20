# API Boiletplate

Setup an environment with the file `server.env` containing following the same structure of `server.env.example`

#### Local Postgres

I will be creating a user "postgres" and a database "mococa_api_v2"

Login to database with user `postgres`

```bash
psql -U postgres
```

Create a database called `mococa_api_v2`

```bash
CREATE DATABASE mococa_api_v2;
```

Select the database you just created

```bash
\c mococa_api_v2;
```

Add the UUID extension

```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Exit psql

```bash
\q
```

Should be good enough to set it up properly
