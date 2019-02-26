
[EsaExporter](https://github.com/NAL-6295/EsaExporter)のgolang版

This is export from esa.io and restructure to local.

esa.io is https://esa.io

## export data and create data
- esa.io target team all article datas.
  - json
  - markdown
- Image files by img tag
- Modify img tag url for local filesystem.
- Create index file.Name is index.md location exported root path.

### sample
```
goEsaExporter -mode=md -root="d:\\" -team="teamname" -token="api token"
```

### help
```
goEsaExporter -h
```

### Local file borwsing
1. Install below extention.'Markdown viewer'
https://chrome.google.com/webstore/detail/markdown-viewer/ckkdlimhmcjmikdlpkmbgfkaikojcbjk?hl=ja&gl=JP
1. Set 'on' allow access to file URLs.
1. Open index.md,it is location exported root path.
