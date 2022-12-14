# Denco + Redis = <3
This is a containerized template for Denco projects which use redis caching behind an nginx proxy.

### Prerequisites: 
- Docker

### Project Structure 
```
├── docker-compose.yaml
├── README.md
├── README-master.md
├── syncreadme.py
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
└── goserver
    ├── go.mod
    ├── server.go
    ├── Dockerfile
    ├── go.sum
    └── README.md
```

### [`docker-compose.yaml`](docker-compose.yaml)
```
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

#### Expected Result: 
```
CONTAINER ID   IMAGE                        COMMAND                  CREATED         STATUS                   PORTS                NAMES
82dc15b9e9b4   denco-redis_nginx            "/docker-entrypoint.…"   2 minutes ago   Up 2 minutes             0.0.0.0:80->80/tcp   proxy
fe1ba0d07433   denco-redis:latest           "app"                    2 minutes ago   Up 2 minutes (healthy)                        denco
ab1615279d9e   redis:latest                 "docker-entrypoint.s…"   3 minutes ago   Up 2 minutes (healthy)   6379/tcp             redis
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

### Build
**warning**: do not ship this for production without adding your preferred authentication methods. there is no additional security provided by this template.

#### Docker:
```
docker compose build
```
#### Kompose (kubernetes):
```
kompose convert -f docker-compose.yaml
```
#### Kompose (openshift):
```
kompose convert -f docker-compose.yaml --provider=openshift
```

### Additional
Maintain synchronized documentation using "README-master". Run the following within the main directory to track changes to your project: 
```
python syncreadme.py
```
