package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	// "io/ioutil"

  "github.com/alex-src/vault/config"
  "github.com/alex-src/vault/server"
  "github.com/alex-src/vault/vault"
  "github.com/GeertJohan/go.rice"
  // "github.com/alex-src/vault/helper/mlock"
)

var (
	cfg            *config.Config
	cfgPath        string
	devMode        bool
	devVaultCh     chan struct{}
	err            error
	nomadTokenFile string
	printVersion   bool
	wrappingToken  string
	newWrappingToken  string
	staticAssets   *rice.Box
)

func init() {
	// customized help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMessage)
	}

	// cmd line args
	flag.BoolVar(&devMode, "dev", false, "Set to true to save time in development. DO NOT SET TO TRUE IN PRODUCTION!!")
	flag.BoolVar(&printVersion, "version", false, "Display vault's version and exit")
	flag.StringVar(&wrappingToken, "token", "", "Token generated from approle (must be wrapped!)")
	flag.StringVar(&nomadTokenFile, "nomad-token-file", "", "If you are using Nomad, this file should contain a secret_id")
	flag.StringVar(&cfgPath, "config", "", "The path of the deployment config HCL file")
}

func main() {
	// if --version, print and exit success
	flag.Parse()
	if printVersion {
		log.Println(versionString)
		os.Exit(0)
	}

	// if dev mode, run a localhost dev vault instance
	// if devMode {
///////////////
// need 1 wrapping token or 3/5 unseal tokens
//////////////
	var newUnsealTokens []string 
	var unsealTokens [5]string
	wrappingToken = "d3551e46-fa64-e44f-bf19-20aa522fb2b0"
	unsealTokens[0] =
		"229968db3c6db2d75bdd0067ebee0367612525915c3d90bd29f1e7ba43fb976063"
	unsealTokens[1] =
		"00072b4132791677df0171f76ddca7f621f288978cfd86fc69508886a65dd33fff"
	unsealTokens[2] =
		"68e38dac27c870f5bc3b196d99d6b7117f806ae91bb5372ad489e5de2eea18a9bb"
	unsealTokens[3] =
		"72ff435f39f84cdd12411dd7c27eb08c6ab54c8310df34ef3d9574c6cb00ac56c4"
	unsealTokens[4] =
		"1a1be5b22c492a5f717b754d3674a06b34c7aefd87978539804c199e43b767c080"
		cfg, devVaultCh, newUnsealTokens, newWrappingToken, err = config.LoadConfigDev()
////////////////
		// cfg, devVaultCh, unsealTokens, wrappingToken, err = config.LoadConfigDev()
		log.Println("\n[INFO ]: OLD wrapping token: " + wrappingToken)
		log.Println("[INFO ]: OLD unseal tokens:")
		fmt.Println(unsealTokens, "\n")
		///
		log.Println("\n[INFO ]: NEW config: \n")
		fmt.Println(cfg.Vault, "\n")
		log.Println("[INFO ]: NEW wrapping token: " + newWrappingToken)
		log.Println("[INFO ]: NEW unseal tokens:\n" + strings.Join(newUnsealTokens, "\n"))
		fmt.Println(devInitString, "\n")
	// } else {
	// 	cfg, err = config.LoadConfigFile(cfgPath)
	// }
	// if err != nil {
	// 	log.Fatalf("[ERROR]: Launching vault: %s", err.Error())
	// }

	// if !cfg.DisableMlock {
	// 	if err := mlock.LockMemory(); err != nil {
	// 		log.Fatalf(mlockError, err.Error())
	// 	}
	// }

	// configure vault server settings and token
	vault.SetConfig(cfg.Vault)
	log.Println("[INFO ]: Vault configs :\n")
	fmt.Println(cfg.Vault, "\n")

	// if bootstrapping options are provided, do so immediately
	if wrappingToken != "" {
		if err := vault.Bootstrap(wrappingToken); err != nil {
			log.Fatalf("[ERROR]: Bootstrapping vault %s", err.Error())
		} else {
			log.Println("[INFO ]: Bootstrapping vault Success\n")
		}
	}
	
	// else if nomadTokenFile != "" {
	// 	raw, err := ioutil.ReadFile(nomadTokenFile)
	// 	if err != nil {
	// 		log.Fatalf("[ERROR]: Could not read token file: %s", err.Error())
	// 	}
	// 	if err := vault.BootstrapRaw(string(raw)); err != nil {
	// 		log.Fatalf("[ERROR]: Bootstrapping vault: %s", err.Error())
	// 	}
	// }

	// start listener
  Rice := rice.Config{
        LocateOrder: []rice.LocateMethod{rice.LocateFS, rice.LocateEmbedded, rice.LocateAppended},
       }

	// if !devMode {
		staticAssets, err = Rice.FindBox("public")
		if err != nil {
			log.Fatalf("[ERROR]: Static assets not found. Build them with npm first. \n%s", err.Error())
		}
	// }
	go server.StartListener(*cfg.Listener, staticAssets)
	fmt.Printf(versionString + initString, "\n")

	// wait for shutdown signal, and cleanup after
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	log.Println("\n\n==> vault shutdown triggered")

	// shut down vault dev server, if it was initialized
	if devVaultCh != nil {
		close(devVaultCh)
	}

	// shut down listener, with a hard timeout
	server.StopListener(5 * time.Second)

	// extra grace time
	time.Sleep(time.Second)
	os.Exit(0)
}

const versionString = "vault version: v0.0.3"

const devInitString = `

---------------------------------------------------
Starting local vault instance...
Your unseal token and root token can be found above
`

const initString = `
vault successfully bootstrapped to vault

  .
  ...             ...
  .........       ......
   ...........   ..........
     .......... ...............
     .............................
      .............................
         ...........................
        ...........................
        ..........................
        ...... ..................
      ......    ...............
     ..        ..      ....
    .                 ..


`

const mlockError = `
Failed to use mlock to prevent swap usage: %s

vault uses mlock similar to Vault. See here for details:
https://www.vaultproject.io/docs/configuration/index.html#disable_mlock

To enable mlock without launching vault as root:
sudo setcap cap_ipc_lock=+ep $(readlink -f $(which vault))

To disable mlock entirely, set disable_mlock to "1" in config file
`

const helpMessage = `Usage: vault [options]
See https://github.com/Caiyeon/vault/wiki for details

Required Arguments:

  -config=config.hcl      The deployment config file
                          See https://github.com/Caiyeon/vault/blob/master/config/sample.hcl
                          for a full list of options

Optional Arguments:

  -token=<uuid>           A wrapping token which contains a secret_id
                          Can be provided after launch, on Login page
                          Generate with 'vault write -f -wrap-ttl=5m auth/approle/role/vault/secret-id'

  -nomad-token-file       A path to a file containing a raw token.
                          Not recommended unless approle is unavailable,
						  in the case of Nomad for example.

  -version                Print the version and exit

  -dev                    Launch vault in dev mode
                          A localhost dev vault instance will be launched
`

