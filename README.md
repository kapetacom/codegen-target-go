# Kapeta Language Target for Go

Provides Kapeta language support for Go

Will use the most commonly used / effective libraries for Go such as:

-   Echo as HTTP server

This language target is only available for Service blocks.

To lean more about Kapeta, see https://kapeta.com or https://docs.kapeta.com

Features

[X] Rest server
[X] Rest API
[] JWT Support
[] Postgres
[] Redis
[] Mongo

## Project Structure
The structure of this language target is as follows:

### `src/web`
Contains the browser-based source used by the Kapeta App when configuring the language target.

### `src/target`
Contains the node-based source for generating code for this language target.

### `templates`
Contains the templates used by the target to generate code. These are written using the [Handlebars](https://handlebarsjs.com/) template language.

See https://github.com/kapetacom/codegen-target for more information about the template syntax.

## Features

**Note:** In your block see the ```kapeta.md``` file for more information specifically about the code that is generated for your block.

### Linting
Will add eslint support - lint your code using `npm run lint`.

### REST Clients
Will generate REST clients for all consumed REST APIs

### REST APIs
Will generate REST API routes for all provided REST APIs

### JWT Provider
Will generate a basic keystore provider for creating and signing JWT tokens

## Changes and Suggestions

If you wish to change the templates or code being generated - consider either opening a PR to if you feel it could be universally beneficial or fork the project and make your own changes - which you can then publish to Kapeta as your own language target.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
