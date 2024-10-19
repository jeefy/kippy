package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/jeefy/kippy/pkg/config"
	"github.com/jeefy/kippy/pkg/events"
	"github.com/jeefy/kippy/pkg/notify"
	"github.com/jeefy/kippy/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var HeartBeat types.KippyHeartbeat
var HeartBeatLock sync.Mutex

var Cmd = &cobra.Command{
	Use:  "kippy",
	Long: "Easy Kubernetes monitoring and alerting.",
	RunE: run,
}

var args struct {
	debug          bool
	heartBeatKey   string
	heartBeatURL   string
	kubeconfig     string
	email          string
	discordWebhook string
	slackWebhook   string
	genericWebhook string
	sendgridAPIKey string
}

func init() {
	flags := Cmd.Flags()

	flags.BoolVar(
		&args.debug,
		"debug",
		false,
		"Enable debug logging",
	)
	flags.StringVar(
		&args.heartBeatKey,
		"heartBeatKey",
		"",
		"API Key for kippy",
	)
	flags.StringVar(
		&args.heartBeatURL,
		"heartBeatURL",
		"",
		"URL for kippy",
	)
	flags.StringVar(
		&args.email,
		"email",
		"",
		"Email address to send notifications",
	)
	flags.StringVar(
		&args.sendgridAPIKey,
		"sendgrid-api-key",
		"",
		"API Key for SendGrid for Emails / SMS",
	)
	flags.StringVar(
		&args.discordWebhook,
		"discordWebhook",
		"",
		"Discord channel webhook URL",
	)
	flags.StringVar(
		&args.slackWebhook,
		"slackWebhook",
		"",
		"Slack channel webhook URL",
	)
	flags.StringVar(
		&args.genericWebhook,
		"genericWebhook",
		"",
		"Custom webhook URL to receive notifications",
	)
	if home := homedir.HomeDir(); home != "" {
		flags.StringVar(
			&args.kubeconfig,
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"Path to the kubeconfig",
		)
	} else {
		flags.StringVar(
			&args.kubeconfig,
			"kubeconfig",
			"",
			"Path to the kubeconfig",
		)
	}

	viper.BindPFlags(flags)
}

func run(cmd *cobra.Command, argv []string) error {
	// creates the in-cluster config
	log.Println("Welcome to Kippy! Your simple Kubernetes insurance policy.")
	config.LoadConfig(cmd)
	kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Failed to create in-cluster config: %v", err)
		var kubeconfigPath *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfigPath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfigPath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}

		// use the current context in kubeconfig
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", *kubeconfigPath)

		if err != nil {
			log.Printf("Failed to create Kubernetes go-client: %v", err)
		}
	}

	if kubeconfig == nil {
		panic("Failed to create Kubernetes go-client")

	}

	log.Println("Kubeconfig found! Now watching events...")

	go func() {
		events.WatchEvents(kubeconfig, &HeartBeat, &HeartBeatLock)
	}()
	for {
		// TODO do something here
		time.Sleep(60 * time.Second)
		go notify.SendNotification(&HeartBeat, &HeartBeatLock)
	}
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
