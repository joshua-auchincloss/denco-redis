# Denco + Redis = <3
This is a containerized template for Denco projects using redis caching behing an nginx proxy.


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
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
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
        - DENCO_PORT=80
    networks:
      - caching
    depends_on:
      - redis
    environment:
      - REDIS_HOST=redis
      - REDIS_PORT=6379
      - DENCO_PORT=8080

    healthcheck:
      test: ["CMD", "curl", "-fsSL", "http://localhost:8080"]
      interval: 30s
      timeout: 5s
      retries: 5

  nginx:
    container_name: proxy
    build:
      context: ./nginx
      args:
        - NGINX_PORT=80
    restart: always
    networks:
      - caching
    depends_on:
      - denco
      - redis
    ports:
      - 80:80
    healthcheck:
      test: ["CMD", "curl", "-fsSL", "http://localhost/"]
      interval: 30s
      timeout: 5s
      retries: 5

  redis:
    image: redis:latest
    container_name: redis
    restart: always
    command: redis-server --save 20 1 --loglevel warning
    volumes:
      - cache:/data
    networks:
      - caching
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 30s
      timeout: 5s
      retries: 5

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
CONTAINER ID   IMAGE                        COMMAND                  CREATED          STATUS                            PORTS                NAMES
48c86288207a   denco-redis_nginx            "/docker-entrypoint.…"   7 seconds ago    Up 2 seconds (health: starting)   0.0.0.0:80->80/tcp   proxy
22bb88c2f124   j000000/denco-redis:latest   "app"                    8 seconds ago    Up 4 seconds (health: starting)                        denco
a95710152acf   redis:latest                 "docker-entrypoint.s…"   34 minutes ago   Up 5 seconds (health: starting)   6379/tcp             redis
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
