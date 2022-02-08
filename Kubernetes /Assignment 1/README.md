# Assignment 1 - Kubernetes (K8s)

## Assignment Description :

Create K8s cluster - Internal Postgres service + Spring Boot Application connected with Postgres service.

---

## K8s Structure :

1. Database-Secret (for DB credentials)
2. Database-Deployment (postgres)
3. Database-Service (Internal Service - allowing other pods to connect to the database)
4. Classroom-Config (ConfigMap for DB URL)
5. Classroom-Deployment (spring boot app)
6. Classroom-Service (External Service - for accessing the spring boot app)

or 

6. Internal Service + Ingress

---

### Images for the assignment

1. Classroom-service image (Docker - Assignment 2 )

        $ docker build -t classroom-service . (building image from Dockerfile)
        $ eval $(minikube docker-env)
        $ minikube image load classroom-service

2. Postgres image 

### Execution

    $ kubectl delete -f database-secrete.yaml && kubectl delete -f database-deployment.yaml && kubectl delete -f database-service.yaml && kubectl delete -f classroom-config.yaml && kubectl delete -f classroom-deployment.yaml && kubectl delete -f classroom-service.yaml

    $ kubectl apply -f database-secrete.yaml && kubectl apply -f database-deployment.yaml && kubectl apply -f database-service.yaml && kubectl apply -f classroom-config.yaml && kubectl apply -f classroom-deployment.yaml
    
    $ kubectl apply -f classroom-service.yaml

    $ minikube service classroom-service (To assign external IP address to the service) - http://192.168.49.2:32123 [ERROR - connection refused for docker driver]

    $ kubectl get all

---

