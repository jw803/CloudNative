### deploy httpserver
```
kubectl create ns module12
kubectl label ns module12 istio-injection=enabled
kubectl create -f  deployment-service*.yaml -n module12
kubectl apply -f nginx.yaml -n module12
```
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
kubectl create -n istio-system secret tls paul-credential --key=paul.io.key --cert=paul.io.crt
kubectl create -f istio-specs.yaml -n module12
```

### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   10.105.227.152
```
### access the httpserver via ingress
```
curl --resolve httpsserver.paul.io:443:10.105.227.152 https://httpsserver.paul.io/healthz -v -k
```
