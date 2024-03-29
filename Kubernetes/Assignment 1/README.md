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
    
    OR

6. classroom-Service as internal service + Ingress (Ingress + Ingress Controller)
(In general service provider provides Cloud Load Balancer which redirects traffic to the ingress controller)


---

### Images for the assignment

1. Classroom-service image (Docker - Assignment 2 )

        $ docker build -t classroom-service . (building image from Dockerfile)
        $ eval $(minikube docker-env)
        $ minikube image load classroom-service

2. Postgres image 

        $ minikube image load postgres (if local image exits)

---

### Commands for the assignment

    $ kubectl delete -f classroom-config.yaml && kubectl delete -f database-secrete.yaml && kubectl delete -f database-deployment.yaml && kubectl delete -f database-service.yaml && kubectl delete -f classroom-deployment.yaml && kubectl delete -f classroom-service.yaml && kubectl delete -f classroom-ingress.yaml

    $ kubectl apply -f classroom-config.yaml && kubectl apply -f database-secrete.yaml && kubectl apply -f database-deployment.yaml && kubectl apply -f database-service.yaml && kubectl apply -f classroom-deployment.yaml
    
* M-1 : Using external service

        $ kubectl apply -f classroom-service.yaml

        $ minikube service classroom-service (To assign external IP address to the service)

    * To Verify : [ERROR - connection refused for docker driver]

            $ curl -Is http://172.17.0.12:32123/api/classroom/ (External Service)

* M-2 : Using Ingress :

        $ minikube addons enable ingress (for ingress controller)

        $ kubectl get pods -n ingress-nginx Or kubectl get pods -n kube-system (To check if ingress controller is running)

        $ kubectl apply -f classroom-service.yaml

        $ kubectl apply -f classroom-ingress.yaml

    * Add ingress Address with hostname to /etc/hosts

    * To Verify :

            $ curl -Is http://classroom-example.com/api/classroom/ (Ingress + Service)

---

