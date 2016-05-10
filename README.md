# To evaluate a solution

Just call it with your csv-file as the single argument:

```
./hackathon test.csv
80% correct
```

# To generate a new answers file

```
./hackathon --hash-file answers.csv
```

This will create a file called answers.hashed. Commit and push that file:

```
git add aswers.hashed
git commit -m "New updated answers"
git push
```

DON'T COMMIT ANSWERS.CSV

The automatic build will create a new binary for you. If you want to do it
manually:

```
go get -u github.com/jteeuwen/go-bindata/...  # only once
go-bindata answers.hashed
go build
```
