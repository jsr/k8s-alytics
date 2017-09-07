package main

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var resources = map[string]string{
	"Pod": "\"kind: Pod\"",
	"ReplicationController":     "\"Kind: ReplicationController\"",
	"Service":                   "\"apiVersion: v1\" \"Kind: Service\" NOT ServiceAccount",
	"ConfigMap":                 "\"Kind: ConfigMap\"",
	"Secret":                    "\"Kind: Secret\"",
	"PersistentVolumeClaim":     "\"Kind: PersistentVolumeClaim\"",
	"Volume":                    "\"apiVersion: v1\" \"Kind: Volume\"",
	"LimitRange":                "\"Kind: LimitRange\"",
	"Namespace":                 "\"Kind: Namespace\"",
	"PersistentVolume":          "\"Kind: PersistentVolume\" NOT persistentvolumeclaim",
	"ResourceQuota":             "\"Kind: ResourceQuota\"",
	"ServiceAccount":            "\"Kind: ServiceAccount\"",
	"Job":                       "\"apiVersion: batch/v1\" \"Kind: Job\"",
	"StorageClass":              "\"Kind: StorageClass\"",
	"HorizontalPodAutoscaler":   "\"Kind: HorizontalPodAutoscaler\"",
	"SubjectAccessReview":       "\"Kind: SubjectAccessReview\"",
	"TokenReview":               "\"Kind: TokenReview\"",
	"NetworkPolicy":             "\"Kind: NetworkPolicy\"",
	"DaemonSet":                 "\"Kind: DaemonSet\"",
	"Deployment":                "\"Kind: Deployment\"",
	"ReplicaSet":                "\"Kind: ReplicaSet\"",
	"StatefulSet":               "\"Kind: StatefulSet\"",
	"Ingress":                   "\"Kind: Ingress\"",
	"ThirdPartyResource":        "\"Kind: ThirdPartyResource\"",
	"PodSecurityPolicy":         "\"Kind: PodSecurityPolicy\"",
	"APIService":                "\"Kind: APIService\"",
	"CertificateSigningRequest": "\"Kind: CertificateSigningRequest\"",
	"ClusterRole":               "\"Kind: ClusterRole\" NOT clusterrolebinding",
	"ClusterRoleBinding":        "\"Kind: ClusterRoleBinding\"",
	"Role":                      "\"Kind: Role\" NOT rolebinding",
	"RoleBinding":               "\"Kind: RoleBinding\"",
	"CronJob":                   "\"Kind: CronJob\"",
	"ExternalAdmissionHookConfiguration": "\"Kind: ExternalAdmissionHookConfiguration\"",
	"InitializerConfiguration":           "\"Kind: InitializerConfiguration\"",
	"Volume - emptyDir":                  "\"Kind: Pod\" volume emptyDir",
	"Volume - hostPath":                  "\"Kind: Pod\" volume hostPath",
	"Volume - gcePersistentDisk":         "\"Kind: Pod\" volume gcePersistentDisk",
	"Volume - awsElasticBlockStore":      "\"Kind: Pod\" volume awsElasticBlockStore",
	"Volume - nfs":                       "\"Kind: Pod\" volume nfs",
	"Volume - iscsi":                     "\"Kind: Pod\" volume iscsi",
	"Volume - fc":                        "\"Kind: Pod\" volume fc",
	"Volume - flocker":                   "\"Kind: Pod\" volume flocker",
	"Volume - glusterfs":                 "\"Kind: Pod\" volume glusterfs",
	"Volume - rbd":                       "\"Kind: Pod\" volume rbd",
	"Volume - cephfs":                    "\"Kind: Pod\" volume cephfs",
	"Volume - gitRepo":                   "\"Kind: Pod\" volume gitRepo",
	"Volume - secret":                    "\"Kind: Pod\" volume secret",
	"Volume - persistentVolumeClaim":     "\"Kind: Pod\" volume persistentVolumeClaim",
	"Volume - downwardAPI":               "\"Kind: Pod\" volume downwardAPI",
	"Volume - projected":                 "\"Kind: Pod\" volume projected",
	"Volume - azureFileVolume":           "\"Kind: Pod\" volume azureFileVolume",
	"Volume - azureDisk":                 "\"Kind: Pod\" volume azureDisk",
	"Volume - vsphereVolume":             "\"Kind: Pod\" volume vsphereVolume",
	"Volume - Quobyte":                   "\"Kind: Pod\" volume Quobyte",
	"Volume - PortworxVolume":            "\"Kind: Pod\" volume PortworxVolume",
	"Volume - ScaleIO":                   "\"Kind: Pod\" volume ScaleIO",
	"Volume - local":                     "\"Kind: Pod\" volume local",
}

func readToken() string {
	path := os.Getenv("HOME") + "/.github_token"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed reading your github token")
		fmt.Fprintln(os.Stderr, "1. Visit https://github.com/blog/1509-personal-api-tokens")
		fmt.Fprintln(os.Stderr, "2. Create a personal API token to use with this tool")
		fmt.Fprintln(os.Stderr, "3. Store the token in a file called $HOME/.github_token")
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: readToken()},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for resource, query := range resources {
		query += " language:yaml language:json"

		var result *github.CodeSearchResult
		var err error

		for {
			result, _, err = client.Search.Code(context.Background(), query, nil)
			if err != nil {
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}

		url := "https://github.com/search?q=" + url.QueryEscape(query) + "&type=Code"
		fmt.Printf("%s, %d, %s\n", resource, result.GetTotal(), url)
	}
}
