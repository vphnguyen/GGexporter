package entities


import  "encoding/xml"

type MpointsOfExtract struct {
    XMLName                   xml.Name `xml:"mpoints"`
    Text                      string   `xml:",chardata"`
    Xsi                       string   `xml:"xsi,attr"`
    NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
    Name                      string   `xml:"name,attr"`
    Process                   struct {
        Text            string `xml:",chardata"`
        Name            string `xml:"name,attr"`
        Type            string `xml:"type,attr"`
        Mode            string `xml:"mode,attr"`
        Status          string `xml:"status,attr"`
        ProcessID       string `xml:"process-id"`
        StartTime       string `xml:"start-time"`
        PortNumber      string `xml:"port-number"`
        FirstMessage    string `xml:"first-message"`
        LastMessage     string `xml:"last-message"`
        ConfigurationEr struct {
            Text          string `xml:",chardata"`
            GgsVersion    string `xml:"ggs-version"`
            InputType     string `xml:"input-type"`
            OutputType    string `xml:"output-type"`
            BuildDatabase string `xml:"build-database"`
            BuildPlatform string `xml:"build-platform"`
            BuildDate     string `xml:"build-date"`
        } `xml:"configuration-er"`
        DatabaseInOut struct {
            Text         string `xml:",chardata"`
            DbName       string `xml:"db-name"`
            DbType       string `xml:"db-type"`
            DbVersion    string `xml:"db-version"`
            DbServerName string `xml:"db-server-name"`
            DbGlobalName string `xml:"db-global-name"`
        } `xml:"database-in-out"`
        TrailOutput []struct {
            Text                   string `xml:",chardata"`
            TrailName              string `xml:"trail-name"`
            TrailPath              string `xml:"trail-path"`
            Hostname               string `xml:"hostname"`
            PortNumber             string `xml:"port-number"`
            PortType               string `xml:"port-type"`
            IoWriteCount           string `xml:"io-write-count"`
            IoWriteBytes           string `xml:"io-write-bytes"`
            TrailWriteTime         string `xml:"trail-write-time"`
            TrailLobBytes          string `xml:"trail-lob-bytes"`
            TrailBufferFlushes     string `xml:"trail-buffer-flushes"`
            TrailBufferFlushTime   string `xml:"trail-buffer-flush-time"`
            TrailBytesFlushed      string `xml:"trail-bytes-flushed"`
            TrailLastBufferFlushTs string `xml:"trail-last-buffer-flush-ts"`
            TrailRba               string `xml:"trail-rba"`
            TrailSeq               string `xml:"trail-seq"`
            TrailMaxBytes          string `xml:"trail-max-bytes"`
        } `xml:"trail-output"`
        PositionEr struct {
            Text            string `xml:",chardata"`
            LowWatermarkLag string `xml:"low-watermark-lag"`
            LowWatermarkTs  struct {
                Text string `xml:",chardata"`
                Nil  string `xml:"nil,attr"`
            } `xml:"low-watermark-ts"`
            HighWatermarkLag string `xml:"high-watermark-lag"`
            HighWatermarkTs  struct {
                Text string `xml:",chardata"`
                Nil  string `xml:"nil,attr"`
            } `xml:"high-watermark-ts"`
            LastOperationLag string `xml:"last-operation-lag"`
            LastOperationTs  string `xml:"last-operation-ts"`
            LastCheckpointTs string `xml:"last-checkpoint-ts"`
            InputCheckpoint  string `xml:"input-checkpoint"`
            OutputCheckpoint string `xml:"output-checkpoint"`
            EndOfFile        string `xml:"end-of-file"`
            TrailTimesAtEof  string `xml:"trail-times-at-eof"`
        } `xml:"position-er"`
        StatisticsExtract struct {
            Text                     string `xml:",chardata"`
            MappedTotalInserts       string `xml:"mapped-total-inserts"`
            MappedTotalUpdates       string `xml:"mapped-total-updates"`
            MappedTotalDeletes       string `xml:"mapped-total-deletes"`
            MappedTotalUpserts       string `xml:"mapped-total-upserts"`
            MappedTotalUnsupported   string `xml:"mapped-total-unsupported"`
            MappedTotalTruncates     string `xml:"mapped-total-truncates"`
            TotalExecutedDdls        string `xml:"total-executed-ddls"`
            TotalExecutedProcedures  string `xml:"total-executed-procedures"`
            TotalDiscards            string `xml:"total-discards"`
            TotalIgnores             string `xml:"total-ignores"`
            TotalConversionErrors    string `xml:"total-conversion-errors"`
            TotalConversionTruncates string `xml:"total-conversion-truncates"`
            TotalRowFetchAttempts    string `xml:"total-row-fetch-attempts"`
            TotalRowFetchFailures    string `xml:"total-row-fetch-failures"`
        } `xml:"statistics-extract"`
        ProcessPerformance struct {
            Text               string `xml:",chardata"`
            ProcessStartTime   string `xml:"process-start-time"`
            ProcessID          string `xml:"process-id"`
            ThreadCount        string `xml:"thread-count"`
            HandleCount        string `xml:"handle-count"`
            CpuTime            string `xml:"cpu-time"`
            KernelTime         string `xml:"kernel-time"`
            UserTime           string `xml:"user-time"`
            IoReadCount        string `xml:"io-read-count"`
            IoWriteCount       string `xml:"io-write-count"`
            IoOtherCount       string `xml:"io-other-count"`
            IoReadBytes        string `xml:"io-read-bytes"`
            IoWriteBytes       string `xml:"io-write-bytes"`
            IoOtherBytes       string `xml:"io-other-bytes"`
            PageFaults         string `xml:"page-faults"`
            WorkingSetSize     string `xml:"working-set-size"`
            PeakWorkingSetSize string `xml:"peak-working-set-size"`
            PrivateBytes       string `xml:"private-bytes"`
        } `xml:"process-performance"`
        ThreadPerformance []struct {
            Text               string `xml:",chardata"`
            ThreadID           string `xml:"thread-id"`
            ThreadName         string `xml:"thread-name"`
            ThreadFunction     string `xml:"thread-function"`
            ThreadStartTime    string `xml:"thread-start-time"`
            ThreadCurrentStack string `xml:"thread-current-stack"`
            CpuTime            string `xml:"cpu-time"`
            KernelTime         string `xml:"kernel-time"`
            UserTime           string `xml:"user-time"`
            ThreadState        string `xml:"thread-state"`
        } `xml:"thread-performance"`
        StatisticsTableExtract struct {
            Text                  string `xml:",chardata"`
            TableName             string `xml:"table-name"`
            MappedTotalInserts    string `xml:"mapped-total-inserts"`
            MappedTotalUpdates    string `xml:"mapped-total-updates"`
            MappedTotalDeletes    string `xml:"mapped-total-deletes"`
            MappedTotalTruncates  string `xml:"mapped-total-truncates"`
            TotalIgnores          string `xml:"total-ignores"`
            TotalDiscards         string `xml:"total-discards"`
            TotalRowFetchAttempts string `xml:"total-row-fetch-attempts"`
            TotalRowFetchFailures string `xml:"total-row-fetch-failures"`
        } `xml:"statistics-table-extract"`
        CacheStatistics struct {
            Text                      string `xml:",chardata"`
            TotalObjectsInCache       string `xml:"total-objects-in-cache"`
            TotalObjects              string `xml:"total-objects"`
            TotalObjectsActive        string `xml:"total-objects-active"`
            TotalObjectsCommitted     string `xml:"total-objects-committed"`
            MaxActiveObjects          string `xml:"max-active-objects"`
            TimesBufferOverflowed     string `xml:"times-buffer-overflowed"`
            TimesGetNextFromFile      string `xml:"times-get-next-from-file"`
            TimesGetLastFromFile      string `xml:"times-get-last-from-file"`
            TimesSmallBuffForcedOut   string `xml:"times-small-buff-forced-out"`
            TimesRetrieved            string `xml:"times-retrieved"`
            TotalNumberOfQHits        string `xml:"total-number-of-q-hits"`
            TotalNumberOfQMisses      string `xml:"total-number-of-q-misses"`
            TotalNumberOfQPuts        string `xml:"total-number-of-q-puts"`
            TotalNumberOfQTries       string `xml:"total-number-of-q-tries"`
            TotalNumberOfQEntries     string `xml:"total-number-of-q-entries"`
            MaxNumberOfQEntries       string `xml:"max-number-of-q-entries"`
            TotalMunmap               string `xml:"total-munmap"`
            TotalCnnblAttempts        string `xml:"total-cnnbl-attempts"`
            TotalCnnblSuccess         string `xml:"total-cnnbl-success"`
            TotalCnnblMbufs           string `xml:"total-cnnbl-mbufs"`
            TotalFileCacheRequests    string `xml:"total-file-cache-requests"`
            TotalFileCacheEntries     string `xml:"total-file-cache-entries"`
            TotalFileCachePlaced      string `xml:"total-file-cache-placed"`
            MaxQlength                string `xml:"max-qlength"`
            MaxProcessed              string `xml:"max-processed"`
            TimesWaitSignaled         string `xml:"times-wait-signaled"`
            TimesFileCacheNotNeeded   string `xml:"times-file-cache-not-needed"`
            TimesRequestorNeededFc    string `xml:"times-requestor-needed-fc"`
            TotalObjectsInFileCache   string `xml:"total-objects-in-file-cache"`
            TotalFileCacheBytesToDisk string `xml:"total-file-cache-bytes-to-disk"`
            TimesCacheFlushed         string `xml:"times-cache-flushed"`
            MaxMemoryUsage            string `xml:"max-memory-usage"`
            AverageMemoryUsage        string `xml:"average-memory-usage"`
        } `xml:"cache-statistics"`
        SuperpoolStats struct {
            Text                string   `xml:",chardata"`
            CurrentVmUsed       string   `xml:"current-vm-used"`
            HardPageout         string   `xml:"hard-pageout"`
            MaxVmUsed           []string `xml:"max-vm-used"`
            CurrentMmapAnon     string   `xml:"current-mmap-anon"`
            CurrentMmapFiles    string   `xml:"current-mmap-files"`
            CurrentMmapRecycles string   `xml:"current-mmap-recycles"`
            VmDefaultAllocation string   `xml:"vm-default-allocation"`
            VmMaxAllocation     string   `xml:"vm-max-allocation"`
            SoftPageoutSize     string   `xml:"soft-pageout-size"`
            InitialVmIncrement  string   `xml:"initial-vm-increment"`
            MinObjToPageout     string   `xml:"min-obj-to-pageout"`
            MinVmToPageout      string   `xml:"min-vm-to-pageout"`
            MmapGrainularity    string   `xml:"mmap-grainularity"`
            CurrentDiskUse      string   `xml:"current-disk-use"`
            CurrentFileQueueLen string   `xml:"current-file-queue-len"`
            ActiveFileQueue     string   `xml:"active-file-queue"`
        } `xml:"superpool-stats"`
        QueueBucket []struct {
            Text                   string `xml:",chardata"`
            QueueBucketID          string `xml:"queue-bucket-id"`
            QueueBucketSize        string `xml:"queue-bucket-size"`
            QueueBucketQHits       string `xml:"queue-bucket-q-hits"`
            QueueBucketCurrLen     string `xml:"queue-bucket-curr-len"`
            QueueBucketMaxLen      string `xml:"queue-bucket-max-len"`
            QueueBucketAvgLen      string `xml:"queue-bucket-avg-len"`
            QueueBucketCanniblized string `xml:"queue-bucket-canniblized"`
        } `xml:"queue-bucket"`
        BrStatus struct {
            Text                      string `xml:",chardata"`
            BrCurrentStatus           string `xml:"br-current-status"`
            NextCheckpointTimestamp   string `xml:"next-checkpoint-timestamp"`
            LastCheckpointTimestamp   string `xml:"last-checkpoint-timestamp"`
            LastCheckpointNumber      string `xml:"last-checkpoint-number"`
            CheckpointIntervalSeconds string `xml:"checkpoint-interval-seconds"`
            ForceCheckpointTime       string `xml:"force-checkpoint-time"`
            TotalNumObjectsPersisted  string `xml:"total-num-objects-persisted"`
            TotalStateBytesPersisted  string `xml:"total-state-bytes-persisted"`
            TotalDataBytesPersisted   string `xml:"total-data-bytes-persisted"`
            OutstandingNumObjects     string `xml:"outstanding-num-objects"`
            OutstandingStateBytes     string `xml:"outstanding-state-bytes"`
            OutstandingDataBytes      string `xml:"outstanding-data-bytes"`
        } `xml:"br-status"`
    } `xml:"process"`
} 
