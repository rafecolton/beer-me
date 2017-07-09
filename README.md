# beer-me

## Installation

Checksums available on the [releases page](https://github.com/rafecolton/beer-me/releases)

```bash
export OUTDIR=/usr/local/bin && \
  curl -sL https://github.com/rafecolton/beer-me/releases/download/v0.2.1/beer-me-v0.2.1 \
  -o $OUTDIR/beer-me && chmod +x $OUTDIR/beer-me
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
tap git-duet/tap

# install brews
brew tmux
brew reattach-to-user-namespace
brew git-duet
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
