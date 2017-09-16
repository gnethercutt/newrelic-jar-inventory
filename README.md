# New Relic Infrastructure Integration for Jar inventory

Reports the manifest of Java dependencies embedded in an Uberjar to New Relic Infrastructure inventory.

## Requirements

Requires access to execute the `/usr/bin/jar` binary from the Java runtime environment.
Expects examined jars to assembled with jar files that conform to the naming scheme `lib/some-java-package-1.2.3.jar`

## Configuration

The user under which the newrelic-infra agent is running needs read access to any jarfile to be inspected.

## Installation

Place the jarinventory-config.yaml in /var/db/newrelic-infra/integrations.d
Place the jarinventory-definition.yaml in /var/db/newrelic-infra/custom-integrations/
Place the jarinventory binary in /var/db/newrelic-infra/custom-integrations/bin and make it executable.

## Usage

### Development/testing
`./bin/jarinventory -pretty -jarfile /path/to/jar`

### Arguments
#### `-jarfile` (MANDATORY)
Complete path to the jarfile to be inspected

#### `-jrepath` (OPTIONAL)
Path to the directory containing the jar binary, /usr/bin by default

## Compatibility

* Supported OS: Linux, MacOs
* jarinventory versions: 0.1.0

