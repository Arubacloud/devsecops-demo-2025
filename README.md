<p align="center">
  <img src="logo.png" alt="Logo" width="200"/>
</p>

# DevSecOps2025 - Demo

This repository demonstrates a modern DevSecOps workflow for deploying a secure WordPress application and static website using Kubernetes, Docker, GitHub Actions, Vault, and Kustomize.

## Features

- **Static Website**: Built with Docker and Nginx, automated CI/CD via GitHub Actions.
- **WordPress Deployment**: Containerized WordPress with custom plugins, deployed to Kubernetes.
- **Kubernetes Manifests**: Managed with Kustomize for environment-specific overlays.
- **Secrets Management**: Secure integration with HashiCorp Vault and External Secrets Operator.
- **TLS & Ingress**: Automated TLS via cert-manager, flexible ingress configuration.
- **CI/CD Automation**: GitHub Actions for build, test, and deployment pipelines.

## Structure

- `static-site/` — Static website source and Dockerfile.
- `kaas/` — WordPress Docker setup and custom plugins.
- `deploy/` — Kubernetes manifests (apps, base, overlays).
- `setup/` — Guides for Vault, External Secrets Operator, and other tools.
- `.github/workflows/` — CI/CD pipeline definitions.

## Getting Started

1. Build and push images using GitHub Actions.
2. Deploy to Kubernetes using Kustomize overlays.
3. Integrate Vault and External Secrets Operator for secure secrets management.

## Technologies

- Docker, Nginx, WordPress
- Kubernetes, Kustomize, cert-manager
- HashiCorp Vault, External Secrets Operator
- GitHub Actions, ArgoCD

## Cloud Provider

 [ArubaCloud](https://arubacloud.com)

**Services:**
- [CloudServer](https://kb.arubacloud.com/en/computing/cloud-servers.aspx)
- [KaaS](https://kb.arubacloud.com/cmp/en/container/kubernetes.aspx)
- [DbaaS](https://kb.arubacloud.com/en/database/cloud-dbaas.aspx)
- [Virtual Private Network](https://kb.arubacloud.com/cmp/en/networking/vpc-networks.aspx)
- [Subnet](https://kb.arubacloud.com/cmp/en/networking/vpc-networks/subnet.aspx)
- [Security Group](https://kb.arubacloud.com/cmp/en/networking/vpc-networks/security-group.aspx)
- [Elastic Ip](https://kb.arubacloud.com/cmp/en/networking/elastic-ip.aspx)
- [Load Balancer](https://kb.arubacloud.com/cmp/en/networking/load-balancer/description.aspx)
- [BlockStorage](https://kb.arubacloud.com/cmp/en/storage/block-storage.aspx)
- [Container Registry](https://kb.arubacloud.com/en/storage/object-storage.aspx)
- [Object Storage](https://kb.arubacloud.com/en/storage/object-storage.aspx)
