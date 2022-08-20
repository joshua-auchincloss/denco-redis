# Denco + Redis = <3
This is a containerized template for Denco projects using redis caching. 


### Prerequisites: 
- Docker

### Project Structure 
```
.
├── goserver
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── server.go
├── docker-compose.yaml
└── README.md
```

### [`docker-compose.yaml`](docker-compose.yaml)
```
version: "3.8"

services:
  denco: 
    container_name: denco
    image: j000000/denco-redis:latest
    build: 
      context: ./goserver
      args:
        - DENCOPORT=8080
    networks:
      - caching
    ports: 
      - 8080:8080
    depends_on:
      - redis
    environment:
      - REDIS_HOST=redis
      - REDIS_PORT=6379
      - DENCOPORT=8080

  redis: 
    image: redis:latest
    container_name: redis
    restart: always 
    command: redis-server --save 20 1 --loglevel warning
    volumes: 
      - cache:/data
    expose: 
      - 6379:6379
    networks: 
      - caching

volumes:
  cache:
    driver: local

networks:
  caching:
```

### Usage
```
gh repo clone joshua-auchincloss/denco-redis
cd denco-redis-master
docker compose up -d --build
```

Expected Result: 
```
> docker ps
CONTAINER ID   IMAGE                        COMMAND                  CREATED          STATUS          PORTS                    NAMES
17372fe978af   j000000/denco-redis:latest   "app"                    13 seconds ago   Up 11 seconds   0.0.0.0:8080->8080/tcp   denco
2b18c654bef4   redis:latest                 "docker-entrypoint.s…"   13 seconds ago   Up 12 seconds   0/tcp, 6379/tcp          redis
```


### Local Development
1. Clone the repo & cd into the main directory (see [usage](#usage))
2. Start redis within docker: 
```
docker pull redis:latest
docker run --name redisdev -d redis --p 6379:6379
```
3. cd into the server project: 
```
cd goserver
```
4. Run the server: 
```
go run .
```
