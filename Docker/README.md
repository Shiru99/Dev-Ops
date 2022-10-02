# Dev-Ops

## Docker Commands:

| Command Name                                          | Command                                           |      
| :--------------------------------------------------- | :----------------------------------------------- |
|  Docker Version                                       | `$ docker --version`                                | 
|  Docker Info                                          | `$ docker info`                                     |
|  List  Docker Images                                  | `$ docker images`                                   |
|  Run the Docker Image (pull+run)                      | `$ docker run <image_name>` <br> `$ docker run -it -d -e <env_var>=<value> -p <port_num_machine>:<port_num_container> --name <container_name> <image_name>` <br> (it : Interactive mode, d : Detached mode, -e : Env variable, -p : Port mapping) |
|  Lists all the docker containers (running)            | `$ docker ps`                                      |
|  List all the docker containers                       | `$ docker ps -a`                                   |
| Search for an image in Docker Hub                     | `$ docker search <image_name>`                      |
|  Start the docker container                           | `$ docker start <container_id>`                     |
|  Stop the docker container                            | `$ docker stop <container_id>`                      |
|  Restart the docker container                         | `$ docker restart <container_id>`                   |
|  Stop the docker container immediately                | `$ docker kill <container_id>`                      |
|  Remove the docker container                          | `$ docker rm <container_id>`                        | 
|  Remove all the docker containers                     | `$ docker rm $(docker ps -a -q)`                    | 
|  Remove the docker image                              | `$ docker rmi <image_id>`                           |
|  Remove all the docker images                         | `$ docker rmi $(docker images -q)`                  |
| Access the docker container's bash | `$ docker exec -it <container_id> /bin/bash` |
|  Access the docker container's bash as root           | `$ docker exec -it -u root <container_id> /bin/bash` |
| Save a new docker image (on the local system) | `$ docker commit <container_id> <image_name>` |
| Login into docker hub | `$ docker login` |
| Push the docker image to docker hub | `$ docker push <image_name>` |
| Pull the docker image from docker hub | `$ docker pull <image_name>` |
| List all the docker networks | `$ docker network ls` |
| Create a new docker network | `$ docker network create <network_name>` |
| Connect a docker container to a network | `$ docker network connect <network_name> <container_id>` |
| Disconnect a docker container from a network | `$ docker network disconnect <network_name> <container_id>` |
| Remove a docker network | `$ docker network rm <network_name>` |
| List all the docker volumes | `$ docker volume ls` |
| Create a new docker volume | `$ docker volume create <volume_name>` |
| Connect a docker container to a volume | `$ docker volume connect <volume_name> <container_id>` |
| Disconnect a docker container from a volume | `$ docker volume disconnect <volume_name> <container_id>` |
| Remove a docker volume | `$ docker volume rm <volume_name>` |
| logs of the docker container | `$ docker logs <container_id>` |
| history of the docker image | `$ docker history <image_name>` |
| inspect the docker container | `$ docker inspect <container_id>` |
| copy files from the docker container to the local system | `$ docker cp <container_id>:<path_in_container> <path_in_local_system>` |
| copy files from the local system to the docker container | `$ docker cp <path_in_local_system> <container_id>:<path_in_container>` |
| export the docker container as a tar file | `$ docker export <container_id> > <tar_file_name>` |
| Build the docker image from the Dockerfile | `$ docker build -t <image_name> .` |
| logout from docker hub | `$ docker logout` |


## Docker Compose Commands:

| Command Name                                          | Command                                           |
| :--------------------------------------------------- | :----------------------------------------------- |
|  Docker Compose Version                               | `$ docker-compose --version`                       |
|  Docker Compose Build                                  | `$ docker-compose build`                           |
|  Docker Compose Up                                    | `$ docker-compose up`                              |
|  Docker Compose Up (detached mode)                    | `$ docker-compose up -d`                           |
|  Docker Compose Down                                  | `$ docker-compose down`                            |
|  Docker Compose Stop                                  | `$ docker-compose stop`                            |
|  Docker Compose Start                                 | `$ docker-compose start`                           |
|  Docker Compose Restart                               | `$ docker-compose restart`                         |
|  Docker Compose Logs                                  | `$ docker-compose logs`                            |
