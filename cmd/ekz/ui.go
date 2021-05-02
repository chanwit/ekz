package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the UI",
	Long:  "The UI command start the EKZ-UI.",
	RunE:  uiCmdRun,
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func waitForServices(services []string, timeOut time.Duration) error {
	var depChan = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(len(services))
	go func() {
		for _, s := range services {
			go func(s string) {
				defer wg.Done()
				for {
					_, err := net.Dial("tcp", s)
					if err == nil {
						return
					}
					time.Sleep(1 * time.Second)
				}
			}(s)
		}
		wg.Wait()
		close(depChan)
	}()

	select {
	case <-depChan: // services are ready
		return nil
	case <-time.After(timeOut):
		return fmt.Errorf("services aren't ready in %s", timeOut)
	}
}

func uiCmdRun(cmd *cobra.Command, args []string) error {
	uiImage := "quay.io/ekz-io/ekz-webui:latest"
	logger.Waitingf("pulling the UI image: %s ...", uiImage)
	if err := script.Exec("docker", "pull", uiImage).Run(); err != nil {
		return err
	}

	go func() {
		waitForServices([]string{"localhost:8080"}, 20*time.Second)
		openBrowser("http://localhost:8080")
	}()
	logger.Successf("EKZ UI started at http://localhost:8080 ...")
	logger.Waitingf("press Ctrl + C to stop the UI.")
	err := script.Exec("docker", "run",
		"--rm", "--network=host",
		"-v", expandKubeConfigFile()+":/root/.kube/config",
		uiImage,
	).Run()

	return err
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
