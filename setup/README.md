# How to setup

## 0. Register a Self Hosted Github Runner

[here](github-runner/)

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
## 5. Have a look

* [Wordpress Website](https://devsecops2025-arubacloud.com)
* [Wordpress Backend](https://devsecops2025-arubacloud.com/wp-admin)
* [ArgoCD](https://argocd.devsecops2025-arubacloud.com)
* [Vault](https://vault.devsecops2025-arubacloud.com)