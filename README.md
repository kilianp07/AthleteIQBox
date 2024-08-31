<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="20%" alt="ATHLETEIQBOX-logo">
</p>
<p align="center">
    <h1 align="center">ATHLETEIQBOX</h1>
</p>
<p align="center">
    <em>Utility tests for resilience 2. Convert, decode with precision 3. JSON agility & compatibility 4. Sophisticated log control 5. Easy log output confg. 6. Tracks, blinks elegantly 7. Continuously regenerative workflow</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/kilianp07/AthleteIQBox?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/kilianp07/AthleteIQBox?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/kilianp07/AthleteIQBox?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/kilianp07/AthleteIQBox?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br>

#####  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
    - [ Prerequisites](#-prerequisites)
    - [ Installation](#-installation)
    - [ Usage](#-usage)
    - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

Converters, json, utils/converter and logs, which defineloggerconfiguration including logging levels. The athleeQBoxinit iscentral entrypoint integrating GPCS data, Bluetooth and interactive buttons tracking to start the servlce on enchnaced perrmoncemonitoring. Moreover,generate-changellodworkflows, located GitHub Workflowss dirvresifies self-updatino in the Git repsiotnyathleeQ Box chanclogs are autmatically committed for neve pull request.

---

##  Features



---

##  Repository Structure

```sh
└── AthleteIQBox/
    ├── .github
    │   └── workflows
    ├── CHANGELOG.md
    ├── box.json
    ├── buttons
    │   ├── button.go
    │   ├── configuration.go
    │   └── manager.go
    ├── cliff.toml
    ├── conf.go
    ├── data
    │   ├── position.go
    │   └── position_test.go
    ├── entrypoint
    │   └── entrypoint.go
    ├── envrc-example
    ├── go.mod
    ├── go.sum
    ├── gps
    │   ├── configuration.go
    │   ├── gps.go
    │   ├── reader
    │   └── recorder
    ├── main.go
    ├── makefile
    ├── renovate.json
    ├── tests
    │   └── json
    ├── transmitter
    │   ├── conf.go
    │   ├── factory.go
    │   ├── factory_test.go
    │   ├── services
    │   └── transmitter.go
    └── utils
        ├── converters.go
        ├── decoder.go
        ├── decoder_test.go
        ├── json.go
        ├── json_test.go
        └── logger
```

---

##  Modules

<details closed><summary>.</summary>

| File | Summary |
| --- | --- |
| [envrc-example](https://github.com/kilianp07/AthleteIQBox/blob/main/envrc-example) | Streamlines build processes by offering default environment configuration via envrc-example file, expediting setup tailored to targeted devices and operating environments, enhancing toolchain extensibility. |
| [go.mod](https://github.com/kilianp07/AthleteIQBox/blob/main/go.mod) | This go module (github.com/kilianp07/AthleteIQBox) forms the basis of the Athlete IQ projects architecture. It serves to manage external dependencies for the software such as payment, GPS handling, Wireless Communication libraries and more – enriching functionality throughout various modules in theRepository hierarchy. |
| [box.json](https://github.com/kilianp07/AthleteIQBox/blob/main/box.json) | Sets up Wi-Fi (devicename= AthleteIQBOX) and GPS functionalities for user customization-Defines GPS parameters such as reading input interval, log storage method etc. InitializesLogger captures essential inFo with a specific filter Assigns onboard buttons to specific SWitCheS (gpio2' starts, gpio3 Stops) |
| [renovate.json](https://github.com/kilianp07/AthleteIQBox/blob/main/renovate.json) | This `renovate.json` sets up auto-maintenance using Renovatebot to ensure the dependency status of AthleteIQBox remains compatible during its lifecycle. Key advantages include keeping libraries current with ease, optimised by extending suggested changes drawn from the official Schema. This proactive tact enforces development quality and reliability while conserving time commitments in nurturing the projects longevity and overall ecosystem performance. |
| [makefile](https://github.com/kilianp07/AthleteIQBox/blob/main/makefile) | This powerful makefile initiates binary compilation (`BINARY_NAME=athleteiqbox`` `MAIN_GO_PATH=./main.go... go build)`. Customizes the compile process for Raspberry Pi-ready application using Docker, CGO, and Pi Go Build Images. Simplifies deployment by automating artifact upload & scripts. Enables local testing to validate development effectively within the architecture stack using precise directives (run`test`) efficiently streamlines execution on target systems with single command deployments during development (run `raspberry:@ssh $(SSH_TARGET)...`), or forces removal after modification with an assured kill approach. (kill``)) The makefile also cleans up obsoletes binaries in the project progression loop (clean:`@rm-rfs...`) making this architecture highly versatile, adaptable to development pipeline enhancements while retaining precision in production deployment (docker-based builds are referenced implicitly). |
| [conf.go](https://github.com/kilianp07/AthleteIQBox/blob/main/conf.go) | Structurally configures key transmission parameters within AthleteIQBox project. Essentially houses the Conformance structure integrating directly with transmitter subsystem, parsable as JSON data. |
| [go.sum](https://github.com/kilianp07/AthleteIQBox/blob/main/go.sum) | Button.go` or a specific tracker). These services, working in consensus with configurable settings defined in various subdirctories like `configuration.go`, establish connections between AthleteIQBox server software and fitness trackers. Noteworthy is the presence of `management.go`, responsible for aggregating data and processing tasks related to workout data synchronization. The flow management takes care throughout entire functionality involving fetching insights about athlete performance progress. The configuration setup is centralized utilizing a single, easy-to-read `box.json`. This JSON provides core properties needed customizing API access permissions of the connected fitness wearables. Users can freely change or tune configurations if needed (`conf.go` manages configuration reading from the JSON file as necessary components inside main application logic). Lastly, this architecture employs `changetab.md`, allowing an overview about new functionality incorporated into distinct releases for convenient developer & user interactions while staying engaged and following updates more seamlessly.. |
| [cliff.toml](https://github.com/kilianp07/AthleteIQBox/blob/main/cliff.toml) | This code serves as `cliff.toml`configuration file for AthleteIQBox, aligning with git-cliff format to organize and document changes efficiently. The configuration sets up header, footer, templates for Git commit history & changelog footer, postprocessors like URL or linter replacement, and committee structuring per task category. Integrated as git repository convention with Conventional Commits style for consistent logging approach within the project. |
| [main.go](https://github.com/kilianp07/AthleteIQBox/blob/main/main.go) | The main.go file initiates execution within AthleteIQBox. It defines the entry point and sets up core arguments using cobra. With the provided--conf flag, it reads essential configuration for executing processes. Upon parsing flags and loading the config, it kicks off the central process via the entrypoint module, marking the start of software utilization in the complete pipeline within AthleteIQBox. |

</details>

<details closed><summary>buttons</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/buttons/configuration.go) | Manages input system setup in this sports training software architecture. Important components encapsulated within `buttons/configuration.go` comprise the definition of electronic switches that serve various functionalities through specifying GPIO pins corresponding to physical buttons. |
| [button.go](https://github.com/kilianp07/AthleteIQBox/blob/main/buttons/button.go) | This `Go` module within AthleteIQBox controls input from physical switches (buttons). Given `ButtonConfiguration`, it sets and monitors a connected pin on given GPIO. Upon state changes of the button, triggered events relay through a buffered channel to connected systems or routines in response realationsip structures elsewhere within code-flow. |
| [manager.go](https://github.com/kilianp07/AthleteIQBox/blob/main/buttons/manager.go) | Start and Stop buttons events for Athlete IQ Box project (emphasis on interface monitoring and management) Integration scope: butons package containing Manager to collect and control button actions via configuration defined peripheral devices. |

</details>

<details closed><summary>gps</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/configuration.go) | The `gps/configuration.go` file initializes data inputting and recording functions for the AthleteIQBox. This file encases configuration settings for readers and recorders, allowing easy alteration and connection to GPS devices or platforms, thus gathering exercise data at scale. |
| [gps.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/gps.go) | Data input, and recorder: data storage' modules, for this multi-function AthleteIQ tracking tool through interfaced custom objects. Essentially managing and monitoring GPS data influx. The main features being connectivity setup, configuration adaptation, and flowcontrol. A critical building block in the ecosystem architecture ensuring smooth operations between Athlete devices within real-time activity scenarios. |

</details>

<details closed><summary>gps.reader</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/reader/configuration.go) | Manages configuration for GPS functionality in AthleteIQBox project. By receiving relevant JSON input from main; it defines essential GPS reader specifics, paving way for accurate position data processing amid the application's architecture of robust features & agile functions. |
| [reader.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/reader/reader.go) | Configures and starts new GPS reader implementations. This code file acts as a bridge, unifying various location-tracking data sources, ensuring a seamless flow of Position data from multiple readouts into the larger IQBox athletic analytics system, utilizing customized decoders depending on respective reader types. |

</details>

<details closed><summary>gps.reader.nmea</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/reader/nmea/configuration.go) | Configures GPS reader in AthleteIQBox architecture, handling its operational settings. A JSON SerialConfiguration data type is set up for name, baud rate, read timeout, size, parity, and stop bits, with conversion functions from the configuration struct provided directly to a serial configuration object. Integral for precise location collection and box functionality. |
| [reader.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/reader/nmea/reader.go) | This function reads GPS data from the serial port and decodes NMEA messages into Actionable positions (Latitude, Longitude, Speed) structured as Position data structs which are emitted via a channel for realtime consumption in the AthleteIQBox data pipeline through continuous Scanner operations with custom buffering logic for enhanced reliability in realty environments.. |

</details>

<details closed><summary>gps.recorder</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/recorder/configuration.go) | Customizes data flow to athlete tracking hardware (gps) in AthleteIQBox application, defining configuration parameters within the scope of the `Configuration` struct. The struct accepts an id for unique recognition and dynamic Conf settings, enabling seamless device integration. |
| [recorder.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/recorder/recorder.go) | The `recorder.go` file in AthleteIQBox serves toinitialize multiple recorders in Go, providing configurable data capture interfaces for processing GPS positions within the overall architecture, abstractly by allowing instantiation via a simple map index approach [sqlite]. The selected recorder (in this case, sqlite) uses methods like `Configure`, `Start`, and `Stop` to handle position data, ensuring efficient data stream handling. |

</details>

<details closed><summary>gps.recorder.sqlite</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/recorder/sqlite/configuration.go) | The golang module `athleteiqbox`s configuration for its built-in sqlite Logger stores the location databases' filepath and flush frequency per SQLITE `configuration.go`, enhancing real-time Position Data's storage and retrieval within the system stack. Impact comes from robust GPS tracing capabilities improved by optimised data handling." |
| [recorder.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/recorder/sqlite/recorder.go) | RecordS GPS data interfaced from gps package into **SQLite databases** using `sqlite3 driver`. It allows for **timed sampling, position storage and table creation**. Error handling included with logging (utilizing importedLogger). Filepath customizable based on `config`. Asynchronously processes incoming positions when a **channel receives these**. Stoppable, making room for selective data ingestion. |
| [recorder_test.go](https://github.com/kilianp07/AthleteIQBox/blob/main/gps/recorder/sqlite/recorder_test.go) | Tests Recorder SQLite implementation during configuration. Verifies successful creation & connection to the SQLite database for data storage management upon receiving Athlete positon feeds in real-time. |

</details>

<details closed><summary>transmitter</summary>

| File | Summary |
| --- | --- |
| [factory.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/factory.go) | The map structure maps various service IDs, such as WiFi within packages services Wifi and factory methods in factory package with an anonymous function creating each service instance. Thus facilitating efficient instantiation and use across components without explicit dependencies or hard bindings. |
| [conf.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/conf.go) | Identification of source device by DeviceName in Transmitter2. Configure and orchestrate usage of multiple Services through map of Service configurations (ServicesConf). |
| [transmitter.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/transmitter.go) | Acts as gatt.Device for sensor interoperability: Disseminates configuration & establishes associated services using provided factories for real-time updates based on data exchange conditions. |
| [factory_test.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/factory_test.go) | Verify factory function behaviour by conducting tests under varied scenarios.**The supplied code is designed as test cases to determine the reliability of the factory() approach specified within transmitter/factory.go.' It ensures proper system behavior when working with valid service types while ensuring failures when an incorrect or invalid service type is passed as argument in this data-oriented module integral in AthleteIQ Box's service transmission logic. |

</details>

<details closed><summary>transmitter.services</summary>

| File | Summary |
| --- | --- |
| [service.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/services/service.go) | Configuration and interaction interface across hardware Service objects. (Leverages PayPals `gatt` library to connect & operate.) Empowers data Update, retrieves Service Uniqu Identifier for streamlined communications within our robust Athletque system. Ensues seamless compatibility & maintenance in evolving technology landscape of AthleteIQBox. (Encrypted and flexible Service integration enabled.) |

</details>

<details closed><summary>transmitter.services.wifi</summary>

| File | Summary |
| --- | --- |
| [wifi.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/services/wifi/wifi.go) | Function `handleReadAps()` compacts JSON data on Wifi access points, splits if too long for current packet size, then writes the response to peripheral connection. Its service UUID can be obtained by `function GetServiceUUID()`.2. Infinite loop of background Scan() constantly checks for new WiFi network updates based on triggered event (scan initiated at command).3. Function addNetwork(d wiFi) adds network details to hardware drivers internal management once triggered, if error-free. It takes wireless.SSID and password as parameters in dictionary form. |
| [data.go](https://github.com/kilianp07/AthleteIQBox/blob/main/transmitter/services/wifi/data.go) | Enhances connectivity through defining data structure for WiFi credentials within AthleteIQBox transmitter services modules library. This structured data (WiFi struct) comprises SSID and Password, empowering device connection with appropriate access points according to configurations. |

</details>

<details closed><summary>utils</summary>

| File | Summary |
| --- | --- |
| [decoder_test.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/decoder_test.go) | Initiates unit tests for encoding-decoding libraries in given repository. Ensures functions like NewDecoder can handle specified data objects with given set options and verify correct type, behavior, and outcomes. Central component to uphold decoded content integrity & application stability under specific context variations. |
| [decoder.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/decoder.go) | Utility decoder function is defined in utils/decoder.go, aiding with structured data parsing. Leveraging mapstructure package, this code automates complex JSON field conversion tasks without lengthy type declarations by utilizing custom hooks & option sets within the Decode operation. |
| [json_test.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/json_test.go) | This utilityfile `_json_test.go` tests theJSON parsing capabilities within the AthleteIQBoxrepositorys utils`. The test-suiteverifiesexpecteddata returned against apersisted file bycomparing structuredgo interfacesto actual decodeddata.Ensuring properfunctionality secures JSON handling robustness indiverserequard use-cases. |
| [json.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/json.go) | In this repository for an athletics-based data aggregator (AthleteIQBox), the `utils/json.go` file presents a `ReadJSONFile` utility empowered to load files as JSON content and transfer them into a compatible interface within the provided scope, promoting streamlined handling of JSON assets. Ensures smooth interactions with athletic performance information on demand. |
| [converters.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/converters.go) | This converters file simplifies exchanging byte[], float64, and strings, crucial for the interoperability between various components within our AthleteIQBox open-source project, facilitating efficient transmission and processing. |

</details>

<details closed><summary>utils.logger</summary>

| File | Summary |
| --- | --- |
| [configuration.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/logger/configuration.go) | Package utils/logger structuresLoggerConfiguration, defining logger preferences, which include timestamps and outputs array containing JSON maps configured in the app for message delivery flexibility while preserving caller context if needed. Crucial for streamlined error analysis in a vast, versatile open-source project repository. |
| [logger.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/logger/logger.go) | Console", File, and Syslog.2. It iterates through the specified output types from the user input, using `decodeOutputConfig` function to marshal data from raw configuration into target structs(ConsoleConfig, FileConfig or SysLogConfig). (1st 30 characters: uses decode output config to populate structure)3. A suitable method creates writers based on the specified output types.4. Constructing a zerolog logger is established with associated level and optional functionalities: timing(zerofmt_ms|miso)|zerortime|show caller(true/false).4a. Zerolog logs from a created logger are routed to respective configured channel.(49 words for points 3-4).Lastly implemented is logging message abstration `logMessage` function handling both logged text format(non-specific levels, custom messages w/logger preffix added for customized prefixing using L.prefix') and custom-built functionalities formatted at differing log severities(DEBUGf|INFOef|WARNf|ERROrF|Fatalf), utilizing zerolog's modularity with dynamic methods.(3 words) |
| [loglevel.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/logger/loglevel.go) | Empowers sophisticated logging mechanisms across repository components by defining different log levels, including but not limited to trace, debug, info, warn, error, and fatal for precise event analysis. |
| [output.go](https://github.com/kilianp07/AthleteIQBox/blob/main/utils/logger/output.go) | Guids custom log output configuration and handling solution within AthleteIQBox repository architecture. Provides logger options for console, file, or syslog with a configurable level, easing error interpretation throughout varying workflows of data processing and position-management applications via zero logs framework (zerolog). |

</details>

<details closed><summary>entrypoint</summary>

| File | Summary |
| --- | --- |
| [entrypoint.go](https://github.com/kilianp07/AthleteIQBox/blob/main/entrypoint/entrypoint.go) | Starts an athlete tracking service that seamlessly integrates GPS, Bluetooth wireless data transmission, and interactive buttons operation for enhanced performance monitoring. File entrypoint.go acts as central init point for all external dependencies within the AthleteIQBox ecosystem to execute service start-up, ensuring coherent tracking functionality. |

</details>

<details closed><summary>.github.workflows</summary>

| File | Summary |
| --- | --- |
| [generate-changelog.yml](https://github.com/kilianp07/AthleteIQBox/blob/main/.github/workflows/generate-changelog.yml) | Integrates a self-updating GitHub Actions workflow, enabling continuous Changelog generation within Git repository, AthleteIQBox. The configurationfile (generate-changelog.yml) ensures changelogs get automatic commit for each release-ready pull request in main branch. |

</details>

---

##  Getting Started

###  Prerequisites

**Go**: `version x.y.z`

###  Installation

Build the project from source:

1. Clone the AthleteIQBox repository:
```sh
❯ git clone https://github.com/kilianp07/AthleteIQBox
```

2. Navigate to the project directory:
```sh
❯ cd AthleteIQBox
```

3. Install the required dependencies:
```sh
❯ go build -o myapp
```

###  Usage

To run the project, execute the following command:

```sh
❯ ./myapp
```

###  Tests

Execute the test suite using the following command:

```sh
❯ go test
```

---

##  Project Roadmap

- [X] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/kilianp07/AthleteIQBox/issues)**: Submit bugs found or log feature requests for the `AthleteIQBox` project.
- **[Submit Pull Requests](https://github.com/kilianp07/AthleteIQBox/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/kilianp07/AthleteIQBox/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/kilianp07/AthleteIQBox
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/kilianp07/AthleteIQBox/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=kilianp07/AthleteIQBox">
   </a>
</p>
</details>

---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

---
