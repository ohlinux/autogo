// Copyright 2012 polaris(studygolang.com). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

import (
    "os/exec"
)

var (
    makeTplFile       = AUTOGO_ROOT + "/templates/make_linux.tpl"
    installFileName   = "install.sh"
    binanryFileSuffix = ""

    installCmd = "./" + installFileName // 编译时传给Command的名称
)

//send SIGUSR1 signal to restart 
func (this *Project) SendRestartSignal() error {
    cmd := exec.Command("killall", "-SIGUSR1", this.MainFile)
    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}
