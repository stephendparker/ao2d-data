package assets

import (
	"github.com/pkg/errors"
	log  "github.com/sirupsen/logrus"
	"github.com/stephendparker/ao2d-data/utils"
	"github.com/stephendparker/ao2d-data/model"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"os/exec"
	"unicode"
)

type AssetDataFile struct {
	XMLFilePath string
	className   string
}

func VerifyXsdChecksums(target string) (updatedChecksums map[string]string, err error) {

	gameDataFolder := target + GameDataPath

	updatedChecksums = make(map[string]string)

	filepath.Walk(gameDataFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "xsd") {

			className := strings.TrimSuffix(info.Name(), ".xsd")
			xmlFileName := gameDataFolder + "\\" + className + ".xml"

			// if xml file exists
			if _, err := os.Stat(xmlFileName); !os.IsNotExist(err) {

				var checksum string
				if checksum, err = utils.GenerateChecksum(path); err != nil {
					return errors.Wrap(err, "Unable to generate checksum: "+path)
				}

				// create a list of checksums to warn user there now have more fields
				if (checksum != model.XsdMap[className]) {
					updatedChecksums[className] = checksum
				}
			}
		}
		return nil
	})

	modelPath := target + ModelPath
	err = utils.CreateDirIfNotExist(modelPath)
	if err != nil {
		return
	}

	createChecksumFile(modelPath, updatedChecksums)

	return
}

/*
   go get github.com/gnewton/chidley
   cd %GOPATH%\src\github.com\gnewton\chidley
   go install
 */
func GenerateClassFiles(target string) (error) {

	gameDataFolder := target + GameDataPath

	dataFiles := []AssetDataFile{}

	filepath.Walk(gameDataFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "xsd") {

			className := strings.TrimSuffix(info.Name(), ".xsd")
			xmlFileName := gameDataFolder + "\\" + className + ".xml"

			// if xml file exists
			if _, err := os.Stat(xmlFileName); !os.IsNotExist(err) {

				dataFile := AssetDataFile{XMLFilePath: xmlFileName, className: className}
				dataFiles = append(dataFiles, dataFile)
			}
		}
		return nil
	})

	modelPath := target + ModelPath
	err := utils.CreateDirIfNotExist(modelPath)
	if err != nil {
		return errors.Wrap(err, "Unable create model directory.")
	}

	dataFiles = append(dataFiles, AssetDataFile{XMLFilePath: gameDataFolder + "\\localization.xml", className: "localization"})
	dataFiles = append(dataFiles, AssetDataFile{XMLFilePath: gameDataFolder + "\\territorytypes.xml", className: "territorytypes"})
	dataFiles = append(dataFiles, AssetDataFile{XMLFilePath: gameDataFolder + "\\cluster\\world.xml", className: "world"})

	for _, element := range dataFiles {
		err = generateClassFile(element, modelPath)
		if err != nil {
			return errors.Wrap(err, "Unable generate class file.")
		}
	}

	return nil
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func generateClassFile(dataFile AssetDataFile, outputPath string) (err error) {
	// chidley -e Asset -G resourcedistpresets.xml > AssetResourcedistpresets.go

	packageName := LcFirst(dataFile.className)
	packageLocation := outputPath + "\\" + packageName

	err = utils.CreateDirIfNotExist(packageLocation)
	if err != nil {
		return errors.Wrap(err, "Unable create model directory.")
	}

	goFileName := packageLocation + "\\asset" + UcFirst(dataFile.className) + ".go"

	log.Debug("Creating go file: " + goFileName)

	cmd := os.ExpandEnv("${GOPATH}\\src\\github.com\\gnewton\\chidley\\chidley.exe")
	args := []string{"-e", "Asset", "-G", dataFile.XMLFilePath}

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Create(goFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, "package "+packageName)
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, string(out))
	return
}

func createChecksumFile(modelPath string, updatedChecksums map[string]string) {
	file, err := os.Create(modelPath + "\\model-checksums.go")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "package model")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "var XsdMap = map[string]string{")
	for k, element := range updatedChecksums {
		fmt.Fprintln(file, "\t\""+k+"\":\""+element+"\",")
	}
	fmt.Fprintln(file, "}")
}
