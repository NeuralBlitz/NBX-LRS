NeuralBlitz v50 Go Documentation
================================

This directory contains the Sphinx-based documentation for NeuralBlitz
v50 Go implementation.

Prerequisites
-------------

- Python 3.8+
- pip

Installation
------------

Install the required dependencies:

.. code-block:: bash

   pip install -r requirements.txt

Building Documentation
----------------------

Build HTML documentation:

.. code-block:: bash

   make html

The built documentation will be in ``_build/html/``.

Serve documentation locally:

.. code-block:: bash

   make serve

Then open http://localhost:8000 in your browser.

Other Build Targets
-------------------

.. code-block:: bash

   make latex    # Build LaTeX
   make pdf      # Build PDF
   make epub     # Build EPUB
   make man      # Build manual pages

Development
-----------

Watch for changes and auto-rebuild:

.. code-block:: bash

   make watch

Check documentation quality:

.. code-block:: bash

   make check

Clean build files:

.. code-block:: bash

   make clean

Documentation Structure
-----------------------

.. code-block:: text

   docs/
   ├── conf.py              # Sphinx configuration
   ├── index.rst            # Main index
   ├── Makefile             # Build automation
   ├── requirements.txt     # Python dependencies
   ├── structure.rst        # Project structure
   │
   ├── architecture/        # Architecture documentation
   │   ├── overview.rst
   │   ├── systems.rst
   │   └── quantum.rst
   │
   ├── api/                 # API reference
   │   ├── systems.rst
   │   ├── quantum.rst
   │   └── common.rst
   │
   ├── guides/              # User guides
   │   ├── getting_started.rst
   │   ├── installation.rst
   │   ├── quickstart.rst
   │   ├── configuration.rst
   │   └── deployment.rst
   │
   ├── examples/            # Example documentation
   │   ├── quantum_neurons.rst
   │   ├── multi_reality.rst
   │   ├── consciousness.rst
   │   └── agents.rst
   │
   └── reference/           # Technical reference
       ├── glossary.rst
       ├── changelog.rst
       ├── contributing.rst
       └── license.rst

Contributing to Documentation
-----------------------------

1. Write documentation in reStructuredText format
2. Follow the existing structure and style
3. Build and test locally before submitting
4. Run ``make check`` to verify quality
5. Submit a pull request

ReadTheDocs Integration
-----------------------

This documentation is designed to be hosted on ReadTheDocs.

Configuration:

1. Connect your GitHub repository to ReadTheDocs
2. Configure the webhook for automatic builds
3. Set the documentation type to "Sphinx"
4. Configure the Python version to 3.8+

The ``.readthedocs.yml`` configuration file is included in the repository.

License
-------

Documentation is released under the MIT License.
