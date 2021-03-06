package ui

import "github.com/flanksource/karina-ui/pkg/api"

type Cluster struct {
	ID           int                   `json:"id"`
	Name         string                `json:"name"`
	ItemIcons    []ItemIcon            `json:"itemicons"`
	Properties   []Property            `json:"properties"`
	Indicators   []Indicator           `json:"indicators"`
	CanaryChecks []api.Check           `json:"canaryChecks"`
	Metrics      map[string]api.Metric `json:"metrics"`
}

type ItemIcon struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
	Label string `json:"label"`
	Count int    `json:"count"`
}

type Property struct {
	ID         int    `json:"id"`
	Service    string `json:"service"`
	Icon       string `json:"icon"`
	Label      string `json:"label"`
	Badge      int    `json:"badge"`
	BadgeColor string `json:"badgecolor"`
}

type Indicator struct {
	ID             string          `json:"id"`
	IndicatorIcon  string          `json:"indicator_icon_lg"`
	IndicatorIcons []IndicatorIcon `json:"indicator_icons"`
}

type IndicatorIcon struct {
	ID      string `json:"id"`
	Color   string `json:"color"`
	Shape   string `json:"shape"`
	Overlay string `json:"overlay"`
}
