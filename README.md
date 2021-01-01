# Unifi throughput

Unifi throughput is a small piece of software to monitor the throughput of a Unifi controller in a terminal


![](https://i.imgur.com/GwyhKLg.png?raw=true)

![](https://i.imgur.com/U0PnujC.png?raw=true)

## Dependencies
- libncurses 6 needs to be installed, if you have an older version of ncurses, you'll need to build it yourself.

## Installation

Just download the correct archive from the [release page](https://github.com/lamarios/unifi-throughput-ncurses/releases/latest), extract it and run
```
./unifi-throughput -create-config
```
to generate the default configuration file in ~/.config/unifi-throughput and edit the newly created config

then 
```
./unifi-throughput
```
or
```
./unifi-throughput -config=path_of_config_file.toml
```

## Build

```shell
git clone https://github.com/lamarios/unifi-throughput-ncurses
cd unifi-throughput-ncurses
make
sudo make install
```

## Usage

```
  -config string
    	External configuration file location (default "~/.config/unifi-throughput/config.toml")
  -create-config
    	Create the default config file ~/.config/unifi-throughput/config.toml THIS WILL OVERWRITE YOUR CURRENT CONFIG AT THE DEFAULT LOCATION
  -version
    	Show version

```

You can also press any key to switch from circle to bar display mode
