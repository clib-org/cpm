package main

type ProjectBuilder struct {
	p    string
	conf PackageConfig
}

func NewProjectBuilder(p string, conf PackageConfig) ProjectBuilder {
	var ret ProjectBuilder

	ret.p = p
	ret.conf = conf

	return ret
}

func (p *ProjectBuilder) build() bool {
	println(":: Building Project ", p.p)

	config, ok := p.conf.data["projects"].(map[string]interface{})[p.p].(map[string]interface{})

	if !ok {
		panic("Project " + p.p + " not found")
	}

	pType := config["type"].(string)
	switch pType {

	case "binary":

	case "library":

	case "test":

	case "external-test":

	default:
		panic("Unknown project type: " + pType)
	}

	return true
}
