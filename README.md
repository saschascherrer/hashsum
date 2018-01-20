# hashsum
command line tool for calculating and checking hashsums of files.

  [Installation](#installation)
| [Usage](#usage)
| [Bugs](#bugs)
| [License](#license)

## Installation
### From Source
  * Make sure you have the go programming language available on
    your operating system. If you need to install them, please check out
    the [project's website](https://golang.org/) for documentation on
    how to [download](https://golang.org/dl/) and
    [configure](https://golang.org/doc/install) Go for your system.
  * [Download the source](https://github.com/saschascherrer/hashsum/archive/master.zip) or
    clone the repository. Make sure the source is located within your
    ``$GOPATH/src`` directory.
    Reccomendation is ``$GOPATH/src/github.com/saschascherrer/hashsum``
  * Build the binary using ``go build`` in the project root directory
    (``$GOPATH/src/github.com/saschascherrer/hashsum``).
    If you need to crosscompile, go has you covered.  
    Example on linux for windows:  
    ``$> GOOS=windows GOARCH=amd64 go build -o hashsum.exe main.go``
### Download binaries
  * The binaries are built with the standard go compiler on a linux system.
    The Windows binary is crosscompiled. This means, that the binaries contain
    everything needed for the program to run, so there are no dependencies.
  * Go to the [Releases](https://github.com/saschascherrer/hashsum/releases)
    Page on GitHub and select the binary for your platform.
  * Make sure the hashsum of the binary matches the provided hash.  
    Hint: Using hashsum to check itself cannot ensure you get what you expect.

## Usage
```
hashsum [-a md5|sha1|sha256|sha512] [-o <outfile>] <file> [<file> [...]]
	calculate hashsums over the specified files, optionally writing it to <outfile>
```
```
hashsum [-a md5|sha1|sha256|sha512] -r <hashstring> <file>
	calculate hashsum of specified file and compare it to the provided hashsum
```

Planned:
```
hashsum [-a md5|sha1|sha256|sha512] -c <sumfile> [<file>, [<file>[, ...]]]
  compares priviously generated hashsum files with present files. Usually
	checks every entry of the sumfile, but can be limited to files of interest
	by specifying them as arguments
```

## Bugs
Bugs can and do happen. If you are misfortunate enough to experience any of them
with the hashsum tool, I would be happy to hear what went wrong. So please feel
free to report any bugs to the
[GitHub Issue Tracker](https://github.com/saschascherrer/hashsum/issues).
You are also encouraged but not obliged to look into the source
(it is not too bulky) and submit bugfixes yourself.

## License
This helper tool is licenced under the MIT License  
If you commit changes to this repository, you must sign-off
your commit to confirm your right to make the contribution
(see [DeveloperCertificate.org](https://developercertificate.org/)).
This can easily be done by using ``git commit -s`` to the commit.
This also means, you must not submit anonymous contributions.
