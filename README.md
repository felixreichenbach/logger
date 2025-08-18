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
