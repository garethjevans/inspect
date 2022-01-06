package kube

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/garethjevans/inspect/pkg/util"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// load auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Kuber type to interact with k8s.
type Kuber struct {
}

func (k *Kuber) loadClientConfig() (*rest.Config, error) {
	po := clientcmd.NewDefaultPathOptions()
	if po == nil {
		return nil, errors.New("unable to get kube config path options")
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", po.GlobalFile)
	if err != nil {
		return nil, err
	}
	// for testing purposes one can enable tracing of Kube REST API calls
	traceKubeAPI := os.Getenv("TRACE_KUBE_API")
	if traceKubeAPI == "1" || traceKubeAPI == "on" {
		restConfig.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
			return &Tracer{RoundTripper: rt}
		}
	}

	return restConfig, nil
}

// GetImagesForNamespace lists all images within a namespace.
func (k *Kuber) GetImagesForNamespace(namespace string) ([]string, error) {
	config, err := k.loadClientConfig()
	if err != nil {
		return nil, err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	images := []string{}

	for _, pod := range pods.Items {
		for _, ic := range pod.Spec.InitContainers {
			if !util.Contains(images, ic.Image) {
				images = append(images, ic.Image)
			}
		}
		for _, c := range pod.Spec.Containers {
			if !util.Contains(images, c.Image) {
				images = append(images, c.Image)
			}
		}
	}
	return images, nil
}
