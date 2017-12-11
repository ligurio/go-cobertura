package cobertura

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type CoberturaReport struct {
	XMLName         xml.Name  `xml:"coverage,attr,omitempty"`
	LinesValid      int       `xml:"lines-valid,attr,omitempty"`
	LinesCovered    int       `xml:"lines-covered,attr,omitempty"`
	LineRate        int       `xml:"line-rate,attr,omitempty"`
	BranchesValid   int       `xml:"branches-valid,attr,omitempty"`
	BranchesCovered int       `xml:"branches-covered,attr,omitempty"`
	Timestamp       int       `xml:"timstamp,attr,omitempty"`
	Complexity      int       `xml:complexity,attr,omitempty`
	Version         int       `xml:version,attr,omitempty`
	//Sources         []Source  `xml:"source"`
	//Packages        []Package `xml:"package"`
}

type Source struct {
	XMLName xml.Name    `xml:"source"`
	Source  InnerResult `xml:"source,omitempty"`
}

type Package struct {
	XMLName    xml.Name `xml:"package"`
	Name       int      `xml:"name,attr,omitempty"`
	LineRate   int      `xml:"line-rate,attr,omitempty"`
	BranchRate int      `xml:"branch-rate,attr,omitempty"`
	Classes    []Class  `xml:"class"`
}

type Class struct {
	XMLName xml.Name `xml:"package"`
	// TODO: class properties
	Name       string   `xml:"attr,name,omitempty"`
	Filename   string   `xml:"attr,filename,omitempty"`
	LineRate   int      `xml:"attr,line-rate,omitempty"`
	BranchRate int      `xml:"attr,branch-rate,omitempty"`
	Methods    []Method `xml:"method"`
	Lines      []Line   `xml:"line"`
}

type Line struct {
	XMLName xml.Name `xml:"property"`
	Number  int      `xml:"attr,number"`
	Hits    int      `xml:"attr,hits"`
	Branch  bool     `xml:"attr,branch"`
}

type Method struct {
	XMLName   xml.Name `xml:"property"`
	Name      string   `xml:"attr,name"`
	Hits      int      `xml:"attr,name"`
	Signature string   `xml:"attr,name"`
	Lines     []Line   `xml:"lines"`
}

type InnerResult struct {
	Value string `xml:",innerxml"`
}

func NewParser(r io.Reader) (*CoberturaReport, error) {

	var report = new(CoberturaReport)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal([]byte(buf), &report)
	if err != nil {
		return nil, err
	}

	return report, nil
}
