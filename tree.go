package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/borodean/asciitree"
)

type Sites map[string][]string
type Nodes map[string][]Node

type Node struct {
	Name     string
	Children []Node
}

func main() {
	sites := make(Sites)
	trees := make(Nodes)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 10 {
			continue
		}
		if strings.Contains(line, "?") {
			line = strings.Split(line, "?")[0]
		}
		line = strings.TrimRight(line, "/")
		u, err := url.Parse(line)
		if err != nil {
			continue
		}
		port := u.Port()

		if port == "" {
			if u.Scheme == "https" {
				port = "443"
			} else {
				port = "80"
			}

		}
		host := u.Hostname() + ":" + port
		if u.Path != "/" {
			sites[u.Scheme+"://"+host] = append(sites[u.Scheme+"://"+host], host+u.Path)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	for site, pths := range sites {
		for i := range pths {
			trees[site] = AddToTree(trees[site], strings.Split(sites[site][i], "/"))
		}
	}

	for site, tree := range trees {
		utree := asciitree.NewDir(site)
		for _, n := range tree[0].Children {
			parseNode(utree, n)
		}

		utree.Sort(asciitree.WithDirsFirst(true))
		fmt.Println(utree)
		fmt.Println()
	}

}

func parseNode(ascnode *asciitree.Node, node Node) {
	newn := ascnode.AddDir(node.Name)
	for _, c := range node.Children {
		parseNode(newn, c)
	}
}

func AddToTree(root []Node, names []string) []Node {
	if len(names) > 0 {
		var i int
		for i = 0; i < len(root); i++ {
			if root[i].Name == names[0] { //already in tree
				break
			}
		}
		if i == len(root) {
			root = append(root, Node{Name: names[0]})
		}
		root[i].Children = AddToTree(root[i].Children, names[1:])
	}
	return root

}
