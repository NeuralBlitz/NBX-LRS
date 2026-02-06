# NeuralBlitz v50 Go Documentation Makefile
# ==========================================

# Variables
SPHINXOPTS    ?=
SPHINXBUILD   ?= sphinx-build
SOURCEDIR     = .
BUILDDIR      = _build

# Default target
.PHONY: help
help:
	@echo "NeuralBlitz v50 Go Documentation"
	@echo "================================"
	@echo ""
	@echo "Available targets:"
	@echo "  make html       - Build HTML documentation"
	@echo "  make latex      - Build LaTeX documentation"
	@echo "  make pdf        - Build PDF documentation"
	@echo "  make epub       - Build EPUB documentation"
	@echo "  make man        - Build manual pages"
	@echo "  make doctest    - Run doctests"
	@echo "  make linkcheck  - Check external links"
	@echo "  make spelling   - Check spelling"
	@echo "  make clean      - Clean build files"
	@echo "  make serve      - Serve documentation locally"
	@echo "  make watch      - Watch and rebuild on changes"

# Build HTML documentation
.PHONY: html
html:
	$(SPHINXBUILD) -b html $(SOURCEDIR) $(BUILDDIR)/html $(SPHINXOPTS)
	@echo ""
	@echo "Build succeeded! Open file://$(PWD)/$(BUILDDIR)/html/index.html"

# Build LaTeX documentation
.PHONY: latex
latex:
	$(SPHINXBUILD) -b latex $(SOURCEDIR) $(BUILDDIR)/latex $(SPHINXOPTS)

# Build PDF documentation
.PHONY: pdf
pdf: latex
	cd $(BUILDDIR)/latex && $(MAKE) all-pdf
	@echo ""
	@echo "PDF generated at $(BUILDDIR)/latex/neuralblitzv50go.pdf"

# Build EPUB documentation
.PHONY: epub
epub:
	$(SPHINXBUILD) -b epub $(SOURCEDIR) $(BUILDDIR)/epub $(SPHINXOPTS)

# Build manual pages
.PHONY: man
man:
	$(SPHINXBUILD) -b man $(SOURCEDIR) $(BUILDDIR)/man $(SPHINXOPTS)

# Run doctests
.PHONY: doctest
doctest:
	$(SPHINXBUILD) -b doctest $(SOURCEDIR) $(BUILDDIR)/doctest $(SPHINXOPTS)

# Check external links
.PHONY: linkcheck
linkcheck:
	$(SPHINXBUILD) -b linkcheck $(SOURCEDIR) $(BUILDDIR)/linkcheck $(SPHINXOPTS)

# Check spelling
.PHONY: spelling
spelling:
	$(SPHINXBUILD) -b spelling $(SOURCEDIR) $(BUILDDIR)/spelling $(SPHINXOPTS)

# Clean build files
.PHONY: clean
clean:
	rm -rf $(BUILDDIR)/*

# Serve documentation locally
.PHONY: serve
serve: html
	@echo "Starting server at http://localhost:8000"
	@cd $(BUILDDIR)/html && python3 -m http.server 8000

# Watch and rebuild
.PHONY: watch
watch:
	@echo "Watching for changes..."
	@while true; do \
		inotifywait -e modify,create,delete -r $(SOURCEDIR) 2>/dev/null || true; \
		echo "Changes detected, rebuilding..."; \
		$(MAKE) html; \
	done

# Build all formats
.PHONY: all
all: html latex epub

# Deploy to ReadTheDocs (local simulation)
.PHONY: deploy
deploy: clean html
	@echo "Documentation ready for deployment"
	@echo "Upload $(BUILDDIR)/html to your hosting service"

# Install dependencies
.PHONY: install
install:
	pip install sphinx sphinx-autobuild sphinx-rtd-theme

# Generate API documentation from source
.PHONY: apidoc
apidoc:
	@echo "Generating API documentation..."
	sphinx-apidoc -o api ../pkg ../pkg/**/*_test.go

# Quick build for development
.PHONY: quick
quick:
	$(SPHINXBUILD) -b html -D exclude_patterns=[] $(SOURCEDIR) $(BUILDDIR)/html $(SPHINXOPTS)

# Check documentation quality
.PHONY: check
check: linkcheck spelling doctest
	@echo "Documentation quality checks complete"
