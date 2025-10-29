# HashiCorp Vault Installation Guide for Aruba Managed Kubernetes

This guide provides step-by-step instructions to deploy HashiCorp Vault in HA mode on Aruba Managed Kubernetes, using persistent storage and optional audit logging.

## Prerequisites
- Aruba Managed Kubernetes cluster
- A Cinder-compatible StorageClass (e.g., `cinder`, `cinder-csi`)
- Helm installed on your local machine

## Installation Steps

1. **Add the HashiCorp Helm repository:**
   ```bash
   helm repo add hashicorp https://helm.releases.hashicorp.com
   helm repo update
   ```

2. **Install Vault with your custom values:**
   ```bash
    helm upgrade --install vault hashicorp/vault -n vault --create-namespace -f setup/2.\ vault/vault-values.yaml
   ```

3. **Verify the deployment:**
   ```bash
   kubectl get pods -n vault
   kubectl get svc -n vault
   ```

## Post-Install: Initialize & Unseal Vault

After deploying Vault, you must initialize and unseal it. For a quick start, exec into a Vault pod:

```bash
kubectl -n vault exec -it statefulset/vault -c vault -- /bin/sh
```
Inside the pod, run:
```bash
vault operator init -key-shares=1 -key-threshold=1
# Store the unseal key and root token securely!
vault operator unseal <UNSEAL_KEY>
export VAULT_TOKEN=<ROOT_TOKEN>
```

## Enable KV v2 and Write a Sample Secret
```bash
vault secrets enable -path=kv kv-v2
vault kv put kv/app/demo username='demo-user' password='s3cr3tP@ss'
```

## Enable Kubernetes Auth in Vault
```bash
vault auth enable kubernetes
# Configure Vault to talk to the cluster (uses in-cluster service account of the vault pod)
vault write auth/kubernetes/config \
  kubernetes_host="https://${KUBERNETES_PORT_443_TCP_ADDR}:443" \
  token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
  kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```

## Create a Policy for External Secrets Operator (ESO)
Create a least-privilege policy to allow read access to the needed secret paths:
```bash
cat > /tmp/pol-eso-app.hcl <<'HCL'
path "kv/data/app/*" {
  capabilities = ["read"]
}
HCL
vault policy write eso-app /tmp/pol-eso-app.hcl
```

## Create a Role Bound to the ESO Service Account
Bind the policy to the service account ESO will use:
```bash
vault write auth/kubernetes/role/eso-app \
  bound_service_account_names=external-secrets \
  bound_service_account_namespaces=external-secrets \
  policies=eso-app \
  ttl=24h
```

> **Note:** Replace `kv/data/app/*` with your exact secret paths. You can create multiple roles/policies for different apps as needed.

## Notes
- For audit logging, set `auditStorage.enabled: true` and configure the storage class and size.
- The UI is enabled by default and exposed on port 8200.
- For production, ensure you secure Vault and configure access controls.

## References
- [Vault Helm Chart Documentation](https://github.com/hashicorp/vault-helm)
- [Aruba Managed Kubernetes Documentation](https://kb.arubacloud.com/cmp/en/container/kubernetes.aspx)
