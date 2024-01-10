# Deployment

## Secrets

Privaten & öffentlichen Schlüssel erzeugen:

```
openssl ecparam -name prime256v1 -genkey -noout -out private.key

openssl ec -in private.key -pubout -out public.key
```

Erstellen der Secrets für den privaten und öffentlichen Schlüssel:
```
 kubectl create secret generic jwt-private-key --from-file private.key

 kubectl create secret generic jwt-public-key --from-file public.key
```

Erstellen des Secret für das PostgresDB Passwort:
```
kubectl create secret generic postgres-password --from-literal=POSTGRES_PASSWORD=password
```

Verifizieren, ob Secrets erfolgreich angelegt wurden:
```
kubectl get secrets

NAME                TYPE     DATA   AGE
jwt-private-key     Opaque   1      0m1s
jwt-public-key      Opaque   1      0m1s
postgres-password   Opaque   1      0m1s
```

## Deploy

```
kubectl create -f . --recursive
```


