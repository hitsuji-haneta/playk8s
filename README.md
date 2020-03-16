### Decrypt the secrets
```
kubesec decrypt -i manifests/encrypted-secret.yml -o manifests/secret.yml
```

### apply
```
kubectl apply -f manifests
```
=> listening on [localhost:80](http://localhost)