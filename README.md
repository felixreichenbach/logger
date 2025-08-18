# Module my-logger

This is a simple module to demonstrate Viam logging capabilities.

## Model felixr:my-logger:my-logger

A simple sensor component which logs "debug", "info", "warning" and "error" messages upon calling the Readings() API.

### Configuration

The following attribute template can be used to configure this model:

```json
{}
```

#### Attributes

The following attributes are available for this model:

| Name              | Type       | Inclusion    | Description                    |
| ----------------- | ---------- | ------------ | ------------------------------ |
| ~~`attribute_1`~~ | ~~float~~  | ~~Required~~ | ~~Description of attribute 1~~ |
| ~~`attribute_2`~~ | ~~string~~ | ~~Optional~~ | ~~Description of attribute 2~~ |

#### Example Configuration

```json
{}
```

### DoCommand

With the DoCommand() you can trigger log entries at different levels.

#### Example DoCommand

```json
{
  "message": "info" // one of [debug, info, warn, error or fatal]
}
```

## Viam Log Configuration

- Viam Server

  -log-file string
  write logs to a file with log rotation
  `-log-file log.txt`

  -debug
  Outputs debug log messages and also sends them to app.viam.com

- Configuration
  https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#logging

- - top level / machine
    Easiest option add `"debug":true` to the top level of your machine configuration. Server automatically set's log level accordingly on startup. Will require restart! Will overwrite all subsettings!
    https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#debugging

    top level `log` array with regex pattern or level definition

- - Module
    `"log_level":"debug"`

- - Component
    via menu or through resource json config
    `"log_configuration":{"level": "debug"}`

Log Deduplication
https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#disable-log-deduplication
At root level in the machine config:
`"disable_log_deduplication": true`
