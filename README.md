# k8s-playground

Kubernetes learning

## Allow CORS

```yaml
  # on the backend application
  filters:
    - type: ResponseHeaderModifier
      responseHeaderModifier:
        add:
          - name: Access-Control-Allow-Origin
            value: "http://localhost:8000"
          - name: Access-Control-Allow-Methods
            value: "GET, POST, PUT, DELETE, OPTIONS"
          - name: Access-Control-Allow-Headers
            value: "Content-Type, Authorization"
          - name: Access-Control-Max-Age
            value: "86400"
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
          ...
```

## Headlamp UI

```
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace kube-system -l "app.kubernetes.io/name=headlamp,app.kubernetes.io/instance=my-headlamp" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace kube-system $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace kube-system port-forward $POD_NAME 8080:$CONTAINER_PORT
2. Get the token using
  kubectl create token my-headlamp --namespace kube-system
```