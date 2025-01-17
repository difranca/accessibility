package report

import (
	"html"
	"strings"

	"github.com/inclunet/accessibility/pkg/accessibility"
)

type AccessibilityReport struct {
	Checks         []accessibility.AccessibilityCheck
	Domain         string
	Html           string
	HtmlReportPath string
	JsonReportPath string
	ReportFile     string
	Summary        map[string]AccessibilitySummary
	Title          string
	Url            string
}

func (r *AccessibilityReport) AddCheck(accessibilityCheck accessibility.AccessibilityCheck) {
	r.Checks = append(r.Checks, r.GetLineNumber(accessibilityCheck))
	r.UpdateSummary(accessibilityCheck)
}

func (r *AccessibilityReport) UpdateSummary(accessibilityCheck accessibility.AccessibilityCheck) {
	total, _ := r.Summary["total"]
	total.Update(accessibilityCheck)
	r.Summary["total"] = total
	Summary, _ := r.Summary[accessibilityCheck.Element]
	Summary.Update(accessibilityCheck)
	r.Summary[accessibilityCheck.Element] = Summary
}

func (r *AccessibilityReport) GenerateSummary() {
	for _, accessibilityCheck := range r.Checks {
		r.UpdateSummary(accessibilityCheck)
	}
}

func (r *AccessibilityReport) GetLineNumber(accessibilityCheck accessibility.AccessibilityCheck) accessibility.AccessibilityCheck {
	text := html.UnescapeString(accessibilityCheck.Text)
	offset := strings.Index(r.Html, text)

	if offset > 0 {
		accessibilityCheck.Line = strings.Count(r.Html[:offset], "\n")
	}

	return accessibilityCheck
}

func (r *AccessibilityReport) NewSummary() {
	r.Summary = make(map[string]AccessibilitySummary)
}

func NewAccessibilityReport(url string, reportFile string, lang string, reportPath string, title string) AccessibilityReport {
	return AccessibilityReport{
		Url:        url,
		Title:      title,
		ReportFile: reportFile,
		Summary:    make(map[string]AccessibilitySummary),
	}
}
