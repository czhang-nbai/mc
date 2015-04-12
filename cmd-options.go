/*
 * Modern Copy, (C) 2014,2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/minio-io/cli"
)

// List of commands
var (
	cpCmd = cli.Command{
		Name:  "cp",
		Usage: "Copy objects and files",
		//		Description: "Copy files and objects recursively across object storage and filesystems",
		Action: doCopyCmd,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "recursive, r",
				Usage: "Recursively crawl a given directory or bucket",
			},
		},
		CustomHelpTemplate: `NAME:
   mc {{.Name}} - {{.Usage}}

USAGE:
   mc {{.Name}}{{if .Flags}} [ARGS...]{{end}} SOURCE TARGET [TARGET...] {{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
EXAMPLES:
   1. Copy an object from Amazon S3 object storage to local fileystem.
      $ mc {{.Name}} https://s3.amazonaws.com/jukebox/klingon_opera_aktuh_maylotah.ogg wakeup.ogg

   2. Copy a bucket recursive from Minio object storage to Amazon S3 object storage
      $ mc {{.Name}} --recursive http://localhost:9000/photos/burningman2011 https://s3.amazonaws.com/burningman/

   3. Copy a local folder to Minio object storage and Amazon S3 object storage
      $ mc {{.Name}} --recursive backup/ http://localhost:9000/archive/ https://s3.amazonaws.com/archive/

`,
	}

	lsCmd = cli.Command{
		Name:  "ls",
		Usage: "List files and objects",
		//		Description: `List files and objects recursively on object storage and fileystems`,
		Action: doListCmd,
		CustomHelpTemplate: `NAME:
   mc {{.Name}} - {{.Usage}}

USAGE:
   mc {{.Name}} TARGET {{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}

EXAMPLES:
   1. List objects on Minio object storage
      $ mc {{.Name}} http://localhost:9000/backup/
      2015-03-28 12:47:50 PDT      51.00 MB 2006-Jan-1/backup.tar.gz
      2015-03-31 14:46:33 PDT      55.00 MB 2006-Mar-1/backup.tar.gz

   2. List buckets on Amazon S3 object storage
      $ mc {{.Name}} https://s3.amazonaws.com/
      2015-01-20 15:42:00 PST               rom
      2015-01-15 00:05:40 PST               zek

`,
	}

	mbCmd = cli.Command{
		Name:  "mb",
		Usage: "Make a bucket",
		//		Description: `Create a bucket on object storage or a folder on filesystem`,
		Action: doMakeBucketCmd,
		CustomHelpTemplate: `NAME:
   mc {{.Name}} - {{.Usage}}

USAGE:
   mc {{.Name}} TARGET {{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}

EXAMPLES:
   1. Create a bucket on Amazon S3 object storage
      $ mc {{.Name}} https://s3.amazonaws.com/public-photo-store

   2. Create a bucket on Minio object storage
      $ mc {{.Name}} http://localhost:9000/mongodb-backup

`,
	}

	configCmd = cli.Command{
		Name:  "config",
		Usage: "Generate configuration \"" + mustGetMcConfigPath() + "\" file.",
		//		Description: Configure minio client configuration.
		//
		//   NOTE: that the configure command only writes values to the config file.
		//   It does not use any configuration values from the environment variables.`,
		Action: doConfigCmd,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "accesskeyid, a",
				Value: "",
				Usage: "AWS access key ID",
			},
			cli.StringFlag{
				Name:  "secretkey, s",
				Value: "",
				Usage: "AWS secret access key",
			},
			cli.StringFlag{
				Name:  "alias",
				Value: "",
				Usage: "Add aliases into config",
			},
			cli.BoolFlag{
				Name:  "completion",
				Usage: "Generate bash completion \"" + mustGetMcBashCompletionFilename() + "\" file.",
			},
		},
	}
	updateCmd = cli.Command{
		Name:        "update",
		Usage:       "Check for new software updates",
		Description: "",
		Action:      doUpdateCmd,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "yes, y",
				Usage: "Download and update local binary",
			},
		},
	}
)

var options = []cli.Command{
	cpCmd,
	lsCmd,
	mbCmd,
	configCmd,
	updateCmd,
}

var (
	flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable HTTP tracing",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "Disable chatty output, such as the progress bar",
		},
	}
)

var (
	mcBashCompletion = `#!/bin/bash

_mc_completion() {
    local cur prev opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}

complete -F _mc_completion mc
`
)
