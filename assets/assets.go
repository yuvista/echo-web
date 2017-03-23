// Code generated by go-bindata.
// sources:
// assets/css/bootstrap-combined.min.css
// assets/css/bootstrap-theme.css
// assets/css/bootstrap-theme.css.map
// assets/css/bootstrap-theme.min.css
// assets/css/bootstrap.css
// assets/css/bootstrap.css.map
// assets/css/bootstrap.min.css
// assets/css/jquery.treetable.css
// assets/css/jquery.treetable.theme.default.css
// assets/fonts/glyphicons-halflings-regular.eot
// assets/fonts/glyphicons-halflings-regular.svg
// assets/fonts/glyphicons-halflings-regular.ttf
// assets/fonts/glyphicons-halflings-regular.woff
// assets/fonts/glyphicons-halflings-regular.woff2
// assets/img/favicon.ico
// assets/img/icon_brand.png
// assets/js/app/jwt/jwt.js
// assets/js/app/treetable/treetable-page.js
// assets/js/lib/bootstrap/bootstrap.js
// assets/js/lib/bootstrap/bootstrap.min.js
// assets/js/lib/bootstrap/npm.js
// assets/js/lib/jquery/jquery-2.1.3.js
// assets/js/lib/jquery/jquery-2.1.3.min.js
// assets/js/lib/jquery-form/jquery.form.3.51.js
// assets/js/lib/jquery-treetable/jquery.treetable.js
// assets/js/lib/require-setup.js
// assets/js/lib/require.js
// assets/js/tools/r.js
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assetsCssBootstrapCombinedMinCss reads file data from disk. It returns an error on failure.
func assetsCssBootstrapCombinedMinCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap-combined.min.css"
	name := "assets/css/bootstrap-combined.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapThemeCss reads file data from disk. It returns an error on failure.
func assetsCssBootstrapThemeCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap-theme.css"
	name := "assets/css/bootstrap-theme.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapThemeCssMap reads file data from disk. It returns an error on failure.
func assetsCssBootstrapThemeCssMap() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap-theme.css.map"
	name := "assets/css/bootstrap-theme.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapThemeMinCss reads file data from disk. It returns an error on failure.
func assetsCssBootstrapThemeMinCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap-theme.min.css"
	name := "assets/css/bootstrap-theme.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapCss reads file data from disk. It returns an error on failure.
func assetsCssBootstrapCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap.css"
	name := "assets/css/bootstrap.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapCssMap reads file data from disk. It returns an error on failure.
func assetsCssBootstrapCssMap() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap.css.map"
	name := "assets/css/bootstrap.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssBootstrapMinCss reads file data from disk. It returns an error on failure.
func assetsCssBootstrapMinCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/bootstrap.min.css"
	name := "assets/css/bootstrap.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssJqueryTreetableCss reads file data from disk. It returns an error on failure.
func assetsCssJqueryTreetableCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/jquery.treetable.css"
	name := "assets/css/jquery.treetable.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssJqueryTreetableThemeDefaultCss reads file data from disk. It returns an error on failure.
func assetsCssJqueryTreetableThemeDefaultCss() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/css/jquery.treetable.theme.default.css"
	name := "assets/css/jquery.treetable.theme.default.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontsGlyphiconsHalflingsRegularEot reads file data from disk. It returns an error on failure.
func assetsFontsGlyphiconsHalflingsRegularEot() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/fonts/glyphicons-halflings-regular.eot"
	name := "assets/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontsGlyphiconsHalflingsRegularSvg reads file data from disk. It returns an error on failure.
func assetsFontsGlyphiconsHalflingsRegularSvg() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/fonts/glyphicons-halflings-regular.svg"
	name := "assets/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontsGlyphiconsHalflingsRegularTtf reads file data from disk. It returns an error on failure.
func assetsFontsGlyphiconsHalflingsRegularTtf() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/fonts/glyphicons-halflings-regular.ttf"
	name := "assets/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontsGlyphiconsHalflingsRegularWoff reads file data from disk. It returns an error on failure.
func assetsFontsGlyphiconsHalflingsRegularWoff() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/fonts/glyphicons-halflings-regular.woff"
	name := "assets/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontsGlyphiconsHalflingsRegularWoff2 reads file data from disk. It returns an error on failure.
func assetsFontsGlyphiconsHalflingsRegularWoff2() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/fonts/glyphicons-halflings-regular.woff2"
	name := "assets/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgFaviconIco reads file data from disk. It returns an error on failure.
func assetsImgFaviconIco() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/img/favicon.ico"
	name := "assets/img/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgIcon_brandPng reads file data from disk. It returns an error on failure.
func assetsImgIcon_brandPng() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/img/icon_brand.png"
	name := "assets/img/icon_brand.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsAppJwtJwtJs reads file data from disk. It returns an error on failure.
func assetsJsAppJwtJwtJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/app/jwt/jwt.js"
	name := "assets/js/app/jwt/jwt.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsAppTreetableTreetablePageJs reads file data from disk. It returns an error on failure.
func assetsJsAppTreetableTreetablePageJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/app/treetable/treetable-page.js"
	name := "assets/js/app/treetable/treetable-page.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibBootstrapBootstrapJs reads file data from disk. It returns an error on failure.
func assetsJsLibBootstrapBootstrapJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/bootstrap/bootstrap.js"
	name := "assets/js/lib/bootstrap/bootstrap.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibBootstrapBootstrapMinJs reads file data from disk. It returns an error on failure.
func assetsJsLibBootstrapBootstrapMinJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/bootstrap/bootstrap.min.js"
	name := "assets/js/lib/bootstrap/bootstrap.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibBootstrapNpmJs reads file data from disk. It returns an error on failure.
func assetsJsLibBootstrapNpmJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/bootstrap/npm.js"
	name := "assets/js/lib/bootstrap/npm.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibJqueryJquery213Js reads file data from disk. It returns an error on failure.
func assetsJsLibJqueryJquery213Js() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/jquery/jquery-2.1.3.js"
	name := "assets/js/lib/jquery/jquery-2.1.3.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibJqueryJquery213MinJs reads file data from disk. It returns an error on failure.
func assetsJsLibJqueryJquery213MinJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/jquery/jquery-2.1.3.min.js"
	name := "assets/js/lib/jquery/jquery-2.1.3.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibJqueryFormJqueryForm351Js reads file data from disk. It returns an error on failure.
func assetsJsLibJqueryFormJqueryForm351Js() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/jquery-form/jquery.form.3.51.js"
	name := "assets/js/lib/jquery-form/jquery.form.3.51.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibJqueryTreetableJqueryTreetableJs reads file data from disk. It returns an error on failure.
func assetsJsLibJqueryTreetableJqueryTreetableJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/jquery-treetable/jquery.treetable.js"
	name := "assets/js/lib/jquery-treetable/jquery.treetable.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibRequireSetupJs reads file data from disk. It returns an error on failure.
func assetsJsLibRequireSetupJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/require-setup.js"
	name := "assets/js/lib/require-setup.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsLibRequireJs reads file data from disk. It returns an error on failure.
func assetsJsLibRequireJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/lib/require.js"
	name := "assets/js/lib/require.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsToolsRJs reads file data from disk. It returns an error on failure.
func assetsJsToolsRJs() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/assets/js/tools/r.js"
	name := "assets/js/tools/r.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/css/bootstrap-combined.min.css": assetsCssBootstrapCombinedMinCss,
	"assets/css/bootstrap-theme.css": assetsCssBootstrapThemeCss,
	"assets/css/bootstrap-theme.css.map": assetsCssBootstrapThemeCssMap,
	"assets/css/bootstrap-theme.min.css": assetsCssBootstrapThemeMinCss,
	"assets/css/bootstrap.css": assetsCssBootstrapCss,
	"assets/css/bootstrap.css.map": assetsCssBootstrapCssMap,
	"assets/css/bootstrap.min.css": assetsCssBootstrapMinCss,
	"assets/css/jquery.treetable.css": assetsCssJqueryTreetableCss,
	"assets/css/jquery.treetable.theme.default.css": assetsCssJqueryTreetableThemeDefaultCss,
	"assets/fonts/glyphicons-halflings-regular.eot": assetsFontsGlyphiconsHalflingsRegularEot,
	"assets/fonts/glyphicons-halflings-regular.svg": assetsFontsGlyphiconsHalflingsRegularSvg,
	"assets/fonts/glyphicons-halflings-regular.ttf": assetsFontsGlyphiconsHalflingsRegularTtf,
	"assets/fonts/glyphicons-halflings-regular.woff": assetsFontsGlyphiconsHalflingsRegularWoff,
	"assets/fonts/glyphicons-halflings-regular.woff2": assetsFontsGlyphiconsHalflingsRegularWoff2,
	"assets/img/favicon.ico": assetsImgFaviconIco,
	"assets/img/icon_brand.png": assetsImgIcon_brandPng,
	"assets/js/app/jwt/jwt.js": assetsJsAppJwtJwtJs,
	"assets/js/app/treetable/treetable-page.js": assetsJsAppTreetableTreetablePageJs,
	"assets/js/lib/bootstrap/bootstrap.js": assetsJsLibBootstrapBootstrapJs,
	"assets/js/lib/bootstrap/bootstrap.min.js": assetsJsLibBootstrapBootstrapMinJs,
	"assets/js/lib/bootstrap/npm.js": assetsJsLibBootstrapNpmJs,
	"assets/js/lib/jquery/jquery-2.1.3.js": assetsJsLibJqueryJquery213Js,
	"assets/js/lib/jquery/jquery-2.1.3.min.js": assetsJsLibJqueryJquery213MinJs,
	"assets/js/lib/jquery-form/jquery.form.3.51.js": assetsJsLibJqueryFormJqueryForm351Js,
	"assets/js/lib/jquery-treetable/jquery.treetable.js": assetsJsLibJqueryTreetableJqueryTreetableJs,
	"assets/js/lib/require-setup.js": assetsJsLibRequireSetupJs,
	"assets/js/lib/require.js": assetsJsLibRequireJs,
	"assets/js/tools/r.js": assetsJsToolsRJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"bootstrap-combined.min.css": &bintree{assetsCssBootstrapCombinedMinCss, map[string]*bintree{}},
			"bootstrap-theme.css": &bintree{assetsCssBootstrapThemeCss, map[string]*bintree{}},
			"bootstrap-theme.css.map": &bintree{assetsCssBootstrapThemeCssMap, map[string]*bintree{}},
			"bootstrap-theme.min.css": &bintree{assetsCssBootstrapThemeMinCss, map[string]*bintree{}},
			"bootstrap.css": &bintree{assetsCssBootstrapCss, map[string]*bintree{}},
			"bootstrap.css.map": &bintree{assetsCssBootstrapCssMap, map[string]*bintree{}},
			"bootstrap.min.css": &bintree{assetsCssBootstrapMinCss, map[string]*bintree{}},
			"jquery.treetable.css": &bintree{assetsCssJqueryTreetableCss, map[string]*bintree{}},
			"jquery.treetable.theme.default.css": &bintree{assetsCssJqueryTreetableThemeDefaultCss, map[string]*bintree{}},
		}},
		"fonts": &bintree{nil, map[string]*bintree{
			"glyphicons-halflings-regular.eot": &bintree{assetsFontsGlyphiconsHalflingsRegularEot, map[string]*bintree{}},
			"glyphicons-halflings-regular.svg": &bintree{assetsFontsGlyphiconsHalflingsRegularSvg, map[string]*bintree{}},
			"glyphicons-halflings-regular.ttf": &bintree{assetsFontsGlyphiconsHalflingsRegularTtf, map[string]*bintree{}},
			"glyphicons-halflings-regular.woff": &bintree{assetsFontsGlyphiconsHalflingsRegularWoff, map[string]*bintree{}},
			"glyphicons-halflings-regular.woff2": &bintree{assetsFontsGlyphiconsHalflingsRegularWoff2, map[string]*bintree{}},
		}},
		"img": &bintree{nil, map[string]*bintree{
			"favicon.ico": &bintree{assetsImgFaviconIco, map[string]*bintree{}},
			"icon_brand.png": &bintree{assetsImgIcon_brandPng, map[string]*bintree{}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"app": &bintree{nil, map[string]*bintree{
				"jwt": &bintree{nil, map[string]*bintree{
					"jwt.js": &bintree{assetsJsAppJwtJwtJs, map[string]*bintree{}},
				}},
				"treetable": &bintree{nil, map[string]*bintree{
					"treetable-page.js": &bintree{assetsJsAppTreetableTreetablePageJs, map[string]*bintree{}},
				}},
			}},
			"lib": &bintree{nil, map[string]*bintree{
				"bootstrap": &bintree{nil, map[string]*bintree{
					"bootstrap.js": &bintree{assetsJsLibBootstrapBootstrapJs, map[string]*bintree{}},
					"bootstrap.min.js": &bintree{assetsJsLibBootstrapBootstrapMinJs, map[string]*bintree{}},
					"npm.js": &bintree{assetsJsLibBootstrapNpmJs, map[string]*bintree{}},
				}},
				"jquery": &bintree{nil, map[string]*bintree{
					"jquery-2.1.3.js": &bintree{assetsJsLibJqueryJquery213Js, map[string]*bintree{}},
					"jquery-2.1.3.min.js": &bintree{assetsJsLibJqueryJquery213MinJs, map[string]*bintree{}},
				}},
				"jquery-form": &bintree{nil, map[string]*bintree{
					"jquery.form.3.51.js": &bintree{assetsJsLibJqueryFormJqueryForm351Js, map[string]*bintree{}},
				}},
				"jquery-treetable": &bintree{nil, map[string]*bintree{
					"jquery.treetable.js": &bintree{assetsJsLibJqueryTreetableJqueryTreetableJs, map[string]*bintree{}},
				}},
				"require-setup.js": &bintree{assetsJsLibRequireSetupJs, map[string]*bintree{}},
				"require.js": &bintree{assetsJsLibRequireJs, map[string]*bintree{}},
			}},
			"tools": &bintree{nil, map[string]*bintree{
				"r.js": &bintree{assetsJsToolsRJs, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

