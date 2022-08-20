# Denco + Redis = <3
This is a containerized template for Denco projects using redis caching. 


### Prerequisites: 
- Docker


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
