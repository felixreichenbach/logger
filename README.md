# Module my-logger

This is a simple module to demonstrate Viam logging capabilities.

## Model felixr:my-logger:my-logger

A simple sensor component which logs "debug", "info", "warning" and "error" messages upon calling the Readings() API.

### Configuration

The module doesn't require any configuration.

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

#### Machine Level

1. Add `"debug":true` to the top level of your machine JSON configuration. This is analog adding the Viam server startup parameter. This will require a Viam server restart and overwrites all subsettings!
2. For more granularity you can use the top level "log" element:
   ```JSON
    "log": [
        {
            "pattern": "rdk.services.*",
            "level": "debug"
        }
    ]
   ``` 

#### Module Level

To [enable debug logs for a whole module](https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#debugging) you can add the following to your module JSON configuration:

```json
"log_level": "debug"
```

#### Component Level

To enable debug log level for a single component you can either enable it via context menue or add this to the component JSON config:

```json
"log_configuration":{
  "level": "debug"
}
```

### Log Deduplication

[Viam-server deduplicates log message](https://docs.viam.com/operate/reference/viam-server/?source=searchResultItem#disable-log-deduplication) that are deemed noisy. A log is deemed noisy if it has been output 3 times in the past 10 seconds.
As this can complicate debugging issues, you can disable it as follows.
Add the following line to the root level in the machine config:
```json
"disable_log_deduplication": true
```

