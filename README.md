# To evaluate a solution

Just call it with your csv-file as the single argument:

```
./hackathon test.csv
Precision 0.081
Recall: 0.082
F1-score: 0.082
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
