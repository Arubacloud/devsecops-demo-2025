# 🚀 Install Argo CD on Kubernetes (kubectl apply + Ingress exposed)

This guide will install **Argo CD** using official YAML manifests and expose the UI via HTTPS Ingress.

---

```bash
kubectl create namespace argocd
```

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Wait for pods to start:
```bash
kubectl -n argocd get pods
```

Enable insecure mode (ARGOCD will serve HTTP - TLS handled by Ingress)
```bash
kubectl -n argocd patch deployment argocd-server \
  --type='json' \
  -p='[{
    "op": "add",
    "path": "/spec/template/spec/containers/0/args/-",
    "value": "--insecure"
  }]'

```

Verify args:
```bash
kubectl -n argocd get deploy argocd-server -o jsonpath="{.spec.template.spec.containers[0].args}"
```

Expose ARGO CD via Ingress
```bash
kubectl apply -f ingress.yaml
```

Register the GithubRepo [Ref](https://argo-cd.readthedocs.io/en/stable/user-guide/private-repositories/)

