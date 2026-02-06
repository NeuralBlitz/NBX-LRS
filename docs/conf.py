# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the
# documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

import os
import sys

# Add the Go project's source path to sys.path
# sys.path.insert(0, os.path.abspath('.'))

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

project = "NeuralBlitz v50 Go"
copyright = "2026, NeuralBlitz Team"
author = "NeuralBlitz Team"
release = "50.0.0"

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.viewcode",
    "sphinx.ext.napoleon",
    "sphinx.ext.todo",
    "sphinx.ext.coverage",
    "sphinx.ext.mathjax",
]

templates_path = ["_templates"]
exclude_patterns = ["_build", "Thumbs.db", ".DS_Store"]

# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

html_theme = "alabaster"
html_static_path = ["_static"]
html_logo = "_static/neuralblitz-logo.png"

html_theme_options = {
    "github_repo": "neuralblitz-v50",
    "github_user": "neuralblitz",
    "show_related": True,
}

html_context = {
    "github_user": "neuralblitz",
    "github_repo": "neuralblitz-v50",
    "github_version": "main",
    "doc_path": "docs",
}

# -- Extension configuration -------------------------------------------------

# Napoleon settings
napoleon_google_docstring = True
napoleon_include_init_with_doc = True
napoleon_include_private_with_doc = True
napoleon_include_special_with_doc = True

# Autodoc settings
autodoc_default_options = {
    "members": True,
    "undoc-members": True,
    "show-inheritance": True,
}

# -- Go domain configuration -----------------------------------------------

# Since Sphinx doesn't natively support Go, we use go-langdomain
# or configure manual cross-references
