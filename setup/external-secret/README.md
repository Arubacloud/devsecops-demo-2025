# External Secrets Operator Installation Guide for Aruba Managed Kubernetes

This guide explains how to install the External Secrets Operator on Aruba Managed Kubernetes using Helm.

## Prerequisites
- Aruba Managed Kubernetes cluster
- Helm installed on your local machine

## Installation Steps

1. **Add the External Secrets Helm repository:**
   ```bash
   helm repo add external-secrets https://charts.external-secrets.io
   helm repo update
   ```

2. **Install the External Secrets Operator:**
   ```bash
   helm upgrade --install external-secrets external-secrets/external-secrets \
     -n external-secrets \
     --create-namespace \
     --set installCRDs=true
   ```
   - The `--create-namespace` flag ensures the namespace exists.
   - The `--set installCRDs=true` flag installs the required Custom Resource Definitions.

3. **Verify the installation:**
   ```bash
   kubectl get pods -n external-secrets
   kubectl get crds | grep external-secrets
   ```

4. **Create a Service Account:**
   ```bash
    kubectl apply -f serviceaccount.yaml
   ```

## Configure Vault

Open a shell **inside the Vault pod** or local machine with Vault CLI configured:

```bash
kubectl -n vault exec -it vault-0 -- sh
export VAULT_ADDR="http://127.0.0.1:8200"
export VAULT_TOKEN="<root-or-admin-token>"
```

Enable Kubernetes auth:
```bash
vault auth enable -path=kubernetes kubernetes
```

Configure auth to talk to Kubernetes API:
```bash
vault write auth/kubernetes/config \
  kubernetes_host="https://kubernetes.default.svc" \
  kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt \
  token_reviewer_jwt=@/var/run/secrets/kubernetes.io/serviceaccount/token
```
Create a read policy (KV v2 support):
```bash
# kv-read.hcl
path "kv/data/*" {
  capabilities = ["read"]
}
```
apply
```bash
vault policy write kv-read kv-read.hcl
```

Create Vault role for ESO:
```bash
vault write auth/kubernetes/role/eso-app \
  bound_service_account_names="external-secrets" \
  bound_service_account_namespaces="external-secrets" \
  token_policies="default,kv-read" \
  token_ttl="1h"
```

## Deploy

```bash
kubectl apply -f rbac.yaml
```
Restart ESO to apply permissions:
```bash
kubectl -n external-secrets rollout restart deploy/external-secrets
```

Create a ClusterSecretStore:
```bash
kubectl deploy -f secrets-store.yaml
```

Verify readiness:
```bash
kubectl get clustersecretstore vault -o yaml | yq '.status.conditions'
# ✅ Should show Ready=True
```

Store a Secret in Vault
```bash
vault kv put kv/app/wordpress \
  WORDPRESS_DB_USER="wp" \
  WORDPRESS_DB_PASSWORD="super-secret" \
  WORDPRESS_DB_HOST="mariadb.default.svc:3306"
```
Then Create the External Secret Resource to sync Secrets from Hashicorp Vault ...

Example:
```yaml
apiVersion: external-secrets.io/v1
kind: ExternalSecret
metadata:
  name: wordpress-db-secret
  namespace: default
spec:
  refreshInterval: 1h
  secretStoreRef:
    kind: ClusterSecretStore
    name: vault
  target:
    name: wordpress-db-secret
    creationPolicy: Owner
  data:
    - secretKey: WORDPRESS_DB_USER
      remoteRef:
        key: app/wordpress
        property: WORDPRESS_DB_USER
    - secretKey: WORDPRESS_DB_PASSWORD
      remoteRef:
        key: app/wordpress
        property: WORDPRESS_DB_PASSWORD
    - secretKey: WORDPRESS_DB_HOST
      remoteRef:
        key: app/wordpress
        property: WORDPRESS_DB_HOST
```

## Notes
- For production, review the [External Secrets Operator documentation](https://external-secrets.io/) for advanced configuration and integration with secret backends (e.g., Vault, AWS Secrets Manager).
- Ensure your cluster has network access to your chosen secret backend.

## References
- [External Secrets Operator Documentation](https://external-secrets.io/)
- [Aruba Managed Kubernetes Documentation](https://kb.arubacloud.com/cmp/en/container/kubernetes.aspx)
