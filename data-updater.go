package main

import (
	"github.com/stephendparker/ao2d-data/assets"
	log "github.com/sirupsen/logrus"
	"github.com/Gurpartap/logrus-stack"
	"os"
	"flag"
)

const (
	StagingFolderName = "ao2d-staging"
)

var (
	generateClasses         = false
	cleanupInstallationFile = true
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	log.AddHook(logrus_stack.StandardHook())
	log.SetFormatter(&log.TextFormatter{})

	flag.BoolVar(&generateClasses, "G", generateClasses, "Generate class files")
	flag.BoolVar(&cleanupInstallationFile, "C", cleanupInstallationFile, "Cleanup installation file after data unzip")

}

func main() {

	stagingLocation := os.TempDir() + "\\" + StagingFolderName

	version, downloadUrl, err := assets.GetLatestVersionDetails()
	if err != nil {
		log.Errorln(err)
		return
	}

	versionDataLocation := stagingLocation + "\\" + version

	err = assets.RetrieveAssetFiles(downloadUrl, versionDataLocation, cleanupInstallationFile)
	if err != nil {
		log.Errorln(err)
		return
	}

	updatedChecksums, err := assets.VerifyXsdChecksums(versionDataLocation)
	if err != nil {
		log.Errorln(err)
		return
	}

	if generateClasses {
		err := assets.GenerateClassFiles(versionDataLocation)
		if err != nil {
			log.Error(err)
			return
		}
	}

	_, err = assets.LoadData(versionDataLocation)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug("dataFileLocation: ", versionDataLocation)

	for k, _ := range updatedChecksums {
		log.Warn("Updated XSD checksum, consider updating go class: " + k)
	}

	log.Debug("Complete.")
}
