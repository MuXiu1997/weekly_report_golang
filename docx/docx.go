package docx

import (
	"flag"
	"github.com/unidoc/unioffice/document"
	"regexp"
)

var rule *regexp.Regexp

func init() {
	//github.com/unidoc/unioffice 需要 license
	//使用 名为 test.v 的 flag , 可使此库 不对输出文件加水印
	var t string
	flag.StringVar(&t, "test.v", "", "")
	flag.Parse()
	_ = t

	rule = regexp.MustCompile("^{([A-Za-z_]+)}$")
}

func New(templateFileName, exportFileName string) *Renderer {
	r := new(Renderer)
	r.TemplateFileName = templateFileName
	r.ExportFileName = exportFileName
	return r
}

type Renderer struct {
	TemplateFileName string
	ExportFileName   string
	Map              map[string]string
}

func (dr *Renderer) Render() error {
	d, err := document.Open(dr.TemplateFileName)
	if err != nil {
		return err
	}

	for _, p := range d.Paragraphs() {
		for _, r := range p.Runs() {
			t := r.Text()
			ss := rule.FindStringSubmatch(t)
			if len(ss) == 0 {
				continue
			}
			k := ss[1]
			v, ok := dr.Map[k]
			if ok {
				r.ClearContent()
				r.AddText(v)
			}
		}
	}

	err = d.SaveToFile(dr.ExportFileName)
	if err != nil {
		return err
	}
	return nil
}

func (dr *Renderer) Set(k, v string) {
	if dr.Map == nil {
		dr.Map = map[string]string{}
	}
	dr.Map[k] = v
}
