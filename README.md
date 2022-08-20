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
1. Clone the repo & cd into the main directory
2. Start redis within docker: 
```
docker pull redis:latest
docker run --name redisdev -d redis --p 6379:6379
```
3. Run go server for testing: 
```
go run .
```
