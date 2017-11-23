package assets

import (
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/pkg/errors"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/stephendparker/ao2d-data/utils"
)

const (
	downloadPage          = "https://albiononline.com/en/download"
	manifestURL           = "https://live.albiononline.com/autoupdate/manifest.xml"
	versionIndex          = 4
	unzipSubFolderName    = "unzip"
	unzipCompleteFileName = "unzip-complete.txt"
	androidApkFileName    = "Albion-Online.apk"
)

func GetLatestVersionDetails() (gameVersion, downloadUrl string, err error) {

	log.Debug("reading download page for andriod link and version")
	resp, err := http.Get(downloadPage)
	if err != nil {
		err = errors.Wrap(err, "failed to get html download page")
		return
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read download page")
		return
	}

	htmlParts := strings.Split(string(html), "\"")

	for _, v := range htmlParts {

		if strings.HasSuffix(v, androidApkFileName) {
			version, errGetVersion := getGameVersion(v)
			if err != nil {
				err = errors.Wrap(errGetVersion, "unable to get version")
				return
			}
			downloadUrl = v
			gameVersion = version
			return
		}
	}

	err = errors.New("unable to find download version in html, page format may have changed.")
	return
}

func getGameVersion(andriodUrl string) (string, error) {
	urlParts := strings.Split(andriodUrl, "/")

	if len(urlParts) > versionIndex {

		return urlParts[versionIndex], nil
	}
	return "", errors.New("unable to pattern match andrion link to find version")
}

func RetrieveAssetFiles(url, target string, cleanupInstallationFile bool) (error) {

	unzipPath := target + "\\" + unzipSubFolderName
	unzipCompleteFlagFile := target + "\\" + unzipCompleteFileName
	installFileLocation := target + "\\" + androidApkFileName
	gameDataAssetPath := unzipPath + GameDataAssetsPath
	gameDataDestPath := target + "\\GameData"

	err := utils.CreateDirIfNotExist(unzipPath)
	if err != nil {
		return errors.Wrap(err, "unable create temp directory.")
	}
	if utils.FileExists(unzipCompleteFlagFile) == false {

		if utils.FileExists(installFileLocation) == false {
			log.Debug("downloading installation file to: " + installFileLocation)
			err = utils.DownloadFile(installFileLocation, "http:"+url)
			if err != nil {
				return errors.Wrap(err, "unable to download installation file.")
			}
		} else {
			log.Debug("installation file already exists, skipping unzip.")
		}

		log.Debug("unzipping installation file.")
		_, err = utils.Unzip(installFileLocation, GameDataAssetsPath, unzipPath)
		if err != nil {
			return errors.Wrap(err, "unable unzip installation file.")
		}

		log.Debug("copying game data folder.")
		err = utils.CopyDir(gameDataAssetPath, gameDataDestPath)
		if err != nil {
			return errors.Wrap(err, "Unable to copy game data files out of unzip folder.")
		}

		log.Debug("cleaning up unzip folder.")
		err = os.RemoveAll(unzipPath)
		if err != nil {
			return errors.Wrap(err, "Unable to delete unzip folder.")
		}

		log.Debug("Installation unzip complete, creating flag file prevent re-unzipping.")
		_, err := os.OpenFile(unzipCompleteFlagFile, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			return errors.Wrap(err, "unable to create unzip flag file.")
		}

		if cleanupInstallationFile {
			log.Debug("removing installation file.")
			err = os.Remove(installFileLocation)
			if err != nil {
				return errors.Wrap(err, "unable to cleanup installation file")
			}
		}

	} else {
		log.Debug("unzip file already exists, skipping unzip.")
	}

	return nil
}
