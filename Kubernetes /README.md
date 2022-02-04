# Kubernetes (K8s)

## Kubernetes Components

### Master Node (Kubernetes Control Plane)

Master Node Processes:

* API Server - cluster gateway (entrypoint) + authentication 
* Scheduler - schedules pods onto the nodes (depending on load of nodes)
* Controller Manager - detects changes in the cluster (dying of pods, etc) and makes changes to the cluster
* Etcd - manages the cluster's storage (key-value data store for cluster state i.e. pod status, health, etc)

for multiple master nodes, the API Server is load balanced & ETCD is distributed storage across all masters

---

### Worker Nodes (Kubernetes Compute Nodes)

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
    * Load balancing

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
