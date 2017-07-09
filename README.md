# beer-me

## Installation

Checksums available on the [releases page](https://github.com/rafecolton/beer-me/releases)

```bash
curl -sL https://github.com/rafecolton/beer-me/releases/download/v0.1.0/beer-me-v0.1.0 \
  -o /usr/local/bin/beer-me && chmod +x /usr/local/bin/beer-me
```

or

```bash
go get github.com/rafecolton/beer-me
```

## Usage

``` bash
# example usage: beer-me <Caskfile>

# default Caskfile is "Caskfile"
beer-me

# or specify a Caskfile
beer-me Caskfile.local
```

## Example Caskfile

Syntax is Gemfile-like

```ruby
# comments like this

# add taps
tap caskroom/cask

# install brews
brew tmux
brew reattach-to-user-namespace
brew mercurial
brew gnu-sed # comments can go here too
brew gnu-tar
brew autoenv

# blank lines are also OK

# install casks
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
