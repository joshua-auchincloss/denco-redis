# Denco + Redis = <3
This is a containerized template for Denco projects which use redis caching behind an nginx proxy.

### Prerequisites: 
- Docker

### Project Structure 
```
{TREE}
```

### [`docker-compose.yaml`](docker-compose.yaml)
```
{COMPOSE-FILE}
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
