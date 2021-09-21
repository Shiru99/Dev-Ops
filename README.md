# Dev-Ops

## Docker Commands:

<!-- | Command Name                                          | Command                                           |      
| :---------------------------------------------------: | :-----------------------------------------------: |
|  Docker Version                                       | $ docker --version                                | 
|  Docker Pull Image                                    | $ docker pull <image_name>                        | 
|  List  Docker Images                                  | $ docker images                                   |
|  Run the Docker Image (pull+run)                      | $ docker run <image_name>                         |
|  Lists all the docker containers (running)            | $ docker ps                                       |
|  List all the docker containers                       | $ docker ps -a                                    |
|  Start the docker container                           | $ docker start <container_id>                     |
|  Stop the docker container                            | $ docker stop <container_id>                      |
|  Restart the docker container                         | $ docker restart <container_id>                   |
|  Stop the docker container immediately                | $ docker kill <container_id>                      |
|  Remove the docker container                          | $ docker rm <container_id>                        | 
|  Remove all the docker containers                     | $ docker rm $(docker ps -a -q)                    | 
|  Remove the docker image                              | $ docker rmi <image_id>                           | -->



1. Docker Version
```
$ docker --version
```

2. Docker Pull Image
```
$ docker pull <image_name>
```

3. Docker Images
```
$ docker images
```

4. Run the Docker Image (pull+run)
```
$ docker run <image_name>
```

5. Lists all the docker containers (running)
```
$ docker ps
```

6. List all the docker containers
```
$ docker ps -a
```

7. Start the docker container
```
$ docker start <container_id>
```

8. Stop the docker container
```
$ docker stop <container_id>
```

9. Restart the docker container
```
$ docker restart <container_id>
```

10. Stop the docker container immediately
```
$ docker kill <container_id>
``` 

11. Remove the docker container
``` 
$ docker rm <container_id>
```

12. Remove all the docker containers
```
$ docker rm $(docker ps -a -q)
```

13. Remove the docker image
``` 
$ docker rmi <image_id>
```

14. Access the docker container and run commands inside the container
```
$ docker exec -it <container_id> /bin/bash
```

15. Remove all the docker images
```
$ docker rmi $(docker images -q)
```

16. Save a new docker image (on the local system)
```
$ docker commit <container_id> <image_name>
```

17. Login into docker hub
```
$ docker login
``` 

18. Upload a docker image (on DockerHub)
```
$ docker push <image_name>
```

19. Get detailed information about docker installed
```
$ docker info
```

20. Copy a file from a docker container to the local system
```
$ docker cp <container_id>:/path/to/file /path/to/file
```

21. history of a docker image
```
$ docker history <image_name>
```

22. logs of the docker container
```
$ docker logs <container_id>
```

23. Search for a docker image on DockerHub
```
$ docker search <image_name>
```

24. Update container configurations
```
$ docker update --help
$ docker update <container_id>
```

25. Build an image form a Docker file
```
$ docker build -t <image_name> .
```

26. Exports a container’s filesystem as a tar archive
```
$ docker export <container_id> > <image_name>.tar
```

27. Attaches to a running container
```
$ docker attach <container_id>
```

25. lists the details of all the network
```
$ docker network ls
```

26. Install a docker plugin
```
$ docker plugin install <plugin_name>
```

27. Create a volume which docker container will use to store data
```
$ docker volume create <volume_name>

$ docker volume ls
```

28. Logging out from DockerHub
```
$ docker logout
```