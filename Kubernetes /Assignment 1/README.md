# Assignment 1 - Kubernetes (K8s)

## Assignment Description :

Create K8s cluster - Internal Postgres service + Spring Boot Application connected with Postgres service.

---

## K8s Structure :

1. Database-Secret (for DB credentials)
2. Database-Deployment (postgres)
3. Database-Service (allowing other pods to connect to the database)
3. ConfigMap (for DB URL)
4. Classroom-Deployment (spring boot app - classroom-service)
5. External Service

---

### Images for the assignment

1. Classroom-service image (Docker - Assignment 2 )

        $ docker build -t classroom-service . (building image from Dockerfile)
        $ eval $(minikube docker-env)
        $ minikube image load classroom-service

2. Postgres image 

### Execution

    $ kubectl apply -f database-secrete.yaml

    $ kubectl apply -f database-deployment.yaml

    $ kubectl apply -f database-service.yaml

    $ kubectl get all

---

