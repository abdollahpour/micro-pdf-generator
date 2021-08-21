You can simply run micro-pdf-generator as a severless application on [Knative](knative.dev) and manage your resource more efficiently. For example:

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: micro-pdf-generator
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/minScale: "1"
    spec:
      containers:
        - image: abdollahpour/micro-pdf-generator
```

For more information check [knative serving](https://knative.dev/docs/serving/) documentations.
