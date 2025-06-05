# Sharingan DevOps

This directory contains configurations and tools for deploying the Sharingan application monitoring system to AWS using Kubernetes and Helm.

## Contents

- `.env.example` - Example environment variables for AWS and Kubernetes configuration
- Kubernetes manifests for application deployment
- Helm charts for simplified deployment
- Terraform configurations for AWS infrastructure
- Deployment scripts and utilities

## Infrastructure Overview

Sharingan is deployed on AWS using the following services:

- **EKS (Elastic Kubernetes Service)**: For container orchestration
- **ECR (Elastic Container Registry)**: For storing container images
- **RDS (Relational Database Service)**: For PostgreSQL databases
- **CloudWatch**: For monitoring and logging
- **S3**: For static asset storage and backups
- **IAM**: For AWS resource access control

## Kubernetes Configuration

The Kubernetes configuration includes:

- Deployment manifests for all services
- Service definitions for internal communication
- Ingress configurations for external access
- ConfigMaps and Secrets for application configuration
- Volume claims for persistent storage

## Helm Charts

Helm charts provide templated Kubernetes resources for easier deployment:

```bash
# Deploy the application using Helm
helm install sharingan ./devops/helm/sharingan -f values-<environment>.yaml
```

## AWS Setup

### Prerequisites

- AWS CLI configured with appropriate permissions
- `kubectl` for Kubernetes cluster management
- Helm for chart deployment

### Deploying to AWS

1. Set up the AWS infrastructure using Terraform:

```bash
cd devops/terraform
terraform init
terraform apply -var-file=environments/<env>.tfvars
```

2. Configure kubectl to interact with the EKS cluster:

```bash
aws eks update-kubeconfig --name sharingan-<environment> --region <aws-region>
```

3. Deploy the application using Helm:

```bash
helm install sharingan ./devops/helm/sharingan -f values-<environment>.yaml
```

## Environment Configuration

Copy the example environment file to create your configuration:

```bash
cp .env.example .env
```

Update the necessary values for your AWS environment.

## Security Considerations

- AWS IAM roles with least-privilege permissions
- Kubernetes RBAC for cluster access control
- Network policies to restrict pod-to-pod communication
- Secrets management using AWS Secrets Manager or Kubernetes Secrets
- TLS termination at the load balancer level
- Regular security scanning of container images
