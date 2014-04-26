package goassets

import (
	"html/template"
)

type GoAssets struct {
	jsAggregator IAggregator
	cssAggregator IAggregator
}

func NewGoAssets(jsAggregator IAggregator, cssAggregator IAggregator) *GoAssets {
	goAssets := &GoAssets{}
	goAssets.Init(jsAggregator, cssAggregator)
	return goAssets
}

func (this *GoAssets) Init(jsAggregator IAggregator, cssAggregator IAggregator){
	this.jsAggregator = jsAggregator
	this.cssAggregator = cssAggregator
}

func (this *GoAssets) IncludeJs(nameOrPath string) template.HTML {
	scripts, err := this.jsAggregator.Aggregate(nameOrPath)
	
	if err != nil { 
		panic(err) 
	}
	
	return template.HTML(scripts) 
}

func (this *GoAssets) IncludeCss(nameOrPath string) template.HTML {
	scripts, err := this.cssAggregator.Aggregate(nameOrPath)
	
	if err != nil { 
		panic(err) 
	}
	
	return template.HTML(scripts) 
}
