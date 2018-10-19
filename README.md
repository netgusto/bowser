# ⭐ Bowser - Dynamic Default Browser Switch

Bowser is a MacOS application that selects the browser to open for a URL based on rules you define.

My use case is to have 2 default browsers : one for development (Chrome), and one for surf (Safari).

## Install from binary releases

Download the latest binary release (`.dmg`) here: https://github.com/netgusto/bowser/releases

Open the `.dmg` image, drag Bowser to your `/Applications` folder.

Then follow steps described in "Setup as default browser" below.

## Install from source

The installation from source requires the Apple clang environment (XCode) and Go 1.8+

```sh
$ git clone https://github.com/netgusto/bowser
$ cd bowser
$ make install
```

Then follow steps described in "Setup as default browser" below.

## Setup as default browser

Bowser has to be defined as default browser to operate.

Once Bowser is installed in `/Applications`, go to your Mac **System Preferences** > **General** tab, and choose **Bowser** in the list of browsers.

## Configuration

During first run, bowser will create a default config file at `~/.config/bowser/config.yml`.

```yml
debug: false
browsers:
- alias: Default
  app: Safari
```

This default config sets Safari as the default browser. You may now edit the file to add browsers and set regex rules corresponding to your likings.

Example setup:

```yml
debug: false

browsers:
- alias: Default
  app: Safari

- alias: Dev
  app: Google Chrome
  match:
  - ^https?://127.0.0.1
  - ^https?://localhost  
```

**debug** set to true will forward debug messages to the syslog. Enable only for debugging purposes.

## License

See the LICENSE file.
