# beer-me

## Installation

```bash
go get github.com/rafecolton/beer-me
```

## Usage

``` bash
# example usage: beer-me <Caskfile>

beer-me Caskfile
```

## Example Caskfile

```ruby
# comments like this
brew caskroom/cask/brew-cask

brew tmux
brew reattach-to-user-namespace
brew mercurial
brew gnu-sed # comments can go here too
brew gnu-tar
brew autoenv

# blank lines are also OK

cask caffeine
cask vagrant
cask virtualbox
cask google-chrome
cask iterm2
cask quicksilver
cask sizeup
cask mailplane
cask spotify
cask karabiner
```
