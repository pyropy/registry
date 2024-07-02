### Example requests

#### Upload file and register it in the file registry

Example request:
```bash
curl --location 'localhost:8000/v1/files' \
--form 'filePath="/tmp/coin.png"' \
--form 'file=@"/path/to/file"'
```

Example response:
```json
{
  "cid": "/ipfs/bafkreihxhjf56iyjmxshsjsd5hjeqzpsmt2ftsltr6ykdat55kemzyo5uu"
}
```

#### Retrieve file CID from the file registry

Example request:
```bash
curl --location 'http://localhost:8000/v1/files?filePath=%2Ftmp%2Fcoin.png'
```

Example response:

```json
{
  "cid": "/ipfs/bafkreihxhjf56iyjmxshsjsd5hjeqzpsmt2ftsltr6ykdat55kemzyo5uu"
}
```