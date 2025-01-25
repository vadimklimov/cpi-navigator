# Cloud Integration Navigator

[![Latest release](https://img.shields.io/github/release/vadimklimov/cpi-navigator)](https://github.com/vadimklimov/cpi-navigator/releases/latest)
[![Codacy](https://app.codacy.com/project/badge/Grade/95069cb0a3f24e43a91978cd1d5bbeb2)](https://app.codacy.com/gh/vadimklimov/cpi-navigator/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![License](https://img.shields.io/github/license/vadimklimov/cpi-navigator)](https://opensource.org/license/MIT)

_Cloud Integration Navigator_, or _CPI Navigator_ for short, is a terminal-based application to browse and explore content packages and integration artifacts within an SAP Cloud Integration (a part of SAP Integration Suite) tenant's workspace.

![cpi-navigator](https://github.com/vadimklimov/cpi-navigator/assets/14906963/fbff81f8-56d6-4849-8aea-4d4b98a6dc25)

## Requirements

### SAP Cloud Integration

CPI Navigator makes use of open (public) APIs of SAP Cloud Integration to fetch the required information about content packages and integration artifacts - namely, [Integration Packages (Design)](https://api.sap.com/api/IntegrationContent/resource/Integration_Packages_Design) of the Integration Content API. The API is OAuth-protected and supports the client credentials flow. This authentication mechanism is employed by CPI Navigator to authenticate and authorize API calls made to the SAP Cloud Integration tenant.

To enable CPI Navigator to access the necessary APIs of the SAP Cloud Integration tenant, it is necessary to create an OAuth client for it in the SAP Business Technology Platform subaccount where the corresponding subscription for SAP Integration Suite has been created. The steps for creating an OAuth client for SAP Cloud Integration vary depending on the environment, whether it is Neo or Cloud Foundry.

> [!IMPORTANT]
> CPI Navigator has been tested with SAP Cloud Integration provisioned in a Cloud Foundry environment. Use it with SAP Cloud Integration provisioned in a Neo environment at your own risk.
>
> In alignment with SAP's strategic direction to phase out support for a Neo environment in the SAP Business Technology Platform over the long term and to focus future innovations on the multi-cloud foundation (see [SAP Note 3351844](https://me.sap.com/notes/3351844)), it is important to note that there are no plans to enable CPI Navigator to support functionalities exclusive to SAP Cloud Integration in a Neo environment.

#### SAP Cloud Integration in Cloud Foundry environment

In a Cloud Foundry environment, a service instance represents an OAuth client - hence, a service instance and a service instance key for it must be created.

1. Create a service instance for the `Process Integration Runtime` service using the `api` plan. The `WorkspacePackagesRead` role and the `Client Credentials` grant type must be selected when configuring service instance parameters.

2. Create a service instance key of the `ClientId/Secret` type for the above-mentioned service instance.

> [!NOTE]
> For information about steps involved in service instance and service instance key creation for the Process Integration Runtime service, refer to the [SAP Help documentation](https://help.sap.com/docs/integration-suite/sap-integration-suite/creating-service-instance-and-service-key-for-inbound-authentication).
> For further details about using an OAuth client credentials grant when calling APIs of SAP Cloud Integration in a Cloud Foundry environment, refer to the [SAP Help documentation](https://help.sap.com/docs/integration-suite/sap-integration-suite/oauth-with-client-credentials-grant-for-api-clients).

#### SAP Cloud Integration in Neo environment

1. Register an OAuth client for the application of the tenant management node of SAP Cloud Integration (the subscription name ends with `tmn`) using the `Client Credentials` authorization grant.

2. Assign the user with name `oauth_client_<client ID>` to the `WebToolingWorkspace.Read` role for the application of the tenant management node of SAP Cloud Integration (the application name ends with `tmn`).

> [!NOTE]
> For further details about using an OAuth client credentials grant when calling APIs of SAP Cloud Integration in a Neo environment, refer to the [SAP Help documentation](https://help.sap.com/docs/cloud-integration/sap-cloud-integration/setting-up-oauth-inbound-authentication-with-client-credentials-grant-for-api-clients).

### Terminal

The terminal application must support True Color.

> [!CAUTION]
> Failure to ensure True Color compatibility may result in discrepancies in colour rendering and distorted colours, compromising the visual experience.

### Networking

The machine where CPI Navigator is run, must have Internet connection to the corresponding SAP Cloud Integration tenant to enable the application to access the necessary APIs.

Connections from the CPI Navigator to the SAP Cloud Integration tenant are HTTP-based and TLS-secured.

> [!NOTE]
> For information about IP addresses that are associated with SAP Cloud Integration, see [SAP Note 2808441](https://me.sap.com/notes/2808441).
> For further details, refer to the following SAP Help documentation:
>
> - [IP addresses information](https://help.sap.com/docs/btp/sap-business-technology-platform/regions-and-api-endpoints-available-for-cloud-foundry-environment) for tenants provisioned in a **Cloud Foundry** environment,
> - [IP addresses information](https://help.sap.com/docs/btp/sap-btp-neo-environment/regions-and-hosts-available-for-neo-environment) for tenants provisioned in a **Neo** environment.
>
> IP addresses associated with SAP Cloud Integration are specific to the infrastructure provider and the region where the corresponding SAP Cloud Integration tenant is provisioned. This principle also extends to the broader SAP Integration Suite and other services within the SAP Business Technology Platform.

## Installation

### Download binary

Download a binary compatible with your operating system / architecture from the [Releases](https://github.com/vadimklimov/cpi-navigator/releases) page.

### Install from source

Use the `go install` command to compile and install the application:

`go install github.com/vadimklimov/cpi-navigator@latest`

> [!IMPORTANT]
> Go must be installed on the machine.
> Official binary distributions are available at the [Go downloads area](https://go.dev/dl/).
> For installation instructions, refer to the [Go documentation](https://go.dev/doc/install).

## Configuration

For CPI Navigator to function correctly, configuration parameters must be provided in the application's configuration file using the YAML syntax.

### Configuration file location

By default, CPI Navigator searches for the configuration file named `config.yaml` or `config.yml` in the current directory and in the `.config/cpi-navigator` directory in the user's home directory.

The default location and the file name can be overwritten using the `--config` flag and providing a custom location of the configuration file.

### Configuration parameters

The following configuration parameters are supported:

The `tenant` configuration section.

| Parameter     | Description                                                                                                                                    |
| ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| webui_url     | WebUI URL (SAP Integration Suite home page URL)                                                                                                |
| base_url      | Base URL. _In a Cloud Foundry environment, can be found in the service instance key: the `url` attribute in the `oauth` section_               |
| token_url     | Token URL. _In a Cloud Foundry environment, can be found in the service instance key: the `tokenurl` attribute in the `oauth` section_         |
| client_id     | Client ID. _In a Cloud Foundry environment, can be found in the service instance key: the `clientid` attribute in the `oauth` section_         |
| client_secret | Client secret. _In a Cloud Foundry environment, can be found in the service instance key: the `clientsecret` attribute in the `oauth` section_ |
| name          | _(optional)_ Tenant name (alias) to be displayed in the status bar. If not provided, the tenant's subdomain is used                            |

The `ui` configuration section.

| Parameter | Description                                                                                    |
| --------- | ---------------------------------------------------------------------------------------------- |
| layout    | _(optional)_ Layout. Valid values: `normal` (default), `compact` (no title bar and status bar) |

The `ui` configuration section supports the following subsections for pane customization:

- `packages_pane`: configures the content packages pane
- `artifacts_pane`: configures the integration artifacts pane

Each pane subsection supports the following parameters:

| Parameter  | Description                                                                                    |
| ---------- | ---------------------------------------------------------------------------------------------- |
| sort_field | _(optional)_ Sort field. Refer to [Sort fields](#sort-fields) for the list of supported fields |
| sort_order | _(optional)_ Sort order. Valid values: `asc` (ascending) (default), `desc` (descending)        |

#### Sort fields

- Content packages pane: `ID` (default), `Version`, `Name`, `ShortText`, `Description`, `Vendor`, `PartnerContent`, `Mode`, `UpdateAvailable`, `SupportedPlatform`, `Products`, `Keywords`, `Countries`, `Industries`, `LineOfBusiness`, `ResourceID`, `CreatedBy`, `CreationDate`, `ModifiedBy`, `ModifiedDate`
- Integration artifacts pane: `ID` (default), `Version`, `PackageID`, `Name`, `Description`, `CreatedBy`, `CreatedAt`, `ModifiedBy`, `ModifiedAt`

### Examples

Below are examples of a `config.yaml` file.

Minimum configuration (required parameters only):

```yaml
tenant:
  webui_url: https://<subdomain>.integrationsuite.cfapps.<region>.hana.ondemand.com
  base_url: https://<subdomain>.it-cpi<xxxxx>.cfapps.<region>.hana.ondemand.com/api/v1
  token_url: https://<subdomain>.authentication.<region>.hana.ondemand.com/oauth/token
  client_id: xxxxxxxxxx
  client_secret: xxxxxxxxxx
```

Full configuration (all supported parameters):

```yaml
tenant:
  name: Sandbox
  webui_url: https://<subdomain>.integrationsuite.cfapps.<region>.hana.ondemand.com
  base_url: https://<subdomain>.it-cpi<xxxxx>.cfapps.<region>.hana.ondemand.com/api/v1
  token_url: https://<subdomain>.authentication.<region>.hana.ondemand.com/oauth/token
  client_id: xxxxxxxxxx
  client_secret: xxxxxxxxxx
ui:
  layout: normal
  packages_pane:
    sort_field: Name
    sort_order: asc
  artifacts_pane:
    sort_field: ModifiedAt
    sort_order: desc
```

## Usage

Run `cpi-navigator` in your terminal to start CPI Navigator.

### Command flags

The following command flags are supported:

| Long flag   | Short flag | Description                     | Possible values                 | Default value                                      |
| ----------- | ---------- | ------------------------------- | ------------------------------- | -------------------------------------------------- |
| --config    | -c         | Set configuration file location | _/path/to/config.yaml_          | ./config.yaml, ~/.config/cpi-navigator/config.yaml |
| --log-level | -l         | Set log level                   | debug, info, warn, error, fatal | info                                               |
| --version   | -v         | Show version information        |                                 |                                                    |
| --help      | -h         | Show help information           |                                 |                                                    |

### Key bindings

The following key bindings are supported:

| Key binding  | Description                                                                             |
| ------------ | --------------------------------------------------------------------------------------- |
| Tab          | Switch an active pane (switch between content packages and integration artifacts panes) |
| Enter        | Display integration artifacts in the selected content package                           |
| ↑ / ↓        | Navigate to the previous/next item within the active pane                               |
| ← / →        | Navigate to the previous/next tab in the integration artifacts pane                     |
| q / Ctrl + C | Quit the application                                                                    |
| l            | Toggle layout (switch between normal and compact layouts)                               |
| r            | Refresh items in the active pane                                                        |
| o            | Open the selected content package or integration artifact in Web UI                     |

## Notes

### Window size

The application window uses fixed width and height, and doesn't support window resizing yet. For the optimal visual experience, the following minimum terminal window size is recommended:

- Width: 155 characters
- Height: 50 characters

> [!IMPORTANT]
> The terminal window size is defined in characters, not pixels.

### Colour themes

The current version of CPI Navigator uses the Mocha flavor (colour palette) of the [Catppuccin](https://catppuccin.com) theme and doesn't allow customizing of the active colour theme or selection of an alternative colour theme yet.
