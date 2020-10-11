## structures
folders are grouped by category and date
- category-based codes will be added as soon as possible as they are completed, and might be referred to for several sessions.
- date-based codes will be added as necessary, should we need to break down sessions into smaller examples

## how to use
you can actually use it right in the repo, just `cd` to the relevant directory and follow the instructions in the README inside

### dates
- to not mess up with the git tree, it would be better to check out a new branch each time you want to play around with the codes
- i.e for work on 2020-10-12, codes will be inside folder named `20201012`
```
git checkout -b 20201012
cd 20201012
... // make changes and try stuffs
... // finished making changes
git add . && git commit
git push -u origin 20201012
```
- make sure to checkout master branch and update periodically before/after working on other branches

### category
generally, it would be better to create a new repository, copying the whole folder, but the above methods work as well. some extra things to note:
- go.mod file is shared for the whole repo, `go mod init` will be required
- local import uses fully-qualified URL for the repository and will need to be modified as well (recursive replace should do the trick)
