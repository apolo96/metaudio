# Metaudio CLI & API

Metaudio extracts the metadata from the audio file.

## How to Install CLI using Homebrew

[Click here to go for Install Guide](https://github.com/apolo96/homebrew-metaudio)


## Starting for Developers

Clone this repository, and run the below commands:


### Install thridparty modules

```bash
go mod tidy
```

### Start API:

```bash
go run ./cmd/api
```

### Testing CLI

```bash
go test ./cmd/cli/helpers -v
```

### Build CLI Program.

Build a Free version binary

```bash
go build -o metaudio ./cmd/cli 
```

Build a Pro version binary

```bash
go build -tags pro -o metaudio ./cmd/cli
```

## Usage CLI

Upload audio file to initialize the metadata extract process

```bash
./metaudio upload -filename ./cmd/cli/helpers/data_test/audio.mp3
```

Get audio metadata by ID

```bash
./metaudio get -id 38843709-96c9-4f10-b535-43786f58f234
```

List all audios metadata

```bash
./metaudio list
```

## Usage API

Upload audio file to initialize the metadata extract process

```bash
curl -X POST 'http://localhost:8000/upload' --form 'file=@"beatdoctor.mp3"'
```

To get the metadata of the audio file, you should Copy the ID returned by the command before.

```bash
curl 'http://localhost:8000/request/ID'
```

To List audio metadata run the below command.

```bash
curl 'http://localhost:8000/list' 
```


Use the storage interface allow swap a storage privider easify, like change MySql to Postgresql
