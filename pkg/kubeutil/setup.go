package kubeutil

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// SetUp deploys pods to the already made EC2 instances.
func SetUp(podCount string) error {

	// Create absolute file path to deployment yaml
	// // home, present := os.LookupEnv("HOME")
	// if present == false {
	// 	return errors.New("Home variable is not set. Can't find kubernetes/deployment.yaml")
	// }
	filePath := fmt.Sprint("kubernetes/deployment.yaml")

	// Appy deployment
	out, err := exec.Command("sudo", "kubectl", "apply", "-f", filePath).Output()
	if err != nil {
		return fmt.Errorf("Could not kubectl apply podCount==%s to filepath: %s", podCount, filePath)
	}

	log.Printf("-Successful Deployment-\n%s\n", out)

	// kubectl scale deployment.v1.apps/collider-deployment --relicas=```podCount```
	scaleOut, err := exec.Command("sudo", "kubectl", "scale", "deployment.v1.apps/collider-deployment", "--replicas", podCount).Output()
	if err != nil {
		return fmt.Errorf("Could not scale deployment to match user request podCount=%s", podCount)
	}

	time.Sleep(time.Duration(60) * time.Second)
	log.Printf("-Sucessful Scaling-\n%s\n", scaleOut)

	return nil
}
