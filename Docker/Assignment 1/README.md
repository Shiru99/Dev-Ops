# Docker Container

## Creating a containerized spring boot application :

Simple Hello World spring boot application


## Application Structure

* Service I

    $ java -jar hello-docker-world-0.0.1-SNAPSHOT.jar

---

### To Build Image :

    $ docker build -f Dockerfile -t hello-docker-world .
    
### Run the image (container) :

    $ docker run --name hello-world-service -it -d -p 12345:12345 hello-docker-world 

    http://localhost:12345/api/hello-world/

### Stop the container :

    $ docker stop <container_id>