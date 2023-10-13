
# WASM REPO PROJECT

## Project Structure:

The overal project structure looks like this

```bash
$ tree

/wasm-repo
|-- /api
|   |-- /handlers
|   |-- /models
|   |-- /middleware
|   |-- main.go
|
|-- /web
|   |-- /src
|   |   |-- /components
|   |   |-- /views
|   |   |-- /assets
|   |   |-- /store
|   |   |-- App.vue
|   |   |-- main.js
|   |-- package.json
|
|-- /wasmctl
|   |-- /commands
|   |-- main.go
|
|-- /shared
|   |-- /utils
|
|-- README.md
|-- .gitignore
```

## Explanation:

- /api: This directory will contain the Go code for your REST API server.
  - /handlers: Contains the request handlers for your CRUD operations.
  - /models: Contains data structures and database models.
  - /middleware: Contains any middleware, such as authentication or logging.
  - main.go: The entry point for your API server.
- /web: This directory will contain your Vue.js web application.
  - /src: Contains the source code for your Vue.js application.
    - /components: Vue components that can be reused across different views.
    - /views: Vue components that represent entire views or pages.
    - /assets: Static assets like images or fonts.
    - /store: Contains Vuex store files.
    - App.vue: The main Vue component.
    - main.js: The entry point for your Vue.js application.
  - package.json: Lists the dependencies for your Vue.js application.
- /wasmctl: This directory will contain the Go code for your command-line tool.
  - /commands: Contains different commands for your CLI.
  - main.go: The entry point for your command-line tool.
- /shared: Contains any code or utilities shared between the API and the CLI.
- README.md: Documentation for your project.
- .gitignore: List of files and directories that should not be tracked by Git.

## Development Steps:

In order to devide the project in to digestible chunks, we defined the following stages

### Backend Development:

Set up a Go server with necessary routes for CRUD operations.
Implement database models and handlers for WASM binaries and metadata.
Add middleware for authentication, logging, etc.

### Frontend Development:

Set up a Vue.js project using Vue CLI.
Create components for listing, uploading, updating, and deleting WASM binaries and metadata.
Use Axios or Fetch API to make requests to your backend.

### CLI Development:

Set up a basic Go CLI structure.
Implement commands for CRUD operations on WASM binaries and metadata.

### Testing:

Write unit tests for your API handlers and CLI commands.
Write end-to-end tests for your Vue.js application.

### Deployment:

Deploy your API to a server or cloud provider.
Build and deploy your Vue.js application to a static hosting service or cloud provider.
Provide installation instructions for your CLI tool.

### Documentation:

Document the API endpoints, request/response formats, and any authentication requirements.
Provide a user guide for the Vue.js application.
Document the installation and usage of the CLI tool.
# wasm-repo
