
# This make file uses composition to keep things KISS and easy.
# In the boilerpalte make files dont do any includes, because you will create multi permutations of possibilities.



# git include
include ./boilerplate/core/help.mk
include ./boilerplate/core/os.mk
include ./boilerplate/core/gitr.mk
include ./boilerplate/core/go.mk


## Print all settings
print: ## print

	
	$(MAKE) os-print
	
	$(MAKE) gitr-print

	$(MAKE) go-print
	

## So high
high: ## high
	@echo i wanna get...
