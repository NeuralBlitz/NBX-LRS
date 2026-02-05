# Environment Documentation

Generated: Thu Feb 05 2026

## 1. System Overview

### Hardware & OS
- **Hostname**: b8d297606b41
- **Operating System**: Linux (NixOS-based container)
- **Kernel Version**: 6.14.11 #1-NixOS SMP PREEMPT_DYNAMIC
- **Architecture**: x86_64 GNU/Linux
- **Platform**: Replit (riker cluster, interactive subcluster)

### Resource Allocation
- **Memory**: 62 GB total
  - Used: 12 GB
  - Free: 17 GB
  - Buff/Cache: 34 GB
  - Available: 50 GB
- **Storage**:
  - Workspace: 256 GB total (9.4 GB used, 245 GB free)
  - Root: 74 GB total (37 GB used, 35 GB free)
  - Nix Store: 32 GB (14 MB used)
  - tmpfs: 16 GB (115 MB used)

## 2. Development Environment

### Runtime Versions
- **Python**: 3.11.14
- **Node.js**: v22.22.0
- **npm**: 10.9.4
- **Nix Channel**: stable-24_11

### Package Managers Available
- pip (Python)
- poetry (Python)
- uv (Python - v0.9.5)
- npm/node (Node.js)
- pnpm (v10.26.1)
- yarn (v1.22.22)
- bun (v1.2.16)

### Build Tools
- TypeScript 5.6.3
- esbuild 0.25.0
- tsx 4.20.5
- vite 7.3.0
- drizzle-kit 0.31.8

### Code Quality Tools
- black (Python formatter)
- ruff (Python linter)
- mypy (Python type checker)
- prettier (JS/TS formatter)
- pre-commit hooks

## 3. Project Structure

### Root Directory Contents
```
/home/runner/workspace/
├── AGENTS.md                    # NeuralBlitz system documentation (320KB)
├── bun.lock                     # Bun package manager lockfile
├── lrs_agents/                  # LRS-Agents Python framework
├── NB-Ecosystem/               # Full-stack web application
├── neuralblitz-v50/            # NeuralBlitz v50 components
├── opencode-lrs-agents-nbx/    # LRS integration bridge
├── lrs_agents/                 # Active Inference agents (Python)
├── quantum_optimization.py     # Quantum algorithms research
├── quantum_temp.py             # Quantum temp module
├── pyproject.toml              # Python project config
├── uv.lock                     # UV package manager lockfile
├── node_modules/               # Node.js dependencies
└── .cache/                     # Build/tool caches
```

## 4. Major Projects

### 4.1 LRS-Agents (lrs_agents/)
**Purpose**: Resilient AI agents via Active Inference
**Language**: Python 3.9+
**Framework**: LangChain, LangGraph

**Key Components**:
- `lrs/core/` - Core Active Inference implementation
  - `precision.py` - Precision tracking (Beta distributions)
  - `free_energy.py` - Free energy calculations
  - `lens.py` - Tool abstraction and composition
  - `registry.py` - Tool management
- `lrs/integration/` - Framework integrations (LangChain, AutoGPT, OpenAI)
- `lrs/monitoring/` - Dashboard and logging
- `lrs/multi_agent/` - Multi-agent coordination
- `lrs/benchmarks/` - Benchmark implementations

**Dependencies**:
- langchain >= 0.1.0
- langchain-anthropic >= 0.1.0
- langchain-openai >= 0.0.5
- langgraph >= 0.0.20
- pydantic >= 2.5.0
- numpy >= 1.24.0
- scipy >= 1.11.0

**Test Coverage**: 95%+

**Build Commands**:
```bash
pytest tests/ -v --cov=lrs          # Run tests with coverage
black lrs tests                      # Format code
ruff check lrs tests --fix          # Lint and fix
mypy lrs --ignore-missing-imports   # Type check
```

### 4.2 NB-Ecosystem
**Purpose**: Full-stack web application platform
**Stack**: React + TypeScript + Express + PostgreSQL

**Tech Stack**:
- **Frontend**: React 18.3.1, TypeScript 5.6.3, Tailwind CSS 3.4.17
- **Backend**: Express 4.22.1, Drizzle ORM 0.39.3
- **Database**: PostgreSQL (via pg driver)
- **UI Components**: Radix UI primitives (40+ components)
- **State**: TanStack Query 5.60.5, React Hook Form 7.55.0
- **Validation**: Zod 3.24.2
- **Build**: Vite 7.3.0, tsx, esbuild

**Key Features** (from MVP_FEATURE_SET_AND_USER_STORIES.md):
- AI-powered research tools
- Multi-agent coordination
- Real-time collaboration
- Advanced analytics dashboard

**Scripts**:
```bash
npm run dev          # Development mode (tsx server/index.ts)
npm run build        # Production build
npm run start        # Production start
npm run db:push      # Database migrations
```

### 4.3 NeuralBlitz v50 (neuralblitz-v50/)
Advanced research components and computational systems

### 4.4 Quantum Optimization (quantum_optimization.py)
Research implementations for quantum algorithms

## 5. Configuration Files

### 5.1 Replit Configuration (.replit)
```toml
entrypoint = "index.js"
modules = ["nodejs-22", "python-3.11"]
hidden = [".config", "package-lock.json"]

[nix]
channel = "stable-24_11"
packages = ["libxcrypt"]

[deployment]
run = ["node", "index.js"]
deploymentTarget = "autoscale"

[agent]
expertMode = true

[[ports]]
localPort = 8000
externalPort = 80
```

### 5.2 Python Project (pyproject.toml)
```toml
[project]
name = "repl-nix-workspace"
version = "0.1.0"
requires-python = ">=3.11"
dependencies = []
```

## 6. Git Repository Status

### Branches
- **Current**: main
- **Local Branches**: main, new-main
- **Remote Branches**:
  - origin/main
  - origin/new-main
  - origin/NeuralBlitz-patch-1
  - gitsafe-backup/main

### Recent Changes
**Added**:
- quantum_optimization.py
- quantum_temp.py
- AGENTS.md (renamed from annt.md)
- NB-Ecosystem/
- opencode-lrs-agents-nbx/

**Modified**:
- lrs_agents/
- neuralblitz-v50/

**Deleted** (cleanup in progress):
- Multiple documentation files
- Test coverage reports
- Python demo scripts
- Server files (index.js, server.py, main.py)
- docs/ directory
- htmlcov/ directory

## 7. Environment Variables

### Key Variables
- `HOME=/home/runner`
- `USER=runner`
- `REPLIT_USER=youngriggs37`
- `REPLIT_USERID=39439480`
- `REPLIT_SESSION=youngriggs37-lcdf`
- `REPLIT_DEV_DOMAIN=c9293dcd-8b5e-4edb-a3b4-151933bf56c0-00-tbhogh34tgtp.riker.replit.dev`
- `REPLIT_DB_URL` - Configured for Replit KV store
- `REPLIT_HELIUM_ENABLED=true`
- `REPLIT_GITSAFE_ENABLED=true`

### Python Environment
- `PYTHONPATH=/nix/store/...`
- `PYTHONUSERBASE=/home/runner/workspace/.pythonlibs`

### Nix Configuration
- `NIX_PATH=nixpkgs=/home/runner/.nix-defexpr/channels/nixpkgs-stable-24_11`
- `NIX_PROFILES=/nix/var/nix/profiles/default /home/runner/.nix-profile`
- `NIXPKGS_ALLOW_UNFREE=1`

## 8. Installed Dependencies

### Node.js Packages (Root)
- ajv 8.17.1 (JSON schema validation)
- ajv-draft-04 1.0.0
- bash-language-server 5.6.0
- fast-glob 3.3.3
- lodash 4.17.21
- prettier 3.6.2
- yaml-language-server
- Various TypeScript definition packages

### NB-Ecosystem Dependencies
**Production (84 packages)**:
- Express ecosystem (express, express-session, passport)
- React ecosystem (react, react-dom, react-hook-form, react-query)
- Radix UI primitives (40+ components)
- Drizzle ORM + PostgreSQL
- Tailwind CSS + animation libraries
- Math rendering (katex, rehype-katex)
- Charts (recharts)
- Form validation (zod)

**Development (24 packages)**:
- TypeScript + type definitions
- Vite + plugins
- Tailwind CSS tooling
- esbuild

## 9. Key Files and Artifacts

### Documentation
- `AGENTS.md` (320KB) - NeuralBlitz Absolute Codex
- `lrs_agents/README.md` (30KB) - LRS-Agents documentation
- `lrs_agents/IMPLEMENTATION_SUMMARY.md` - Implementation details
- `NB-Ecosystem/README.md` (12KB) - NB-Ecosystem documentation
- `NB-Ecosystem/MVP_FEATURE_SET_AND_USER_STORIES.md` - Product specs
- `NB-Ecosystem/PHASE1_TECHNICAL_SPECS.md` - Technical specifications

### Configuration
- `lrs_agents/pyproject.toml` - Python project config
- `NB-Ecosystem/package.json` - Node.js dependencies
- `NB-Ecosystem/docker-compose.yml` - Docker orchestration
- `NB-Ecosystem/drizzle.config.ts` - Database config
- `.replit` - Replit environment config

### Source Code
- `lrs_agents/lrs/` - Python package source
- `NB-Ecosystem/server/` - Express backend
- `NB-Ecosystem/client/` - React frontend
- `quantum_optimization.py` - Quantum algorithms

## 10. Deployment Configuration

### Replit Deployment
- **Target**: Autoscale
- **Entry Point**: node index.js
- **Port Mapping**: 8000 (local) → 80 (external)
- **Domain**: c9293dcd-8b5e-4edb-a3b4-151933bf56c0-00-tbhogh34tgtp.riker.replit.dev

### Docker Support
- `NB-Ecosystem/docker-compose.yml` - Multi-service orchestration
- `NB-Ecosystem/Dockerfile.frontend` - Frontend container
- `lrs_agents/docker/` - LRS-Agents Docker configs

### Kubernetes
- `lrs_agents/k8s/` - Kubernetes deployment manifests

---

## Notes

1. **GitSafe** is enabled for version control safety
2. **Helium** is enabled for AI/ML features on Replit
3. The workspace appears to be in a cleanup phase with many files deleted
4. Multiple Python projects coexist with different dependency managers (pip, poetry, uv)
5. The AGENTS.md file contains extensive NeuralBlitz system documentation
