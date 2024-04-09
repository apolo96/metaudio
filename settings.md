# VS Code Settings

## Gopls support build flags

Go to VSCode settings to find gopls, and edit the settings.json file.

Then add the following code: 

```json
"build.buildFlags": [
      "-tags=pro"
]
```

`pro` is a custom build flag in this project.

For more info, you can visit gopls doc: https://github.com/golang/tools/blob/master/gopls/doc/settings.md

