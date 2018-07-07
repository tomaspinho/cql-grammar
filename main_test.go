package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

const DIRNAME = "./test/antlr4/cql3/"

func TestParse(t *testing.T) {

	dir, err := ioutil.ReadDir(DIRNAME)

	if err != nil {
		panic("Cannot run tests, cql3 dir does not exist. Error:" + err.Error())
	}

	for _, f := range dir {
		fName := f.Name()

		split := strings.Split(fName, ".")
		fNameNoExt := string(split[0])
		fExt := string(split[1])

		// Skip parseTree
		if fExt != "cql" {
			continue
		}

		p := initParser(DIRNAME + fName)
		tree := p.Statements()

		treeS := tree.ToStringTree([]string{}, p)

		expectedTreeBytes, err := ioutil.ReadFile(DIRNAME + fNameNoExt + ".parseTree")

		if err != nil {
			panic("Error reading parsedTree: " + err.Error())
		}

		expectedTree := string(expectedTreeBytes)
		lines := strings.Split(expectedTree, "\n")

		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}

		expectedTree = strings.Join(lines, " ")

		if treeS != expectedTree {
			println(treeS)
			println(expectedTree)
			t.Error("Parsed tree does not match expected tree for file: " + DIRNAME + fName)
		}
	}
}
