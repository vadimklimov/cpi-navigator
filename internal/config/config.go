package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Config struct {
	Tenant *Tenant `mapstructure:"tenant"`
	UI     *UI     `mapstructure:"ui"`
}

type Tenant struct {
	Name         string   `mapstructure:"name"`
	WebUIURL     *url.URL `mapstructure:"webui_url"`
	BaseURL      *url.URL `mapstructure:"base_url"`
	TokenURL     *url.URL `mapstructure:"token_url"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
}

type UI struct {
	Layout Layout `mapstructure:"layout"`
	Panes  Panes  `mapstructure:",squash"`
}

type Layout string

type Panes struct {
	Packages  PackagesPane  `mapstructure:"packages_pane"`
	Artifacts ArtifactsPane `mapstructure:"artifacts_pane"`
}

type PackagesPane struct {
	Sort Sort `mapstructure:",squash"`
}

type ArtifactsPane struct {
	Sort Sort `mapstructure:",squash"`
}

type Sort struct {
	Field string    `mapstructure:"sort_field"`
	Order SortOrder `mapstructure:"sort_order"`
}

type SortOrder string

const (
	LayoutNormal  Layout = "normal"
	LayoutCompact Layout = "compact"
)

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

const (
	DefaultUserConfigDir  = ".config"
	DefaultAppConfigDir   = "cpi-navigator"
	DefaultConfigFileName = "config"
	DefaultConfigFileExt  = "yaml"
)

var cfg *Config

func Init(configFile string) {
	cfg = &Config{
		Tenant: &Tenant{},
		UI:     &UI{},
	}

	if err := cfg.load(configFile); err != nil {
		log.Fatal("Unable to load configuration", "err", err)
	}

	if err := cfg.checkMandatory(); err != nil {
		log.Fatal("Mandatory configuration parameters were not provided", "err", err)
	}

	cfg.setDefaults()
}

func TenantName() string {
	return cfg.Tenant.Name
}

func TenantWebUIURL() *url.URL {
	return cfg.Tenant.WebUIURL
}

func TenantBaseURL() *url.URL {
	return cfg.Tenant.BaseURL
}

func TenantTokenURL() *url.URL {
	return cfg.Tenant.TokenURL
}

func TenantClientID() string {
	return cfg.Tenant.ClientID
}

func TenantClientSecret() string {
	return cfg.Tenant.ClientSecret
}

func UILayout() Layout {
	return cfg.UI.Layout
}

func UIPackagesPaneSortField() string {
	return cfg.UI.Panes.Packages.Sort.Field
}

func UIPackagesPaneSortOrder() SortOrder {
	return cfg.UI.Panes.Packages.Sort.Order
}

func UIArtifactsPaneSortField() string {
	return cfg.UI.Panes.Artifacts.Sort.Field
}

func UIArtifactsPaneSortOrder() SortOrder {
	return cfg.UI.Panes.Artifacts.Sort.Order
}

func (c *Config) load(configFile string) error {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		workDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error determining current (working) directory: %w", err)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error determining user's home directory: %w", err)
		}

		viper.AddConfigPath(workDir)
		viper.AddConfigPath(filepath.Join(homeDir, DefaultUserConfigDir, DefaultAppConfigDir))
		viper.SetConfigName(DefaultConfigFileName)
		viper.SetConfigType(DefaultConfigFileExt)
	}

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configuration: %w", err)
	}

	if err := viper.Unmarshal(&c, viper.DecodeHook(composeDecodeHook())); err != nil {
		return fmt.Errorf("error unmarshalling configuration: %w", err)
	}

	return nil
}

func (c *Config) checkMandatory() error {
	missingConfigParams := make([]string, 0)

	if c.Tenant.WebUIURL == nil {
		missingConfigParams = append(missingConfigParams, "tenant.webui_url")
	}

	if c.Tenant.BaseURL == nil {
		missingConfigParams = append(missingConfigParams, "tenant.base_url")
	}

	if c.Tenant.TokenURL == nil {
		missingConfigParams = append(missingConfigParams, "tenant.token_url")
	}

	if c.Tenant.ClientID == "" {
		missingConfigParams = append(missingConfigParams, "tenant.client_id")
	}

	if c.Tenant.ClientSecret == "" {
		missingConfigParams = append(missingConfigParams, "tenant.client_secret")
	}

	if len(missingConfigParams) > 0 {
		return fmt.Errorf("missing parameters: %s",
			strings.Join(missingConfigParams, ", "),
		)
	}

	return nil
}

func (c *Config) setDefaults() {
	// Set tenant name.
	if c.Tenant.Name == "" {
		c.Tenant.Name = strings.Split(c.Tenant.WebUIURL.Hostname(), ".")[0]
	}

	// Set UI layout.
	layout := Layout(strings.ToLower(string(c.UI.Layout)))

	switch layout {
	case LayoutNormal, LayoutCompact:
		c.UI.Layout = layout
	default:
		c.UI.Layout = LayoutNormal
	}

	// Set content packages sort field.
	if c.UI.Panes.Packages.Sort.Field == "" {
		c.UI.Panes.Packages.Sort.Field = "ID"
	}

	// Set content packages sort order.
	packagesSortOrder := SortOrder(strings.ToLower(string(c.UI.Panes.Packages.Sort.Order)))

	switch packagesSortOrder {
	case SortOrderAscending, SortOrderDescending:
		c.UI.Panes.Packages.Sort.Order = packagesSortOrder
	default:
		c.UI.Panes.Packages.Sort.Order = SortOrderAscending
	}

	// Set integration artifacts sort field.
	if c.UI.Panes.Artifacts.Sort.Field == "" {
		c.UI.Panes.Artifacts.Sort.Field = "ID"
	}

	// Set integration artifacts sort order.
	artifactsSortOrder := SortOrder(strings.ToLower(string(c.UI.Panes.Artifacts.Sort.Order)))

	switch artifactsSortOrder {
	case SortOrderAscending, SortOrderDescending:
		c.UI.Panes.Artifacts.Sort.Order = artifactsSortOrder
	default:
		c.UI.Panes.Artifacts.Sort.Order = SortOrderAscending
	}
}
