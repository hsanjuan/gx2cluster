package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	cid "github.com/ipfs/go-cid"
	client "github.com/ipfs/ipfs-cluster/api/rest/client"
	multiaddr "github.com/multiformats/go-multiaddr"

	gx "github.com/whyrusleeping/gx/gxutil"
)

type pinArgs struct {
	hash *cid.Cid
	name string
}

var (
	peer string
	user string
	pw   string
	pnet string
	ssl  bool
)

func init() {
	flag.StringVar(&peer, "peer", "/ip4/127.0.0.1/tcp/9094", "multiaddress of the IPFS Cluster API")
	flag.StringVar(&user, "user", "", "basic auth username")
	flag.StringVar(&pw, "pw", "", "basic auth pw")
	flag.StringVar(&pnet, "pnet", "", "pnet key")
	flag.BoolVar(&ssl, "ssl", false, "enable ssl")
}

func main() {
	flag.Parse()

	var pm *gx.PM

	gxcfg, err := gx.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pm, err = gx.NewPM(gxcfg)
	if err != nil {
		log.Fatal(err)
	}

	root, err := gx.GetPackageRoot()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(root, gx.PkgFileName)

	var pkg gx.Package
	err = gx.LoadPackageFile(&pkg, path)
	if err != nil {
		log.Fatal(err)
	}

	var deps []string

	depmap, err := pm.EnumerateDependencies(&pkg)
	if err != nil {
		log.Fatal(err)
	}

	for k := range depmap {
		deps = append(deps, k)
	}

	sort.Strings(deps)

	var pins []*pinArgs

	for _, d := range deps {
		var dpkg gx.Package
		err := gx.LoadPackage(&dpkg, pkg.Language, d)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("package %s not found. Install it first with gx install\n", d)
			}
			log.Fatal(err)
		}

		ci, err := cid.Decode(d)
		if err != nil {
			log.Fatal(err)
		}

		pins = append(pins, &pinArgs{
			hash: ci,
			name: fmt.Sprintf("%s-%s", dpkg.Name, dpkg.Version),
		})
	}

	cfg := &client.Config{
		Username: user,
		Password: pw,
		SSL:      ssl,
	}
	addr, err := multiaddr.NewMultiaddr(peer)
	if err != nil {
		log.Fatal(err)
	}

	cfg.APIAddr = addr

	// Uncomment when supported
	// if strings.Contain(peer, "/ipfs/") {
	// 	cfg.PeerAddr = addr
	// 	if pnet != "" {
	// 		secret, err := hex.DecodeString(pnet)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		cfg.ProtectorKey = secret
	// 	}
	// }

	c, err := client.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pins {
		fmt.Printf("pinning: %s\t%s", p.hash, p.name)
		err = c.Pin(p.hash, 0, 0, p.name)
		if err != nil {
			log.Println(err)
		}
	}
}
