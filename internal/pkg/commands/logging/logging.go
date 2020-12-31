package logging

import (
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/commands/logging/azureblob"
	"github.com/integralist/customcli/internal/pkg/commands/logging/bigquery"
	"github.com/integralist/customcli/internal/pkg/commands/logging/cloudfiles"
	"github.com/integralist/customcli/internal/pkg/commands/logging/datadog"
	"github.com/integralist/customcli/internal/pkg/commands/logging/digitalocean"
	"github.com/integralist/customcli/internal/pkg/commands/logging/elasticsearch"
	"github.com/integralist/customcli/internal/pkg/commands/logging/ftp"
	"github.com/integralist/customcli/internal/pkg/commands/logging/gcs"
	"github.com/integralist/customcli/internal/pkg/commands/logging/googlepubsub"
	"github.com/integralist/customcli/internal/pkg/commands/logging/heroku"
	"github.com/integralist/customcli/internal/pkg/commands/logging/honeycomb"
	"github.com/integralist/customcli/internal/pkg/commands/logging/https"
	"github.com/integralist/customcli/internal/pkg/commands/logging/kafka"
	"github.com/integralist/customcli/internal/pkg/commands/logging/logentries"
	"github.com/integralist/customcli/internal/pkg/commands/logging/loggly"
	"github.com/integralist/customcli/internal/pkg/commands/logging/logshuttle"
	"github.com/integralist/customcli/internal/pkg/commands/logging/openstack"
	"github.com/integralist/customcli/internal/pkg/commands/logging/papertrail"
	"github.com/integralist/customcli/internal/pkg/commands/logging/s3"
	"github.com/integralist/customcli/internal/pkg/commands/logging/scalyr"
	"github.com/integralist/customcli/internal/pkg/commands/logging/sftp"
	"github.com/integralist/customcli/internal/pkg/commands/logging/splunk"
	"github.com/integralist/customcli/internal/pkg/commands/logging/sumologic"
	"github.com/integralist/customcli/internal/pkg/commands/logging/syslog"
	"github.com/integralist/customcli/internal/pkg/print"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"azureblob":     "Manipulate Fastly service version Azure Blob Storage logging endpoints",
	"bigquery":      "Manipulate Fastly service version BigQuery logging endpoints",
	"cloudfiles":    "Manipulate Fastly service version Cloudfiles logging endpoints",
	"datadog":       "Manipulate Fastly service version Datadog logging endpoints",
	"digitalocean":  "Manipulate Fastly service version DigitalOcean Spaces logging endpoints",
	"elasticsearch": "Manipulate Fastly service version Elasticsearch logging endpoints",
	"ftp":           "Manipulate Fastly service version FTP logging endpoints",
	"gcs":           "Manipulate Fastly service version GCS logging endpoints",
	"googlepubsub":  "Manipulate Fastly service version Google Cloud Pub/Sub logging endpoints",
	"heroku":        "Manipulate Fastly service version Heroku logging endpoints",
	"honeycomb":     "Manipulate Fastly service version Honeycomb logging endpoints",
	"https":         "Manipulate Fastly service version HTTPS logging endpoints",
	"kafka":         "Manipulate Fastly service version Kafka logging endpoints",
	"logentries":    "Manipulate Fastly service version Logentries logging endpoints",
	"loggly":        "Manipulate Fastly service version Loggly logging endpoints",
	"logshuttle":    "Manipulate Fastly service version Logshuttle logging endpoints",
	"openstack":     "Manipulate Fastly service version OpenStack logging endpoints",
	"papertrail":    "Manipulate Fastly service version Papertrail logging endpoints",
	"s3":            "Manipulate Fastly service version S3 logging endpoints",
	"scalyr":        "Manipulate Fastly service version Scalyr logging endpoints",
	"sftp":          "Manipulate Fastly service version SFTP logging endpoints",
	"splunk":        "Manipulate Fastly service version Splunk logging endpoints",
	"sumologic":     "Manipulate Fastly service version Sumologic logging endpoints",
	"syslog":        "Manipulate Fastly service version Syslog logging endpoints",
}

// Help prints the help info for the command, then exits.
//
// NOTE: we extract this logic into an exported function because we want to
// reuse most of the logic from the top-level 'help' command.
// e.g. fastly help <command>
func Help(cmds map[string]string, args []string, err error) {
	example := "fastly logging <command> [<args> ...]"
	description := "Manipulate Fastly service version logging endpoints"
	print.Help(example, description, "COMMANDS", Cmds)
	print.ErrorContext(args, err)
	os.Exit(1)
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		Help(Cmds, args[2:], err)
	}
	fmt.Println("logging cmd:", cmd)

	switch cmd {
	case "azureblob":
		azureblob.Run(args)
	case "bigquery":
		bigquery.Run(args)
	case "cloudfiles":
		cloudfiles.Run(args)
	case "datadog":
		datadog.Run(args)
	case "digitalocean":
		digitalocean.Run(args)
	case "elasticsearch":
		elasticsearch.Run(args)
	case "ftp":
		ftp.Run(args)
	case "gcs":
		gcs.Run(args)
	case "googlepubsub":
		googlepubsub.Run(args)
	case "heroku":
		heroku.Run(args)
	case "honeycomb":
		honeycomb.Run(args)
	case "https":
		https.Run(args)
	case "kafka":
		kafka.Run(args)
	case "logentries":
		logentries.Run(args)
	case "loggly":
		loggly.Run(args)
	case "logshuttle":
		logshuttle.Run(args)
	case "openstack":
		openstack.Run(args)
	case "papertrail":
		papertrail.Run(args)
	case "s3":
		s3.Run(args)
	case "scalyr":
		scalyr.Run(args)
	case "sftp":
		sftp.Run(args)
	case "splunk":
		splunk.Run(args)
	case "sumologic":
		sumologic.Run(args)
	case "syslog":
		syslog.Run(args)
	}
}
