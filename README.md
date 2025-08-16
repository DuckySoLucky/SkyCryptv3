# SkyCrypt v3 Monorepo

This is the monorepo for SkyCrypt v3, containing both frontend and backend as git submodules.

## Structure

- `frontend/` - SkyCrypt v2 frontend (feat/go-api branch) from [SkyCryptWebsite/SkyCryptv2](https://github.com/SkyCryptWebsite/SkyCryptv2/tree/feat/go-api/)
- `backend/` - SkyCrypt Backend from [DuckySoLucky/SkyCrypt-Backend](https://github.com/DuckySoLucky/SkyCrypt-Backend)

## Getting Started

To clone this repository with all submodules:

```bash
git clone --recurse-submodules https://github.com/your-username/SkyCryptv3.git
```

If you've already cloned without submodules, initialize them:

```bash
git submodule update --init --recursive
```

## Working with Submodules

To update submodules to their latest commits:

```bash
git submodule update --remote
```

To pull changes in a specific submodule:

```bash
cd frontend
git pull origin main
cd ..
git add frontend
git commit -m "Update frontend submodule"
```
