package config

import (
    "fsnotify"
    "io/ioutil"
    "logger"
    "project"
    "simplejson"
    "time"
)

// Watch 监控配置文件
func Watch(configFile string) error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    eventNum := make(chan int)
    go func() {
        for {
            i := 0
        GetEvent:
            for {
                select {
                case <-watcher.Event:
                    i++
                case <-time.After(200e6):
                    break GetEvent
                }
            }
            if i > 0 {
                eventNum <- i
            }
        }
    }()

    go func() {
        for {
            select {
            case <-eventNum:
                logger.LogNoticef("[INFO] ReloadConfig...")
                Load(configFile)
            }
        }
    }()

    return watcher.Watch(configFile)
}

// Load加载解析配置文件
func Load(configFile string) error {
    content, err := ioutil.ReadFile(configFile)
    if err != nil {
        logger.LogFatalf("读配置文件错误")
        return err
    }
    allConfig, err := simplejson.NewJson(content)
    if err != nil {
        logger.LogFatalf("配置文件格式错误")
        return err
    }
    middleJs, err := allConfig.Array()
    if err != nil {
        logger.LogFatalf("配置文件格式错误")
        return err
    }
    for i, length := 0, len(middleJs); i < length; i++ {
        oneProject := allConfig.GetIndex(i)
        name := oneProject.Get("name").MustString()
        root := oneProject.Get("root").MustString()
        goWay := oneProject.Get("go_way").MustString()
        deamon := oneProject.Get("deamon").MustBool(true)
        mainFile := oneProject.Get("main").MustString()
        depends := oneProject.GetStringSlice("depends")
        project.Watch(name, root, goWay, mainFile, deamon, depends...)
    }
    return err
}
