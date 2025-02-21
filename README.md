# Plugin System Overview

## Draft Version 1.0

### Concept and Workflow

The plugin system is designed with a focus on security, flexibility, and ease of use. Below is the proposed workflow for plugin integration and execution:

1. **Plugin Connection Request**: Plugins will initiate a connection request to the main application.

2. **Execution Permission Request**: Upon connection, plugins will request `exec` permissions to execute specific commands. Examples include:
   - `docker ps -a`
   - `docker run -d %@`

3. **User-Driven Permission Granting**: The user will review and grant the necessary permissions to the plugin based on the requested actions.

4. **Command Execution**: Once permissions are granted, the plugin will proceed to execute the approved commands.

5. **Result Handling**: The application will handle the results of the plugin's execution, including returning output, errors, and progress updates.

6. **Logging**: All plugin activities will be logged by the application for auditing and troubleshooting purposes.

### Open Permissions

For the initial phase, only `exec` permissions will be made available to ensure a controlled environment.

### Feedback and Contributions

We are open to suggestions and contributions to refine this plugin system. Your feedback is invaluable in shaping the future of our plugin ecosystem.

### Contact Information

For any inquiries or to provide feedback, please reach out to us at:

- Email: [admin@ssh2.app](mailto:admin@ssh2.app)
- Reply-To: [ssh2.app@gmail.com](mailto:ssh2.app@gmail.com)

*Please ensure you are whitelisted in our system to receive prompt responses.*

Join us in building a powerful and secure plugin system that empowers developers and enhances application functionality.