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

## Viam Log Configuration Possibilities

### Via Viam Server

Viam Server has a variety of interesting settings for debugging which you can set directly when starting the service. 
Two options you will likely find useful:

- `-log-file <file name>` Write logs to a file with log rotation
- `-debug` Outputs debug log messages and also sends them to app.viam.com

### Via Machine Configuration

Besides directly adding debug logging through the Viam Server binary, further options are available through the machine configuration.

#### Top Level / Machine

1. Add `"debug":true` to the top level of your machine JSON configuration. This is analog adding the Viam server startup parameter. This will require a Viam server restart and overwrites all subsettings!
2. For more granularity you can use the top level "log" element:
   ```JSON
   "log":[{
     "pattern": "rdk.services.*",
     "level": "debug"}]
   ``` 
   

#### Module
    `"log_level":"debug"`
     https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#debugging

#### Component
    via menu or through resource json config
    `"log_configuration":{"level": "debug"}`

### Log Deduplication
https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#disable-log-deduplication
At root level in the machine config:
`"disable_log_deduplication": true`
