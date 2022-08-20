# Denco + Redis = <3
This is a containerized template for Denco projects using redis caching. 


### Prerequisites: 
- Docker


### Usage
```
gh repo clone joshua-auchincloss/denco-redis
cd denco-redis-master
docker compose up
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
