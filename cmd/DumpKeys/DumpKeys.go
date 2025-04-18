package DumpKeys

import (
	"os"
	"strings"

	"github.com/SpoofIMEI/BucketDump/cmd/ErrorCheck"
	"github.com/SpoofIMEI/BucketDump/cmd/Log"
	"github.com/SpoofIMEI/BucketDump/cmd/RetUrl"
)

func Dump(keys []string, AWSUrl string, saveDir string) {
	for _, key := range keys {
		Log.Msg("Dumping:"+key, "info")
		keySegments := strings.Split(key, "/")
		os.MkdirAll(saveDir+strings.Join(keySegments[:len(keySegments)-1], "/"), 0600)

		keyUrl := AWSUrl + key

		content, err := RetUrl.Get(keyUrl)
		ErrorCheck.Check(err, 0)

		if stat, err := os.Stat(saveDir + key); err == nil {
			if stat.IsDir() {
				continue
			}
		}
		outputHandle, err := os.OpenFile(saveDir+key, os.O_CREATE|os.O_WRONLY, 0600)
		ErrorCheck.Check(err, 1)

		_, err = outputHandle.Write(content)
		ErrorCheck.Check(err, 1)
	}
}
