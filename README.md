# k8s-manifest-sync

## Goals
  + Pull all kubernetes manifests from Kubernetes to local disk in desired format (JSON, Yaml, etc...)
  + Push changes from manifest files on local disk to Kubernetes with diff and confirmation
  + Backup & Restore entire Kubernets clusters resources (pods, rcs, deployments, etc...)

## Links
  + [Kubernetes Go Client](https://github.com/phutchins/client-go)
  + [Example Implementation of Kubernetes Go Client](https://github.com/timoreimann/kubernetes-goclient-example/blob/master/client.go)
