# i18n Test harness

This shows how to use the i18n GoogleSheet tool.

The GoogleSheets are here:
https://drive.google.com/drive/folders/1SvB8gOFuvD1IF8baU63Wcp0XTOVT1iSV?usp=sharing

- see the config sheet for what config to use.

## TODO

0. DONE: Move to getcouragenow.org drive.

1. finish config.

2. Make a simple simple Hugo site.

- https://github.com/bep/docuapi

	- Same as https://github.com/linode/docs, https://www.linode.com/docs/

## ROADMAP

- Flip over to NOT using Google sheets.

	- Data Store is the Git tree !
	- Trans tool is whatever you want it to be.

- Review of the auto translation. Since its just a PR, a PR bot can get another translator to check their work.

	- The Data format MUST model the human overrides so they are not lost.

- Providence. So from the GUI shows exactly where the translation data came from.

- Cross Check.

	- The GUI and the Translation data are decoupled. Need to tool to compare both sides and tell us what is mis-linked, and what translations are missing

	- Git Bot that does this all the time.
		- Sink is github issues.

- Strong type codegen.

	- From the JSON we code gen the ARB. We still wan to run off non arb when in Providence mode.

- GUI so you do not need VSCODE etc

	- Flutter GUI that represents the Translation data model.
	- GRPC API !!
		- Data Sources:
			- FS (stage 1) ( for devs)
			- Git Proxy (stage 2) (for complete non devs) that talks to the MASTER, holds the users changes and then pushes  into the master for the user as the users PR.


## Hosting

## Mage & CI

Mage is used for local and CI builds and deploys.

Just tell Github actions to install golang and our bootstrap and then its easy.

