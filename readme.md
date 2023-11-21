docker run --name mrm-postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=mydatabase -p 0.0.0.0:5432:5432 -d postgres

https://raw.githubusercontent.com/dkzhang/mrm/master/docker-compose.yml