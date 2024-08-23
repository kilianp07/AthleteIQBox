# Package `gps` Documentation

The `gps` package provides a flexible framework for handling GPS data from various sources and formats, such as NMEA. It allows the dynamic creation, configuration, and management of GPS data readers, enabling easy integration with different GPS hardware and protocols.

## Overview

The `gps` package centers around the `Reader` interface, which all GPS readers must implement. This package facilitates the creation of GPS readers based on user-provided configurations, allowing for the seamless integration of different GPS data formats.

Dependencies:
- **mapstructure**: Used for decoding configuration structures.
  ```shell
  go get github.com/go-viper/mapstructure/v2
  ```
- **go-nmea**: For parsing NMEA sentences.
  ```shell
  go get github.com/adrianmo/go-nmea
  ```
- **serial**: For handling serial port communication.
  ```shell
  go get github.com/tarm/serial
  ```

## Types

### `Reader` Interface

The `Reader` interface defines the methods that any GPS reader implementation must provide:

```go
type Reader interface {
    Configure() error         // Configures the GPS reader using the provided settings.
    Start() error             // Starts reading GPS data.
    Stop() error              // Stops reading GPS data.
    RuntimeErr() chan error   // Returns a channel for runtime errors.
    Conf() any                // Returns the configuration structure of the reader.
}
```

### `Configuration` Struct

The `Configuration` struct specifies the type of GPS reader and its configuration:

```go
type Configuration struct {
    Id   string `json:"id"`   // Identifier for the reader type (e.g., "nmea").
    Conf any    `json:"conf"` // Configuration settings for the reader.
}
```

## Functions

### `New(conf Configuration) (Reader, error)`

The `New` function creates and configures a new GPS reader based on the provided configuration:

```go
func New(conf Configuration) (Reader, error)
```

- **Parameters:**
  - `conf`: A `Configuration` struct that includes the `Id` of the reader type and its configuration.

- **Returns:**
  - `Reader`: A configured instance of the requested GPS reader.
  - `error`: An error if the reader creation or configuration fails.

- **Example Usage:**

```go
conf := gps.Configuration{
    Id:   "nmea",
    Conf: map[string]interface{}{
        "Period": "1s",
        "serialConfig": map[string]interface{}{
            "Name": "/dev/ttyUSB0",
            "Baud": 9600,
        },
    },
}

reader, err := gps.New(conf)
if err != nil {
    log.Fatalf("Failed to create GPS reader: %v", err)
}

if err := reader.Start(); err != nil {
    log.Fatalf("Failed to start GPS reader: %v", err)
}
```

## NMEA Reader Implementation

The `nmea` package provides an example of a GPS reader implementation that parses NMEA sentences from a serial port.

### `nmea.Reader` Struct

The `nmea.Reader` struct implements the `Reader` interface and reads NMEA GPS data from a serial port.

```go
type Reader struct {
    conf   *Configuration
    period time.Duration

    runCh chan bool
    errCh chan error

    currentPoint Point
}
```

### `nmea.Reader` Methods

- **`New()`**: Creates a new `nmea.Reader`.

  ```go
  func New() *Reader {
      return &Reader{}
  }
  ```

- **`Conf()`**: Returns the configuration structure.

  ```go
  func (r *Reader) Conf() any {
      return &r.conf
  }
  ```

- **`Configure()`**: Configures the reader using the provided settings.

  ```go
  func (r *Reader) Configure() error {
      // Example configuration handling
  }
  ```

- **`Start()`**: Starts the GPS data reading process. Opens the serial port and begins reading and parsing NMEA sentences.

  ```go
  func (r *Reader) Start() error {
      // Example start implementation
  }
  ```

- **`Stop()`**: Stops the GPS data reading process.

  ```go
  func (r *Reader) Stop() error {
      r.runCh <- false
      return nil
  }
  ```

- **`RuntimeErr()`**: Returns a channel for runtime errors.

  ```go
  func (r *Reader) RuntimeErr() chan error {
      return r.errCh
  }
  ```

- **`run(s *serial.Port)`**: Internal method that continuously reads from the serial port, parses NMEA sentences, and updates the current GPS point.

  ```go
  func (r *Reader) run(s *serial.Port) {
      // Example run implementation
  }
  ```

### Example NMEA Sentence Handling

In the `run` method, NMEA sentences such as `GGA`, `RMC`, and others are parsed, and relevant data is logged or used to update the current GPS point:

```go
switch sentence.DataType() {
case nmea.TypeGGA:
    // Handle GGA sentence
case nmea.TypeRMC:
    // Handle RMC sentence
case nmea.TypeGLL:
    // Handle GLL sentence
case nmea.TypeGSA:
    // Handle GSA sentence
case nmea.TypeGSV:
    // Handle GSV sentence
case nmea.TypeVTG:
    // Handle VTG sentence
}
```

## Extending the `gps` Package

To support additional GPS formats, implement the `Reader` interface for the new format and register it in the package's `init` function:

```go
func init() {
    readers["new_format"] = func() Reader {
        return newFormat.New()
    }
}
```
