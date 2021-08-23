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

In more complex scenarios, for example, when you want to trigger PDF generating using [Knative eventing](https://knative.dev/docs/eventing/) and store it somewhere, you can use micro-pdf-generator as a sidecar. For example:

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: pdf-renderer
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/minScale: "1"
        networking.knative.dev/visibility: "cluster-local"
        autoscaling.knative.dev/target: "4"
    spec:
      containers:
        - image: tooltime/pdf-renderer
          ports:
            - containerPort: 8080
        - image: abdollahpour/micro-pdf-generator:v0.1.1
          env:
            - name: MPG_PORT
              value: "7070"
```
