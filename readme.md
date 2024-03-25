# MetAudio CLI & API

Metaudio extract the metadata from audio file.

## Starting for Developers

Clone this repository, and run the below commands:

```bash
go mod tidy
```
Starting localhost API:

```bash
cd cmd/api 
```

```bash
go run .
```

Build CLI Program.


```bash
go build -o metaudio cmd/cli/main.go
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
