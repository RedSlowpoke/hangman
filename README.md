# hangman
A repo for training on writing proper (almost) program

## The task
Checklist:
- [x] Description in `README.MD` using Markdown.
- [x] An open License in `LICENSE`.
- [ ] Declaration of dependencies and building.
- [ ] Static code analyzer.
- [ ] Test.
- [ ] CI (?Travis?).
- [ ] There is a code coverage analyzer (?CodeCov?).

Anything that requires automation (build, test, lint) should be possible done w/o any additional steps or with one command in console

## Description
The game provides only CLI  
The rules are:
- you have a random word guessed for you
- you may guess one letter case insensitive, if it's not right you lose one live
- you play till you make too many mistakes or you guess a full word
