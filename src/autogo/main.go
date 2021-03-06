// Copyright 2012 polaris(studygolang.com). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
    "config"
    "flag"
    "logger"
    "os"
    "runtime"
)

var configFile string
var AUTOGO_ROOT = os.Getenv("AUTOGO_ROOT")
var AUTOGO_CMD = os.Getenv("AUTOGO_CMD")

func init() {
    logger.SetLogLevel(2048)
    runtime.GOMAXPROCS(runtime.NumCPU())
    if AUTOGO_ROOT == "" {
        logger.LogFatalf("Please use the autogo shell cmd,or check the cmd '" + AUTOGO_CMD + "' is right or not.")
    }
    logger.Logf("Start autogo root " + AUTOGO_ROOT)
    flag.StringVar(&configFile, "f", AUTOGO_ROOT+"/config/projects.json", "配置文件：需要监听哪些工程")
    flag.Parse()
}

func main() {
    config.Load(configFile)
    config.Watch(configFile)
    select {}
}
