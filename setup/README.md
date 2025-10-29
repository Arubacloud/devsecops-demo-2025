# How to setup


## 1. Install ArgoCD

[here](argocd/)

---
## 2. Install Hashicorp Vault

[here](vault/)

---
## 3. Install External Secret Operator

[here](external-secret/)

---
## 4. Deploy ArgoCD Root Application

Create the ARGO CD Application
```bash
kubectl apply -f setup/argocd/root.yaml
```
