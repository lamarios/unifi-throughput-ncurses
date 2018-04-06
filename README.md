# Unifi throughput

Unifi throughput is a small piece of software to monitor the throughput of a Unifi controller in a terminal


![](https://i.imgur.com/GwyhKLg.png?raw=true)

![](https://i.imgur.com/U0PnujC.png?raw=true)

## Dependencies

- ncurse needs to be installed

## Installation

Just download the correct archive from the [release page](https://github.com/lamarios/unifi-throughput-ncurses/releases/latest), extract it and run
```
./unifi-throughput -create-config
```
to generate the default configuration file in ~/.config/unifi-throughput

then 
```
./unifi-throughput
```
or
```
./unifi-throughput -config=path_of_config_file.toml
```


## Usage

```
  -config string
    	External configuration file location (default "/home/pi/.config/unifi-throughput/config.toml")
  -create-config
    	Create the default config file /home/pi/.config/unifi-throughput/config.toml THIS WILL OVERWRITE YOUR CURRENT CONFIG AT THE DEFAULT LOCATION
  -version
    	Show version

```

You can also press any key to switch from circle to bar display mode
