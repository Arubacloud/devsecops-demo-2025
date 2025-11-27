# Actalis ACME Issuer for cert-manager  
Use Actalis CA to issue TLS certificates in Kubernetes

This repository provides instructions for installing **cert-manager** and configuring it to use the **Actalis ACME Certificate Authority** through the official Helm chart:

👉 **https://github.com/Arubacloud/helm-charts/tree/main/charts/actalis-cert-manager**

The `actalis-cert-manager` chart automatically deploys an `Issuer` or `ClusterIssuer` configured to talk with Actalis’ ACME endpoint using EAB (External Account Binding).

---

## Prerequisites

- Kubernetes 1.20+
- `kubectl` and `helm` installed
- A DNS domain pointing to your cluster
- cert-manager installed (see below)
- Actalis ACME account with:
  - ACME directory URL
  - EAB **Key ID (kid)**
  - EAB **HMAC Key** (base64url-encoded, **without "=" padding**)

---

## 1. Install cert-manager

Install cert-manager using Helm (recommended):

```bash
kubectl create namespace cert-manager

helm install cert-manager oci://quay.io/jetstack/charts/cert-manager \
  -n cert-manager \
  --set crds.enabled=true
