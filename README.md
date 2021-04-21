# Jenkins Tracer

[![License](https://img.shields.io/badge/license-MIT-red)](LICENSE)
[![Version](https://img.shields.io/badge/version-0.1.0-blue)](CHANGELOG.md)
![Maintained](https://img.shields.io/badge/maintained-yes-brightgreen)

Jenkins tracer is used to record all the jenkins job variables like record the build duration, build variables, repository metadata, etc.

All job environments variable will be send to elasticsearch cluster.

## Notes

Please run this binary in the end of post jenkins pipeline.

## Fields

- All environment variables
- Build duration
- Build start
- Build end
- Build result

## Kibana Dashboard and Visualization

You can import the example dashboard and visualization in `kibana` directory.

## Maintainer

- Misbahul Ardani