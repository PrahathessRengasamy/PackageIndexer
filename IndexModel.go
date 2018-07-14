package PackageIndexer

import (
	"sync"
)

//Package Index logic

// Package and its properties.
type Package struct {
	Name         string
	Dependencies []string
}


var package_lock = &sync.RWMutex{}

// package_index holds all of the packages for the package tree.
var package_index = make(map[string][]string)

// Verifies if the given Package's dependencies exist in the package index.

func HasDependencies(pkg Package) bool {
	package_lock.RLock()
	defer package_lock.RUnlock()

	deps := pkg.Dependencies
	for d := range deps {
		p := Package{Name: deps[d], Dependencies: []string{}}
		if IsPresent(p) {
			return true
		}
	}
	return false
}

//  Is this package a dependency in another ie. is the given package is a transitive dependency
func TransitiveDep(pkg Package) bool {
	package_lock.RLock()
	defer package_lock.RUnlock()

	for _, value := range package_index {
		for d := range value {
			if value[d] == pkg.Name {
				return true
			}
		}
	}
	return false
}

// checks if the package is within the index

func IsPresent(pkg Package) bool {
	package_lock.RLock()
	defer package_lock.RUnlock()

	for key := range package_index {
		if key == pkg.Name {
			return true
		}
	}
	return false
}

// Add to index
func AddPackage(pkg string, deps []string) {
	package_lock.Lock()
	p := Package{Name: pkg, Dependencies: deps}
	package_index[p.Name] = p.Dependencies
	package_lock.Unlock()
}

// Remove from index
func DeletePackage(pkg string) {
	package_lock.Lock()
	delete(package_index, pkg)
	package_lock.Unlock()
}
