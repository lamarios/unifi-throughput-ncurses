# Unifi throughput

Unifi throughput is a small piece of software to monitor the throughput of a Unifi controller in a terminal


## Dependencies

- ncurse needs to be installed

## Installation

Just download the correct archive from the [release page](https://github.com/lamarios/unifi-throughput-ncurses/releases), extract it and run
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
