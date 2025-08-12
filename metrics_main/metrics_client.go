package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	pm "catch/metrics_main/metrics"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	localConfig, err := os.ReadFile("metrics_main/config.local")
	if err != nil {
		log.Fatal(err)
	}

	lineSep := "\r\n"
	lineIndex1 := bytes.Index(localConfig, []byte(lineSep))
	if lineIndex1 == -1 {
		lineSep = "\n"
		lineIndex1 = bytes.Index(localConfig, []byte(lineSep))
	}
	lineSeparatorLen := len(lineSep)

	host := string(localConfig[len("host="):lineIndex1])
	token := string(localConfig[lineIndex1+lineSeparatorLen+len("token="):])

	conn, err := grpc.NewClient("localhost:9890", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error creating metrics client:", err)
	}
	defer conn.Close()

	client := pm.NewCatchMetricsClient(conn)

	allMetrics := []func(ctx context.Context, in *pm.GetRequest, opts ...grpc.CallOption) (*pm.MetricsReply, error){
		client.GetNumberOfActiveUsers,
		client.GetGameBooferSize,
		client.GetGameBooferGamesAvailable,
	}
	var metric *pm.MetricsReply
	var b strings.Builder
	var req *http.Request
	httpClient := &http.Client{}
	for {
		b.WriteString(`{"resourceMetrics":[{"scopeMetrics":[{"metrics":[`)
		for i, fun := range allMetrics {
			metric, err = fun(context.Background(), &pm.GetRequest{})
			if err != nil {
				log.Fatal("error getting metric GetNumberOfActiveUsers:", err)
			}
			b.WriteString(fmt.Sprintf(`{
              "name": "%s",
              "description": "",
              "gauge": {
                "dataPoints": [
                  {
                    "asInt": %d,
                    "timeUnixNano": %d
                  }
                ]
              }
            }`, metric.MetricName, metric.MetricValue, time.Now().UnixNano()))
			if i < len(allMetrics)-1 {
				b.WriteString(",")
			}
		}
		b.WriteString(`]}]}]}`)

		req, err = http.NewRequest("POST", host, bytes.NewBuffer([]byte(b.String())))
		if err != nil {
			log.Fatal("error while sending metrics:", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		b.Reset()
		time.Sleep(60 * time.Second)
	}

}
