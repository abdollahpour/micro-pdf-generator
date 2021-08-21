You can simply run micro-pdf-generator on kubernetes.

Create a file `mpg.yml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: micro-pdf-generator
  labels:
    app: micro-pdf-generator
spec:
  selector:
    matchLabels:
      app: micro-pdf-generator
  template:
    metadata:
      labels:
        app: micro-pdf-generator
    spec:
      containers:
        - name: micro-pdf-generator
          image: abdollahpour/micro-pdf-generator:0.1.1
          ports:
            - name: http
              containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: micro-pdf-generator
spec:
  selector:
    app: micro-pdf-generator
  ports:
    - protocol: TCP
      port: 80
      targetPort: http
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: micro-pdf-generator
  # For SSL you need cert-manager and letsencrypt-prod ClusterIssuer
  # annotations:
  #  cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  # tls:
  #  - hosts:
  #      - your-domain.com
  #    secretName: micro-pdf-generator-cert
  rules:
    - host: your-domain.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: micro-pdf-generator
                port:
                  number: 80
```

Change your `your-domain.com` with your desired domain name. You can also active SSL if you have already [cert-manager](https://cert-manager.io) setup.

And then run `kubectl apply -f mpg.yml`
