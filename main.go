package main

import (
  "fmt"
  "log"
  "os"
  "path"
  "flag"
  "time"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
  //"k8s.io/client-go/1.5/rest"
)

const (
  defaultServer string = "http://127.0.0.1:8001"
  serverEnvVar  string = "SERVER"
  tokenEnvVar   string = "TOKEN"
  caFileENvVar  string = "CA_FILE"
)

var logger *log.Logger
var (
  kubeconfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")
)

func init() {
  logger = log.New(os.Stdout, "", 0)
}

func usage(msg string) {
  logger.Printf("%s\n\n", msg)
  fmt.Fprintf(os.Stderr, "usage: %s version | pull [type] [name]", path.Base(os.Args[0]))
  os.Exit(1)
}

func main() {
  logger.Printf("Welcome to ksync!")

  flag.Parse()

  // Uses the current context in kubeconfig
  config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
  if err != nil {
    panic(err.Error())
  }

  // Creates the clientset
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  for {
    pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil {
      panic(err.Error())
    }

    fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
    time.Sleep(10 * time.Second)
  }

}

