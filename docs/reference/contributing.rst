Contributing Guide
==================

Thank you for your interest in contributing to NeuralBlitz v50 Go!
This document provides guidelines for contributing to the project.

Getting Started
---------------

1. **Fork the Repository**

   Fork the repository on GitHub and clone your fork locally:

   .. code-block:: bash

      git clone https://github.com/YOUR_USERNAME/neuralblitz-v50-go.git
      cd neuralblitz-v50-go

2. **Set Up Development Environment**

   Ensure you have Go 1.21+ installed:

   .. code-block:: bash

      go version  # Should show go1.21 or higher

3. **Install Dependencies**

   .. code-block:: bash

      go mod download

4. **Run Tests**

   Make sure all tests pass before making changes:

   .. code-block:: bash

      go test ./pkg/systems/... ./pkg/quantum/...

Development Workflow
--------------------

1. **Create a Branch**

   .. code-block:: bash

      git checkout -b feature/your-feature-name

2. **Make Changes**

   - Write clean, idiomatic Go code
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed

3. **Test Your Changes**

   .. code-block:: bash

      go test ./pkg/... -v
      go build ./pkg/...

4. **Commit Changes**

   .. code-block:: bash

      git add .
      git commit -m "feat: add your feature description"

   Follow conventional commits format:

   - ``feat:`` - New feature
   - ``fix:`` - Bug fix
   - ``docs:`` - Documentation changes
   - ``test:`` - Test changes
   - ``refactor:`` - Code refactoring
   - ``perf:`` - Performance improvements

5. **Push and Create Pull Request**

   .. code-block:: bash

      git push origin feature/your-feature-name

   Then create a pull request on GitHub.

Code Style
----------

Follow standard Go conventions:

- Use ``gofmt`` for formatting
- Use ``golint`` for linting
- Use ``go vet`` for static analysis
- Follow Effective Go guidelines

Example:

.. code-block:: go

   // Good
   func NewQuantumAgent(agentID string, state QuantumState) *QuantumAgent {
       return &QuantumAgent{
           AgentID:            agentID,
           ConsciousnessLevel: state,
           CoherenceFactor:    1.0,
       }
   }

   // Bad
   func NewQuantumAgent(agentID string,state QuantumState)*QuantumAgent{
       return &QuantumAgent{AgentID:agentID,ConsciousnessLevel:state,CoherenceFactor:1.0}
   }

Testing Guidelines
------------------

- Write unit tests for all new functions
- Aim for >80% code coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions
- Use parallel testing for concurrent code

Example test:

.. code-block:: go

   func TestNewQuantumAgent(t *testing.T) {
       tests := []struct {
           name      string
           agentID   string
           state     QuantumState
           wantState QuantumState
       }{
           {
               name:      "create aware agent",
               agentID:   "agent_1",
               state:     StateAWARE,
               wantState: StateAWARE,
           },
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               agent := NewQuantumAgent(tt.agentID, tt.state)
               if agent.ConsciousnessLevel != tt.wantState {
                   t.Errorf("expected %v, got %v", tt.wantState, agent.ConsciousnessLevel)
               }
           })
       }
   }

Documentation
-------------

- Update docstrings for all public functions
- Update README.rst if adding major features
- Add examples to docs/examples/
- Update API documentation in docs/api/

Commit Message Format
---------------------

.. code-block:: text

   <type>: <subject>

   <body>

   <footer>

Types:

- ``feat`` - New feature
- ``fix`` - Bug fix
- ``docs`` - Documentation only
- ``style`` - Formatting, missing semi colons, etc
- ``refactor`` - Code change that neither fixes a bug nor adds a feature
- ``perf`` - Code change that improves performance
- ``test`` - Adding missing tests
- ``chore`` - Changes to the build process or auxiliary tools

Example:

.. code-block:: text

   feat: add quantum teleportation support

   Implements quantum teleportation between entangled agents
   using Bell state measurements and classical communication.

   Closes #123

Pull Request Process
--------------------

1. Ensure all tests pass
2. Update documentation
3. Add entry to CHANGELOG.rst
4. Request review from maintainers
5. Address review feedback
6. Squash commits if requested

Reporting Issues
----------------

When reporting issues, please include:

- Go version (``go version``)
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Error messages
- Code samples if applicable

Feature Requests
----------------

For feature requests:

1. Check existing issues first
2. Describe the use case
3. Explain why current solutions are inadequate
4. Suggest possible implementations

Security Issues
---------------

For security issues, please email security@neuralblitz.org instead of
using the public issue tracker.

Code of Conduct
---------------

- Be respectful and inclusive
- Welcome newcomers
- Accept constructive criticism
- Focus on what's best for the community
- Show empathy towards others

License
-------

By contributing, you agree that your contributions will be licensed
under the MIT License.

Questions?
----------

- GitHub Discussions: https://github.com/neuralblitz/neuralblitz-v50-go/discussions
- Discord: https://discord.gg/neuralblitz
- Email: contributors@neuralblitz.org

Thank you for contributing!
