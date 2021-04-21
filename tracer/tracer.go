package tracer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/misbahulard/jenkins-tracer/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	envs := getEnvVars()

	buildUrl := fmt.Sprintf("%v", envs["BUILD_URL"])
	getJenkinsBuild(buildUrl, &envs)

	// add timestamp
	envs["@timestamp"] = time.Now()

	// check and set build user
	if envs["BUILD_USER"] == "" {
		envs["BUILD_USER"] = "webhook"
	}

	log.Debugf("Environments: %+v", envs)

	// marshal json
	data, err := json.Marshal(envs)
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		os.Exit(1)
	}

	log.Debugf("Json data: %s", data)

	// index name
	var idxNm string
	if viper.GetString("elasticsearch.index") == "" {
		idxNm = "jenkins"
	} else {
		idxNm = viper.GetString("elasticsearch.index")
	}

	indexName := fmt.Sprintf("%s-%s", idxNm, time.Now().Format("2006.01"))
	log.Infof("Index name: %s", indexName)

	payload := bytes.NewBuffer(data)
	res, err := config.EsClient.Index(indexName, payload)
	if err != nil {
		log.Errorf("Error when send data to elasticsearch cluster: %s", err)
		os.Exit(1)
	}

	if res.StatusCode == 201 {
		log.Info("Data has been send to elasticsearch cluster")
	}
}

func getEnvVars() map[string]interface{} {
	keywords := []string{"BUILD_ID", "BUILD_NUMBER", "EXECUTOR_NUMBER"}
	envs := make(map[string]interface{})

	for _, env := range os.Environ() {
		str := strings.Split(env, "=")
		k, v := str[0], str[1]

		envs[k] = v
	}

	// parse string to int if it in the keyword list
	for k, v := range envs {
		for _, kword := range keywords {
			if k == kword {
				envs[k], _ = strconv.Atoi(fmt.Sprintf("%s", v))
			}
		}
	}

	return envs
}

func getJenkinsBuild(buildUrl string, envs *map[string]interface{}) {
	url := buildUrl + "api/json"

	client := resty.New().SetDebug(viper.GetBool("log.debug"))

	jenkinsResponse := &JenkinsResponse{}

	res, err := client.
		SetRetryCount(3).
		R().
		SetBasicAuth(viper.GetString("jenkins.username"), viper.GetString("jenkins.token")).
		SetHeader("Content-Type", "application/json").
		SetResult(&jenkinsResponse).
		Get(url)

	if err != nil {
		log.Errorf("Failed to get jenkins build status: %s", err)
		os.Exit(1)
	}

	if res.StatusCode() != 200 {
		log.Error("Failed to get jenkins build status: %s", res.Body())
	}

	log.Infof("Build result: %v", jenkinsResponse.Result)

	buildStart := time.Unix(jenkinsResponse.Timestamp/1000, 0)
	log.Infof("Build start: %v", buildStart)

	buildEnd := time.Unix(time.Now().Unix(), 0)
	log.Infof("Build end: %v", buildEnd)

	buildDuration := buildEnd.Sub(buildStart)
	log.Infof("Build duration: %v", buildDuration)
	log.Infof("Build duration in seconds: %v", buildDuration.Seconds())

	(*envs)["BUILD_RESULT"] = jenkinsResponse.Result
	(*envs)["BUILD_START"] = buildStart
	(*envs)["BUILD_END"] = buildEnd
	(*envs)["BUILD_DURATION"] = buildDuration.String()
	(*envs)["BUILD_DURATION_SECONDS"] = buildDuration.Seconds()
}
