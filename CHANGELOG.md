# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.3] - 2023-03-30

### Added

- Registration will be every `n` minutes to be sure that module is registered
- Registration logic

## [1.0.2] - 2023-03-24

### Added

- `AMQP` connection and created receiver
- Telegram info in db
- Updating telegram information dynamically from `telegram-module`

### Changed

- Sorts by name, not by ID

## [1.0.1] - 2023-03-22

### Added

- Registration in `orchestrator`
- Requests to send roles

## [1.0.0] - 2023-03-15

### Added

- Database.
- API handlers.

[1.0.0]: https://gitlab.com/distributed_lab/acs/orchestrator/-/tree/feature/requests_action_filter
[1.0.0]: https://github.com/distributed_lab/acs/identity-svc/compare/develop...feature/add_registration_roles
