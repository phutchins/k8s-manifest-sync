# k8s-manifest-sync

## Goals
  + Pull all kubernetes manifests from Kubernetes to local disk in desired format (JSON, Yaml, etc...)
  + Push changes from manifest files on local disk to Kubernetes with diff and confirmation
  + Backup & Restore entire Kubernets clusters resources (pods, rcs, deployments, etc...)
