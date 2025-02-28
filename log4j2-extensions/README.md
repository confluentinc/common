StructuredJsonLayout Plugin Usage Guide
=======================================

This guide provides instructions on how to use the `StructuredJsonLayout` plugin in a Log4j2 YAML configuration file.

## Prerequisites

Ensure you have the `log4j2-extensions` dependency included in your project. Add the following to your `pom.xml` if you are using Maven:

```xml
<dependency>
    <groupId>io.confluent</groupId>
    <artifactId>confluent-log4j2-extensions</artifactId>
</dependency>
```

## Configuration

Below is an example of how to configure the StructuredJsonLayout plugin in a Log4j2 YAML configuration file

```yaml
Configuration:
  status: warn
  appenders:
    console:
      name: Console
      target: SYSTEM_OUT
      StructuredJsonLayout:
        Property:
          - name: "property1"
            value: "value1"
          - name: "property2"
            value: "value2"
  loggers:
    root:
      level: info
      appenderRefs:
        - ref: Console
```

If you are using `SerializableSchemaAndValue`, then the following properties can be provided to `StructuredJsonLayout` to migrate to log4j2:-

```yaml
  StructuredJsonLayout:
    Property:
      - name: schemas.enable
        value: false
```
