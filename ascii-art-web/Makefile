# Variables
IMAGE_NAME=ascii-art-web
CONTAINER_NAME=dockerize
PORT=8080

# Build Docker image
build:
	docker image build -t $(IMAGE_NAME) .

# Run container in detached mode
run:
	docker container run -d -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

# Enter the running container shell
exec:
	docker exec -it $(CONTAINER_NAME) sh 

# Stop and remove container
stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

# Restart a previously stopped container
start:
	docker start $(CONTAINER_NAME)  

# Rebuild everything
rebuild: stop build run

# View logs
logs:
	docker logs -f $(CONTAINER_NAME)

# Remove unused objects
clean:
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

# Run in foreground
run-fg:
	docker container run -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

# Show running containers and images
status:
	docker ps -a
	docker images

# Open browser to localhost (cross-platform)
open:
ifeq ($(OS),Windows_NT)
	@powershell.exe start http://localhost:8080
else
	xdg-open http://localhost:$(PORT) || echo "Open browser manually: http://localhost:$(PORT)"
endif

# Rebuild, run, and open browser
serve: rebuild open
