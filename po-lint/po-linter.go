package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/ministryofjustice/cloud-platform-environments/pkg/authenticate"
	"github.com/prometheus-operator/prometheus-operator/pkg/admission"
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	"github.com/prometheus/prometheus/model/rulefmt"
	githubaction "github.com/sethvargo/go-githubactions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func prDir(token, githubrepo, githubref string) ([]string, error) {
	if token == "" {
		return nil, errors.New("you must have a valid GitHub token")
	}

	client, err := authenticate.GitHubClient(token)
	if err != nil {
		return nil, err
	}

	//repo user and repo name
	githubrepoS := strings.Split(githubrepo, "/")
	owner := githubrepoS[0]
	repo := githubrepoS[1]

	// get pr owner
	githubrefS := strings.Split(githubref, "/")
	branch := githubrefS[2]
	bid, _ := strconv.Atoi(branch)

	repos, _, _ := client.PullRequests.ListFiles(context.Background(), owner, repo, bid, nil)

	var fileDir []string
	for _, repo := range repos {
		if strings.Contains(*repo.Filename, "kubectl_deploy") {
			fileDir = append(fileDir, *repo.Filename)
		}
	}
	return fileDir, nil

}

func polint(dir []string) {
	files := dir

	for _, filename := range files {
		log.SetPrefix(fmt.Sprintf("%s: ", filename))
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		var meta metav1.TypeMeta

		err = yaml.Unmarshal(content, &meta)
		if err != nil {
			log.Fatal(err)
		}

		switch meta.Kind {
		case v1.AlertmanagersKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var alertmanager v1.Alertmanager
			err = decoder.Decode(&alertmanager)
			if err != nil {
				log.Fatalf("alertmanager is invalid: %v", err)
			}
		case v1.PrometheusesKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var prometheus v1.Prometheus
			err = decoder.Decode(&prometheus)
			if err != nil {
				log.Fatalf("prometheus is invalid: %v", err)
			}
		case v1.PrometheusRuleKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var rule v1.PrometheusRule
			err = decoder.Decode(&rule)
			if err != nil {
				log.Fatalf("prometheus rule is invalid: %v", err)
			}
			err = validateRules(content)
			if err != nil {
				log.Fatalf("prometheus rule validation failed: %v", err)
			}
		case v1.ServiceMonitorsKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var serviceMonitor v1.ServiceMonitor
			err = decoder.Decode(&serviceMonitor)
			if err != nil {
				log.Fatalf("serviceMonitor is invalid: %v", err)
			}
		case v1.PodMonitorsKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var podMonitor v1.PodMonitor
			err = decoder.Decode(&podMonitor)
			if err != nil {
				log.Fatalf("podMonitor is invalid: %v", err)
			}
		case v1.ProbesKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var probe v1.Probe
			if err := decoder.Decode(&probe); err != nil {
				log.Fatalf("probe is invalid: %v", err)
			}
		case v1.ThanosRulerKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var thanosRuler v1.ThanosRuler
			err = decoder.Decode(&thanosRuler)
			if err != nil {
				log.Fatalf("thanosRuler is invalid: %v", err)
			}
		case v1alpha1.AlertmanagerConfigKind:
			j, err := yaml.YAMLToJSON(content)
			if err != nil {
				log.Fatalf("unable to convert YAML to JSON: %v", err)
			}

			decoder := json.NewDecoder(bytes.NewBuffer(j))
			decoder.DisallowUnknownFields()

			var alertmanagerConfig v1alpha1.AlertmanagerConfig
			err = decoder.Decode(&alertmanagerConfig)
			if err != nil {
				log.Fatalf("alertmanagerConfig is invalid: %v", err)
			}
		default:
			log.Print("MetaType is unknown to linter. Not in Alertmanager, Prometheus, PrometheusRule, ServiceMonitor, PodMonitor, Probe, ThanosRuler, AlertmanagerConfig")
		}
	}
	githubaction.SetOutput("po-linter", "true")
}

func validateRules(content []byte) error {
	rule := &admission.PrometheusRules{}
	err := yaml.Unmarshal(content, rule)
	if err != nil {
		return fmt.Errorf("unable load prometheus rule: %w", err)
	}
	rules, errorsArray := rulefmt.Parse(rule.Spec.Raw)
	if len(errorsArray) != 0 {
		for _, err := range errorsArray {
			log.Println(err)
		}
		return errors.New("rules are not valid")
	}
	if len(rules.Groups) == 0 {
		return errors.New("no group found")
	}
	for _, group := range rules.Groups {
		if len(group.Rules) == 0 {
			return fmt.Errorf("no rules found in group: %s: %w", group.Name, err)
		}
	}
	return nil
}

var (
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
)

func main() {
	d, err := prDir(*token, *githubrepo, *githubref)
	if err != nil {
		log.Fatal(err)
	}
	polint(d)
}
