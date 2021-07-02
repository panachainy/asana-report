echo $(cat Makefile | head -1 | grep VERSION= | cut -d '=' -f 2)
