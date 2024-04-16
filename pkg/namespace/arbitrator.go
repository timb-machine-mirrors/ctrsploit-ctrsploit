package namespace

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

type Arbitrator interface {
	Arbitrate(namespace Namespace) (namespaceLevel container.NamespaceLevel, err error)
	PrerequisitesSatisfied() (satisfied bool, err error)
}

var (
	ErrPrerequisiteNotSatisfied = errors.Errorf("not support because of prerequisite not satisfied")
)

func Arbitrate(arbitrator Arbitrator, namespace Namespace) (namespaceLevel container.NamespaceLevel, err error) {
	satisfied, err := arbitrator.PrerequisitesSatisfied()
	if err != nil {
		return
	}
	if satisfied {
		namespaceLevel, err = arbitrator.Arbitrate(namespace)
	} else {
		err = ErrPrerequisiteNotSatisfied
	}
	return
}

func CheckNamespaceLevel(arbitrator Arbitrator) (result map[string]container.NamespaceLevel, names []string, err error) {
	result = map[string]container.NamespaceLevel{}
	satisfied, err := arbitrator.PrerequisitesSatisfied()
	if err != nil {
		return
	}
	namespaces, names, err := ParseNamespaces()
	if err != nil {
		return
	}
	if satisfied {
		for _, namespace := range namespaces {
			result[namespace.Name], err = arbitrator.Arbitrate(namespace)
			if err != nil {
				err = nil
				continue
			}
		}
	} else {
		err = ErrPrerequisiteNotSatisfied
	}
	return
}

func GetNamespaceLevel(arbitrator Arbitrator, ns string) (level container.NamespaceLevel, err error) {
	if !CheckNamespaceValid(ns) {
		err = fmt.Errorf("ns %s not valid", ns)
		awesome_error.CheckWarning(err)
		return
	}
	namespaceLevels, _, err := CheckNamespaceLevel(arbitrator)
	if err != nil {
		return
	}
	level, ok := namespaceLevels[ns]
	if !ok {
		err = fmt.Errorf("ns %s not support", ns)
		awesome_error.CheckWarning(err)
		return
	}
	return
}

func CheckNamespaceValid(ns string) (valid bool) {
	_, valid = container.NamespaceMapName2Type[ns]
	return
}
