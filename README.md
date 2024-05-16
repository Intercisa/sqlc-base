DB: 
    running on my raspberryPi4: 
        docker run -d --name postgresql -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -v $(pwd)/data:/var/lib/postgresql/data --restart always postgres

SQL-GEN: 
    docker pull sqlc/sqlc
    docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

BEFORE RUN: 
    go get github.com/jackc/pgx/v5  