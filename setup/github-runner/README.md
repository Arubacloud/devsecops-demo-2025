# GitHub Self-Hosted Runner Setup (Aruba Cloud Ubuntu)

This guide explains how to configure an Ubuntu server on Aruba Cloud to register as a self-hosted GitHub Actions runner.

## Prerequisites

- Aruba Cloud Ubuntu server (SSH access)
- GitHub repository with admin access

## Steps

1. **Install Required Packages**
   ```bash
   sudo apt update
   sudo apt install -y curl git
   ```

2. **Create a Runner User (optional, for security)**
   ```bash
   sudo adduser --disabled-password --gecos "" githubrunner
   sudo usermod -aG docker githubrunner
   ```

3. **Download GitHub Runner**
   ```bash
   mkdir -p ~/actions-runner && cd ~/actions-runner
   curl -o actions-runner-linux-x64-2.314.1.tar.gz -L https://github.com/actions/runner/releases/download/v2.314.1/actions-runner-linux-x64-2.314.1.tar.gz
   tar xzf actions-runner-linux-x64-2.314.1.tar.gz
   ```

4. **Register the Runner**
   - Go to your GitHub repo: `Settings > Actions > Runners > New self-hosted runner`
   - Copy the registration command (with your repo and token), e.g.:
     ```bash
     ./config.sh --url https://github.com/<owner>/<repo> --token <TOKEN>
     ```

5. **Start the Runner**
   ```bash
   ./run.sh
   ```

6. **(Optional) Run as a Service**
   ```bash
   sudo ./svc.sh install
   sudo ./svc.sh start
   ```

## Notes

- Ensure the runner user has access to Docker if needed for your workflows.
- Set any required environment variables before starting the runner.

## References

- [GitHub Docs: Adding self-hosted runners](https://docs.github.com/en/actions/hosting-your-own-runners/adding-self-hosted-runners)
