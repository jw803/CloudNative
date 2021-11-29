#!/bin/bash

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=skyraker.com/O=skyraker"
kubectl create secret tls hellion-tls --cert=tls.crt --key=tls.key --dry-run=client -oyaml > secret.yaml