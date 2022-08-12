package main

import (
	"fmt"

	secretproviderclient "sigs.k8s.io/secrets-store-csi-driver/pkg/client/clientset/versioned"
)

func main() {
	var client secretproviderclient.Interface

	fmt.Print(client)
}
