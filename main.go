package main

import (
   "fmt"
    "time"
    "math/rand"
    // "net/http"
    // "github.com/prometheus/client_golang/prometheus/promhttp"
    // "log"
    "GGexporter/entities"
    "GGexporter/services"
)

var (
    mgroups entities.MGroups
    mpointsofextract []entities.MpointsOfExtract
    mpointsofmgr entities.MpointsOfMGR
    mpointsofpmsrvr entities.MpointsOfPMSRVR
)



func main() {
    //=========
    services.GetGGRunningInstances(&mgroups)
    //=====
    services.GetGGRunningExtractInstances(&mgroups,&mpointsofextract)    
    services.GetGGRunningMGRInstances(&mgroups,&mpointsofmgr)    
    services.GetGGRunningPMSRVRInstances(&mgroups,&mpointsofpmsrvr) 

    //=====


    fmt.Printf("%+v\n \n",mgroups)
    



    foo := newFooCollector()
    registry := prometheus.NewRegistry() 
    registry.MustRegister(foo)


    services.ActivieHTTPHandler(registry)



    








}