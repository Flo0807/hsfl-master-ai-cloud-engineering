# Deployment

## Prerequisites

Make sure secrets are created before deploying the application.

Create private and public key for JWT authentication:
```
openssl ecparam -name prime256v1 -genkey -noout -out private.key

openssl ec -in private.key -pubout -out public.key
```

Create the secrets for the private and public key:
```
 kubectl create secret generic jwt-private-key --from-file private.key

 kubectl create secret generic jwt-public-key --from-file public.key
```

Create the secret for the postgres password:
```
kubectl create secret generic postgres-password --from-literal=POSTGRES_PASSWORD=password
```

Verify that all secrets are created:
```
kubectl get secrets

NAME                TYPE     DATA   AGE
jwt-private-key     Opaque   1      0m1s
jwt-public-key      Opaque   1      0m1s
postgres-password   Opaque   1      0m1s
```

## Deploy on Minikube

Make sure [Minikube](https://minikube.sigs.k8s.io/docs/start/) is properly installed and all secrets are created.

Run `minikube start` to start the cluster.

Enable ingress addon:

```
minikube addons enable ingress
```

Deploy the application:

```
kubectl apply -f k8s/manifests --recursive
```

Only in docker desktop: Run `minikube tunnel` to expose the ingress controller to the host. (see [Ingres documentation](https://minikube.sigs.k8s.io/docs/start/))

You can now access the application at `http:localhost:80`

## Deploy on other Kubernetes Cluster

Make sure you have a working Kubernetes cluster with an ingress controller running. See https://spacelift.io/blog/kubernetes-ingress on how to set up an ingress controller. In addition, make sure all secrets are created.

Deploy the application:

```
kubectl apply -f k8s/manifests --recursive
```

Run `kubectl get ingress` to get the ingress IP address.

Example:

```
❯ kubectl get ingress
NAME            CLASS   HOSTS   ADDRESS        PORTS   AGE
proxy-ingress   nginx   *       192.168.49.2   80      13m
```

Access the application at `http://<ingress-ip>:80`.

## Monitoring

### Access grafana

Use the following command to get the grafana port:

```
❯ kubectl get service/grafana --namespace=monitoring
NAME      TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
grafana   NodePort   10.110.91.191   <none>        3000:32702/TCP   17h
```

When using minikube, run `minikube service grafana --namespace=monitoring` to open grafana in the browser.

Default grafana credentials are `admin` and `admin`.
