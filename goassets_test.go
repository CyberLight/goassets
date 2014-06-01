package goassets_test

import (
	"github.com/CyberLight/goassets"
	"github.com/CyberLight/goassets/test_utils"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"strings"
	"testing"
)

const (
	CrLf = "\r\n"

	ScriptTagRegexp = "^<script src=\".*\\.js\\?ver=\\d+\"></script>$"
	CssTagRegexp    = "^<link type=\"text\\/css\" rel=\"stylesheet\" href=\".*\\.css\\?ver=\\d+\">$"
)

var (
	utils         = test_utils.NewTestUtils()
	jsAggregator  = goassets.NewJsAggregator()
	cssAggregator = goassets.NewCssAggregator()
)

func TestGoassetsJsInclude(t *testing.T) {
	Convey("Given 3 js files in assets directory", t, func() {
		utils.RemoveAll("static")
		utils.CreateFolder("static", t)
		utils.CreateFiles("static/file%v.js", 3, t)

		Convey("When I set folder name to IncludeJs", func() {
			jsAssetsHtml := goassets.NewGoAssets().IncludeJs("static")

			Convey("Then I should get mapped Html script tags for all 3 files", func() {
				arrayWithoutRightRn := strings.TrimRight(string(jsAssetsHtml), CrLf)
				actualJsLines := strings.Split(arrayWithoutRightRn, CrLf)
				So(len(actualJsLines), ShouldEqual, 3)
				for i := 0; i < len(actualJsLines); i++ {
					matched, err := regexp.MatchString(ScriptTagRegexp, actualJsLines[i])
					if err != nil {
						t.Fatalf("Error %v", err)
					}
					So(true, ShouldEqual, matched)
				}
			})
		})

		Reset(func() {
			utils.RemoveAll("static")
		})
	})

	Convey("Given one js file inside some directory", t, func() {
		utils.CreateFiles("file%v.js", 1, t)
		Convey("When I set file path to JsInclude", func() {
			jsAssetsHtml := goassets.NewGoAssets().IncludeJs("file1.js")

			Convey("Then I should get mapped Html script tag for one file", func() {
				arrayWithoutRightRn := strings.TrimRight(string(jsAssetsHtml), CrLf)
				actualJsLines := strings.Split(arrayWithoutRightRn, CrLf)
				So(len(actualJsLines), ShouldEqual, 1)
				matched, err := regexp.MatchString(ScriptTagRegexp, actualJsLines[0])
				if err != nil {
					t.Fatalf("Error %v", err)
				}
				So(true, ShouldEqual, matched)
			})
		})

		Reset(func() {
			utils.RemoveAll("file1.js")
		})
	})

	Convey("Given not exists assets folder or file name", t, func() {
		Convey("When I set not exists folder or file name to IncludeJs", func() {
			Convey("Then I should get error information", func() {
				So(func() { goassets.NewGoAssets().IncludeJs("NotExistsFolder") }, ShouldPanic)
			})
		})
	})
}

func TestGoassetsCssInclude(t *testing.T) {
	Convey("Given 3 css files in assets directory", t, func() {
		utils.RemoveAll("static")
		utils.CreateFolder("static", t)
		utils.CreateFiles("static/style%v.css", 3, t)

		Convey("When I set folder name to IncludeCss", func() {
			cssAssetsHtml := goassets.NewGoAssets().IncludeCss("static")

			Convey("Then I should get mapped Html rel link tags for all 3 css files", func() {
				arrayWithoutRightRn := strings.TrimRight(string(cssAssetsHtml), CrLf)
				actualCssLines := strings.Split(arrayWithoutRightRn, CrLf)
				So(len(actualCssLines), ShouldEqual, 3)
				for i := 0; i < len(actualCssLines); i++ {
					matched, err := regexp.MatchString(CssTagRegexp, actualCssLines[i])
					if err != nil {
						t.Fatalf("Error %v", err)
					}
					So(true, ShouldEqual, matched)
				}
			})
		})

		Reset(func() {
			utils.RemoveAll("static")
		})
	})

	Convey("Given 1 css file", t, func() {
		utils.RemoveAll("style1.css")
		utils.CreateFiles("style%v.css", 1, t)

		Convey("When I set css file name to IncludeCss", func() {
			cssAssetsHtml := goassets.NewGoAssets().IncludeCss("style1.css")

			Convey("Then I should get mapped Html rel link tags for one css file", func() {
				arrayWithoutRightRn := strings.TrimRight(string(cssAssetsHtml), CrLf)
				actualCssLines := strings.Split(arrayWithoutRightRn, CrLf)
				matched, err := regexp.MatchString(CssTagRegexp, actualCssLines[0])
				if err != nil {
					t.Fatalf("Error %v", err)
				}

				So(len(actualCssLines), ShouldEqual, 1)
				So(true, ShouldEqual, matched)
			})
		})

		Reset(func() {
			utils.RemoveAll("style1.css")
		})
	})

	Convey("Given not exists assets folder or file name", t, func() {
		Convey("When I set not exists folder or file name to IncludeCss", func() {
			Convey("Then I should get error information", func() {
				So(func() { goassets.NewGoAssets().IncludeCss("NotExistsFolder") }, ShouldPanic)
			})
		})
	})
}

func TestGoassetsJsAggregator(t *testing.T) {
	fileName := "asset1.js"

	Convey("Given one js assets file", t, func() {
		utils.RemoveAll(fileName)
		utils.CreateFiles("asset%v.js", 1, t)
		Convey("When I try to aggregate js files with wrong regex", func() {
			Convey("Then I should get error information", func() {
				aggregator := goassets.NewAssetAggregator("*.\\.js$", goassets.DefaultScriptTemplate)
				So(func() { aggregator.Aggregate(fileName) }, ShouldPanic)
			})
		})

		Reset(func() {
			utils.RemoveAll(fileName)
		})
	})

	Convey("Given one css file", t, func() {
		utils.RemoveAll(fileName)
		utils.CreateFiles("style%v.css", 1, t)

		Convey("When I try to aggregate css file instead of js", func() {
			Convey("Then I should get empty aggregated data", func() {
				data, _ := jsAggregator.Aggregate("style1.css")
				So("", ShouldEqual, data)
			})
		})

		Reset(func() {
			utils.RemoveAll("style1.css")
		})
	})
}

func TestGoassetsAssetCssAggregator(t *testing.T) {
	Convey("Given one css file", t, func() {
		utils.CreateFiles("teststyle%v.css", 1, t)
		Convey("When I try to aggregate css files with wrong regex", func() {
			Convey("Then I should get error information", func() {
				aggregator := goassets.NewAssetAggregator("*.\\.css$", goassets.DefaultCssTemplate)
				So(func() { aggregator.Aggregate("teststyle1.css") }, ShouldPanic)
			})
		})

		Reset(func() {
			utils.RemoveAll("teststyle1.css")
		})
	})

	Convey("Given one js file instead of css", t, func() {
		utils.RemoveAll("file1.js")
		utils.CreateFiles("file%v.js", 1, t)

		Convey("When I try to aggregate css file instead of js", func() {
			Convey("Then I should get empty aggregated data", func() {
				data, _ := cssAggregator.Aggregate("file1.js")
				So("", ShouldEqual, data)
			})
		})

		Reset(func() {
			utils.RemoveAll("file1.js")
		})
	})

}
