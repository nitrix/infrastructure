package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Target struct{
	Namespace string
	File      string
}

var targets = []Target{
	{ Namespace: "default", File: "default" },
	{ Namespace: "nitrixme", File: "nitrix.me" },
	{ Namespace: "nekohubcom", File: "nekohub.com" },
	{ Namespace: "ingress-nginx", File: "ingress.yml" },
}

var purgeWhitelist = []string{
	"core/v1/ConfigMap",
	"core/v1/Endpoints",
	"core/v1/PersistentVolumeClaim",
	"core/v1/Pod",
	"core/v1/ReplicationController",
	"core/v1/Service",
	"batch/v1/Job",
	"batch/v1beta1/CronJob",
	"apps/v1/DaemonSet",
	"apps/v1/Deployment",
	"apps/v1/ReplicaSet",
	"apps/v1/StatefulSet",
}

func main() {
	for _, target := range targets {
		err := apply(target)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func apply(target Target) error {
	args := []string{
		"apply",
		"--all",
		"--recursive",
		"--namespace", target.Namespace,
		"--prune",
	}

	// args = append(args, "--dry-run")

	for _, w := range purgeWhitelist {
		args = append(args, fmt.Sprintf("--prune-whitelist=%s", w))
	}

	args = append(args, "-f", target.File)

	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}