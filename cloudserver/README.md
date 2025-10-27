# DevSecOps Demo 2025

This repository contains a static website and automation to build and deploy it as a Docker image using GitHub Actions and a self-hosted runner.

## Features
- Static website (HTML/CSS/JS) in `static-site/`
- Dockerfile for Nginx-based container
- GitHub Actions workflow for CI/CD
- Pushes Docker image to: `ha-pvt.accbw.itbg.svc8d.registry.services.testops.arubacloud.com`

## Usage

### 1. Static Website
Edit files in `static-site/` to update the website content.

### 2. Build & Deploy
On push to `main`, the GitHub Actions workflow will:
- Build the Docker image
- Push it to the registry

### 3. Self-hosted Runner Requirements
- Runner must have Docker installed
- Runner user must be in the `docker` group for permissions
- Set environment variables before starting the runner:
  ```bash
  export DOCKER_USERNAME='<< USERNAME >>'
  export DOCKER_PASSWORD='<< PASSWORD >>'
  ```

## Registry Login
The workflow logs in to the registry using credentials from runner environment variables.

## Troubleshooting
- If you see `permission denied` errors, ensure the runner user is in the `docker` group.
- If Docker credentials are not set, export them before starting the runner.

## License
MIT
