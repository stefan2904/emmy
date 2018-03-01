/*
 * Copyright 2017 XLAB d.o.o.
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
 *
 */

package main

import (
	"os"

	"github.com/urfave/cli"
	emmy "github.com/xlab-si/emmy/cmd"
)

// main runs the emmy CLI app.
func main() {
	app := cli.NewApp()
	app.Name = "emmy"
	app.Version = "0.1"
	app.Usage = `A CLI app for running emmy server, emmy clients 
		and examples of proofs offered by the emmy library`
	app.Commands = []cli.Command{emmy.ServerCmd, emmy.ClientCmd}

	app.Run(os.Args)
}
