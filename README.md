# Spica

Simple testing tool for competitive programming problem setters



## config sample
```toml
timelimit=30.0
workers=4

[[languages]]
  ext=".cs"
  compile="mcs /r:System.Numerics.dll -out:a.exe $SRC"
  exec="mono a.exe"

[[languages]]
  ext=".cpp"
  compile="g++ -O2 -std=gnu++11 -o a.out $SRC"
  exec="./a.out"

[[languages]]
  ext=".txt"
  compile="go version"
  exec="cat $SRC"

```
