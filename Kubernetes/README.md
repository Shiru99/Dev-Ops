# Kubernetes (K8s)

## Kubernetes Components

### 1. Master Node (Kubernetes Control Plane)

Master Node Processes:

* API Server - entrypoint of cluster (for k8s configuration) 
* Scheduler - schedules pods onto the nodes (depending on load of nodes)
* Controller Manager - detects changes in the cluster (dying of pods, etc) and makes changes to the cluster
* Etcd - manages the cluster's storage (key-value data store for cluster state i.e. pod status, health, etc)

for multiple master nodes, the API Server is load balanced & ETCD is distributed storage across all masters

---

### 2. Worker Nodes (Kubernetes Compute Nodes)

Worker Node Processes:

* Kubelet - responsible for starting and stopping pods (Etcd + Controller Manager ==> Scheduler ==> Kubelet)
* Kube-proxy - responsible for routing traffic to the correct nodes
* Container runtime - responsible for running containers

Different Components of Kubernetes:

* Pod:

    * Smallest unit of Kubernetes
    * Abstraction over container
    * usually 1 container per pod
    * Pod gets it's own IP address (IP dies with pod)

* Service:

    * Provides permanent IP addresses for pods 
    * common IP address for all pod replicas
    * Load balancer between pod replicas

* Ingress:

    * Allows access to k8s services from outside the Kubernetes cluster by exposing a service to the internet

* ConfigMap:

    * Stores configuration files.

* Secret:

    * Stores sensitive data.

* volume:

    * K8s doesn't manage data persistence.
    * Volume is a way to share data between pods.

* Deployment:

    * Blueprints for pods OR way to deploy a pod
    * Abstraction over pods

* StatefulSet:

    * Like deployment but with state (persistent data) for database services

---

### MiniKube

```
    $ minikube start --driver=docker (While accessing service from external network getting - connection refused)

    Or 

    $ minikube start  --vm-driver none
```

Alternative - https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/
---

## Kubernetes Commands:

1. Kubernetes Status for Nodes (Master + Worker) / Pods (running containers) / ReplicaSets / Deployments / Services / Ingress /  ConfigMap / Secret / StatefulSet
    
        $ kubectl get nodes / pods / replicaSets / deployments / services / ingress / configmap / secret / statefulset

        $ kubectl get pods -o wide (wide format)

        $ kubectl get deployments <deployment-name> -o yaml (ETCD stores deployment info in YAML format)

        $ kubectl get pods <pod-name> -o yaml (details in YAML format)

2. Deploy new deployment
        
        $ kubectl create deployment <pod_name> --image=<image_name>
    <!-- kubectl create deployment nginx-depl --image=nginx:latest -->
    <!-- $ eval $(minikube docker-env unset) / $ minikube docker-env unset -->
        
    From local docker image 

        $ eval $(minikube docker-env)
        $ minikube image load <image-name>  
        $ kubectl run <pod-name> --image=<image-name>  --image-pull-policy=IfNotPresent

3. Edit deployment
    
        $ kubectl edit deployment <deployment_name>

4.  Create/Update deployment using file
    
        $ kubectl apply -f <file_name>

5. Delete deployment using file
    
        $ kubectl delete -f <file_name>

6. Info about a pod / deployment / service / ingress / configmap / secret / statefulset
    
        $ kubectl describe pod <pod_name>

7. Logs for a pod
    
        $ kubectl logs <pod_name>

8. Access to a pod's bash shell
    
        $ kubectl exec -it <pod_name> -- /bin/bash Or -- /bin/sh

9. Delete a pod / deployment / service / ingress / configmap / secret / statefulset
    
        $ kubectl delete pod <pod_name> 

    (until one deletes deployment, pods get recreated automatically)
    
---

## Useful K8s Aliases:

    alias k='kubectl '
    alias kg='kubectl get '
    alias kd='kubectl describe '
    alias kr='kubectl run '

    alias kci='kubectl cluster-info'
    alias krm='kubectl delete '
    alias kex='kubectl exec -i -t '
    alias klo='kubectl logs '

    alias kga='kubectl get all '
    alias kgp='kubectl get pods '
    alias kdp='kubectl describe pod '
    alias kcd='kubectl create deployment '

    alias kaf='kubectl apply -f '
    alias kdf='kubectl delete -f '
    alias krf='kubectl replace -f '

---

## Namespaces
    
namespace is a way to organize your resources by creating virtual clusters inside a cluster

1. default Namespaces :

    * kube-system - system processes like master & kubectl processes (no edit access)
    * kube-public - publicly accessible data (e.g. ConfigMap of cluster-info, etc)
    * kube-node-lease - stores heartbeats (availability of nodes)
    * default - default namespace

            $ kubectl get namespaces
            $ kubectl create namespace <namespace_name>
            $ kubectl delete namespace <namespace_name>
            $ kubectl create -f <file_name> --namespace <namespace_name> 
            OR set in configuration file - metadata.namespace = <namespace_name>

2. Resources like ConfigMaps & Secrets can't be shared across namespaces but services can.
3. Node & volumes are not bound to any namespaces.

        $ kubectl api-resources --namespaced=false

4. Why to use namespaces?

    * Resources grouped in namespaces e.g. database,web-app, monitoring, elastic stack, Nginx-Ingress controller, etc
    * Resources sharing (e.g. database, elastic stack, shared between staging & production OR different production versions)
    * Limiting resources(CPU, RAM, storage) or Limiting access to resources (e.g. only admin can edit database, only admin can edit elastic stack, etc)

5. To switch the namespace

        $ brew install kubectx / kubens
        $ kubens <namespace_name>

---
