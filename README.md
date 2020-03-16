### Decrypt the secrets
```
kubesec decrypt -i secret/encrypted-secret.yml -o manifests/secret.yml
```

### apply
```
kubectl apply -f manifests
```
=> listening on [localhost:80](http://localhost)

### How to make encrypted secrets
```
kubesec encrypt --key=gcp:projects/lastrust-stg/locations/global/keyRings/playk8s/cryptoKeys/kubesec --cleartext -i secret/secret.yml -o secret/encrypted-secret.yml
```