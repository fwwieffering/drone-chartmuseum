package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "drone-chartmuseum-plugin"
	app.Usage = "drone plugin to upload charts to chartmuseum server"
	app.Version = fmt.Sprintf("1.0.0")

	mainFlag := []cli.Flag{
		cli.StringFlag{
			Name:   "repo-url",
			Value:  "",
			Usage:  "chartmuseum server endpoint",
			EnvVar: "PLUGIN_REPO_URL,REPO_URL",
		},
		cli.StringFlag{
			Name:   "chart-path",
			Usage:  "chart path",
			Value:  "",
			EnvVar: "PLUGIN_CHART_PATH,CHART_PATH",
		},
		cli.StringFlag{
			Name:   "charts-dir",
			Value:  "./",
			Usage:  "chart directory",
			EnvVar: "PLUGIN_CHARTS_DIR,CHARTS_DIR",
		},
		cli.StringFlag{
			Name:   "save-dir",
			Value:  "uploads/",
			Usage:  "directory to save chart packages",
			EnvVar: "PLUGIN_SAVE_DIR,SAVE_DIR",
		},
		cli.StringFlag{
			Name:   "previous-commit",
			Usage:  "previous commit id (`COMMIT_SHA`)",
			EnvVar: "PLUGIN_PREVIOUS_COMMIT,PREVIOUS_COMMIT",
		},
		cli.StringFlag{
			Name:   "current-commit",
			Usage:  "current commit id (`COMMIT_SHA`)",
			EnvVar: "PLUGIN_CURRENT_COMMIT,CURRENT_COMMIT",
		},
	}

	app.Action = cli.ActionFunc(defaultAction)
	app.Flags = mainFlag

	return app
}

func defaultAction(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			RepoURL:          c.String("repo-url"),
			ChartsDir:        c.String("charts-dir"),
			ChartPath:        c.String("chart-path"),
			PreviousCommitID: c.String("previous-commit"),
			CurrentCommitID:  c.String("current-commit"),
			SaveDir:          c.String("save-dir"),
		},
	}
	return plugin.exec()
}

func main() {
	app := initApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
