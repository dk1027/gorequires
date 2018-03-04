# gorequires
Checks if something exists

**func Env(key string) string**

Checks if an environment variable is defined and returns its value.
```
value := Env("myEnvironmentVariable")
```

**func File(path string) os.FileInfo**

Checks if a file exists and returns the corresponding FileInfo.
```
fi := File("path/to/file")
```
