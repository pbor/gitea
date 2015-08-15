// +build nobindata

package public

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Unknwon/com"
	"github.com/go-gitea/gitea/modules/setting"
)

func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)

	if com.IsExist(customFile(cannonicalName)) {
		result, err := ioutil.ReadFile(customFile(cannonicalName))

		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}

		return result, nil
	}

	return nil, fmt.Errorf("Asset %s not found", name)
}

func MustAsset(name string) []byte {
	a, err := Asset(name)

	if err != nil {
		panic(fmt.Sprintf("asset: Asset(%s): %s", name, err.Error()))
	}

	return a
}

func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)

	if com.IsExist(customFile(cannonicalName)) {
		return os.Stat(customFile(name))
	}

	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

func AssetNames() []string {
	result := make([]string, 0)

	if com.IsDir(customPath()) {
		files, err := com.StatDir(customPath(), true)

		if err != nil {
			panic(fmt.Sprintf("asset: AssetNames(): %s", err.Error()))
		}

		for _, f := range files {
			if !com.IsSliceContainsStr(result, f) {
				result = append(result, f)
			}
		}
	}

	return result
}

func AssetDir(name string) ([]string, error) {
	result := make([]string, 0)

	if com.IsDir(customFile(name)) {
		files, err := com.StatDir(customFile(name), true)

		if err != nil {
			return nil, fmt.Errorf("Failed to read %s directory", name)
		}

		for _, f := range files {
			if !com.IsSliceContainsStr(result, f) {
				result = append(result, f)
			}
		}
	}

	return result, nil
}

func customPath() string {
	return path.Join(setting.StaticRootPath, "public")
}

func customFile(name string) string {
	return path.Join(customPath(), name)
}
