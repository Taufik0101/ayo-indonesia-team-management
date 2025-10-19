# Backend Team Management

## Instructions

### 1. Clone Project

```bash
git clone https://github.com/Taufik0101/ayo-indonesia-team-management.git
```

### 2. Instalation

```bash
go get .

copy .env.example .env / cp .env.example .env

make database in supabase, because i save uploaded logo file to supabase storage

get credential for database connection

get credential for storage connection
- SUPABASE_URL = can use project settings -> data api -> choose URL inside Project URL
- SUPABASE_ACCESS_KEY = can generate from storage -> settings -> new access key, create random name, it will return access and secret
- SUPABASE_SECRET_KEY = can generate from storage -> settings -> new access key, create random name, it will return access and secret
- SUPABASE_REGION = can generate from storage -> settings -> Region

RUN Migration and Seed use wsl or mac can directly use terminal

run make migrate-up database="postgres://username:password@host:port/db_name?sslmode=disable" (change based on your credential)

import postman collection to your local postman

go run main.go