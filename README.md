# GoEnv
Manage GoLang virtual enviroments.

## Installation

```bash
go get -u github.com/moisespsena/go-goenv/goenv
```

## Usage

### Init new enviroment

```bash
goenv init env1
```

### List enviroments

```bash
goenv ls
```

### Activate repository:

```bash
eval $(goenv activate env1)
```

or (see to [Database](#database) section)

```bash
source $(goenv db)/env1/activate
```

#### Deactivate it

```bash
deactivate
```

### Remove repository:

Move to trash directory (`DB_DIR/.trash`, see to [Database](#database) section):
```bash
goenv rm env1
```

Remove permanently:
```bash
goenv rm -p env1
```

### Database

Get database path:

```bash
goenv db
```

#### Custom database path

##### With command paramenter
For use custom database path, set the `-d` flag:

```bash
goenv -d ~/custom-db args...
```

Examples:

```bash
goenv -d ~/custom-db init env2
goenv -d ~/custom-db ls
```

##### With enviroment variable

Set the enviroment variable:
 
```bash
export GOENVDB=~/custom-db
```

or

```bash
GOENVDB=~/custom-db goenv args...
```

## Thank's!

By [Moises P. Sena](https://github.com/moisespsena).